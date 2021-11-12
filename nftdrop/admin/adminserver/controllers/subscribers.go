// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/subscribers"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/users"
)

var (
	// ErrSubscribers is an internal error type for subscribers controller.
	ErrSubscribers = errs.Class("subscribers controller error")
)

// SubscribersTemplates holds all subscribers related templates.
type SubscribersTemplates struct {
	List *template.Template
}

// Subscribers is a mvc controller that handles all subscribers related views.
type Subscribers struct {
	log logger.Logger

	subscribers *subscribers.Service

	templates SubscribersTemplates
}

// NewSubscribers is a constructor for subscribers controller.
func NewSubscribers(log logger.Logger, subscribers *subscribers.Service, templates SubscribersTemplates) *Subscribers {
	subscribersController := &Subscribers{
		log:         log,
		subscribers: subscribers,
		templates:   templates,
	}

	return subscribersController
}

// List is an endpoint that will provide a web page with subscribers page.
func (controller *Subscribers) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		err   error
		limit int
		page  int
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

	subscribersPage, err := controller.subscribers.List(ctx, cursor)
	if err != nil {
		controller.log.Error("could not list subscribers", ErrSubscribers.Wrap(err))
		http.Error(w, "could not list subscribers", http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, subscribersPage)
	if err != nil {
		controller.log.Error("could not execute list subscriber template", ErrSubscribers.Wrap(err))
		http.Error(w, "could not execute list subscriber template", http.StatusInternalServerError)
		return
	}
}

// Delete is an endpoint that deletes subscriber.
func (controller *Subscribers) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	email := params["email"]
	if email == "" {
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	}

	err := controller.subscribers.Delete(ctx, email)
	if err != nil {
		controller.log.Error("could not delete subscriber by email", ErrSubscribers.Wrap(err))
		if users.ErrNoUser.Has(err) {
			http.Error(w, "no subscriber with such email", http.StatusNotFound)
			return
		}
		http.Error(w, "could not delete subscriber", http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/subscribers", http.MethodGet)
}
