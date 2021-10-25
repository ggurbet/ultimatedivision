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

	_, err = controller.clubs.Create(ctx, id)
	if err != nil {
		controller.log.Error("could not create club", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/clubs/"+id.String(), http.MethodGet)
}

// Get is an endpoint that provides a web page with users club.
func (controller *Clubs) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	userID, err := uuid.Parse(params["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	club, err := controller.clubs.Get(ctx, userID)
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

	err = controller.templates.List.Execute(w, club)
	if err != nil {
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

	_, err = controller.clubs.CreateSquad(ctx, id)
	if err != nil {
		controller.log.Error("could not create squad", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/clubs/"+id.String()+"/squad", http.MethodGet)
}

// GetSquad is an endpoint that provides a web page with squad.
func (controller *Clubs) GetSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	clubID, err := uuid.Parse(params["clubId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	squad, err := controller.clubs.GetSquad(ctx, clubID)
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

	err = controller.templates.ListSquads.Execute(w, squad)
	if err != nil {
		controller.log.Error("could not parse template", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
		squad, err := controller.clubs.GetSquad(ctx, clubID)
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

		err = controller.templates.UpdateSquad.Execute(w, squad)
		if err != nil {
			controller.log.Error("could not parse template", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err = r.ParseForm()
		if err != nil {
			http.Error(w, "could not parse form", http.StatusBadRequest)
			return
		}

		newFormation, err := strconv.Atoi(r.FormValue("formation"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
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

		err = controller.clubs.UpdateSquad(ctx, squadID, clubs.Formation(newFormation), clubs.Tactic(newTactic), newCaptainID)
		if err != nil {
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

	squadCards, err := controller.clubs.GetSquadCards(ctx, squadID)
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

	err = controller.templates.ListSquadCards.Execute(w, squadCards)
	if err != nil {
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
		err = controller.templates.AddCard.Execute(w, squadID)
		if err != nil {
			controller.log.Error("could not parse template", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err = r.ParseForm()
		if err != nil {
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

		err = controller.clubs.AddSquadCard(ctx, squadID, clubs.SquadCard{
			Position: clubs.Position(position),
			CardID:   cardID,
		})
		if err != nil {
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

		err = controller.templates.UpdateCardPosition.Execute(w, updateCardPosition)
		if err != nil {
			controller.log.Error("could not parse template", ErrClubs.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err = r.ParseForm()
		if err != nil {
			http.Error(w, "could not parse form", http.StatusBadRequest)
			return
		}

		newPosition, err := strconv.Atoi(r.FormValue("position"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = controller.clubs.UpdateCardPosition(ctx, squadID, cardID, clubs.Position(newPosition))
		if err != nil {
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

	err = controller.clubs.Delete(ctx, squadID, cardID)
	if err != nil {
		controller.log.Error("could not delete card", ErrClubs.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/clubs/squad/"+squadID.String(), http.MethodGet)
}
