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
	Create(ctx context.Context, lootBox LootBox) error
	// Delete deletes opened lootbox by user in db.
	Delete(ctx context.Context, lootBox LootBox) error
}

// LootBox defines lootbox.
type LootBox struct {
	UserID    uuid.UUID `json:"userId"`
	LootBoxID uuid.UUID `json:"id"`
	Name      Type      `json:"name"`
}

// Type defines type of LootBox.
type Type string

const (
	// RegularBox defines regular box type.
	RegularBox Type = "Regular Box"
	// UDReleaseCelebrationBox defines UD Release Celebration Box type.
	UDReleaseCelebrationBox Type = "UD Release Celebration Box"
)

// Config defines configuration for LootBox.
type Config struct {
	Cost     int `json:"cost"`
	CardsNum int `json:"cardsNum"`
	Wood     int `json:"wood"`
	Silver   int `json:"silver"`
	Gold     int `json:"gold"`
	Diamond  int `json:"diamond"`
}
