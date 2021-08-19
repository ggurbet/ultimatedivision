// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/sqlsearchoperators"
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
	var filters []cards.Filters
	urlQuery := r.URL.Query()

	for key, value := range urlQuery {
		for k, v := range sqlsearchoperators.SearchOperators {
			name := key
			action := sqlsearchoperators.EQ

			if strings.HasSuffix(key, k) {
				countName := len(key) - 1 + len(k)
				name = key[:countName]
				action = v
			}

			keyFilter := cards.Filter(key)
			if keyFilter == cards.FilterQuality || keyFilter == cards.FilterDominantFoot || keyFilter == cards.FilterType {
				action = sqlsearchoperators.EQ
			}

			filter := cards.Filters{
				Name:           cards.Filter(name),
				Value:          value[0],
				SearchOperator: action,
			}
			filters = append(filters, filter)
			break
		}
	}

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

// ListByPlayerName is an endpoint that allows will view cards by player name.
func (controller *Cards) ListByPlayerName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var filter cards.Filters
	playerName := r.URL.Query().Get(string(cards.FilterPlayerName))
	if playerName != "" {
		filter = cards.Filters{
			Name:           cards.FilterPlayerName,
			Value:          playerName,
			SearchOperator: sqlsearchoperators.LIKE,
		}
	}

	cardsList, err := controller.cards.ListByPlayerName(ctx, filter)
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
