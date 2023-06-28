// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
	"ultimatedivision/marketplace/bids"
	"ultimatedivision/pkg/auth"
)

var (
	// ErrBids is an internal error type for bids controller.
	ErrBids = errs.Class("bids controller error")
)

// Bids is a mvc controller that handles all bids related views.
type Bids struct {
	log         logger.Logger
	bids        *bids.Service
	marketplace *marketplace.Service
}

// NewBids is a constructor for bids controller.
func NewBids(log logger.Logger, bids *bids.Service, marketplace *marketplace.Service) *Bids {
	bidsController := &Bids{
		log:         log,
		bids:        bids,
		marketplace: marketplace,
	}

	return bidsController
}

// BidResponse defines response for card bids.
type BidResponse struct {
	ID        uuid.UUID `json:"id"`
	LotID     uuid.UUID `json:"lotId"`
	UserID    uuid.UUID `json:"userId"`
	UserName  string    `json:"userName"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

// Bet is an endpoint that place bet of lot.
func (controller *Bids) Bet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrUsers.Wrap(err))
		return
	}

	type request struct {
		LotID  uuid.UUID `json:"lotId"`
		Amount float64   `json:"amount"`
	}

	var req request
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrBids.Wrap(err))
		return
	}

	amount, err := evmsignature.EthereumFloatToWeiBig(req.Amount)
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrBids.Wrap(err))
		return
	}

	bid := bids.Bid{
		LotID:  req.LotID,
		UserID: claims.UserID,
		Amount: *amount,
	}

	if err := controller.marketplace.UpdateShopperIDLot(ctx, bid.LotID, bid.UserID); err != nil {
		if errs.Is(err, bids.ErrSmallAmountOfBid) {
			controller.serveError(w, http.StatusBadRequest, ErrUsers.Wrap(err))
			return
		}
	}

	if err = controller.bids.Create(ctx, bid); err != nil {
		if errs.Is(err, bids.ErrSmallAmountOfBid) {
			controller.serveError(w, http.StatusBadRequest, ErrUsers.Wrap(err))
			return
		}
		controller.log.Error(fmt.Sprintf("could not create bet with lot %x, user %x and amount %v", bid.LotID, bid.UserID, bid.Amount), ErrBids.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrBids.Wrap(err))
		return
	}

	lotData, err := controller.marketplace.GetMakeOfferByCardID(ctx, bid.LotID)
	if err != nil {
		controller.log.Error("there is no such NFT data", ErrMarketplace.Wrap(err))
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
	}

	if err = json.NewEncoder(w).Encode(lotData); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// GetMakeOfferData is an endpoint that return make offer data by lot id.
func (controller *Bids) GetMakeOfferData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	vars := mux.Vars(r)

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrUsers.Wrap(err))
		return
	}

	cardID, err := uuid.Parse(vars["card_id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	lotData, err := controller.bids.GetMakeOfferData(ctx, cardID, claims.UserID)
	if err != nil {
		controller.log.Error("there is no such NFT data", ErrMarketplace.Wrap(err))
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
	}

	if err = json.NewEncoder(w).Encode(lotData); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// ListByLotID is an endpoint that returns bids by lot id.
func (controller *Bids) ListByLotID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	vars := mux.Vars(r)

	lotID, err := uuid.Parse(vars["lotId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrBids.Wrap(err))
		return
	}

	cardBids, err := controller.bids.ListByLotID(ctx, lotID)
	if err != nil {
		controller.log.Error(fmt.Sprintf("could not get lot %x bids", lotID), ErrBids.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrBids.Wrap(err))
		return
	}

	type response struct {
		Bids []BidResponse `json:"bids"`
	}
	var res response

	for _, cardBid := range cardBids {
		bid := BidResponse{
			ID:        cardBid.ID,
			LotID:     cardBid.LotID,
			UserID:    cardBid.UserID,
			UserName:  cardBid.UserName,
			Amount:    evmsignature.WeiBigToEthereumFloat(&cardBid.Amount),
			CreatedAt: cardBid.CreatedAt,
		}

		res.Bids = append(res.Bids, bid)
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		controller.log.Error("failed to write json error response", ErrBids.Wrap(err))
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Bids) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()
	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrBids.Wrap(err))
	}
}
