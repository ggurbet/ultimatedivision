// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"

	"github.com/google/uuid"
)

// Service is handling cards related logic.
//
// architecture: Service
type Service struct {
	cards DB
}

// NewService is a constructor for cards service.
func NewService(cards DB) *Service {
	return &Service{
		cards: cards,
	}
}

// Create add card in DB.
func (service *Service) Create(ctx context.Context, card Card) error {
	return service.cards.Create(ctx, card)
}

// Get returns card from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Card, error) {
	return service.cards.Get(ctx, cardID)
}

// List returns all cards from DB.
func (service *Service) List(ctx context.Context) ([]Card, error) {
	return service.cards.List(ctx)
}

// Delete destroy card in DB.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return service.cards.Delete(ctx, cardID)
}
