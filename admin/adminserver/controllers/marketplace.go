// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/users"
)

var (
	// ErrMarketplace is an internal error type for marketplace controller.
	ErrMarketplace = errs.Class("marketplace controller error")
)

// MarketplaceTemplates holds all marketplace related templates.
type MarketplaceTemplates struct {
	List   *template.Template
	Get    *template.Template
	Create *template.Template
	Bet    *template.Template
}

// Marketplace is a mvc controller that handles all marketplace related views.
type Marketplace struct {
	log logger.Logger

	marketplace *marketplace.Service
	cards       *cards.Service
	users       *users.Service

	templates MarketplaceTemplates
}

// NewMarketplace is a constructor for marketplace controller.
func NewMarketplace(log logger.Logger, marketplace *marketplace.Service, cards *cards.Service, users *users.Service, templates MarketplaceTemplates) *Marketplace {
	marketplaceController := &Marketplace{
		log:         log,
		marketplace: marketplace,
		cards:       cards,
		users:       users,
		templates:   templates,
	}

	return marketplaceController
}

// ListActiveLots is an endpoint that will provide a web page with active lots.
func (controller *Marketplace) ListActiveLots(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		lotsPage    marketplace.Page
		err         error
		limit, page int
	)
	urlQuery := r.URL.Query()
	limitQuery := urlQuery.Get("limit")
	pageQuery := urlQuery.Get("page")

	if limitQuery != "" {
		if limit, err = strconv.Atoi(limitQuery); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if pageQuery != "" {
		if page, err = strconv.Atoi(pageQuery); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	cursor := pagination.Cursor{
		Limit: limit,
		Page:  page,
	}
	lotsPage, err = controller.marketplace.ListActiveLots(ctx, cursor)
	if err != nil {
		controller.log.Error("could not list lots", ErrMarketplace.Wrap(err))
		switch {
		case marketplace.ErrNoLot.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err = controller.templates.List.Execute(w, lotsPage); err != nil {
		controller.log.Error("can not execute list lots template", ErrMarketplace.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetLotByID is an endpoint that will provide a web page with lot by id.
func (controller *Marketplace) GetLotByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lot, err := controller.marketplace.GetLotByID(ctx, id)
	if err != nil {
		controller.log.Error("could not get lot by id", ErrMarketplace.Wrap(err))
		switch {
		case marketplace.ErrNoLot.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err = controller.templates.Get.Execute(w, lot); err != nil {
		controller.log.Error("can not execute get lot template", ErrMarketplace.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateLot is an endpoint that will add lot to database.
func (controller *Marketplace) CreateLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		var (
			err         error
			limit, page int
		)
		urlQuery := r.URL.Query()
		limitQuery := urlQuery.Get("limit")
		pageQuery := urlQuery.Get("page")

		if limitQuery != "" {
			if limit, err = strconv.Atoi(limitQuery); err != nil {
				http.Error(w, ErrCards.Wrap(err).Error(), http.StatusBadRequest)
				return
			}
		}

		if pageQuery != "" {
			if page, err = strconv.Atoi(pageQuery); err != nil {
				http.Error(w, ErrCards.Wrap(err).Error(), http.StatusBadRequest)
				return
			}
		}

		cursor := pagination.Cursor{
			Limit: limit,
			Page:  page,
		}
		cardsListPage, err := controller.cards.List(ctx, cursor)
		if err != nil {
			controller.log.Error("could not list cards", ErrMarketplace.Wrap(err))
			switch {
			case cards.ErrNoCard.Has(err):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		usersList, err := controller.users.List(ctx)
		if err != nil {
			controller.log.Error("could not list users", ErrMarketplace.Wrap(err))
			switch {
			case users.ErrNoUser.Has(err):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		responseCreateLot := marketplace.ResponseCreateLot{
			Cards: cardsListPage,
			Users: usersList,
		}

		if err = controller.templates.Create.Execute(w, responseCreateLot); err != nil {
			controller.log.Error("could not execute create lot template", ErrMarketplace.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		itemIDForm := r.FormValue("itemId")
		strings.ToValidUTF8(itemIDForm, "")
		itemID, err := uuid.Parse(itemIDForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userIDForm := r.FormValue("userId")
		strings.ToValidUTF8(userIDForm, "")
		userID, err := uuid.Parse(userIDForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		startPriceForm := r.FormValue("startPrice")
		strings.ToValidUTF8(startPriceForm, "")
		startPrice, err := strconv.ParseFloat(startPriceForm, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		maxPriceForm := r.FormValue("maxPrice")
		strings.ToValidUTF8(maxPriceForm, "")
		maxPrice, err := strconv.ParseFloat(maxPriceForm, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		periodForm := r.FormValue("period")
		strings.ToValidUTF8(periodForm, "")
		period, err := strconv.Atoi(periodForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		createLot := marketplace.CreateLot{
			ItemID:     itemID,
			UserID:     userID,
			StartPrice: startPrice,
			MaxPrice:   maxPrice,
			Period:     marketplace.Period(period),
		}

		if err := createLot.ValidateCreateLot(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if err := controller.marketplace.CreateLot(ctx, createLot); err != nil {
			controller.log.Error("could not create lot", ErrMarketplace.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/marketplace", "GET")
	}
}

// PlaceBetLot is an endpoint that will add card to database.
func (controller *Marketplace) PlaceBetLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	if vars["id"] == "" {
		http.Error(w, ErrMarketplace.New("id parameter is empty").Error(), http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		_, err := controller.marketplace.GetLotByID(ctx, id)
		if err != nil {
			controller.log.Error("could not list cards", ErrMarketplace.Wrap(err))
			switch {
			case cards.ErrNoCard.Has(err):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		usersList, err := controller.users.List(ctx)
		if err != nil {
			controller.log.Error("could not list users", ErrMarketplace.Wrap(err))
			switch {
			case users.ErrNoUser.Has(err):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		responsePlaceBetLot := marketplace.ResponsePlaceBetLot{
			ID:    id,
			Users: usersList,
		}

		if err = controller.templates.Bet.Execute(w, responsePlaceBetLot); err != nil {
			controller.log.Error("could not execute bet lot template", ErrMarketplace.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userIDForm := r.FormValue("userId")
		strings.ToValidUTF8(userIDForm, "")
		userID, err := uuid.Parse(userIDForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		betAmountForm := r.FormValue("betAmount")
		strings.ToValidUTF8(betAmountForm, "")
		betAmount, err := strconv.ParseFloat(betAmountForm, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		betLot := marketplace.BetLot{
			ID:        id,
			UserID:    userID,
			BetAmount: betAmount,
		}

		if err := betLot.ValidateBetLot(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if err := controller.marketplace.PlaceBetLot(ctx, betLot); err != nil {
			controller.log.Error("could not place bet lot", ErrMarketplace.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/marketplace", "GET")
	}
}
