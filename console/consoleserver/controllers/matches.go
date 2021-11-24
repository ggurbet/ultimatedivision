// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
)

var (
	// ErrMatches is an internal error type for matches controller.
	ErrMatches = errs.Class("matches controller error")
)

// Matches is a mvc controller that handles all matches related views.
type Matches struct {
	log logger.Logger

	matches *matches.Service
}

// NewMatches is a constructor for matches controller.
func NewMatches(log logger.Logger, matches *matches.Service) *Matches {
	matchesController := &Matches{
		log:     log,
		matches: matches,
	}

	return matchesController
}

// serveError replies to request with specific code and error.
func (controller *Matches) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrMatches.Wrap(err))
	}
}
