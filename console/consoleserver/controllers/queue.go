// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/queue"
)

var (
	// ErrQueue is an internal error type for queue controller.
	ErrQueue = errs.Class("queue controller error")
)

// Queue is a mvc controller that handles all queue related views.
type Queue struct {
	log   logger.Logger
	queue *queue.Service
}

// NewQueue is a constructor for queue controller.
func NewQueue(log logger.Logger, queue *queue.Service) *Queue {
	queueController := &Queue{
		log:   log,
		queue: queue,
	}

	return queueController
}

// Create is an endpoint that creates place.
func (controller *Queue) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrQueue.Wrap(err))
		return
	}

	place := queue.Place{
		UserID: claims.UserID,
		Status: queue.StatusSearches,
	}

	if err = controller.queue.Create(ctx, place); err != nil {
		controller.log.Error("could not create place", ErrQueue.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrQueue.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Queue) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrQueue.Wrap(err))
	}
}
