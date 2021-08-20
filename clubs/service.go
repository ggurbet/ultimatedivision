// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/auth"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
)

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// Service is handling users related logic.
//
// architecture: Service
type Service struct {
	clubs DB
	users users.Service
}

// NewService is a constructor for clubs service.
func NewService(clubs DB) *Service {
	return &Service{
		clubs: clubs,
	}
}

// Create creates clubs.
func (service *Service) Create(ctx context.Context) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return userauth.ErrUnauthenticated.Wrap(err)
	}

	nickname, err := service.users.GetNickNameByID(ctx, claims.ID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newClub := Club{
		ID:        uuid.New(),
		OwnerID:   claims.ID,
		Name:      nickname,
		CreatedAt: time.Now().UTC(),
	}

	return ErrClubs.Wrap(service.clubs.Create(ctx, newClub))
}

// CreateSquad creates new squad for club.
func (service *Service) CreateSquad(ctx context.Context, clubID uuid.UUID) error {
	newSquad := Squad{
		ID:     uuid.New(),
		ClubID: clubID,
	}

	return ErrClubs.Wrap(service.clubs.CreateSquad(ctx, newSquad))
}

// Add add new card to the squad of the club.
func (service *Service) Add(ctx context.Context, newSquadCard SquadCard) error {
	return ErrClubs.Wrap(service.clubs.AddSquadCard(ctx, newSquadCard))
}

// Delete deletes card from squad.
func (service *Service) Delete(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID) error {
	return ErrClubs.Wrap(service.clubs.DeleteSquadCard(ctx, squadID, cardID))
}

// UpdateSquad updates tactic and formation of the squad.
func (service *Service) UpdateSquad(ctx context.Context, updatedSquad Squad) error {
	return ErrClubs.Wrap(service.clubs.UpdateTacticFormationCaptain(ctx, updatedSquad))
}

// UpdateCardPosition updates position of card in the squad.
func (service *Service) UpdateCardPosition(ctx context.Context, squadCard SquadCard) error {
	return ErrClubs.Wrap(service.clubs.UpdatePosition(ctx, squadCard.SquadID, squadCard.CardID, squadCard.Position))
}

// GetSquad returns all squads from club.
func (service *Service) GetSquad(ctx context.Context, clubID uuid.UUID) (Squad, []SquadCard, error) {
	squad, err := service.clubs.GetSquad(ctx, clubID)
	if err != nil {
		return Squad{}, nil, ErrClubs.Wrap(err)
	}

	squadCards, err := service.clubs.ListSquadCards(ctx, squad.ID)
	if err != nil {
		return Squad{}, nil, ErrClubs.Wrap(err)
	}

	return squad, squadCards, nil
}

// Get returns user club.
func (service *Service) Get(ctx context.Context) (Club, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return Club{}, userauth.ErrUnauthenticated.Wrap(err)
	}

	userID := claims.ID

	club, err := service.clubs.GetByUserID(ctx, userID)
	return club, ErrClubs.Wrap(err)
}
