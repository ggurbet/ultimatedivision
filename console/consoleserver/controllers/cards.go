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
	"ultimatedivision/users/userauth"
)

var (
	// ErrCards is an internal error type for cards controller.
	ErrCards = errs.Class("cards controller error")
)

const (
	// numberPositionOfURLParameter is a number that shows the position of the url parameter.
	numberPositionOfURLParameter = 0
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
	w.Header().Set("Content-Type", "application/json")
	var (
		cardsList []cards.Card
		err       error
		filters   []cards.Filters
	)
	urlQuery := r.URL.Query()

	if len(urlQuery) > 0 {
		for key, value := range urlQuery {
			filter := cards.Filters{
				Name:           "",
				Value:          value[numberPositionOfURLParameter],
				SearchOperator: "",
			}

			for k, v := range sqlsearchoperators.SearchOperators {
				if strings.HasSuffix(key, k) {
					countName := len(key) - (1 + len(k))
					filter.Name = cards.Filter(key[:countName])
					filter.SearchOperator = v
				}
			}

			keyFilter := cards.Filter(key)
			if keyFilter == cards.FilterQuality || keyFilter == cards.FilterDominantFoot || keyFilter == cards.FilterType {
				filter.Name = cards.Filter(key)
				filter.SearchOperator = sqlsearchoperators.EQ
			}

			if filter.Name == "" {
				controller.serveError(w, http.StatusBadRequest, cards.ErrInvalidFilter.New("invalid name parameter - "+key))
				return
			}

			filters = append(filters, filter)
		}
	}

	if len(filters) > 0 {
		cardsList, err = controller.cards.ListWithFilters(ctx, filters)
	} else {
		cardsList, err = controller.cards.List(ctx)
	}
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		switch {
		case userauth.ErrUnauthenticated.Has(err):
			controller.serveError(w, http.StatusUnauthorized, ErrCards.Wrap(err))
		case cards.ErrNoCard.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrCards.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrCards.Wrap(err))
		}
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
	w.Header().Set("Content-Type", "application/json")

	var filter cards.Filters
	playerName := r.URL.Query().Get(string(cards.FilterPlayerName))
	if playerName == "" {
		controller.serveError(w, http.StatusBadRequest, ErrCards.New("empty player name parameter"))
		return
	}
	filter = cards.Filters{
		Name:           cards.FilterPlayerName,
		Value:          playerName,
		SearchOperator: sqlsearchoperators.LIKE,
	}

	cardsList, err := controller.cards.ListByPlayerName(ctx, filter)
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		switch {
		case userauth.ErrUnauthenticated.Has(err):
			controller.serveError(w, http.StatusUnauthorized, ErrCards.Wrap(err))
		case cards.ErrNoCard.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrCards.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrCards.Wrap(err))
		}
		return
	}

	if err = json.NewEncoder(w).Encode(cardsList); err != nil {
		controller.log.Error("failed to write json response", ErrCards.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Cards) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrCards.Wrap(err))
	}
}
