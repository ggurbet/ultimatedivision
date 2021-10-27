// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
)

var (
	// ErrClubs is an internal error type for clubs controller.
	ErrClubs = errs.Class("clubs controller error")
)

const (
	// minimumPositionValue defines the minimal value of the position in the squad.
	minimumPositionValue clubs.Position = 0
	// maximumPositionValue defines the maximal value of the position in the squad.
	maximumPositionValue clubs.Position = 10
)

// Clubs is a mvc controller that handles all clubs related views.
type Clubs struct {
	log logger.Logger

	clubs *clubs.Service
}

// NewClubs is a constructor for clubs controller.
func NewClubs(log logger.Logger, clubs *clubs.Service) *Clubs {
	clubsController := &Clubs{
		log:   log,
		clubs: clubs,
	}

	return clubsController
}

// UpdateRequest is struct for update body payload.
type UpdateRequest struct {
	ID        uuid.UUID       `json:"squadId"`
	Tactic    clubs.Tactic    `json:"tactic"`
	Captain   uuid.UUID       `json:"captain"`
	Formation clubs.Formation `json:"formation"`
}

// ClubResponse is a struct for response clubs, squad and squadCards.
type ClubResponse struct {
	Clubs      clubs.Club        `json:"clubs"`
	Squad      clubs.Squad       `json:"squad"`
	SquadCards []clubs.SquadCard `json:"squadCards"`
}

// Create is an endpoint that creates new club.
func (controller *Clubs) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrClubs.Wrap(err))
		return
	}

	id, err := controller.clubs.Create(ctx, claims.UserID)
	if err != nil {
		controller.log.Error("could not create club", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(&id); err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// CreateSquad is an endpoint that creates new squad for club.
func (controller *Clubs) CreateSquad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := uuid.Parse(params["clubId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	squadID, err := controller.clubs.CreateSquad(ctx, id)
	if err != nil {
		controller.log.Error("could not create squad", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(&squadID); err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// Get is an endpoint that returns club, squad and squad cards by user id.
func (controller *Clubs) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrClubs.Wrap(err))
		return
	}

	club, err := controller.clubs.Get(ctx, claims.UserID)
	if err != nil {
		controller.log.Error("could not get user club", ErrClubs.Wrap(err))

		if clubs.ErrNoClub.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrClubs.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}

	squad, err := controller.clubs.GetSquad(ctx, club.ID)
	if err != nil {
		controller.log.Error("could not get squad", ErrClubs.Wrap(err))

		if clubs.ErrNoSquad.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrClubs.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}

	squadCards, err := controller.clubs.ListSquadCards(ctx, squad.ID)
	if err != nil {
		controller.log.Error("could not get squad cards", ErrClubs.Wrap(err))

		if clubs.ErrNoSquad.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrClubs.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}

	userTeam := ClubResponse{
		Clubs:      club,
		Squad:      squad,
		SquadCards: squadCards,
	}

	if err = json.NewEncoder(w).Encode(userTeam); err != nil {
		controller.log.Error("failed to write json response", ErrClubs.Wrap(err))
		return
	}
}

// UpdatePosition is an endpoint that updates card position in the squad.
func (controller *Clubs) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)

	cardID, err := uuid.Parse(params["cardId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	var squadCard clubs.SquadCard

	if err = json.NewDecoder(r.Body).Decode(&squadCard); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	if squadCard.Position < minimumPositionValue || squadCard.Position > maximumPositionValue {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.New("invalid value of position"))
		return
	}

	err = controller.clubs.UpdateCardPosition(ctx, squadID, cardID, squadCard.Position)
	if err != nil {
		controller.log.Error("could not update card position", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// UpdateSquad is an endpoint that updates squad tactic, capitan and formation.
func (controller *Clubs) UpdateSquad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)
	var updatedSquad clubs.Squad

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&updatedSquad); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	if err = controller.clubs.UpdateSquad(ctx, squadID, updatedSquad.Formation, updatedSquad.Tactic, updatedSquad.CaptainID); err != nil {
		controller.log.Error("could not update squad", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// Add is an endpoint that adds new cards to the squad.
func (controller *Clubs) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)
	var squadCard clubs.SquadCard

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	cardID, err := uuid.Parse(params["cardId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	squadCard.CardID = cardID

	if err = json.NewDecoder(r.Body).Decode(&squadCard); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
	}

	squadCard.CardID = cardID

	if squadCard.Position < minimumPositionValue || squadCard.Position > maximumPositionValue {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.New("invalid value of position"))
		return
	}

	if err = controller.clubs.AddSquadCard(ctx, squadID, squadCard); err != nil {
		controller.log.Error("could not add card to the squad", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// Delete is an endpoint that removes card from squad.
func (controller *Clubs) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)

	cardID, err := uuid.Parse(params["cardId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	if err = controller.clubs.Delete(ctx, squadID, cardID); err != nil {
		controller.log.Error("could not delete card from the squad", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// ChangeFormation is a method that change formation and card position.
func (controller *Clubs) ChangeFormation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)

	newFormationID, err := strconv.Atoi(params["formationID"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	formation := clubs.Formation(newFormationID)

	if !clubs.Formation(newFormationID).IsValid() {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.New("Formation ID is not correct"))
		return
	}

	squadID, err := uuid.Parse(params["squadId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	data, err := controller.clubs.ChangeFormation(ctx, formation, squadID)
	if err != nil {
		controller.log.Error("could not change formation", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(data); err != nil {
		controller.log.Error("failed to write json response", ErrClubs.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Clubs) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrClubs.Wrap(err))
	}
}
