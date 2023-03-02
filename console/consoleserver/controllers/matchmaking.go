// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/matchmaking"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
)

var (
	// ErrMatchmaking is an internal error type for matchmaking controller.
	ErrMatchmaking = errs.Class("matchmaking controller error")
)

// Matchmaking is a mvc controller that handles all matchmaking related views.
type Matchmaking struct {
	log logger.Logger

	matchmaking *matchmaking.Service
}

// NewMatchmaking is a constructor for matchmaking controller.
func NewMatchmaking(log logger.Logger, matchmaking *matchmaking.Service) *Matchmaking {
	matchmakingController := &Matchmaking{
		log:         log,
		matchmaking: matchmaking,
	}

	return matchmakingController
}

// Create is an endpoint that creates matchmaking of players.
func (controller *Matchmaking) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrMatchmaking.Wrap(err))
		return
	}

	err = controller.matchmaking.Delete(claims.UserID)
	if err != nil {
		if !matchmaking.ErrNoPlayer.Has(err) {
			controller.log.Error(fmt.Sprintf("could not delete old player for user %x", claims.UserID), ErrMatchmaking.Wrap(err))
			controller.serveError(w, http.StatusInternalServerError, ErrMatchmaking.Wrap(err))
			return
		}
	}

	if err = controller.matchmaking.Create(claims.UserID); err != nil {
		controller.log.Error(fmt.Sprintf("could not create player for user %x", claims.UserID), ErrMatchmaking.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMatchmaking.Wrap(err))
		return
	}
}

// serveError replies to request with specific code and error.
func (controller *Matchmaking) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrMatchmaking.Wrap(err))
	}
}
