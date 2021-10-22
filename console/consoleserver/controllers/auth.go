// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
)

// AuthError is a internal error for auth controller.
var AuthError = errs.Class("auth controller error")

// AuthTemplates holds all auth related templates.
type AuthTemplates struct {
	Login          *template.Template
	Register       *template.Template
	ChangePassword *template.Template
}

// Auth is an authentication controller that exposes users authentication functionality.
type Auth struct {
	log      logger.Logger
	userAuth *userauth.Service
	cookie   *auth.CookieAuth

	templates *AuthTemplates
}

// NewAuth returns new instance of Auth.
func NewAuth(log logger.Logger, userAuth *userauth.Service, authCookie *auth.CookieAuth, templates *AuthTemplates) *Auth {
	return &Auth{
		log:       log,
		userAuth:  userAuth,
		cookie:    authCookie,
		templates: templates,
	}
}

// Register creates a new user account.
func (auth *Auth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request users.CreateUserFields

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(err))
		return
	}

	err = auth.userAuth.Register(ctx, request.Email, request.Password, request.NickName, request.FirstName, request.LastName)
	if err != nil {
		auth.log.Error("Unable to register new user", AuthError.Wrap(err))
		auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
		return
	}
}

// ConfirmEmail confirms the email of the user based on the received token.
func (auth *Auth) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	params := mux.Vars(r)
	token := params["token"]
	if token == "" {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(errors.New("Unable to confirm address. Missing token")))
		return
	}
	err := auth.userAuth.ConfirmUserEmail(ctx, token)
	if userauth.ErrPermission.Has(err) {
		auth.log.Error("Permission denied", AuthError.Wrap(err))
		auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(errors.New("permission denied")))
		return
	}
	if err != nil {
		auth.log.Error("Unable to confirm address", AuthError.Wrap(err))
		auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
		return
	}
}

// Login is an endpoint to authorize user and set auth cookie in browser.
func (auth *Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request users.CreateUserFields
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(err))
		return
	}

	if request.Email == "" || request.Password == "" {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(err))
		return
	}

	authToken, err := auth.userAuth.LoginToken(ctx, request.Email, request.Password)
	if err != nil {
		auth.log.Error("could not get auth token", AuthError.Wrap(err))
		switch {
		case users.ErrNoUser.Has(err):
			auth.serveError(w, http.StatusNotFound, AuthError.Wrap(err))
		case userauth.ErrUnauthenticated.Has(err):
			auth.serveError(w, http.StatusUnauthorized, AuthError.Wrap(err))
		default:
			auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
		}

		return
	}

	auth.cookie.SetTokenCookie(w, authToken)
}

// Logout is an endpoint to log out and remove auth cookie from browser.
func (auth *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	auth.cookie.RemoveTokenCookie(w)
}

// RegisterTemplateHandler is web app http handler function.
func (auth *Auth) RegisterTemplateHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "same-origin")

	if err := auth.templates.Register.Execute(w, nil); err != nil {
		auth.log.Error("index template could not be executed", AuthError.Wrap(err))
		return
	}
}

// ChangePassword change users password.
func (auth *Auth) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request users.Password

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(err))
		return
	}

	err = auth.userAuth.ChangePassword(ctx, request.Password, request.NewPassword)
	if err != nil {
		auth.log.Error("Unable to change password", AuthError.Wrap(err))
		switch {
		case users.ErrNoUser.Has(err):
			auth.serveError(w, http.StatusNotFound, AuthError.Wrap(err))
		case userauth.ErrUnauthenticated.Has(err):
			auth.serveError(w, http.StatusUnauthorized, AuthError.Wrap(err))
		default:
			auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
		}

		return
	}

	if err = json.NewEncoder(w).Encode("success"); err != nil {
		auth.log.Error("failed to write json response", ErrUsers.Wrap(err))
		return
	}
}

// ResetPasswordSendEmail send email with token about reset users password.
func (auth *Auth) ResetPasswordSendEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	params := mux.Vars(r)
	userEmail := params["email"]
	if userEmail == "" {
		auth.serveError(w, http.StatusBadRequest, AuthError.New("Unable to reset password. Missing email"))
		return
	}

	err = auth.userAuth.ResetPasswordSendEmail(ctx, userEmail)
	if err != nil {
		auth.log.Error("Unable to change password", AuthError.Wrap(err))
		switch {
		case users.ErrNoUser.Has(err):
			auth.serveError(w, http.StatusNotFound, AuthError.Wrap(err))
		default:
			auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
		}

		return
	}

	if err = json.NewEncoder(w).Encode("success"); err != nil {
		auth.log.Error("failed to write json response", ErrUsers.Wrap(err))
		return
	}
}

// CheckAuthToken checks auth token and sets auth cookie in browser for change users password.
func (auth *Auth) CheckAuthToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	params := mux.Vars(r)
	preAuthToken := params["token"]
	if preAuthToken == "" {
		auth.serveError(w, http.StatusBadRequest, AuthError.New("Unable to reset password. Missing token"))
		return
	}

	err := auth.userAuth.CheckAuthToken(ctx, preAuthToken)
	if err != nil {
		auth.log.Error("Unable to check auth token", AuthError.Wrap(err))
		switch {
		case users.ErrNoUser.Has(err):
			auth.serveError(w, http.StatusNotFound, AuthError.Wrap(err))
		case userauth.ErrUnauthenticated.Has(err):
			auth.serveError(w, http.StatusUnauthorized, AuthError.Wrap(err))
		default:
			auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
		}

		return
	}

	auth.cookie.SetTokenCookie(w, preAuthToken)
}

// ResetPassword reset password and change users password.
func (auth *Auth) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request users.Password

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(err))
		return
	}

	err = auth.userAuth.ResetPassword(ctx, request.NewPassword)
	if err != nil {
		auth.log.Error("Unable to recovery password", AuthError.Wrap(err))
		switch {
		case users.ErrNoUser.Has(err):
			auth.serveError(w, http.StatusNotFound, AuthError.Wrap(err))
		case userauth.ErrUnauthenticated.Has(err):
			auth.serveError(w, http.StatusUnauthorized, AuthError.Wrap(err))
		default:
			auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
		}

		return
	}

	if err = json.NewEncoder(w).Encode("success"); err != nil {
		auth.log.Error("failed to write json response", ErrUsers.Wrap(err))
		return
	}
}

// LoginTemplateHandler is web app http handler function.
func (auth *Auth) LoginTemplateHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "same-origin")

	if err := auth.templates.Login.Execute(w, nil); err != nil {
		auth.log.Error("index template could not be executed", AuthError.Wrap(err))
		return
	}
}

// ChangePasswordTemplateHandler is web app http handler function.
func (auth *Auth) ChangePasswordTemplateHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "same-origin")

	if err := auth.templates.ChangePassword.Execute(w, nil); err != nil {
		auth.log.Error("index template could not be executed", AuthError.Wrap(err))
		return
	}
}

// serveError replies to request with specific code and error.
func (auth *Auth) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		auth.log.Error("failed to write json error response", AuthError.Wrap(err))
	}
}
