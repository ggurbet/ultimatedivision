// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/avatars"
	"ultimatedivision/internal/logger"
)

var (
	// ErrAvatars is an internal error type for avatars controller.
	ErrAvatars = errs.Class("avatars controller error")
)

// AvatarTemplates holds all avatars related templates.
type AvatarTemplates struct {
	Get *template.Template
}

// Avatars is a mvc controller that handles all avatars related views.
type Avatars struct {
	log       logger.Logger
	avatars   *avatars.Service
	templates AvatarTemplates
}

// NewAvatars is a constructor for avatars controller.
func NewAvatars(log logger.Logger, avatars *avatars.Service, templates AvatarTemplates) *Avatars {
	avatarsController := &Avatars{
		log:       log,
		avatars:   avatars,
		templates: templates,
	}

	return avatarsController
}

// Get is an endpoint that will provide a web page with avatar by card id.
func (controller *Avatars) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	cardID, err := uuid.Parse(params["cardId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	avatar, err := controller.avatars.Get(ctx, cardID)
	if err != nil {
		controller.log.Error("could not get avatar by card id", ErrAvatars.Wrap(err))
		switch {
		case avatars.ErrNoAvatar.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = controller.templates.Get.Execute(w, avatar)
	if err != nil {
		controller.log.Error("can not execute get avatar template", ErrAvatars.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
