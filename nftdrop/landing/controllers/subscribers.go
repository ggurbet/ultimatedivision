// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/subscribers"
)

var (
	// ErrSubscribers is an internal error type for subscribers controller.
	ErrSubscribers = errs.Class("subscribers controller error")
)

// Subscribers is a mvc controller that handles all subscribers related views.
type Subscribers struct {
	log logger.Logger

	subscriber *subscribers.Service
}

// NewEmails is a constructor for subscribers controller.
func NewEmails(log logger.Logger, subscribers *subscribers.Service) *Subscribers {
	subscribersController := &Subscribers{
		log:        log,
		subscriber: subscribers,
	}

	return subscribersController
}

// Create is an endpoint that writes subscriber.
func (controller *Subscribers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request subscribers.CreateSubscriberFields

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrSubscribers.Wrap(err))
		return
	}

	err = controller.subscriber.Create(ctx, request.Email)
	if err != nil {
		switch {
		case subscribers.ErrSubscribersDB.Has(err):
			controller.serveError(w, http.StatusBadRequest, errs.New("Address is already in use"))
			return
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrSubscribers.Wrap(err))
			return
		}
	}

	if err = json.NewEncoder(w).Encode("success"); err != nil {
		controller.log.Error("failed to write json response", ErrSubscribers.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Subscribers) serveError(w http.ResponseWriter, status int, err error) {
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
