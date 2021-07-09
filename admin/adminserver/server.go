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

	"ultimatedivision/admin/admins"
	"ultimatedivision/admin/adminserver/controllers"
	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("admin web server error")
)

// Config contains configuration for admin web server.
type Config struct {
	Address   string `help:"server address of the frontend app" devDefault:"127.0.0.1:0" releaseDefault:"127.0.0.1:0"`
	StaticDir string `help:"path to static resources" default:""`
}

// Server represents admin web server.
//
// architecture: Endpoint
type Server struct {
	log    logger.Logger
	config Config

	listener net.Listener
	server   http.Server

	templates struct {
		admin controllers.AdminTemplates
		user  controllers.UserTemplates
		card  controllers.CardTemplates
	}
}

// NewServer is a constructor for admin web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, admins *admins.Service, users *users.Service, cards *cards.Service) (*Server, error) {
	server := &Server{
		log:      log,
		config:   config,
		listener: listener,
	}

	err := server.initializeTemplates()
	if err != nil {
		return nil, Error.Wrap(err)
	}

	router := mux.NewRouter()
	adminsRouter := router.PathPrefix("/admins").Subrouter().StrictSlash(true)
	// managersRouter.Use(server.withAuth) // TODO: implement cookie auth and auth service.
	adminsController := controllers.NewAdmins(log, admins, server.templates.admin)
	adminsRouter.HandleFunc("", adminsController.List).Methods(http.MethodGet)
	adminsRouter.HandleFunc("/create", adminsController.Create).Methods(http.MethodGet, http.MethodPost)
	adminsRouter.HandleFunc("/update/{id}", adminsController.Update).Methods(http.MethodGet, http.MethodPost)

	userRouter := router.PathPrefix("/users").Subrouter().StrictSlash(true)
	// userRouter.Use(server.withAuth) // TODO: implement cookie auth and auth service.
	userController := controllers.NewUsers(log, users, server.templates.user)
	userRouter.HandleFunc("", userController.List).Methods(http.MethodGet)
	userRouter.HandleFunc("/create", userController.Create).Methods(http.MethodGet, http.MethodPost)
	userRouter.HandleFunc("/update/status/{id}", userController.Update).Methods(http.MethodGet, http.MethodPost)
	userRouter.HandleFunc("/delete/{id}", userController.Delete).Methods(http.MethodGet)

	cardsRouter := router.PathPrefix("/cards").Subrouter().StrictSlash(true)
	cardsController := controllers.NewCards(log, cards, server.templates.card)
	cardsRouter.HandleFunc("", cardsController.List).Methods(http.MethodGet)
	cardsRouter.HandleFunc("/create/{userID}", cardsController.Create).Methods(http.MethodGet)
	cardsRouter.HandleFunc("/delete/{id}", cardsController.Delete).Methods(http.MethodGet)

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
	server.templates.user.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "users", "list.html"))
	if err != nil {
		return Error.Wrap(err)
	}
	server.templates.user.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "users", "create.html"))
	if err != nil {
		return Error.Wrap(err)
	}
	server.templates.user.Update, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "users", "update.html"))
	if err != nil {
		return Error.Wrap(err)
	}
	server.templates.admin.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "admins", "list.html"))
	if err != nil {
		return Error.Wrap(err)
	}
	server.templates.card.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "cards", "list.html"))
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}
