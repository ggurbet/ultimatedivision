// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/store"
)

var (
	// ErrStore is an internal error type for store controller.
	ErrStore = errs.Class("store controller error")
)

// StoreTemplates holds all store related templates.
type StoreTemplates struct {
	List   *template.Template
	Update *template.Template
}

// Store is a mvc controller that handles all admins related views.
type Store struct {
	log logger.Logger

	store *store.Service

	templates StoreTemplates
}

// NewStore is a constructor for store controller.
func NewStore(log logger.Logger, store *store.Service, templates StoreTemplates) *Store {
	storeController := &Store{
		log:       log,
		store:     store,
		templates: templates,
	}

	return storeController
}

// List is an endpoint that will provide a web page with all settings.
func (controller *Store) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	store, err := controller.store.List(ctx)
	if err != nil {
		controller.log.Error("could not get store list", ErrStore.Wrap(err))
		http.Error(w, "could not get store list", http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, store)
	if err != nil {
		controller.log.Error("can not execute list store template", ErrStore.Wrap(err))
		http.Error(w, "can not execute list store template", http.StatusInternalServerError)
		return
	}
}

// Update is an endpoint that will update setting of store.
func (controller *Store) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	settingID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "could not parse setting id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		setting, err := controller.store.Get(ctx, settingID)
		if err != nil {
			controller.log.Error("could not get setting", ErrStore.Wrap(err))
			if store.ErrNoSetting.Has(err) {
				http.Error(w, "no setting with such id", http.StatusNotFound)
				return
			}
			http.Error(w, "could not get setting", http.StatusInternalServerError)
			return
		}

		if err = controller.templates.Update.Execute(w, setting); err != nil {
			controller.log.Error("could not execute update store template", ErrStore.Wrap(err))
			http.Error(w, "could not execute update store template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not get store form", http.StatusBadRequest)
			return
		}

		cardsAmount, err := strconv.Atoi(r.FormValue("cardsAmount"))
		if err != nil {
			http.Error(w, "invalid cardsAmount", http.StatusBadRequest)
			return
		}

		isRenewal, err := strconv.ParseBool(r.FormValue("isRenewal"))
		if err != nil {
			http.Error(w, "invalid isRenewal", http.StatusBadRequest)
			return
		}

		dateRenewal, err := time.Parse("2006-01-02T15:04", r.FormValue("dateRenewal"))
		if err != nil {
			http.Error(w, "invalid dateRenewal", http.StatusBadRequest)
			return
		}

		setting := store.Setting{
			ID:          settingID,
			CardsAmount: cardsAmount,
			IsRenewal:   isRenewal,
			DateRenewal: dateRenewal,
		}

		if err = controller.store.Update(ctx, setting); err != nil {
			controller.log.Error("could not update store status", ErrStore.Wrap(err))
			http.Error(w, "could not update store status", http.StatusInternalServerError)
			return
		}
		Redirect(w, r, "/store", http.MethodGet)
	}
}
