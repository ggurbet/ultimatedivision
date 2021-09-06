// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
)

var (
	// ErrLootBoxes is an internal error type for lootboxes controller.
	ErrLootBoxes = errs.Class("lootboxes controller error")
)

// LootBoxes is a mvc controller that handles all lootboxes related views.
type LootBoxes struct {
	log logger.Logger

	lootBoxes *lootboxes.Service
}

// NewLootBoxes is a constructor for lootboxes controller.
func NewLootBoxes(log logger.Logger, lootBoxes *lootboxes.Service) *LootBoxes {
	lootBoxesController := &LootBoxes{
		log:       log,
		lootBoxes: lootBoxes,
	}

	return lootBoxesController
}

// Create is an endpoint that creates new lootbox for user.
func (controller *LootBoxes) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var lootBox lootboxes.LootBox

	if err := json.NewDecoder(r.Body).Decode(&lootBox); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrLootBoxes.Wrap(err))
		return
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrLootBoxes.Wrap(err))
		return
	}

	userLootBox, err := controller.lootBoxes.Create(ctx, lootBox.Type, claims.UserID)
	if err != nil {
		controller.log.Error("could not create loot box for user", ErrLootBoxes.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrLootBoxes.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(userLootBox); err != nil {
		controller.log.Error("could not response with json", ErrLootBoxes.Wrap(err))
		return
	}
}

// Open is an endpoint that opens user lootbox.
func (controller *LootBoxes) Open(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrLootBoxes.Wrap(err))
		return
	}

	params := mux.Vars(r)
	lootboxID, err := uuid.Parse(params["id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrLootBoxes.New("lootbox id is missing or invalid"))
		return
	}

	cards, err := controller.lootBoxes.Open(ctx, claims.UserID, lootboxID)
	if err != nil {
		controller.log.Error("could not open loot box", ErrLootBoxes.Wrap(err))
		switch {
		case lootboxes.ErrNoLootBox.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrLootBoxes.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrLootBoxes.Wrap(err))
		}
		return
	}

	if err = json.NewEncoder(w).Encode(cards); err != nil {
		controller.log.Error("could not encode lootbox cards", ErrLootBoxes.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *LootBoxes) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrLootBoxes.Wrap(err))
	}
}
