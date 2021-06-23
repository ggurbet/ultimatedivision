// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package ultimatedivision

import (
	"context"

	"ultimatedivision/admin/admins"
	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	//Admins provides access to admins db.
	Admins() admins.DB
	// Users provides access to users db.
	Users() users.DB

	// Cards provides access to cards db.
	Cards() cards.DB

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

	// exposes admins relates logic.
	Admins struct {
		Service *admins.Service
	}

	// exposes users related logic.
	Users struct {
		Service *users.Service
	}

	// exposes cards related logic.
	Cards struct {
		Service *cards.Service
	}
}

// New is a constructor for ultimatedivision Peer.
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

	{ // admins setup
		peer.Admins.Service = admins.NewService(
			peer.Database.Admins(),
		)
	}

	{ // cards setup
		peer.Cards.Service = cards.NewService(
			peer.Database.Cards(),
		)
	}

	return peer, nil
}
