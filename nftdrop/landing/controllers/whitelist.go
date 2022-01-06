// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BoostyLabs/evmsignature"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/whitelist"
)

var (
	// ErrWhitelist is an internal error type for whitelist controller.
	ErrWhitelist = errs.Class("whitelist controller error")
)

// Whitelist is a mvc controller that handles all whitelist related views.
type Whitelist struct {
	log logger.Logger

	whitelist *whitelist.Service
}

// NewWhitelist is a constructor for whitelist controller.
func NewWhitelist(log logger.Logger, whitelist *whitelist.Service) *Whitelist {
	whitelistController := &Whitelist{
		log:       log,
		whitelist: whitelist,
	}

	return whitelistController
}

// Get is an endpoint that returns password by wallet address.
func (controller *Whitelist) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	address := evmsignature.Address(params["address"])

	if !address.IsValidAddress() {
		controller.serveError(w, http.StatusBadRequest, ErrWhitelist.New("invalid address"))
	}

	transactionValue, err := controller.whitelist.GetByAddress(ctx, address)
	if err != nil {
		controller.log.Error("could don't get transaction value", ErrWhitelist.Wrap(err))

		if whitelist.ErrNoWallet.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrWhitelist.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrWhitelist.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(transactionValue); err != nil {
		controller.log.Error("could not transaction value with json", ErrWhitelist.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrWhitelist.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Whitelist) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrWhitelist.Wrap(err))
	}
}
