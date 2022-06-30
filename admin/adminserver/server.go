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
	"ultimatedivision/admin/adminserver/controllers"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/clubs"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/gameplay/queue"
	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/metrics"
	"ultimatedivision/internal/templatefuncs"
	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/seasons"
	"ultimatedivision/store"
	"ultimatedivision/store/lootboxes"
	"ultimatedivision/users"
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
		admin       controllers.AdminTemplates
		user        controllers.UserTemplates
		card        controllers.CardTemplates
		avatar      controllers.AvatarTemplates
		auth        controllers.AuthTemplates
		lootbox     controllers.LootBoxesTemplates
		marketplace controllers.MarketplaceTemplates
		club        controllers.ClubsTemplates
		queue       controllers.QueueTemplates
		divisions   controllers.DivisionsTemplates
		match       controllers.MatchesTemplate
		store       controllers.StoreTemplates
	}

	cards.PercentageQualities
}

// NewServer is a constructor for admin web server.
func NewServer(config Config, log logger.Logger, listener net.Listener, authService *adminauth.Service,
	admins *admins.Service, users *users.Service, cards *cards.Service, percentageQualities cards.PercentageQualities,
	avatars *avatars.Service, marketplace *marketplace.Service, lootboxes *lootboxes.Service, clubs *clubs.Service,
	queue *queue.Service, divisions *divisions.Service, matches *matches.Service, seasons *seasons.Service, store *store.Service, metric *metrics.Metric) (*Server, error) {
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
	adminsController := controllers.NewAdmins(log, admins, server.templates.admin)
	adminsRouter.HandleFunc("", adminsController.List).Methods(http.MethodGet)
	adminsRouter.HandleFunc("/create", adminsController.Create).Methods(http.MethodGet, http.MethodPost)
	adminsRouter.HandleFunc("/update/{id}", adminsController.Update).Methods(http.MethodGet, http.MethodPost)

	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(server.withAuth)
	userController := controllers.NewUsers(log, users, server.templates.user)
	userRouter.HandleFunc("", userController.List).Methods(http.MethodGet)
	userRouter.HandleFunc("/create", userController.Create).Methods(http.MethodGet, http.MethodPost)
	userRouter.HandleFunc("/update/status/{id}", userController.Update).Methods(http.MethodGet, http.MethodPost)
	userRouter.HandleFunc("/delete/{id}", userController.Delete).Methods(http.MethodGet)

	cardsRouter := router.PathPrefix("/cards").Subrouter()
	cardsRouter.Use(server.withAuth)
	cardsController := controllers.NewCards(log, cards, server.templates.card, percentageQualities)
	cardsRouter.HandleFunc("", cardsController.List).Methods(http.MethodGet)
	cardsRouter.HandleFunc("/create/{userId}", cardsController.Create).Methods(http.MethodGet)
	cardsRouter.HandleFunc("/delete/{id}", cardsController.Delete).Methods(http.MethodGet)

	avatarsRouter := router.PathPrefix("/avatars").Subrouter()
	avatarsRouter.Use(server.withAuth)
	avatarsController := controllers.NewAvatars(log, avatars, server.templates.avatar)
	avatarsRouter.HandleFunc("/{cardId}", avatarsController.Get).Methods(http.MethodGet)

	marketplaceRouter := router.PathPrefix("/marketplace").Subrouter()
	marketplaceRouter.Use(server.withAuth)
	marketplaceController := controllers.NewMarketplace(log, marketplace, cards, users, server.templates.marketplace)
	marketplaceRouter.HandleFunc("", marketplaceController.ListActiveLots).Methods(http.MethodGet)
	marketplaceRouter.HandleFunc("/get/{id}", marketplaceController.GetLotByID).Methods(http.MethodGet)
	marketplaceRouter.HandleFunc("/create", marketplaceController.CreateLot).Methods(http.MethodGet, http.MethodPost)
	marketplaceRouter.HandleFunc("/bet/{id}", marketplaceController.PlaceBetLot).Methods(http.MethodGet, http.MethodPost)

	lootBoxesRouter := router.PathPrefix("/lootboxes").Subrouter()
	lootBoxesRouter.Use(server.withAuth)
	lootBoxesController := controllers.NewLootBoxes(log, lootboxes, server.templates.lootbox)
	lootBoxesRouter.HandleFunc("", lootBoxesController.List).Methods(http.MethodGet)
	lootBoxesRouter.HandleFunc("/create/{id}", lootBoxesController.Create).Methods(http.MethodGet, http.MethodPost)
	lootBoxesRouter.HandleFunc("/open/{userId}/{lootboxId}", lootBoxesController.Open).Methods(http.MethodGet)

	clubsRouter := router.PathPrefix("/clubs").Subrouter()
	clubsRouter.Use(server.withAuth)
	clubsController := controllers.NewClubs(log, clubs, server.templates.club)
	clubsRouter.HandleFunc("/create/{userId}", clubsController.Create).Methods(http.MethodGet)
	clubsRouter.HandleFunc("/{userId}", clubsController.List).Methods(http.MethodGet)
	clubsRouter.HandleFunc("/{userId}/{clubId}/update", clubsController.UpdateStatus).Methods(http.MethodGet)
	clubsRouter.HandleFunc("/{clubId}/squad/create", clubsController.CreateSquad).Methods(http.MethodGet)
	clubsRouter.HandleFunc("/{clubId}/squad", clubsController.GetSquadByClubID).Methods(http.MethodGet)
	clubsRouter.HandleFunc("/{clubId}/squad/{squadId}/update", clubsController.UpdateSquad).Methods(http.MethodGet, http.MethodPost)
	clubsRouter.HandleFunc("/squad/{squadId}", clubsController.ListSquadCards).Methods(http.MethodGet)
	clubsRouter.HandleFunc("/squad/{squadId}/squad-cards", clubsController.Add).Methods(http.MethodGet, http.MethodPost)
	clubsRouter.HandleFunc("/squad/{squadId}/squad-cards/{cardId}/update", clubsController.UpdateCardPosition).Methods(http.MethodGet, http.MethodPost)
	clubsRouter.HandleFunc("/squad/{squadId}/squad-cards/{cardId}", clubsController.DeleteCard).Methods(http.MethodGet)

	queueRouter := router.PathPrefix("/queue").Subrouter()
	queueRouter.Use(server.withAuth)
	queueController := controllers.NewQueue(log, queue, server.templates.queue)
	queueRouter.HandleFunc("", queueController.List).Methods(http.MethodGet)
	queueRouter.HandleFunc("/{id}", queueController.Get).Methods(http.MethodGet)

	matchesRouter := router.PathPrefix("/matches").Subrouter()
	matchesRouter.Use(server.withAuth)
	matchesController := controllers.NewMatches(log, matches, server.templates.match, clubs, seasons)
	matchesRouter.HandleFunc("/create", matchesController.Create).Methods(http.MethodGet, http.MethodPost)
	matchesRouter.HandleFunc("", matchesController.ListMatches).Methods(http.MethodGet)
	matchesRouter.HandleFunc("/delete/{id}", matchesController.Delete).Methods(http.MethodGet)
	matchesRouter.HandleFunc("/{id}/goals", matchesController.ListMatchGoals).Methods(http.MethodGet)

	divisionsRouter := router.PathPrefix("/divisions").Subrouter()
	divisionsRouter.Use(server.withAuth)
	divisionsController := controllers.NewDivisions(log, divisions, server.templates.divisions)
	divisionsRouter.HandleFunc("", divisionsController.List).Methods(http.MethodGet)
	divisionsRouter.HandleFunc("/create", divisionsController.Create).Methods(http.MethodGet, http.MethodPost)
	divisionsRouter.HandleFunc("/delete/{id}", divisionsController.Delete).Methods(http.MethodGet)

	storeRouter := router.PathPrefix("/store").Subrouter()
	storeRouter.Use(server.withAuth)
	storeController := controllers.NewStore(log, store, server.templates.store)
	storeRouter.HandleFunc("", storeController.List).Methods(http.MethodGet)
	storeRouter.HandleFunc("/update/{id}", storeController.Update).Methods(http.MethodGet, http.MethodPost)

	router.Handle("/metrics", metric.GetHandler())

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
		return err
	}
	server.templates.user.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "users", "create.html"))
	if err != nil {
		return err
	}
	server.templates.user.Update, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "users", "update.html"))
	if err != nil {
		return err
	}

	server.templates.admin.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "admins", "list.html"))
	if err != nil {
		return err
	}
	server.templates.admin.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "admins", "create.html"))
	if err != nil {
		return err
	}
	server.templates.admin.Update, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "admins", "update.html"))
	if err != nil {
		return err
	}

	server.templates.card.List, err = template.New("list.html").Funcs(template.FuncMap{
		"Iter": templatefuncs.Iter,
		"Inc":  templatefuncs.Inc,
		"Dec":  templatefuncs.Dec,
	}).ParseFiles(
		filepath.Join(server.config.StaticDir, "cards", "list.html"),
		filepath.Join(server.config.StaticDir, "cards", "pagination.html"))
	if err != nil {
		return err
	}

	server.templates.avatar.Get, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "avatars", "get.html"))
	if err != nil {
		return err
	}

	server.templates.marketplace.List, err = template.New("list.html").Funcs(template.FuncMap{
		"Iter": templatefuncs.Iter,
		"Inc":  templatefuncs.Inc,
		"Dec":  templatefuncs.Dec,
	}).ParseFiles(
		filepath.Join(server.config.StaticDir, "marketplace", "list.html"),
		filepath.Join(server.config.StaticDir, "marketplace", "pagination.html"))
	if err != nil {
		return err
	}

	server.templates.marketplace.Get, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "marketplace", "get.html"))
	if err != nil {
		return err
	}

	server.templates.marketplace.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "marketplace", "create.html"))
	if err != nil {
		return err
	}

	server.templates.marketplace.Bet, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "marketplace", "bet.html"))
	if err != nil {
		return err
	}

	server.templates.auth.Login, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "auth", "login.html"))
	if err != nil {
		return err
	}

	server.templates.lootbox.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "lootboxes", "list.html"))
	if err != nil {
		return err
	}
	server.templates.lootbox.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "lootboxes", "create.html"))
	if err != nil {
		return err
	}
	server.templates.lootbox.ListCards, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "lootboxes", "listCards.html"))
	if err != nil {
		return err
	}

	server.templates.club.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "clubs", "list.html"))
	if err != nil {
		return err
	}
	server.templates.club.ListSquads, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "clubs", "listSquad.html"))
	if err != nil {
		return err
	}
	server.templates.club.ListSquadCards, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "clubs", "listSquadCards.html"))
	if err != nil {
		return err
	}
	server.templates.club.AddCard, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "clubs", "addCard.html"))
	if err != nil {
		return err
	}
	server.templates.club.UpdateSquad, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "clubs", "updateSquad.html"))
	if err != nil {
		return err
	}
	server.templates.club.UpdateCardPosition, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "clubs", "updateCardPosition.html"))
	if err != nil {
		return err
	}

	server.templates.queue.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "queue", "list.html"))
	if err != nil {
		return err
	}

	server.templates.queue.Get, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "queue", "get.html"))
	if err != nil {
		return err
	}

	server.templates.match.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "matches", "create.html"))
	if err != nil {
		return err
	}
	server.templates.match.List, err = template.New("list.html").Funcs(template.FuncMap{
		"Iter": templatefuncs.Iter,
		"Inc":  templatefuncs.Inc,
		"Dec":  templatefuncs.Dec,
	}).ParseFiles(
		filepath.Join(server.config.StaticDir, "matches", "list.html"),
		filepath.Join(server.config.StaticDir, "matches", "pagination.html"))
	if err != nil {
		return err
	}
	server.templates.match.ListGoals, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "matches", "listGoals.html"))
	if err != nil {
		return err
	}

	server.templates.divisions.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "divisions", "list.html"))
	if err != nil {
		return err
	}
	server.templates.divisions.Create, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "divisions", "create.html"))
	if err != nil {
		return err
	}

	server.templates.store.List, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "store", "list.html"))
	if err != nil {
		return err
	}
	server.templates.store.Update, err = template.ParseFiles(filepath.Join(server.config.StaticDir, "store", "update.html"))
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
