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
	"ultimatedivision/clubs"
	"ultimatedivision/console/consoleserver/controllers"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
	"ultimatedivision/users"
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

	Auth struct {
		CookieName string `json:"cookieName"`
		Path       string `json:"path"`
	} `json:"auth"`
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
func NewServer(config Config, log logger.Logger, listener net.Listener, cards *cards.Service, lootBoxes *lootboxes.Service, clubs *clubs.Service, userAuth *userauth.Service, users *users.Service) *Server {
	server := &Server{
		log:         log,
		config:      config,
		listener:    listener,
		authService: userAuth,
		cookieAuth: auth.NewCookieAuth(auth.CookieSettings{
			Name: config.Auth.CookieName,
			Path: config.Auth.Path,
		}),
	}

	authController := controllers.NewAuth(server.log, server.authService, server.cookieAuth, server.templates.auth)
	userController := controllers.NewUsers(server.log, users)
	cardsController := controllers.NewCards(log, cards)
	clubsController := controllers.NewClubs(log, clubs)
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
	authRouter.Handle("/change-password", server.withAuth(http.HandlerFunc(authController.ChangePassword))).Methods(http.MethodPost)

	profile := apiRouter.PathPrefix("/profile").Subrouter()
	profile.Handle("", server.withAuth(http.HandlerFunc(userController.GetProfile))).Methods(http.MethodGet)

	cardsRouter := router.PathPrefix("/cards").Subrouter()
	cardsRouter.Handle("", server.withAuth(http.HandlerFunc(cardsController.List))).Methods(http.MethodGet)

	clubsRouter := apiRouter.PathPrefix("/clubs").Subrouter()
	clubsRouter.Handle("", server.withAuth(http.HandlerFunc(clubsController.Create))).Methods(http.MethodPost)
	clubsRouter.Handle("", server.withAuth(http.HandlerFunc(clubsController.Get))).Methods(http.MethodGet)
	clubsRouter.Handle("", server.withAuth(http.HandlerFunc(clubsController.UpdateSquad))).Methods(http.MethodPatch)

	squadsRouter := clubsRouter.Path("/squads").Subrouter()
	squadsRouter.Handle("/{clubId}", server.withAuth(http.HandlerFunc(clubsController.Create))).Methods(http.MethodPost)

	squadCardsRouter := squadsRouter.Path("/squad-cards").Subrouter()
	squadCardsRouter.Handle("", server.withAuth(http.HandlerFunc(clubsController.Add))).Methods(http.MethodPost)
	squadCardsRouter.Handle("", server.withAuth(http.HandlerFunc(clubsController.UpdatePosition))).Methods(http.MethodPatch)
	squadCardsRouter.Handle("", server.withAuth(http.HandlerFunc(clubsController.Delete))).Methods(http.MethodDelete)

	lootBoxesRouter := router.PathPrefix("/lootboxes").Subrouter()
	lootBoxesRouter.Handle("", server.withAuth(http.HandlerFunc(lootBoxesController.Create))).Methods(http.MethodPost)
	lootBoxesRouter.Handle("", server.withAuth(http.HandlerFunc(lootBoxesController.Open))).Methods(http.MethodDelete)

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

// withAuth performs initial authorization before every request.
func (server *Server) withAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token, err := server.cookieAuth.GetToken(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		ctx = auth.SetToken(ctx, []byte(token))

		claims, err := server.authService.Authorize(ctx)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		ctx = auth.SetClaims(ctx, claims)

		handler.ServeHTTP(w, r.Clone(ctx))
	})
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
