// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/seasons"
)

var (
	// ErrSeasons is an internal error type for seasons controller.
	ErrSeasons = errs.Class("seasons controller error")
)

// Seasons is a mvc controller that handles all seasons related views.
type Seasons struct {
	log logger.Logger

	seasons *seasons.Service
}

// NewSeasons is a constructor for seasons controller.
func NewSeasons(log logger.Logger, seasons *seasons.Service) *Seasons {
	seasonsController := &Seasons{
		log:     log,
		seasons: seasons,
	}

	return seasonsController
}

// GetCurrentSeasons returns all current seasons.
func (controller *Seasons) GetCurrentSeasons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	seasons, err := controller.seasons.GetCurrentSeasons(ctx)
	if err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrSeasons.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(seasons); err != nil {
		controller.log.Error("failed to write json response", ErrSeasons.Wrap(err))
		return
	}
}

// serveError replies to request with specific code and error.
func (controller *Seasons) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrSeasons.Wrap(err))
	}
}
