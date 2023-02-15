// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package connections

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"
)

// ErrNoConnection indicated that connection does not exist.
var ErrNoConnection = errs.Class("connection does not exist")

// DB is exposing access to connection database.
//
// architecture: DB
type DB interface {
	// Create creates new connection by user id.
	Create(userID uuid.UUID, connection *websocket.Conn) error
	// List returns all connections.
	List() map[uuid.UUID]*websocket.Conn
	// Get gets connection by user id.
	Get(userID uuid.UUID) (*websocket.Conn, error)
	// Delete deletes connection by user id.
	Delete(userID uuid.UUID) error
}
