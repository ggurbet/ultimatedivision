// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/admin/adminauth"
	"ultimatedivision/admin/admins"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
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
		err = auth.loginTemplate.Execute(w, nil)
		if err != nil {
			auth.log.Error("Could not execute login template", AuthError.Wrap(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err = r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		email := r.Form["email"]
		password := r.Form["password"]
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
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case adminauth.ErrUnauthenticated.Has(err):
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
}
