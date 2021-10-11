// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

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
	address := whitelist.Hex(params["address"])

	if !address.IsValidAddress() {
		controller.serveError(w, http.StatusBadRequest, ErrWhitelist.New("invalid address"))
	}

	smartContractWithWhiteList, err := controller.whitelist.GetByAddress(ctx, address)
	if err != nil {
		controller.log.Error("could get password", ErrWhitelist.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrWhitelist.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(smartContractWithWhiteList); err != nil {
		controller.log.Error("could not response with json", ErrWhitelist.Wrap(err))
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
