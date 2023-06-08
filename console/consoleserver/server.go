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
	"github.com/rs/cors"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/cards"
	"ultimatedivision/cards/waitlist"
	"ultimatedivision/clubs"
	"ultimatedivision/console/connections"
	"ultimatedivision/console/consoleserver/controllers"
	"ultimatedivision/gameplay/matchmaking"
	"ultimatedivision/gameplay/queue"
	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/metrics"
	"ultimatedivision/marketplace"
	"ultimatedivision/marketplace/bids"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/seasons"
	"ultimatedivision/store"
	"ultimatedivision/store/lootboxes"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("console web server error")
)

// Config contains configuration for console web server.
type Config struct {
	Address    string `json:"address"`
	StaticDir  string `json:"staticDir"`
	AvatarsDir string `json:"avatarsDir"`

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
func NewServer(config Config, log logger.Logger, listener net.Listener, cards *cards.Service, lootBoxes *lootboxes.Service,
	marketplace *marketplace.Service, bids *bids.Service, clubs *clubs.Service, userAuth *userauth.Service, users *users.Service,
	queue *queue.Service, seasons *seasons.Service, waitList *waitlist.Service, store *store.Service, metric *metrics.Metric,
	currencyWaitList *currencywaitlist.Service, connections *connections.Service, matchmaking *matchmaking.Service) *Server {
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

	authController := controllers.NewAuth(server.log, server.authService, server.cookieAuth, server.templates.auth, metric)
	userController := controllers.NewUsers(server.log, users)
	cardsController := controllers.NewCards(log, cards)
	clubsController := controllers.NewClubs(log, clubs)
	lootBoxesController := controllers.NewLootBoxes(log, lootBoxes)
	marketplaceController := controllers.NewMarketplace(log, marketplace)
	bidsController := controllers.NewBids(log, bids)
	// TODO: now use a new service - matchmaking for the game
	// queueController := controllers.NewQueue(log, queue, connections).
	seasonsController := controllers.NewSeasons(log, seasons)
	waitListController := controllers.NewWaitList(log, waitList)
	storeController := controllers.NewStore(log, store)
	contractCasperController := controllers.NewContractCasper(log, currencyWaitList)
	connectionController := controllers.NewConnections(log, connections)
	matchmakingController := controllers.NewMatchmaking(log, matchmaking)

	router := mux.NewRouter()
	router.HandleFunc("/register", authController.RegisterTemplateHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", authController.LoginTemplateHandler).Methods(http.MethodGet)

	apiRouter := router.PathPrefix("/api/v0").Subrouter()
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authController.Login).Methods(http.MethodPost)

	metamaskRouter := authRouter.PathPrefix("/metamask").Subrouter()
	metamaskRouter.HandleFunc("/register", authController.MetamaskRegister).Methods(http.MethodPost)
	metamaskRouter.HandleFunc("/nonce", authController.Nonce).Methods(http.MethodGet)
	metamaskRouter.HandleFunc("/login", authController.MetamaskLogin).Methods(http.MethodPost)

	velasRouter := authRouter.PathPrefix("/velas").Subrouter()
	velasRouter.HandleFunc("/register", authController.VelasRegister).Methods(http.MethodPost)
	velasRouter.HandleFunc("/nonce", authController.Nonce).Methods(http.MethodGet)
	velasRouter.HandleFunc("/login", authController.VelasLogin).Methods(http.MethodPost)
	velasRouter.HandleFunc("/vaclient", authController.VelasVAClientFields).Methods(http.MethodGet)
	velasRouter.HandleFunc("/register-data/{user_id}", authController.GetVelasData).Methods(http.MethodGet)

	casperRouter := authRouter.PathPrefix("/casper").Subrouter()
	casperRouter.HandleFunc("/register", authController.CasperRegister).Methods(http.MethodPost)
	casperRouter.HandleFunc("/nonce", authController.PublicKey).Methods(http.MethodGet)
	casperRouter.HandleFunc("/login", authController.CasperLogin).Methods(http.MethodPost)

	apiRouter.HandleFunc("/casper/send-tx", contractCasperController.SendTx).Methods(http.MethodPost)

	apiRouter.Handle("/connection", server.withAuth(http.HandlerFunc(connectionController.Connect))).Methods(http.MethodGet)

	authRouter.HandleFunc("/logout", authController.Logout).Methods(http.MethodPost)
	authRouter.HandleFunc("/register", authController.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/email/confirm/{token}", authController.ConfirmEmail).Methods(http.MethodGet)
	authRouter.HandleFunc("/password/{email}", authController.ResetPasswordSendEmail).Methods(http.MethodGet)
	authRouter.HandleFunc("/reset-password/{token}", authController.CheckAuthToken).Methods(http.MethodGet)
	authRouter.Handle("/reset-password", server.withAuth(http.HandlerFunc(authController.ResetPassword))).Methods(http.MethodPatch)
	authRouter.Handle("/change-password", server.withAuth(http.HandlerFunc(authController.ChangePassword))).Methods(http.MethodPost)
	authRouter.Handle("/change-email", server.withAuth(http.HandlerFunc(authController.SendEmailForChangeEmail))).Methods(http.MethodPost)
	authRouter.Handle("/change-email/{token}", server.withAuth(http.HandlerFunc(authController.ChangeEmail))).Methods(http.MethodGet)

	profileRouter := apiRouter.PathPrefix("/profile").Subrouter()
	profileRouter.Use(server.withAuth)
	profileRouter.HandleFunc("", userController.GetProfile).Methods(http.MethodGet)

	metamaskRouterWithAuth := profileRouter.PathPrefix("/metamask").Subrouter()
	metamaskRouterWithAuth.HandleFunc("/wallet", userController.CreateWalletFromMetamask).Methods(http.MethodPatch)
	metamaskRouterWithAuth.HandleFunc("/wallet/change", userController.ChangeWalletFromMetamask).Methods(http.MethodPatch)

	cardsRouter := apiRouter.PathPrefix("/cards").Subrouter()
	cardsRouter.Use(server.withAuth)
	cardsRouter.HandleFunc("", cardsController.List).Methods(http.MethodGet)
	cardsRouter.HandleFunc("/{id}", cardsController.Get).Methods(http.MethodGet)
	cardsRouter.HandleFunc("/status/{id}", cardsController.GetStatus).Methods(http.MethodGet)

	clubsRouter := apiRouter.PathPrefix("/clubs").Subrouter()
	clubsRouter.Use(server.withAuth)
	clubsRouter.HandleFunc("", clubsController.Create).Methods(http.MethodPost)
	clubsRouter.HandleFunc("", clubsController.Get).Methods(http.MethodGet)
	clubsRouter.HandleFunc("/{clubId}", clubsController.UpdateStatus).Methods(http.MethodPatch)

	squadRouter := clubsRouter.PathPrefix("/{clubId}/squads").Subrouter()
	squadRouter.HandleFunc("", clubsController.CreateSquad).Methods(http.MethodPost)
	squadRouter.HandleFunc("/{squadId}", clubsController.UpdateTacticCaptain).Methods(http.MethodPatch)
	squadRouter.HandleFunc("/{squadId}/formation/{formationId}", clubsController.ChangeFormation).Methods(http.MethodPut)

	squadCardsRouter := squadRouter.PathPrefix("/{squadId}/cards").Subrouter()
	squadCardsRouter.HandleFunc("/{cardId}", clubsController.Add).Methods(http.MethodPost)
	squadCardsRouter.HandleFunc("/{cardId}", clubsController.Delete).Methods(http.MethodDelete)
	squadCardsRouter.HandleFunc("/{cardId}", clubsController.UpdatePosition).Methods(http.MethodPatch)

	lootBoxesRouter := apiRouter.PathPrefix("/lootboxes").Subrouter()
	lootBoxesRouter.Use(server.withAuth)
	lootBoxesRouter.HandleFunc("", lootBoxesController.Create).Methods(http.MethodPost)
	lootBoxesRouter.HandleFunc("/{id}", lootBoxesController.Open).Methods(http.MethodPost)

	marketplaceRouter := apiRouter.PathPrefix("/marketplace").Subrouter()
	marketplaceRouter.HandleFunc("", marketplaceController.ListActiveLots).Methods(http.MethodGet)
	marketplaceRouterWithAuth := marketplaceRouter
	marketplaceRouterWithAuth.Use(server.withAuth)
	marketplaceRouterWithAuth.HandleFunc("/{id}", marketplaceController.GetLotByID).Methods(http.MethodGet)
	marketplaceRouterWithAuth.HandleFunc("/end-time/{id}", marketplaceController.IsExpired).Methods(http.MethodGet)
	marketplaceRouterWithAuth.HandleFunc("/price/{card_id}", marketplaceController.GetCurrentPriceByCardID).Methods(http.MethodGet)
	marketplaceRouterWithAuth.HandleFunc("/lot-data/{card_id}", marketplaceController.GetLotData).Methods(http.MethodGet)
	marketplaceRouterWithAuth.HandleFunc("", marketplaceController.CreateLot).Methods(http.MethodPost)
	marketplaceRouterWithAuth.HandleFunc("/bet", marketplaceController.PlaceBetLot).Methods(http.MethodPost)

	apiRouter.HandleFunc("/casper-approve", marketplaceController.GetApproveData).Methods(http.MethodGet)

	bidsRouter := apiRouter.PathPrefix("/bids").Subrouter()
	bidsRouter.Use(server.withAuth)
	bidsRouter.HandleFunc("/offer/{card_id}", bidsController.GetMakeOfferData).Methods(http.MethodGet)
	bidsRouter.HandleFunc("", bidsController.Bet).Methods(http.MethodPost)
	bidsRouter.HandleFunc("/{lotId}", bidsController.ListByLotID).Methods(http.MethodGet)

	queueRouter := apiRouter.PathPrefix("/queue").Subrouter()
	queueRouter.Use(server.withAuth)
	queueRouter.HandleFunc("", matchmakingController.Create).Methods(http.MethodGet)

	seasonsRouter := apiRouter.PathPrefix("/seasons").Subrouter()
	seasonsRouter.Use(server.withAuth)
	seasonsRouter.HandleFunc("/current", seasonsController.GetCurrentSeasons).Methods(http.MethodGet)
	seasonsRouter.HandleFunc("/reward", seasonsController.GetRewardByUserID).Methods(http.MethodGet)
	seasonsRouter.HandleFunc("/reward/tokens", seasonsController.GetValueOfTokensRewardByUserID).Methods(http.MethodGet)
	seasonsRouter.HandleFunc("/statistics/division/{divisionName}", seasonsController.GetAllClubsStatistics).Methods(http.MethodGet)
	seasonsRouter.HandleFunc("/club", seasonsController.UpdatesClubsToNewDivision).Methods(http.MethodPut)

	waitListRouter := apiRouter.PathPrefix("/nft-waitlist").Subrouter()
	waitListRouter.Use(server.withAuth)
	waitListRouter.HandleFunc("", waitListController.Create).Methods(http.MethodPost)

	storeRouter := apiRouter.PathPrefix("/store").Subrouter()
	storeRouter.Use(server.withAuth)
	storeRouter.HandleFunc("/buy", storeController.Buy).Methods(http.MethodPost)

	av := http.FileServer(http.Dir(server.config.AvatarsDir))
	router.PathPrefix("/avatars/").Handler(http.StripPrefix("/avatars/", av))

	fs := http.FileServer(http.Dir(server.config.StaticDir))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	router.PathPrefix("/").HandlerFunc(server.appHandler)

	server.server = http.Server{
		Handler: cors.AllowAll().Handler(router),
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
	// header.Set("X-Content-Type-Options", "nosniff").
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
			return
		}

		claims, err := server.authService.Authorize(ctx, token)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
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
