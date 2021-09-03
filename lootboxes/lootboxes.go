// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"

	"github.com/google/uuid"
)

// DB is exposing access to lootboxes db.
//
// architecture: DB
type DB interface {
	// Create creates lootbox of user in db.
	Create(ctx context.Context, lootBox LootBox) error
	// Delete deletes opened lootbox by user in db.
	Delete(ctx context.Context, lootboxID uuid.UUID) error
}

// LootBox defines lootbox.
type LootBox struct {
	UserID    uuid.UUID `json:"-"`
	LootBoxID uuid.UUID `json:"id"`
	Type      Type      `json:"name"`
}

// Type defines type of LootBox.
type Type string

const (
	// RegularBox defines regular box type.
	RegularBox Type = "Regular Box"
	// UDReleaseCelebrationBox defines UD Release Celebration Box type.
	UDReleaseCelebrationBox Type = "UD Release Celebration Box"
)

// RegularBoxConfig defines configuration for Regular Box.
type RegularBoxConfig struct {
	Cost     int `json:"cost"`
	CardsNum int `json:"cardsNum"`
	Wood     int `json:"wood"`
	Silver   int `json:"silver"`
	Gold     int `json:"gold"`
	Diamond  int `json:"diamond"`
}

// UDReleaseCelebrationBoxConfig defines configuration for UD Release Celebration Box.
type UDReleaseCelebrationBoxConfig struct {
	Cost     int `json:"cost"`
	CardsNum int `json:"cardsNum"`
	Wood     int `json:"wood"`
	Silver   int `json:"silver"`
	Gold     int `json:"gold"`
	Diamond  int `json:"diamond"`
}

// Config defines configuration for lootboxes.
type Config struct {
	RegularBoxConfig              `json:"regular"`
	UDReleaseCelebrationBoxConfig `json:"UDReleaseCelebration"`
}
