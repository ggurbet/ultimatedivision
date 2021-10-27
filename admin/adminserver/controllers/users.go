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

	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
)

var (
	// ErrUsers is an internal error type for users controller.
	ErrUsers = errs.Class("users controller error")
)

// UserTemplates holds all users related templates.
type UserTemplates struct {
	List   *template.Template
	Create *template.Template
	Update *template.Template
}

// Users is a mvc controller that handles all admins related views.
type Users struct {
	log logger.Logger

	users *users.Service

	templates UserTemplates
}

// NewUsers is a constructor for users controller.
func NewUsers(log logger.Logger, users *users.Service, templates UserTemplates) *Users {
	usersController := &Users{
		log:       log,
		users:     users,
		templates: templates,
	}

	return usersController
}

// List is an endpoint that will provide a web page with all users.
func (controller *Users) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := controller.users.List(ctx)
	if err != nil {
		controller.log.Error("could not get users list", ErrUsers.Wrap(err))
		http.Error(w, "could not get users list", http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, users)
	if err != nil {
		controller.log.Error("can not execute list users template", ErrUsers.Wrap(err))
		http.Error(w, "can not execute list users template", http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that will create a new user.
func (controller *Users) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		if err := controller.templates.Create.Execute(w, nil); err != nil {
			controller.log.Error("could not execute create users template", ErrUsers.Wrap(err))
			http.Error(w, "could not execute create users template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "could not get users form", http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "email is empty", http.StatusBadRequest)
			return
		}
		password := r.FormValue("password")
		if password == "" {
			http.Error(w, "password is empty", http.StatusBadRequest)
			return
		}
		nickName := r.FormValue("nickName")
		if nickName == "" {
			http.Error(w, "nick name is empty", http.StatusBadRequest)
			return
		}
		firstName := r.FormValue("firstName")
		if firstName == "" {
			http.Error(w, "first name is empty", http.StatusBadRequest)
			return
		}
		lastName := r.FormValue("lastName")
		if lastName == "" {
			http.Error(w, "last name is empty", http.StatusBadRequest)
			return
		}

		if err := controller.users.Create(ctx, email, password, nickName, firstName, lastName); err != nil {
			controller.log.Error("could not create user", ErrUsers.Wrap(err))
			http.Error(w, "could not create user", http.StatusInternalServerError)
			return
		}
		Redirect(w, r, "/users", http.MethodGet)
	}
}

// Update is an endpoint that will update users status.
func (controller *Users) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	userID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "could not parse user id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := controller.users.Get(ctx, userID)
		if err != nil {
			controller.log.Error("could not get user", ErrUsers.Wrap(err))
			if users.ErrNoUser.Has(err) {
				http.Error(w, "no user with such id", http.StatusNotFound)
				return
			}
			http.Error(w, "could not get user", http.StatusInternalServerError)
			return
		}

		if err = controller.templates.Update.Execute(w, user); err != nil {
			controller.log.Error("could not execute update users template", ErrUsers.Wrap(err))
			http.Error(w, "could not execute update users template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not get users form", http.StatusBadRequest)
			return
		}

		status, err := strconv.Atoi(r.FormValue("status"))
		if err != nil {
			http.Error(w, "invalid status", http.StatusBadRequest)
			return
		}

		if err = controller.users.Update(ctx, users.Status(status), userID); err != nil {
			controller.log.Error("could not update users status", ErrUsers.Wrap(err))
			http.Error(w, "could not update users status", http.StatusInternalServerError)
			return
		}
		Redirect(w, r, "/users", http.MethodGet)
	}
}

// Delete is an endpoint that will delete a user by email.
func (controller *Users) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, "could not parse user id", http.StatusBadRequest)
		return
	}

	if err = controller.users.Delete(ctx, id); err != nil {
		controller.log.Error("could not delete user", ErrUsers.Wrap(err))
		http.Error(w, "could not delete user", http.StatusInternalServerError)
		return
	}
	Redirect(w, r, "/users", http.MethodGet)
}
