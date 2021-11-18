// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"ultimatedivision/admin/admins"

	"github.com/zeebo/errs"

	"ultimatedivision/admin/adminauth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
)

// AuthError is a internal error for auth controller.
var AuthError = errs.Class("auth controller error")

// AuthTemplates holds all auth related templates.
type AuthTemplates struct {
	Login *template.Template
}

// Auth login authentication entity.
type Auth struct {
	log     logger.Logger
	service *adminauth.Service
	cookie  *auth.CookieAuth

	loginTemplate *template.Template
}

// NewAuth returns new instance of Auth.
func NewAuth(log logger.Logger, service *adminauth.Service, authCookie *auth.CookieAuth, templates AuthTemplates) *Auth {
	return &Auth{
		log:           log,
		service:       service,
		cookie:        authCookie,
		loginTemplate: templates.Login,
	}
}

// Login is an endpoint to authorize admin and set auth cookie in browser.
func (auth *Auth) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	switch r.Method {
	case http.MethodGet:
		if err = auth.loginTemplate.Execute(w, nil); err != nil {
			auth.log.Error("could not execute login template", AuthError.Wrap(err))
			http.Error(w, "could not execute login template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not parse login form", http.StatusBadRequest)
			return
		}

		email := r.Form["email"]
		password := r.Form["password"]
		// TODO: change chek
		if len(email) == 0 || len(password) == 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if email[0] == "" || password[0] == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		response, err := auth.service.Token(ctx, email[0], password[0])
		if err != nil {
			auth.log.Error("could not get auth token", AuthError.Wrap(err))
			switch {
			case admins.ErrNoAdmin.Has(err):
				http.Error(w, "could not get auth token", http.StatusNotFound)
			case adminauth.ErrUnauthenticated.Has(err):
				http.Error(w, "could not get auth token", http.StatusUnauthorized)
			default:
				http.Error(w, "could not get auth token", http.StatusInternalServerError)
			}

			return
		}

		auth.cookie.SetTokenCookie(w, response)

		Redirect(w, r, "/admins", "GET")
	}
}

// Logout is an endpoint to log out and remove auth cookie from browser.
func (auth *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	auth.cookie.RemoveTokenCookie(w)
	Redirect(w, r, "/login", "GET")
}
