// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/cards"
	"ultimatedivision/console/consoleserver/controllers"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/users/userauth"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("console web server error")
)

// Config contains configuration for console web server.
type Config struct {
	Address string `json:"address"`
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
		auth *controllers.AuthTemplates
	}
}

// NewServer is a constructor for console web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, cards *cards.Service) (*Server, error) {
	server := &Server{
		log:      log,
		config:   config,
		listener: listener,
	}

	authController := controllers.NewAuth(server.log, server.authService, server.cookieAuth, server.templates.auth)
	cardsController := NewCards(log, cards)

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
