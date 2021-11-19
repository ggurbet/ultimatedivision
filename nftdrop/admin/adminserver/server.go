// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package adminserver

import (
	"context"
	"errors"
	"html/template"
	"net"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/admin/adminauth"
	"ultimatedivision/admin/admins"
	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/templatefuncs"
	"ultimatedivision/nftdrop/admin/adminserver/controllers"
	"ultimatedivision/nftdrop/subscribers"
	"ultimatedivision/nftdrop/whitelist"
	"ultimatedivision/pkg/auth"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("admin web server error")
)

// Config contains configuration for admin web server.
type Config struct {
	Address   string `json:"address"`
	StaticDir string `json:"staticDir"`

	Auth struct {
		CookieName string `json:"cookieName"`
		Path       string `json:"path"`
	} `json:"auth"`
}

// Server represents admin web server.
//
// architecture: Endpoint
type Server struct {
	log    logger.Logger
	config Config

	listener net.Listener
	server   http.Server

	authService *adminauth.Service
	cookieAuth  *auth.CookieAuth

	templates struct {
		auth        controllers.AuthTemplates
		admins      controllers.AdminTemplates
		whitelist   controllers.WhitelistTemplates
		subscribers controllers.SubscribersTemplates
	}
}

// NewServer is a constructor for admin web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, authService *adminauth.Service, admins *admins.Service, whitelist *whitelist.Service, subscribers *subscribers.Service) (*Server, error) {
	server := &Server{
		log:    log,
		config: config,
		cookieAuth: auth.NewCookieAuth(auth.CookieSettings{
			Name: config.Auth.CookieName,
			Path: config.Auth.Path,
		}),
		authService: authService,
		listener:    listener,
	}

	err := server.initializeTemplates()
	if err != nil {
		return nil, Error.Wrap(err)
	}

	router := mux.NewRouter()
	authController := controllers.NewAuth(server.log, server.authService, server.cookieAuth, server.templates.auth)
	router.HandleFunc("/login", authController.Login).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/logout", authController.Logout).Methods(http.MethodGet)

	adminsRouter := router.PathPrefix("/admins").Subrouter()
	adminsRouter.Use(server.withAuth)
	adminsController := controllers.NewAdmins(log, admins, server.templates.admins)
	adminsRouter.HandleFunc("", adminsController.List).Methods(http.MethodGet)
	adminsRouter.HandleFunc("/create", adminsController.Create).Methods(http.MethodGet, http.MethodPost)
	adminsRouter.HandleFunc("/update/{id}", adminsController.Update).Methods(http.MethodGet, http.MethodPost)

	whitelistRouter := router.PathPrefix("/whitelist").Subrouter()
	whitelistRouter.Use(server.withAuth)
	whitelistController := controllers.NewWhitelist(log, whitelist, server.templates.whitelist)
	whitelistRouter.HandleFunc("", whitelistController.List).Methods(http.MethodGet)
	whitelistRouter.HandleFunc("/create", whitelistController.Create).Methods(http.MethodGet, http.MethodPost)
	whitelistRouter.HandleFunc("/delete/{address}", whitelistController.Delete).Methods(http.MethodGet)
	whitelistRouter.HandleFunc("/set-password", whitelistController.SetPassword).Methods(http.MethodGet, http.MethodPost)

	subscribersRouter := router.PathPrefix("/subscribers").Subrouter()
	subscribersRouter.Use(server.withAuth)
	subscribersController := controllers.NewSubscribers(log, subscribers, server.templates.subscribers)
	subscribersRouter.HandleFunc("", subscribersController.List).Methods(http.MethodGet)
	subscribersRouter.HandleFunc("/delete/{email}", subscribersController.Delete).Methods(http.MethodGet)

	server.server = http.Server{
		Handler: router,
	}

	return server, nil
}

// Run starts the server that host webapp and api endpoint.
func (server *Server) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		return server.server.Shutdown(context.Background())
	})
	group.Go(func() error {
		defer cancel()
		err := server.server.Serve(server.listener)
		isCancelled := errs.IsFunc(err, func(err error) bool { return errors.Is(err, context.Canceled) })
		if isCancelled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}

// Close closes server and underlying listener.
func (server *Server) Close() error {
	return server.server.Close()
}

// initializeTemplates initializes and caches templates for managers controller.
func (server *Server) initializeTemplates() (err error) {
	server.templates.auth.Login, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "auth", "login.html"))
	if err != nil {
		return err
	}

	server.templates.admins.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "admins", "list.html"))
	if err != nil {
		return err
	}
	server.templates.admins.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "admins", "create.html"))
	if err != nil {
		return err
	}
	server.templates.admins.Update, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "admins", "update.html"))
	if err != nil {
		return err
	}

	server.templates.whitelist.List, err = template.New("list.html").Funcs(template.FuncMap{
		"Iter": templatefuncs.Iter,
		"Inc":  templatefuncs.Inc,
		"Dec":  templatefuncs.Dec,
	}).ParseFiles(
		filepath.Join(server.config.StaticDir, "whitelist", "list.html"),
		filepath.Join(server.config.StaticDir, "whitelist", "pagination.html"))
	if err != nil {
		return err
	}
	server.templates.whitelist.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "whitelist", "create.html"))
	if err != nil {
		return err
	}
	server.templates.whitelist.SetPassword, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "whitelist", "setPassword.html"))
	if err != nil {
		return err
	}
	server.templates.subscribers.List, err = template.New("list.html").Funcs(template.FuncMap{
		"Iter": templatefuncs.Iter,
		"Inc":  templatefuncs.Inc,
		"Dec":  templatefuncs.Dec,
	}).ParseFiles(
		filepath.Join(server.config.StaticDir, "subscribers", "list.html"),
		filepath.Join(server.config.StaticDir, "subscribers", "pagination.html"))
	if err != nil {
		return err
	}

	return nil
}

// withAuth performs initial authorization before every request.
func (server *Server) withAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		ctxWithAuth := func(ctx context.Context) context.Context {
			token, err := server.cookieAuth.GetToken(r)
			if err != nil {
				controllers.Redirect(w, r, "/login", "GET")
			}

			claims, err := server.authService.Authorize(ctx, token)
			if err != nil {
				controllers.Redirect(w, r, "/login", "GET")
			}

			return auth.SetClaims(ctx, claims)
		}

		ctx = ctxWithAuth(r.Context())

		handler.ServeHTTP(w, r.Clone(ctx))
	})
}
