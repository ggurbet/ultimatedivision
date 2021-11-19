// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/admin/admins"
	"ultimatedivision/internal/logger"
)

var (
	// ErrAdmins is an internal error type for admins controller.
	ErrAdmins = errs.Class("admins controller error")
)

// AdminTemplates holds all admins related templates.
type AdminTemplates struct {
	List   *template.Template
	Create *template.Template
	Update *template.Template
}

// Admins is a mvc controller that handles all admins related views.
type Admins struct {
	log logger.Logger

	admins *admins.Service

	templates AdminTemplates
}

// NewAdmins is a constructor for admins controller.
func NewAdmins(log logger.Logger, admins *admins.Service, templates AdminTemplates) *Admins {
	adminsController := &Admins{
		log:       log,
		admins:    admins,
		templates: templates,
	}

	return adminsController
}

// List is an endpoint that will provide a web page with all admins.
func (controller *Admins) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	admins, err := controller.admins.List(ctx)
	if err != nil {
		controller.log.Error("could not get admins list", ErrAdmins.Wrap(err))
		http.Error(w, "could not get admins list", http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, admins)
	if err != nil {
		controller.log.Error("can not execute list admins template", ErrAdmins.Wrap(err))
		http.Error(w, "can not execute list admins template", http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that creates new admin.
func (controller *Admins) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Create.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute create admins template", ErrAdmins.Wrap(err))
			http.Error(w, "could not execute create admins template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		ctx := r.Context()

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "could not parse admin create form", http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		if email == "" || password == "" {
			http.Error(w, "email or password input is empty", http.StatusBadRequest)
			return
		}

		err = controller.admins.Create(ctx, email, []byte(password))
		if err != nil {
			controller.log.Error("could not create admin", ErrAdmins.Wrap(err))
			http.Error(w, "could not create admin", http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/admins", http.MethodGet)
	}
}

// Update is an endpoint that updates admin.
func (controller *Admins) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	idParam := params["id"]

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		admin, err := controller.admins.Get(ctx, id)
		if err != nil {
			controller.log.Error("could not get admins list", ErrAdmins.Wrap(err))

			if admins.ErrNoAdmin.Has(err) {
				http.Error(w, "no admins with such id", http.StatusNotFound)
				return
			}

			http.Error(w, "could not get admins list", http.StatusInternalServerError)
			return
		}

		err = controller.templates.Update.Execute(w, admin)
		if err != nil {
			controller.log.Error("could not execute update admins template", ErrAdmins.Wrap(err))
			http.Error(w, "could not execute update admins template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err = r.ParseForm()
		if err != nil {
			http.Error(w, "could not parse admin create form", http.StatusBadRequest)
			return
		}

		password := r.FormValue("password")
		if password == "" {
			http.Error(w, "empty field", http.StatusBadRequest)
			return
		}

		err = controller.admins.Update(ctx, id, []byte(password))
		if err != nil {
			controller.log.Error("could not update admin", ErrAdmins.Wrap(err))
			http.Error(w, "could not update admin", http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/admins", http.MethodGet)
	}
}
