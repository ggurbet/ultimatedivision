// Copyright (C) 2021 - 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoGames indicates that game does not exist.
var ErrNoGames = errs.Class("game does not exist")

// DB is exposing access to games db.
//
// architecture: DB.
type DB interface {
	// Create creates game in db.
	Create(ctx context.Context, matchID uuid.UUID, gameInformationInJSON string) error
	// Get returns game information in JSON by match id.
	Get(ctx context.Context, matchID uuid.UUID) (string, error)
	// Update updates game info in the database by match id.
	Update(ctx context.Context, matchID uuid.UUID, gameInformationInJSON string) error
	// Delete deletes game information in JSON.
	Delete(ctx context.Context, gameID uuid.UUID) error
}

// Game defines game data.
type Game struct {
	MatchID  uuid.UUID
	GameInfo []CardIDWithPosition
}

// CardIDWithPosition defines card ID with possible moves cells for game.
type CardIDWithPosition struct {
	CardID   uuid.UUID `json:"cardID"`
	Position int       `json:"position"`
}
