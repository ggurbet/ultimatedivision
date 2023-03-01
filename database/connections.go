// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"

	"ultimatedivision/console/connections"
)

// ensures that connectionDB implements connection.DB.
var _ connections.DB = (*connectionDB)(nil)

// ErrConnections is an error class for connections errors.
var ErrConnections = errs.Class("connections db error")

// connectionDB provides access to connection db.
//
// architecture: Database
type connectionDB struct {
	lock sync.Mutex
	db   *DB
}

// Create creates new connection by user id.
func (connectionDB *connectionDB) Create(userID uuid.UUID, connection *websocket.Conn) error {
	connectionDB.lock.Lock()
	defer connectionDB.lock.Unlock()

	connectionDB.db.connections[userID] = connection

	player, ok := connectionDB.db.players[userID]
	if ok {
		player.Conn = connection
		connectionDB.db.players[userID] = player
	}

	return nil
}

// List returns all connections.
func (connectionDB *connectionDB) List() map[uuid.UUID]*websocket.Conn {
	connectionDB.lock.Lock()
	allConnections := connectionDB.db.connections
	connectionDB.lock.Unlock()

	return allConnections
}

// Get gets connection by user id.
func (connectionDB *connectionDB) Get(userID uuid.UUID) (*websocket.Conn, error) {
	connection, ok := connectionDB.db.connections[userID]
	if !ok {
		return nil, connections.ErrNoConnection.New("no connection by user")
	}

	return connection, nil
}

// Delete deletes connection by user id.
func (connectionDB *connectionDB) Delete(userID uuid.UUID) error {
	_, err := connectionDB.Get(userID)
	if err != nil {
		return ErrConnections.Wrap(err)
	}

	connectionDB.lock.Lock()
	defer connectionDB.lock.Unlock()

	delete(connectionDB.db.connections, userID)

	return nil
}
