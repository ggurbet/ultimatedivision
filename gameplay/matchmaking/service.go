// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matchmaking

import (
	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrMatchmaking indicates that there was an error in the service.
var ErrMatchmaking = errs.Class("matchmaking service error")

// Service is handling matchmaking related logic.
//
// architecture: Service
type Service struct {
	players DB
}

// NewService is a constructor for matchmaking service.
func NewService(players DB) *Service {
	return &Service{
		players: players,
	}
}

// Create creates a player by user.
func (service *Service) Create(userID uuid.UUID) error {
	player := Player{
		UserID: userID,
	}
	return ErrMatchmaking.Wrap(service.players.Create(player))
}

// List returns all players.
func (service *Service) List() map[uuid.UUID]Player {
	return service.players.List()
}

// Get returns player by user.
func (service *Service) Get(userID uuid.UUID) (Player, error) {
	player, err := service.players.Get(userID)
	return player, ErrMatchmaking.Wrap(err)
}

// Delete player by user.
func (service *Service) Delete(id uuid.UUID) error {
	return ErrMatchmaking.Wrap(service.players.Delete(id))
}
