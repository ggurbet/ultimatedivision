// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/divisions"
	"ultimatedivision/internal/logger"
)

var (
	// ErrDivisions is an internal error type for divisions controller.
	ErrDivisions = errs.Class("divisions controller error")
)

// DivisionsTemplates holds all divisions related templates.
type DivisionsTemplates struct {
	List   *template.Template
	Create *template.Template
}

// Divisions is a mvc controller that handles all divisions related views.
type Divisions struct {
	log logger.Logger

	divisions *divisions.Service

	templates DivisionsTemplates
}

// NewDivisions is a constructor for divisions controller.
func NewDivisions(log logger.Logger, divisions *divisions.Service, templates DivisionsTemplates) *Divisions {
	divisionsController := &Divisions{
		log:       log,
		divisions: divisions,
		templates: templates,
	}

	return divisionsController
}

// List is an endpoint that will provide a web page with all divisions.
func (controller *Divisions) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	divisions, err := controller.divisions.List(ctx)
	if err != nil {
		controller.log.Error("could not get divisions list", ErrDivisions.Wrap(err))
		http.Error(w, "could not get divisions list", http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, divisions)
	if err != nil {
		controller.log.Error("can not execute list divisions template", ErrDivisions.Wrap(err))
		http.Error(w, "can not execute list divisions template", http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that will create a new division.
func (controller *Divisions) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Create.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute create divisions template", ErrDivisions.Wrap(err))
			http.Error(w, "could not execute create divisions template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "name is empty", http.StatusBadRequest)
			return
		}

		err = controller.divisions.Create(ctx, name)
		if err != nil {
			controller.log.Error("could not create division", ErrDivisions.Wrap(err))
			http.Error(w, "could not create division", http.StatusInternalServerError)
			return
		}
		Redirect(w, r, "/divisions", http.MethodGet)
	}
}

// Delete is an endpoint that will delete a division by ID.
func (controller *Divisions) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, "could not parse division id", http.StatusBadRequest)
		return
	}

	err = controller.divisions.Delete(ctx, id)
	if err != nil {
		if divisions.ErrNoDivision.Has(err) {
			http.Error(w, "division does not exist", http.StatusNotFound)
			return
		}
		controller.log.Error("could not delete division", ErrDivisions.Wrap(err))
		http.Error(w, "could not delete division", http.StatusInternalServerError)
		return
	}
	Redirect(w, r, "/divisions", http.MethodGet)
}
