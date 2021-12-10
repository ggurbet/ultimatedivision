// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/pagination"
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

// Get is an endpoint that allows to view details of cards.
func (controller *Cards) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrCards.Wrap(err))
		return
	}

	card, err := controller.cards.Get(ctx, id)
	if err != nil {
		controller.log.Error("could not get card", ErrCards.Wrap(err))
		switch {
		case cards.ErrNoCard.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrCards.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrCards.Wrap(err))
		}
		return
	}

	if err = json.NewEncoder(w).Encode(card); err != nil {
		controller.log.Error("failed to write json response", ErrCards.Wrap(err))
		return
	}
}

// List is an endpoint that allows will view cards.
func (controller *Cards) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var (
		cardsListPage cards.Page
		err           error
		filters       cards.SliceFilters
		limit, page   int
	)
	urlQuery := r.URL.Query()
	limitQuery := urlQuery.Get("limit")
	pageQuery := urlQuery.Get("page")

	if limitQuery != "" {
		if limit, err = strconv.Atoi(limitQuery); err != nil {
			controller.serveError(w, http.StatusBadRequest, ErrCards.Wrap(err))
			return
		}
	}

	if pageQuery != "" {
		if page, err = strconv.Atoi(pageQuery); err != nil {
			controller.serveError(w, http.StatusBadRequest, ErrCards.Wrap(err))
			return
		}
	}

	cursor := pagination.Cursor{
		Limit: limit,
		Page:  page,
	}
	playerName := urlQuery.Get(string(cards.FilterPlayerName))

	if playerName == "" {
		if err := filters.DecodingURLParameters(urlQuery); err != nil {
			controller.serveError(w, http.StatusBadRequest, ErrCards.Wrap(err))
		}
		if len(filters) > 0 {
			cardsListPage, err = controller.cards.ListWithFilters(ctx, filters, cursor)
		} else {
			cardsListPage, err = controller.cards.List(ctx, cursor)
		}
	} else {
		filter := cards.Filters{
			Name:           cards.FilterPlayerName,
			Value:          playerName,
			SearchOperator: sqlsearchoperators.LIKE,
		}
		cardsListPage, err = controller.cards.ListByPlayerName(ctx, filter, cursor)
	}
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		switch {
		case cards.ErrNoCard.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrCards.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrCards.Wrap(err))
		}
		return
	}

	if err = json.NewEncoder(w).Encode(cardsListPage); err != nil {
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
