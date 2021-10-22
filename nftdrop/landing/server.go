// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package landing

import (
	"context"
	"errors"
	"html/template"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/landing/controllers"
	"ultimatedivision/nftdrop/subscribers"
	"ultimatedivision/nftdrop/whitelist"
	"ultimatedivision/pkg/ratelimit"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("nftdrop web server error")
)

// Config contains configuration for nftdrop web server.
type Config struct {
	Address   string `json:"address"`
	StaticDir string `json:"staticDir"`
}

// Server represents nftdrop web server.
//
// architecture: Endpoint
type Server struct {
	log    logger.Logger
	config Config

	listener    net.Listener
	server      http.Server
	rateLimiter *ratelimit.RateLimiter

	templates struct {
		index *template.Template
	}
}

// NewServer is a constructor for nftdrop web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, whitelist *whitelist.Service, subscribers *subscribers.Service) *Server {
	server := &Server{
		log:      log,
		config:   config,
		listener: listener,
	}

	// TODO: read from config.
	server.rateLimiter = ratelimit.NewRateLimiter(time.Minute*5, 5, 10000)

	whitelistController := controllers.NewWhitelist(log, whitelist)
	subscribersController := controllers.NewEmails(log, subscribers)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	whitelistRouter := apiRouter.PathPrefix("/whitelist").Subrouter()
	whitelistRouter.Handle("/{address}", server.rateLimit(http.HandlerFunc(whitelistController.Get))).Methods(http.MethodGet)

	SubscribersRouter := apiRouter.PathPrefix("/subscribers").Subrouter()
	SubscribersRouter.Handle("", server.rateLimit(http.HandlerFunc(subscribersController.Create))).Methods(http.MethodPost)

	fs := http.FileServer(http.Dir(server.config.StaticDir))
	router.PathPrefix("/static/").Handler(server.brotliMiddleware(http.StripPrefix("/static", fs)))
	router.PathPrefix("/").Handler(http.HandlerFunc(server.appHandler)).Methods(http.MethodGet)

	server.server = http.Server{
		Handler: router,
	}

	return server
}

// Run starts the server that host webapp and api endpoint.
func (server *Server) Run(ctx context.Context) (err error) {
	err = server.initializeTemplates()
	if err != nil {
		return Error.Wrap(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		return Error.Wrap(server.server.Shutdown(context.Background()))
	})
	group.Go(func() error {
		defer cancel()
		err = server.server.Serve(server.listener)
		isCancelled := errs.IsFunc(err, func(err error) bool { return errors.Is(err, context.Canceled) })
		if isCancelled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return Error.Wrap(err)
	})

	return Error.Wrap(group.Wait())
}

// Close closes server and underlying listener.
func (server *Server) Close() error {
	return Error.Wrap(server.server.Close())
}

// appHandler is web app http handler function.
func (server *Server) appHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	cspValues := []string{
		"default-src 'self'",
		"connect-src 'self'",
		"frame-ancestors 'self'",
		"frame-src 'self'",
		"img-src 'self' data:",
		"font-src 'self'",
		"style-src 'self' 'unsafe-inline'",
		"script-src https://www.googletagmanager.com 'self'",
	}

	featurePolicies := []string{
		"accelerometer 'none'",
		"ambient-light-sensor 'none'",
		"autoplay 'self'",
		"battery 'none'",
		"camera 'none'",
		"display-capture 'none'",
		"document-domain 'none'",
		"encrypted-media 'none'",
		"execution-while-not-rendered 'none'",
		"execution-while-out-of-viewport 'none'",
		"fullscreen 'none'",
		"geolocation 'none'",
		"gyroscope 'none'",
		"layout-animations 'none'",
		"legacy-image-formats 'none'",
		"magnetometer 'none'",
		"microphone 'none'",
		"midi 'none'",
		"navigation-override 'none'",
		"oversized-images 'none'",
		"payment 'none'",
		"picture-in-picture 'none'",
		"publickey-credentials-get 'none'",
		"sync-xhr 'none'",
		"usb 'none'",
		"vr 'none'",
		"wake-lock 'none'",
		"xr-spatial-tracking 'none'",
	}

	header.Set("Content-Security-Policy", strings.Join(cspValues, "; "))
	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("Feature-Policy", strings.Join(featurePolicies, "; "))
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "no-referrer")

	if server.templates.index == nil {
		server.log.Error("index template is not set", nil)
		return
	}

	if err := server.templates.index.Execute(w, nil); err != nil {
		server.log.Error("index template could not be executed", err)
		return
	}
}

// seoHandler is used by web crawlers to improve seo.
func (server *Server) seoHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("Cache-Control", "public,max-age=31536000,immutable")
	header.Set("Content-Type", mime.TypeByExtension(".txt"))
	header.Set("X-Content-Type-Options", "nosniff")

	_, err := w.Write([]byte("User-agent: *\nDisallow:\nDisallow: /cgi-bin/"))
	if err != nil {
		server.log.Error("could not return robots.txt file", err)
	}
}

// initializeTemplates is used to initialize all templates.
func (server *Server) initializeTemplates() (err error) {
	server.templates.index, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "dist", "index.html"))
	if err != nil {
		server.log.Error("dist folder is not generated. use 'npm run build' command", err)
		return err
	}

	return nil
}

// rateLimit is an handler that prevents from multiple requests from single ip address.
// TODO: apply to api requests.
func (server *Server) rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, err := server.getIP(r)
		if err != nil {
			server.log.Error("could not get remote ip", err)
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		server.log.Error(ip, nil)

		isAllowed := server.rateLimiter.IsAllowed(ip, time.Now().UTC())
		if !isAllowed {
			server.log.Debug("rate limit exceeded, ip:" + ip)
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// TODO: place to internal.
func (server *Server) getIP(r *http.Request) (ip string, err error) {
	ip, _, err = net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	if ip == "127.0.0.1" {
		ip = r.Header.Get("X-Real-IP")
	}
	return ip, nil
}

// brotliMiddleware is used to compress static content using brotli to minify resources if browser support such decoding.
func (server *Server) brotliMiddleware(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		isBrotliSupported := strings.Contains(r.Header.Get("Accept-Encoding"), "br")
		if !isBrotliSupported {
			fn.ServeHTTP(w, r)
			return
		}

		info, err := os.Stat(server.config.StaticDir + strings.TrimPrefix(r.URL.Path, "/static") + ".br")
		if err != nil {
			fn.ServeHTTP(w, r)
			return
		}

		extension := filepath.Ext(info.Name()[:len(info.Name())-3])
		w.Header().Set("Content-Type", mime.TypeByExtension(extension))
		w.Header().Set("Content-Encoding", "br")

		newRequest := new(http.Request)
		*newRequest = *r
		newRequest.URL = new(url.URL)
		*newRequest.URL = *r.URL
		newRequest.URL.Path += ".br"

		fn.ServeHTTP(w, newRequest)
	})
}
