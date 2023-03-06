// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matchmaking

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"
)

// ErrNoPlayer indicated that player does not exist.
var ErrNoPlayer = errs.Class("player does not exist")

// DB is exposing access to matchmaking database.
//
// architecture: DB
type DB interface {
	// Create creates new player by user id.
	Create(player Player) error
	// List returns all players.
	List() map[uuid.UUID]Player
	// Get gets player by user id.
	Get(userID uuid.UUID) (Player, error)
	// Delete deletes player by user id.
	Delete(userID uuid.UUID) error
}

// Player describes player entity.
type Player struct {
	UserID  uuid.UUID       `json:"userId"`
	SquadID uuid.UUID       `json:"squadId"`
	Conn    *websocket.Conn `json:"conn"`
	Waiting bool            `json:"waiting"`
}

// Match describes match entity.
type Match struct {
	Player1 *Player
	Player2 *Player
}
