// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package server

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

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/server/controllers"
	"ultimatedivision/nftdrop/whitelist"
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

	listener net.Listener
	server   http.Server

	templates struct {
		index *template.Template
	}
}

// NewServer is a constructor for nftdrop web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, whitelist *whitelist.Service) *Server {
	server := &Server{
		log:      log,
		config:   config,
		listener: listener,
	}

	whitelistController := controllers.NewWhitelist(log, whitelist)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	whitelistRouter := apiRouter.PathPrefix("/whitelist").Subrouter()
	whitelistRouter.HandleFunc("/{address}", whitelistController.Get).Methods(http.MethodGet)

	fs := http.FileServer(http.Dir(server.config.StaticDir))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	router.PathPrefix("/").HandlerFunc(server.appHandler)

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
		err := server.server.Serve(server.listener)
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

	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("Referrer-Policy", "same-origin")

	if server.templates.index == nil {
		server.log.Error("index template is not set", nil)
		return
	}

	if err := server.templates.index.Execute(w, nil); err != nil {
		server.log.Error("index template could not be executed", err)
		return
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
