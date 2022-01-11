// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"
	"sort"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ErrNoLootBox indicates that loot box does not exist.
var ErrNoLootBox = errs.Class("loot box does not exist")

// DB is exposing access to lootboxes db.
//
// architecture: DB
type DB interface {
	// Create creates lootbox of user in db.
	Create(ctx context.Context, lootBox LootBox) error
	// Delete deletes opened lootbox by user in db.
	Delete(ctx context.Context, lootboxID uuid.UUID) error
	// List returns all loot boxes.
	List(ctx context.Context) ([]LootBox, error)
	// Get returns lootbox by user id.
	Get(ctx context.Context, lootboxID uuid.UUID) (LootBox, error)
}

// LootBox defines lootbox.
type LootBox struct {
	UserID    uuid.UUID `json:"-"`
	LootBoxID uuid.UUID `json:"id"`
	Type      Type      `json:"type"`
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

// sortLootBoxCards sorts cards returned from loot box.
func sortLootBoxCards(cards []cards.Card) {
	sort.Slice(cards, func(i, j int) bool {
		sortByQuality := cards[i].Quality.GetValueOfQuality() > cards[j].Quality.GetValueOfQuality()

		if cards[i].Quality.GetValueOfQuality() != cards[j].Quality.GetValueOfQuality() {
			return sortByQuality
		}

		parametersOfCard1 := cards[i].Tactics + cards[i].Physique + cards[i].Technique + cards[i].Offence + cards[i].Defence + cards[i].Goalkeeping
		parametersOfCard2 := cards[j].Tactics + cards[j].Physique + cards[j].Technique + cards[j].Offence + cards[j].Defence + cards[j].Goalkeeping

		return parametersOfCard1 > parametersOfCard2
	})
}
