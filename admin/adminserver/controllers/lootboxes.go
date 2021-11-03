// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
)

// ErrLootBoxes is an internal error type for loot boxes controller.
var ErrLootBoxes = errs.Class("lootboxes controller error")

// LootBoxesTemplates holds all lootboxes related templates.
type LootBoxesTemplates struct {
	List      *template.Template
	Create    *template.Template
	ListCards *template.Template
}

// LootBoxes is a mvc controller that handles all lootboxes related views.
type LootBoxes struct {
	log logger.Logger

	lootboxes *lootboxes.Service

	templates LootBoxesTemplates
}

// NewLootBoxes is a constructor for loot boxes controller.
func NewLootBoxes(log logger.Logger, lootboxes *lootboxes.Service, templates LootBoxesTemplates) *LootBoxes {
	lootBoxesController := &LootBoxes{
		log:       log,
		lootboxes: lootboxes,
		templates: templates,
	}

	return lootBoxesController
}

// Create is an endpoint that creates loot box for user.
func (controller *LootBoxes) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if err = controller.templates.Create.Execute(w, id); err != nil {
			controller.log.Error("could not execute template", ErrLootBoxes.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		lootBoxType := r.FormValue("lootbox")
		if _, err = controller.lootboxes.Create(ctx, lootboxes.Type(lootBoxType), id); err != nil {
			controller.log.Error("could not create loot box", ErrLootBoxes.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Redirect(w, r, "/lootboxes", http.MethodGet)
	}
}

// Open is an endpoint that opens loot box by user.
func (controller *LootBoxes) Open(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	userID, err := uuid.Parse(vars["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lootboxID, err := uuid.Parse(vars["lootboxId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cards, err := controller.lootboxes.Open(ctx, userID, lootboxID)
	if err != nil {
		controller.log.Error("could not open loot box", ErrLootBoxes.Wrap(err))
		switch {
		case lootboxes.ErrNoLootBox.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err = controller.templates.ListCards.Execute(w, cards); err != nil {
		controller.log.Error("could not execute template", ErrLootBoxes.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// List is an endpoint that will provide a web page with all users loot boxes.
func (controller *LootBoxes) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lootBoxes, err := controller.lootboxes.List(ctx)
	if err != nil {
		controller.log.Error("could not list loot boxes", ErrLootBoxes.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = controller.templates.List.Execute(w, lootBoxes); err != nil {
		controller.log.Error("could not execute template", ErrLootBoxes.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
