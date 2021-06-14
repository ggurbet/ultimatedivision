// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package ultimatedivision

import (
	"context"

	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// Users provides access to users db.
	Users() users.DB

	// Close closes underlying db connection.
	Close() error

	// CreateSchema create tables.
	CreateSchema(ctx context.Context) (err error)
}

// Config is the global configuration for ultimatedivision.
type Config struct {
}

// Peer is the representation of a ultimatedivision.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB

	// exposes users related logic.
	Users struct {
		Service *users.Service
	}
}

// NewPeer is a constructor for ultimatedivision Peer.
func New(logger logger.Logger, config Config, db DB, ctx context.Context) (*Peer, error) {
	peer := &Peer{
		Log:      logger,
		Database: db,
	}

	{ // users setup
		peer.Users.Service = users.NewService(
			peer.Database.Users(),
		)
	}

	return peer, nil
}
