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

	if err = controller.templates.List.Execute(w, admins); err != nil {
		controller.log.Error("can not execute list admins template", ErrAdmins.Wrap(err))
		http.Error(w, "can not execute list admins template", http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that creates new admin.
func (controller *Admins) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := controller.templates.Create.Execute(w, nil); err != nil {
			controller.log.Error("could not execute create admins template", ErrAdmins.Wrap(err))
			http.Error(w, "could not execute create admins template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		ctx := r.Context()

		if err := r.ParseForm(); err != nil {
			http.Error(w, "could not parse admin create form", http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "email input is empty", http.StatusBadRequest)
			return
		}

		password := r.FormValue("password")
		if password == "" {
			http.Error(w, "password input is empty", http.StatusBadRequest)
			return
		}

		if err := controller.admins.Create(ctx, email, []byte(password)); err != nil {
			controller.log.Error("could not create admin", ErrAdmins.Wrap(err))
			http.Error(w, "could not create admin", http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "", http.MethodGet)
	}
}

// Update is an endpoint that updates admin.
func (controller *Admins) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := uuid.Parse(params["id"])
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

		if err = controller.templates.Update.Execute(w, admin); err != nil {
			controller.log.Error("could not execute update admins template", ErrAdmins.Wrap(err))
			http.Error(w, "could not execute update admins template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not parse admin create form", http.StatusBadRequest)
			return
		}

		password := r.FormValue("password")
		if password == "" {
			http.Error(w, "password input is empty", http.StatusBadRequest)
			return
		}
		passwordAgain := r.FormValue("password-again")
		if passwordAgain == "" {
			http.Error(w, "password-again input is empty", http.StatusBadRequest)
			return
		}

		if password != passwordAgain {
			http.Error(w, "password mismatch", http.StatusBadRequest)
			return
		}

		if err = controller.admins.Update(ctx, id, []byte(password)); err != nil {
			controller.log.Error("could not update admin", ErrAdmins.Wrap(err))
			http.Error(w, "could not update admin", http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/admins", http.MethodGet)
	}
}
