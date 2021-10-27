// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/pagination"
)

// ErrMatches is an internal error type for matches controller.
var ErrMatches = errs.Class("matches controller error")

// MatchesTemplate holds all matches related templates.
type MatchesTemplate struct {
	List      *template.Template
	ListGoals *template.Template
	Create    *template.Template
}

// Matches is a mvc controller that handles all matches related views.
type Matches struct {
	log logger.Logger

	matches *matches.Service

	templates MatchesTemplate
}

// NewMatches is a constructor for matches controller.
func NewMatches(log logger.Logger, matches *matches.Service, templates MatchesTemplate) *Matches {
	matchesController := &Matches{
		log:       log,
		matches:   matches,
		templates: templates,
	}

	return matchesController
}

// Create is an endpoint that will create match and initiate it.
func (controller *Matches) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Create.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute create matches template", ErrMatches.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		ctx := r.Context()

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user1ID, err := uuid.Parse(r.FormValue("user1"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		squad1ID, err := uuid.Parse(r.FormValue("squad1"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user2ID, err := uuid.Parse(r.FormValue("user2"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		squad2ID, err := uuid.Parse(r.FormValue("squad2"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = controller.matches.Create(ctx, squad1ID, squad2ID, user1ID, user2ID)
		if err != nil {
			controller.log.Error("could not create match", ErrMatches.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "", http.MethodGet)
	}
}

// ListMatches is an endpoint that will provide a web page with matches.
func (controller *Matches) ListMatches(w http.ResponseWriter, r *http.Request) {
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

	matchesPage, err := controller.matches.List(ctx, cursor)
	if err != nil {
		controller.log.Error("could not list matches", ErrMatches.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, matchesPage)
	if err != nil {
		controller.log.Error("could not execute list matches template", ErrMatches.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListMatchGoals is an endpoint that will provide a web page with all goals from match.
func (controller *Matches) ListMatchGoals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	matchID, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	matchGoals, err := controller.matches.ListMatchGoals(ctx, matchID)
	if err != nil {
		controller.log.Error("could not list match goals", ErrMatches.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = controller.templates.ListGoals.Execute(w, matchGoals)
	if err != nil {
		controller.log.Error("could not execute list goals template", ErrMatches.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete is an endpoint that deletes match.
func (controller *Matches) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	matchID, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.matches.Delete(ctx, matchID)
	if err != nil {
		controller.log.Error("could not delete match", ErrMatches.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/matches", http.MethodGet)
}
