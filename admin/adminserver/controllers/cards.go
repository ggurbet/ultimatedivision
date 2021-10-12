// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/pagination"
)

var (
	// ErrCards is an internal error type for cards controller.
	ErrCards = errs.Class("cards controller error")
)

// CardTemplates holds all cards related templates.
type CardTemplates struct {
	List *template.Template
}

// Cards is a mvc controller that handles all cards related views.
type Cards struct {
	log logger.Logger

	cards *cards.Service

	templates CardTemplates

	percentageQualities cards.PercentageQualities
}

// NewCards is a constructor for cards controller.
func NewCards(log logger.Logger, cards *cards.Service, templates CardTemplates, percentageQualities cards.PercentageQualities) *Cards {
	cardsController := &Cards{
		log:                 log,
		cards:               cards,
		templates:           templates,
		percentageQualities: percentageQualities,
	}

	return cardsController
}

// List is an endpoint that will provide a web page with all cards.
func (controller *Cards) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		err         error
		limit, page int
	)
	urlQuery := r.URL.Query()
	limitQuery := urlQuery.Get("limit")
	pageQuery := urlQuery.Get("page")

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	cursor := pagination.Cursor{
		Limit: limit,
		Page:  page,
	}
	listCardsPage, err := controller.cards.List(ctx, cursor)
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		switch {
		case cards.ErrNoCard.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = controller.templates.List.Execute(w, listCardsPage)
	if err != nil {
		controller.log.Error("can not execute list cards template", ErrCards.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that will add card to database.
func (controller *Cards) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	percentageQualities := []int{
		controller.percentageQualities.Wood,
		controller.percentageQualities.Silver,
		controller.percentageQualities.Gold,
		controller.percentageQualities.Diamond,
	}

	userID, err := uuid.Parse(vars["userId"])
	if err != nil {
		http.Error(w, "could not parse user id", http.StatusBadRequest)
		return
	}

	if _, err := controller.cards.Create(ctx, userID, percentageQualities, strconv.Itoa((rand.Intn(99) + 1))); err != nil {
		controller.log.Error("could not create card", ErrCards.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/cards", "GET")
}

// Delete is an endpoint that will destroy record card to database.
func (controller *Cards) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	if vars["id"] == "" {
		http.Error(w, "id parameter is empty", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "could not parse card id", http.StatusBadRequest)
		return
	}
	if err := controller.cards.Delete(ctx, id); err != nil {
		controller.log.Error("could not delete card", ErrCards.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/cards", "GET")
}
