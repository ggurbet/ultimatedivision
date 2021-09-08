// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/sqlsearchoperators"
)

var (
	// ErrMarketplace is an internal error type for marketplace controller.
	ErrMarketplace = errs.Class("marketplace controller error")
)

// Marketplace is a mvc controller that handles all marketplace related views.
type Marketplace struct {
	log         logger.Logger
	marketplace *marketplace.Service
}

// NewMarketplace is a constructor for marketplace controller.
func NewMarketplace(log logger.Logger, marketplace *marketplace.Service) *Marketplace {
	marketplaceController := &Marketplace{
		log:         log,
		marketplace: marketplace,
	}

	return marketplaceController
}

// ListActiveLots is an endpoint that returns active lots list.
func (controller *Marketplace) ListActiveLots(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var (
		listActiveLots []marketplace.Lot
		err            error
		filters        cards.SliceFilters
	)
	urlQuery := r.URL.Query()
	playerName := urlQuery.Get(string(cards.FilterPlayerName))

	if playerName == "" {
		if err := filters.DecodingURLParameters(urlQuery); err != nil {
			controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		}
		if len(filters) > 0 {
			listActiveLots, err = controller.marketplace.ListActiveLotsWithFilters(ctx, filters)
		} else {
			listActiveLots, err = controller.marketplace.ListActiveLots(ctx)
		}
	} else {
		filter := cards.Filters{
			Name:           cards.FilterPlayerName,
			Value:          playerName,
			SearchOperator: sqlsearchoperators.LIKE,
		}
		listActiveLots, err = controller.marketplace.ListActiveLotsByPlayerName(ctx, filter)
	}
	if err != nil {
		controller.log.Error("could not get active lots list", ErrMarketplace.Wrap(err))
		switch {
		case marketplace.ErrNoLot.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrMarketplace.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		}
		return
	}

	if err = json.NewEncoder(w).Encode(listActiveLots); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// GetLotByID is an endpoint that returns lot by id.
func (controller *Marketplace) GetLotByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	lot, err := controller.marketplace.GetLotByID(ctx, id)
	if err != nil {
		controller.log.Error("could not get lot by id", ErrMarketplace.Wrap(err))
		switch {
		case marketplace.ErrNoLot.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrMarketplace.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		}
		return
	}

	responseLot := marketplace.ResponseLot{
		ID:           lot.ID,
		ItemID:       lot.ItemID,
		Type:         lot.Type,
		Status:       lot.Status,
		StartPrice:   lot.StartPrice,
		MaxPrice:     lot.MaxPrice,
		CurrentPrice: lot.CurrentPrice,
		StartTime:    lot.StartTime,
		EndTime:      lot.EndTime,
		Period:       lot.Period,
	}

	if err = json.NewEncoder(w).Encode(responseLot); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// CreateLot is an endpoint that returns lot by id.
func (controller *Marketplace) CreateLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.log.Error("could not authorize user", ErrMarketplace.Wrap(err))
		controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
		return
	}

	var createLot marketplace.CreateLot
	if err = json.NewDecoder(r.Body).Decode(&createLot); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}
	createLot.UserID = claims.UserID

	if err = createLot.ValidateCreateLot(); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
	}

	if err = controller.marketplace.CreateLot(ctx, createLot); err != nil {
		controller.log.Error("could not create lot", ErrMarketplace.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		return
	}
}

// PlaceBetLot is an endpoint that returns lot by id.
func (controller *Marketplace) PlaceBetLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.log.Error("could not authorize user", ErrMarketplace.Wrap(err))
		controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
		return
	}

	var betLot marketplace.BetLot
	if err = json.NewDecoder(r.Body).Decode(&betLot); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}
	betLot.UserID = claims.UserID

	if err = betLot.ValidateBetLot(); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
	}

	if err = controller.marketplace.PlaceBetLot(ctx, betLot); err != nil {
		controller.log.Error("could not place bet lot", ErrMarketplace.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Marketplace) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrMarketplace.Wrap(err))
	}
}
