// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
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
func (service *Service) Create(ctx context.Context, userID uuid.UUID) error {
	newClub := Club{
		ID:        uuid.New(),
		OwnerID:   userID,
		CreatedAt: time.Now().UTC(),
	}

	return service.clubs.Create(ctx, newClub)
}

// CreateSquad creates new squad for club.
func (service *Service) CreateSquad(ctx context.Context, clubID uuid.UUID) error {
	newSquad := Squads{
		ID:     uuid.New(),
		ClubID: clubID,
	}

	return service.clubs.CreateSquad(ctx, newSquad)
}

// Add add new card to the squad of the club.
func (service *Service) Add(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, position Position) error {
	capitanID, err := service.clubs.GetCapitan(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newSquadCard := SquadCards{
		ID:        squadID,
		CardID:    cardID,
		Position:  position,
		CapitanID: capitanID,
	}

	return service.clubs.Add(ctx, newSquadCard)
}

// UpdateSquad updates tactic and formation of the squad.
func (service *Service) UpdateSquad(ctx context.Context, squadID uuid.UUID, tactic Tactic, formation Formation) error {
	updatedSquad := Squads{
		ID:        squadID,
		Formation: formation,
		Tactic:    tactic,
	}

	return service.clubs.UpdateTacticFormation(ctx, updatedSquad)
}

// UpdateCapitan updates capitan in the club.
func (service *Service) UpdateCapitan(ctx context.Context, squadID uuid.UUID, capitanID uuid.UUID) error {
	return service.clubs.UpdateCapitan(ctx, capitanID, squadID)
}

// UpdateCardPosition updates position of card in the squad.
func (service *Service) UpdateCardPosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, position Position) error {
	return service.clubs.UpdatePosition(ctx, squadID, cardID, position)
}

// GetSquad returns all squads from club.
func (service *Service) GetSquad(ctx context.Context, clubID uuid.UUID) (Squads, []SquadCards, error) {
	squad, err := service.clubs.GetSquad(ctx, clubID)
	if err != nil {
		return Squads{}, nil, ErrClubs.Wrap(err)
	}

	squadCards, err := service.clubs.ListSquadCards(ctx, squad.ID)
	if err != nil {
		return Squads{}, nil, ErrClubs.Wrap(err)
	}

	return squad, squadCards, nil
}

// List returns users clubs.
func (service *Service) List(ctx context.Context, userID uuid.UUID) ([]Club, error) {
	return service.clubs.List(ctx, userID)
}
