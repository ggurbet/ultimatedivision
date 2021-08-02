// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrLootBoxes indicates that there was an error in the service.
var ErrLootBoxes = errs.Class("lootboxes service error")

// Service is handling lootboxes related logic.
//
// architecture: Service
type Service struct {
	lootboxes DB
	config    Config
}

// NewService is a constructor for lootboxes service.
func NewService(lootboxes DB, config Config) *Service {
	return &Service{
		lootboxes: lootboxes,
		config:    config,
	}
}

// Create creates LootBox.
func (service *Service) Create(ctx context.Context, userID uuid.UUID, lootBoxID uuid.UUID) error {
	openedLootBox := UserLootBoxes{
		UserID:    userID,
		LootBoxID: lootBoxID,
	}

	err := service.lootboxes.Create(ctx, openedLootBox)

	return ErrLootBoxes.Wrap(err)
}

// Open opens lootbox by user.
func (service *Service) Open(ctx context.Context, userID uuid.UUID, lootBoxID uuid.UUID) error {
	// TODO: call create cards method.
	// TODO: check if user has enough money for lootbox.

	err := service.lootboxes.Delete(ctx, userID, lootBoxID)

	// TODO: return slice of generated cards and error.

	return ErrLootBoxes.Wrap(err)
}
