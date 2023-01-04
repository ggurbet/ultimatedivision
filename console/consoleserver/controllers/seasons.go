// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/divisions"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/seasons"
)

var (
	// ErrSeasons is an internal error type for seasons controller.
	ErrSeasons = errs.Class("seasons controller error")
)

// Seasons is a mvc controller that handles all seasons related views.
type Seasons struct {
	log logger.Logger

	seasons *seasons.Service
}

// NewSeasons is a constructor for seasons controller.
func NewSeasons(log logger.Logger, seasons *seasons.Service) *Seasons {
	seasonsController := &Seasons{
		log:     log,
		seasons: seasons,
	}

	return seasonsController
}

// GetCurrentSeasons returns all current seasons.
func (controller *Seasons) GetCurrentSeasons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	seasons, err := controller.seasons.GetCurrentSeasons(ctx)
	if err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrSeasons.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(seasons); err != nil {
		controller.log.Error("failed to write json response", ErrSeasons.Wrap(err))
		return
	}
}

// GetRewardByUserID returns returns user reward.
func (controller *Seasons) GetRewardByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
		return
	}

	reward, err := controller.seasons.GetRewardByUserID(ctx, claims.UserID)
	if err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrSeasons.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(reward); err != nil {
		controller.log.Error("failed to write json response", ErrSeasons.Wrap(err))
		return
	}
}

// GetAllClubsStatistics returns all clubs statistics by division.
func (controller *Seasons) GetAllClubsStatistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)

	divisionName, err := strconv.Atoi(params["divisionName"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrSeasons.Wrap(err))
		return
	}

	division, err := controller.seasons.GetDivision(ctx, divisionName)
	if err != nil {
		switch {
		case divisions.ErrNoDivision.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrSeasons.New("division does not exist"))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrSeasons.Wrap(err))
		}
		return
	}

	clubsStatistics, err := controller.seasons.GetAllClubsStatistics(ctx, division)
	if err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrSeasons.Wrap(err))
		return
	}

	statistic := seasons.SeasonStatistics{
		Division:   division,
		Statistics: clubsStatistics,
	}

	if err := json.NewEncoder(w).Encode(statistic); err != nil {
		controller.log.Error("failed to write json response", ErrSeasons.Wrap(err))
		return
	}
}

// UpdatesClubsToNewDivision updates clubs to new division.
func (controller *Seasons) UpdatesClubsToNewDivision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	err := controller.seasons.UpdateClubsToNewDivision(ctx)
	if err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrSeasons.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode("OK"); err != nil {
		controller.log.Error("failed to write json response", ErrSeasons.Wrap(err))
		return
	}
}

// serveError replies to request with specific code and error.
func (controller *Seasons) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrSeasons.Wrap(err))
	}
}
