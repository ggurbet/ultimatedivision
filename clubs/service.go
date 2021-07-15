// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// Service is handling users related logic.
//
// architecture: Service
type Service struct {
	clubs DB
}

// NewService is a constructor for clubs service.
func NewService(clubs DB) *Service {
	return &Service{
		clubs: clubs,
	}
}

// Create creates clubs.
func (service *Service) Create(ctx context.Context, club Club) error {
	return service.clubs.Create(ctx, club)
}

// Add add new card to the players of the club.
func (service *Service) Add(ctx context.Context, userID uuid.UUID, card cards.Card, position Position) error {
	capitan, err := service.clubs.GetCapitan(ctx, userID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}
	return service.clubs.Add(ctx, userID, card, capitan, position)
}

// UpdateClub updates tactic and formation the club.
func (service *Service) UpdateClub(ctx context.Context, userID uuid.UUID, tactic Tactic, formation Formation) error {
	newClub := Club{
		UserID:    userID,
		Tactic:    tactic,
		Formation: formation,
	}

	return service.clubs.Update(ctx, newClub)
}

// UpdateCapitan updates capitan in the club.
func (service *Service) UpdateCapitan(ctx context.Context, userID uuid.UUID, capitan uuid.UUID) error {
	return service.clubs.UpdateCapitan(ctx, userID, capitan)
}

// List returns all cards from club.
func (service *Service) List(ctx context.Context, userID uuid.UUID) ([]Player, error) {
	return service.clubs.ListCards(ctx, userID)
}

// Get returns club.
func (service *Service) Get(ctx context.Context, userID uuid.UUID) (Club, error) {
	return service.clubs.GetClub(ctx, userID)
}
