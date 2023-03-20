// Copyright (C) 2021 - 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package games

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
	Create(ctx context.Context, lootBox Game) error
	// List returns all games.
	List(ctx context.Context) ([]Game, error)
	// Get returns game by match id.
	Get(ctx context.Context, gameID uuid.UUID) (Game, error)
	// Delete deletes game by match id from db.
	Delete(ctx context.Context, gameID uuid.UUID) error
}

// Game defines game data.
type Game struct {
	MatchID uuid.UUID
	GameInfo
}

// GameInfo defines game parameters.
type GameInfo struct {
}

// Config defines configuration for game.
type Config struct {
}
