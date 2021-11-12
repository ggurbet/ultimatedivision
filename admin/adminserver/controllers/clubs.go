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

	"ultimatedivision/clubs"
	"ultimatedivision/internal/logger"
)

var (
	// ErrClubs is an internal error type for clubs controller.
	ErrClubs = errs.Class("clubs controller error")
)

// ClubsTemplates holds all clubs related templates.
type ClubsTemplates struct {
	List               *template.Template
	ListSquads         *template.Template
	UpdateSquad        *template.Template
	ListSquadCards     *template.Template
	UpdateCardPosition *template.Template
	AddCard            *template.Template
}

// Clubs is a mvc controller that handles all clubs related views.
type Clubs struct {
	log logger.Logger

	clubs *clubs.Service

	templates ClubsTemplates
}

// NewClubs is a constructor for clubs controller.
func NewClubs(log logger.Logger, clubs *clubs.Service, templates ClubsTemplates) *Clubs {
	clubsController := &Clubs{
		log:       log,
		clubs:     clubs,
		templates: templates,
	}

	return clubsController
}

// Create is an endpoint that creates club.
func (controller *Clubs) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := uuid.Parse(params["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = controller.clubs.Create(ctx, id); err != nil {
		controller.log.Error("could not create club", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/clubs/"+id.String(), http.MethodGet)
}

// List is an endpoint that provides a web page with users clubs.
func (controller *Clubs) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	userID, err := uuid.Parse(params["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	allClubs, err := controller.clubs.ListByUserID(ctx, userID)
	if err != nil {
		controller.log.Error("could not get club", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoClub.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err = controller.templates.List.Execute(w, allClubs); err != nil {
		controller.log.Error("could not parse template", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateSquad is an endpoint that creates squad for club.
func (controller *Clubs) CreateSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := uuid.Parse(params["clubId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = controller.clubs.CreateSquad(ctx, id); err != nil {
		controller.log.Error("could not create squad", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/clubs/"+id.String()+"/squad", http.MethodGet)
}

// GetSquadByClubID is an endpoint that provides a web page with squad.
func (controller *Clubs) GetSquadByClubID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	clubID, err := uuid.Parse(params["clubId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	squad, err := controller.clubs.GetSquadByClubID(ctx, clubID)
	if err != nil {
		controller.log.Error("could not get squad", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoSquad.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err = controller.templates.ListSquads.Execute(w, squad); err != nil {
		controller.log.Error("could not parse template", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateStatus is an endpoint that updates club status.
func (controller *Clubs) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	userID, err := uuid.Parse(params["userId"])
	if err != nil {
		http.Error(w, "could not parse user id", http.StatusBadRequest)
		return
	}

	clubID, err := uuid.Parse(params["clubId"])
	if err != nil {
		http.Error(w, "could not parse club id", http.StatusBadRequest)
		return
	}

	if err = controller.clubs.UpdateStatus(ctx, userID, clubID, clubs.StatusActive); err != nil {
		controller.log.Error("could not change status", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoClub.Has(err):
			http.Error(w, "club does not exist", http.StatusNotFound)
		default:
			http.Error(w, "could not change status", http.StatusInternalServerError)
		}
		return
	}

	Redirect(w, r, "/clubs/"+userID.String(), http.MethodGet)
}

// UpdateSquad is an endpoint that updates squad.
func (controller *Clubs) UpdateSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clubID, err := uuid.Parse(params["clubId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		squad, err := controller.clubs.GetSquadByClubID(ctx, clubID)
		if err != nil {
			controller.log.Error("could not get squad", ErrClubs.Wrap(err))
			switch {
			case clubs.ErrNoSquad.Has(err):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err = controller.templates.UpdateSquad.Execute(w, squad); err != nil {
			controller.log.Error("could not parse template", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not parse form", http.StatusBadRequest)
			return
		}

		newTactic, err := strconv.Atoi(r.FormValue("tactic"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newCaptainID, err := uuid.Parse(r.FormValue("captainId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = controller.clubs.UpdateSquad(ctx, squadID, clubs.Tactic(newTactic), newCaptainID); err != nil {
			controller.log.Error("could not update squad", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/clubs/"+clubID.String()+"/squad", http.MethodGet)
	}
}

// ListSquadCards is an endpoint that list cards of the squad.
func (controller *Clubs) ListSquadCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	squadCards, err := controller.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		controller.log.Error("could not get squad cards", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoSquad.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err = controller.templates.ListSquadCards.Execute(w, squadCards); err != nil {
		controller.log.Error("could not parse template", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Add is an endpoint that adds new card to the squad.
func (controller *Clubs) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if err = controller.templates.AddCard.Execute(w, squadID); err != nil {
			controller.log.Error("could not parse template", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not parse form", http.StatusBadRequest)
			return
		}

		position, err := strconv.Atoi(r.FormValue("position"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cardID, err := uuid.Parse(r.FormValue("cardId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		squadCard := clubs.SquadCard{
			Position: clubs.Position(position),
			CardID:   cardID,
		}

		if err = controller.clubs.AddSquadCard(ctx, squadID, squadCard); err != nil {
			controller.log.Error("could not add card to the squad", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/clubs/squad/"+squadID.String(), http.MethodGet)
	}
}

// UpdateCardPosition is an endpoint that updates card position in the squad.
func (controller *Clubs) UpdateCardPosition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardID, err := uuid.Parse(params["cardId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		updateCardPosition := struct {
			CardID  uuid.UUID
			SquadID uuid.UUID
		}{
			CardID:  cardID,
			SquadID: squadID,
		}

		if err = controller.templates.UpdateCardPosition.Execute(w, updateCardPosition); err != nil {
			controller.log.Error("could not parse template", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not parse form", http.StatusBadRequest)
			return
		}

		newPosition, err := strconv.Atoi(r.FormValue("position"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = controller.clubs.UpdateCardPosition(ctx, squadID, cardID, clubs.Position(newPosition)); err != nil {
			controller.log.Error("could not update card position", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "/clubs/squad/"+squadID.String(), http.MethodGet)
	}
}

// DeleteCard is an endpoint that deletes card from squad.
func (controller *Clubs) DeleteCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cardID, err := uuid.Parse(params["cardId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = controller.clubs.Delete(ctx, squadID, cardID); err != nil {
		controller.log.Error("could not delete card", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/clubs/squad/"+squadID.String(), http.MethodGet)
}
