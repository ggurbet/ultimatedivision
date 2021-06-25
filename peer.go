// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package ultimatedivision

import (
	"context"
	"errors"
	"net"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/admin/admins"
	"ultimatedivision/admin/adminserver"
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
	CreateSchema(ctx context.Context) error
}

// Config is the global configuration for ultimatedivision.
type Config struct {
	Admin adminserver.Config
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

	// Admin web server server with web UI.
	Admin struct {
		Listener net.Listener
		Endpoint *adminserver.Server
	}
}

// New is a constructor for ultimatedivision.Peer.
func New(logger logger.Logger, config Config, db DB) (peer *Peer, err error) {
	peer = &Peer{
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

	{ // admin setup
		peer.Admin.Listener, err = net.Listen("tcp", config.Admin.Address)
		if err != nil {
			return nil, err
		}

		peer.Admin.Endpoint, err = adminserver.NewServer(
			config.Admin,
			logger,
			peer.Admin.Listener,
			peer.Admins.Service,
		)
		if err != nil {
			return nil, err
		}
	}

	return peer, nil
}

// Run runs ultimatedivision.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	// start ultimatedivision servers as a separate goroutines.
	group.Go(func() error {
		return ignoreCancel(peer.Admin.Endpoint.Run(ctx))
	})
	// group.Go(func() error {
	//     return ignoreCancel(peer.Console.Endpoint.Run(ctx))
	// })

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group

	errlist.Add(peer.Admin.Endpoint.Close())
	// errlist.Add(peer.Console.Endpoint.Close())

	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
