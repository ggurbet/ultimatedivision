// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/pkg/pagination"
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
		lotsPage    marketplace.Page
		err         error
		filters     cards.SliceFilters
		limit, page int
	)
	urlQuery := r.URL.Query()
	limitQuery := urlQuery.Get("limit")
	pageQuery := urlQuery.Get("page")
	playerName := urlQuery.Get(string(cards.FilterPlayerName))

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
	if playerName == "" {
		if err := filters.DecodingURLParameters(urlQuery); err != nil {
			controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		}
		if len(filters) > 0 {
			lotsPage, err = controller.marketplace.ListActiveLotsWithFilters(ctx, filters, cursor)
		} else {
			lotsPage, err = controller.marketplace.ListActiveLots(ctx, cursor)
		}
	} else {
		filter := cards.Filters{
			Name:           cards.FilterPlayerName,
			Value:          playerName,
			SearchOperator: sqlsearchoperators.LIKE,
		}
		lotsPage, err = controller.marketplace.ListActiveLotsByPlayerName(ctx, filter, cursor)
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

	var lots []marketplace.Lot
	for _, oneLot := range lotsPage.Lots {
		lot := marketplace.Lot{
			CardID:       oneLot.Card.ID,
			Type:         oneLot.Type,
			UserID:       oneLot.UserID,
			ShopperID:    oneLot.ShopperID,
			Status:       oneLot.Status,
			StartPrice:   oneLot.StartPrice,
			MaxPrice:     oneLot.MaxPrice,
			CurrentPrice: *evmsignature.WeiBigToEthereumBig(&oneLot.CurrentPrice),
			StartTime:    oneLot.StartTime,
			EndTime:      oneLot.EndTime,
			Period:       oneLot.Period,
			Card:         oneLot.Card,
		}
		lots = append(lots, lot)
	}

	lotsListPage := marketplace.Page{
		Lots: lots,
		Page: lotsPage.Page,
	}

	if err = json.NewEncoder(w).Encode(lotsListPage); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// GetLotByID is an endpoint that returns lot by id.
func (controller *Marketplace) GetLotByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	vars := mux.Vars(r)

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

	getLot := struct {
		CardID       uuid.UUID          `json:"cardId"`
		Type         marketplace.Type   `json:"type"`
		Status       marketplace.Status `json:"status"`
		StartPrice   float64            `json:"startPrice"`
		MaxPrice     float64            `json:"maxPrice"`
		CurrentPrice float64            `json:"currentPrice"`
		StartTime    time.Time          `json:"startTime"`
		EndTime      time.Time          `json:"endTime"`
		Period       marketplace.Period `json:"period"`
		Card         cards.Card         `json:"card"`
	}{
		CardID:       lot.Card.ID,
		Type:         lot.Type,
		Status:       lot.Status,
		StartPrice:   roundFloat(evmsignature.WeiBigToEthereumFloat(&lot.StartPrice), 3),
		MaxPrice:     roundFloat(evmsignature.WeiBigToEthereumFloat(&lot.MaxPrice), 3),
		CurrentPrice: roundFloat(evmsignature.WeiBigToEthereumFloat(&lot.CurrentPrice), 3),
		StartTime:    lot.StartTime,
		EndTime:      lot.EndTime,
		Period:       lot.Period,
		Card:         lot.Card,
	}

	if err = json.NewEncoder(w).Encode(getLot); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// GetCurrentPriceByCardID is an endpoint that returns price by card id.
func (controller *Marketplace) GetCurrentPriceByCardID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	vars := mux.Vars(r)

	cardID, err := uuid.Parse(vars["card_id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	cardPrice, err := controller.marketplace.GetCurrentPriceByCardID(ctx, cardID)
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

	if err = json.NewEncoder(w).Encode(cardPrice.Int64()); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// CreateLot is an endpoint that creates lot.
func (controller *Marketplace) CreateLot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
		return
	}

	var createLot marketplace.CreateLot
	if err = json.NewDecoder(r.Body).Decode(&createLot); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	// TODO: remove after adding Casper contract.
	// if _, err = controller.marketplace.GetNFTByCardID(ctx, createLot.CardID); err != nil {
	//  controller.log.Error("there is no such NFT", ErrMarketplace.Wrap(err))
	//  controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
	// } .

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
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
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

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrMarketplace.Wrap(err))
	}
}

// roundFloat floating-point rounding function.
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
