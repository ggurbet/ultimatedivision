// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

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

	"ultimatedivision/cards"
	"ultimatedivision/console/consoleserver/controllers"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
	"ultimatedivision/users/userauth"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("console web server error")
)

// Config contains configuration for console web server.
type Config struct {
	Address   string `json:"address"`
	StaticDir string `json:"staticDir"`
}

// Server represents console web server.
//
// architecture: Endpoint
type Server struct {
	log    logger.Logger
	config Config

	listener net.Listener
	server   http.Server

	authService *userauth.Service
	cookieAuth  *auth.CookieAuth

	templates struct {
		index *template.Template
		auth  *controllers.AuthTemplates
	}
}

// NewServer is a constructor for console web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, cards *cards.Service, lootBoxes *lootboxes.Service) *Server {
	server := &Server{
		log:      log,
		config:   config,
		listener: listener,
	}

	authController := controllers.NewAuth(server.log, server.authService, server.cookieAuth, server.templates.auth)
	cardsController := controllers.NewCards(log, cards)
	lootBoxesController := controllers.NewLootBoxes(log, lootBoxes)

	router := mux.NewRouter()
	router.HandleFunc("/register", authController.RegisterTemplateHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", authController.LoginTemplateHandler).Methods(http.MethodGet)

	apiRouter := router.PathPrefix("/api/v0").Subrouter()
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authController.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/logout", authController.Logout).Methods(http.MethodPost)
	authRouter.HandleFunc("/register", authController.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/email/confirm/{token}", authController.ConfirmEmail).Methods(http.MethodGet)

	cardsRouter := router.PathPrefix("/cards").Subrouter()
	cardsRouter.HandleFunc("", cardsController.List).Methods(http.MethodGet)

	lootBoxesRouter := router.PathPrefix("/lootboxes").Subrouter()
	lootBoxesRouter.HandleFunc("", lootBoxesController.Create).Methods(http.MethodPost)
	lootBoxesRouter.HandleFunc("", lootBoxesController.Open).Methods(http.MethodDelete)

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
	// header.Set("X-Content-Type-Options", "nosniff")
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
