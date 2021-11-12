// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/users"
)

// ErrWaitList is an internal error type for waitList controller.
var ErrWaitList = errs.Class("waitList controller error")

// WaitList is a mvc controller that handles all WaitList related views.
type WaitList struct {
	log      logger.Logger
	waitList *waitlist.Service
}

// NewWaitList is a constructor for WaitList controller.
func NewWaitList(log logger.Logger, waitList *waitlist.Service) *WaitList {
	WaitListController := &WaitList{
		log:      log,
		waitList: waitList,
	}

	return WaitListController
}

// Create is an endpoint that creates nft token.
func (controller *WaitList) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrWaitList.Wrap(err))
		return
	}

	var createNFT waitlist.CreateNFT

	if err = json.NewDecoder(r.Body).Decode(&createNFT); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrWaitList.Wrap(err))
		return
	}
	createNFT.UserID = claims.UserID

	err = controller.waitList.Create(ctx, createNFT)
	if err != nil {
		controller.log.Error("could not create nft token", ErrWaitList.Wrap(err))

		if users.ErrNoUser.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrWaitList.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrWaitList.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *WaitList) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrWaitList.Wrap(err))
	}
}
