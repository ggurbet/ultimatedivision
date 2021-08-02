// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoLootBox indicates that lootbox does not exist.
var ErrNoLootBox = errs.Class("lootbox does not exist")

// DB is exposing access to lootboxes db.
//
// architecture: DB
type DB interface {
	// Create creates lootbox of user in db.
	Create(ctx context.Context, lootBox UserLootBoxes) error
	// Delete deletes opened lootbox by user in db.
	Delete(ctx context.Context, userID uuid.UUID, lootBoxID uuid.UUID) error
}

// LootBox defines lootbox.
type LootBox struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Probability defines probability of getting different types of cards.
type Probability float32

// Config defines configuration for LootBox.
type Config struct {
	Cost     int         `json:"cost"`
	CardsNum int         `json:"cardsNum"`
	Wood     Probability `json:"wood"`
	Silver   Probability `json:"silver"`
	Gold     Probability `json:"gold"`
	Diamond  Probability `json:"diamond"`
}

// UserLootBoxes describes lootbox that user has.
type UserLootBoxes struct {
	UserID    uuid.UUID `json:"userId"`
	LootBoxID uuid.UUID `json:"LootBoxId"`
}
