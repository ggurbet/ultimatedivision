// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
)

var (
	// ErrCards is an internal error type for cards controller.
	ErrCards = errs.Class("cards controller error")
)

// Cards is a mvc controller that handles all cards related views.
type Cards struct {
	log logger.Logger

	cards *cards.Service
}

// NewCards is a constructor for cards controller.
func NewCards(log logger.Logger, cards *cards.Service) *Cards {
	cardsController := &Cards{
		log:   log,
		cards: cards,
	}

	return cardsController
}

// List is an endpoint that allows will view cards.
func (controller *Cards) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var cardsList []cards.Card
	var err error
	var filters cards.SliceFilters
	urlQuery := r.URL.Query()
	tactics := urlQuery.Get(string(cards.Tactics))
	minPhysique := urlQuery.Get(string(cards.MinPhysique))
	maxPhysique := urlQuery.Get(string(cards.MaxPhysique))
	playerName := urlQuery.Get(string(cards.PlayerName))

	filters.Add(cards.Tactics, tactics)
	filters.Add(cards.MinPhysique, minPhysique)
	filters.Add(cards.MaxPhysique, maxPhysique)
	filters.Add(cards.PlayerName, playerName)

	if len(filters) > 0 {
		cardsList, err = controller.cards.ListWithFilters(ctx, filters)
	} else {
		cardsList, err = controller.cards.List(ctx)
	}
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		http.Error(w, "could not get cards list", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(cardsList); err != nil {
		controller.log.Error("failed to write json response", ErrCards.Wrap(err))
		return
	}
}
