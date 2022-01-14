// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/waitlist"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/store"
)

var (
	// ErrStore is an internal error type for store controller.
	ErrStore = errs.Class("store controller error")
)

// Store is a mvc controller that handles store related views.
type Store struct {
	log logger.Logger

	store *store.Service
}

// NewStore is a constructor for store controller.
func NewStore(log logger.Logger, store *store.Service) *Store {
	storeController := &Store{
		log:   log,
		store: store,
	}

	return storeController
}

// Buy is an endpoint that allows to view details of store.
func (controller *Store) Buy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrStore.Wrap(err))
		return
	}

	var createNFT waitlist.CreateNFT

	if err = json.NewDecoder(r.Body).Decode(&createNFT); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrStore.Wrap(err))
		return
	}
	if !createNFT.WalletAddress.IsValidAddress() {
		controller.serveError(w, http.StatusBadRequest, ErrStore.New("wallet address is invalid"))
		return
	}
	createNFT.UserID = claims.UserID

	transaction, err := controller.store.Buy(ctx, createNFT)
	if err != nil {
		switch {
		case cards.ErrNoCard.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrStore.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrStore.Wrap(err))
			controller.log.Error("could not buy card", ErrStore.Wrap(err))
		}
		return
	}

	if err = json.NewEncoder(w).Encode(transaction); err != nil {
		controller.log.Error("failed to write json response", ErrStore.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Store) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrStore.Wrap(err))
	}
}
