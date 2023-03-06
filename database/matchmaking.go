// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"sync"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/matchmaking"
)

// ensures that matchmakingDB implements matchmaking.DB.
var _ matchmaking.DB = (*matchmakingDB)(nil)

// ErrMatchmaking is an error class for matchmaking errors.
var ErrMatchmaking = errs.Class("matchmaking db error")

// matchmakingDB provides access to matchmaking db.
//
// architecture: Database
type matchmakingDB struct {
	lock sync.Mutex
	db   *DBPlayers
}

// Create creates new player by user id.
func (matchmakingDB *matchmakingDB) Create(player matchmaking.Player) error {
	matchmakingDB.lock.Lock()
	defer matchmakingDB.lock.Unlock()

	matchmakingDB.db.players[player.UserID] = player
	return nil
}

// List returns all players.
func (matchmakingDB *matchmakingDB) List() map[uuid.UUID]matchmaking.Player {
	matchmakingDB.lock.Lock()
	allPlayers := matchmakingDB.db.players
	matchmakingDB.lock.Unlock()

	return allPlayers
}

// Get gets player by user id.
func (matchmakingDB *matchmakingDB) Get(userID uuid.UUID) (matchmaking.Player, error) {
	player, ok := matchmakingDB.db.players[userID]
	if !ok {
		return matchmaking.Player{}, matchmaking.ErrNoPlayer.New("no player by user")
	}

	return player, nil
}

// Delete deletes player by user id.
func (matchmakingDB *matchmakingDB) Delete(userID uuid.UUID) error {
	_, err := matchmakingDB.Get(userID)
	if err != nil {
		return ErrMatchmaking.Wrap(err)
	}

	matchmakingDB.lock.Lock()
	defer matchmakingDB.lock.Unlock()

	delete(matchmakingDB.db.players, userID)

	return nil
}
