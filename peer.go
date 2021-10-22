// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package ultimatedivision

import (
	"context"
	"errors"
	"net"
	"net/mail"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/admin/adminauth"
	"ultimatedivision/admin/admins"
	"ultimatedivision/admin/adminserver"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/clubs"
	"ultimatedivision/console/consoleserver"
	"ultimatedivision/console/emails"
	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/auth"
	mail2 "ultimatedivision/pkg/mail"
	"ultimatedivision/queue"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// Admins provides access to admins db.
	Admins() admins.DB

	// Users provides access to users db.
	Users() users.DB

	// Cards provides access to cards db.
	Cards() cards.DB

	// Avatars provides access to avatars db.
	Avatars() avatars.DB

	// Clubs provides access to clubs db.
	Clubs() clubs.DB

	// LootBoxes provides access to lootboxes db.
	LootBoxes() lootboxes.DB

	// Marketplace provides access to marketplace db.
	Marketplace() marketplace.DB

	// Queue provides access to queue db.
	Queue() queue.DB

	// Close closes underlying db connection.
	Close() error

	// CreateSchema create tables.
	CreateSchema(ctx context.Context) error
}

// Config is the global configuration for ultimatedivision.
type Config struct {
	Admins struct {
		Server adminserver.Config `json:"server"`
		Auth   struct {
			TokenAuthSecret string `json:"tokenAuthSecret"`
		} `json:"auth"`
	} `json:"admins"`

	Users struct {
		// Server userserver.Config `json:"server"`
		Auth struct {
			TokenAuthSecret string `json:"tokenAuthSecret"`
		} `json:"auth"`
	} `json:"users"`

	Console struct {
		Server consoleserver.Config `json:"server"`
		Emails emails.Config        `json:"emails"`
	} `json:"console"`

	Cards struct {
		cards.Config
		cards.PercentageQualities `json:"percentageQualities"`
	} `json:"cards"`

	Avatars struct {
		avatars.Config
	} `json:"avatars"`

	LootBoxes struct {
		Config lootboxes.Config `json:"lootBoxes"`
	} `json:"lootBoxes"`

	Marketplace struct {
		marketplace.Config
	} `json:"marketplace"`

	Queue struct {
		queue.Config
	} `json:"queue"`
}

// Peer is the representation of a ultimatedivision.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB
	Sender   mail2.Sender

	// exposes admins relates logic.
	Admins struct {
		Service *admins.Service
		Auth    *adminauth.Service
	}

	// exposes users related logic.
	Users struct {
		Service *users.Service
		Auth    *userauth.Service
	}

	// exposes cards related logic.
	Cards struct {
		Service *cards.Service
	}

	// exposes avatars related logic.
	Avatars struct {
		Service *avatars.Service
	}

	// exposes clubs related logic.
	Clubs struct {
		Service *clubs.Service
	}

	// exposes lootboxes related logic.
	LootBoxes struct {
		Service *lootboxes.Service
	}

	// exposes marketplace related logic.
	Marketplace struct {
		Service            *marketplace.Service
		ExpirationLotChore *marketplace.Chore
	}

	// exposes queue related logic.
	Queue struct {
		Service    *queue.Service
		PlaceChore *queue.Chore
	}

	// Admin web server server with web UI.
	Admin struct {
		Listener net.Listener
		Endpoint *adminserver.Server
	}

	// Console web server server with web UI.
	Console struct {
		Listener     net.Listener
		Endpoint     *consoleserver.Server
		EmailService *emails.Service
	}
}

// New is a constructor for ultimatedivision.Peer.
func New(logger logger.Logger, config Config, db DB) (peer *Peer, err error) {
	peer = &Peer{
		Log:      logger,
		Database: db,
	}

	{ // email setup
		from, err := mail.ParseAddress(config.Console.Emails.From)
		if err != nil {
			logger.Error("email address is not valid", err)
			return nil, err
		}

		sender := mail2.SMTPSender{
			ServerAddress: config.Console.Emails.SMTPServerAddress,
			From:          *from,
			Auth: mail2.LoginAuth{
				Username: config.Console.Emails.PlainLogin,
				Password: config.Console.Emails.PlainPassword,
			},
		}

		peer.Console.EmailService = emails.NewService(
			logger,
			&sender,
			config.Console.Emails,
		)
	}

	{ // users setup
		peer.Users.Service = users.NewService(
			peer.Database.Users(),
		)
		peer.Users.Auth = userauth.NewService(
			peer.Database.Users(),
			auth.TokenSigner{
				Secret: []byte(config.Users.Auth.TokenAuthSecret),
			},
			peer.Console.EmailService,
			logger)
	}

	{ // admins setup
		peer.Admins.Service = admins.NewService(
			peer.Database.Admins(),
		)
		peer.Admins.Auth = adminauth.NewService(
			peer.Database.Admins(),
			auth.TokenSigner{
				Secret: []byte(config.Admins.Auth.TokenAuthSecret),
			},
		)
	}

	{ // Avatars setup
		peer.Avatars.Service = avatars.NewService(
			peer.Database.Avatars(),
			config.Avatars.Config,
		)
	}

	{ // cards setup
		peer.Cards.Service = cards.NewService(
			peer.Database.Cards(),
			config.Cards.Config,
			peer.Avatars.Service,
		)
	}

	{ // clubs setup
		peer.Clubs.Service = clubs.NewService(
			peer.Database.Clubs(),
			peer.Users.Service,
		)
	}

	{ // lootboxes setup
		peer.LootBoxes.Service = lootboxes.NewService(
			config.LootBoxes.Config,
			peer.Database.LootBoxes(),
			peer.Cards.Service,
		)
	}

	{ // marketplace setup
		peer.Marketplace.Service = marketplace.NewService(
			config.Marketplace.Config,
			peer.Database.Marketplace(),
			peer.Users.Service,
			peer.Cards.Service,
		)

		peer.Marketplace.ExpirationLotChore = marketplace.NewChore(
			peer.Log,
			config.Marketplace.Config,
			peer.Database.Marketplace(),
			peer.Users.Service,
			peer.Cards.Service,
		)
	}

	{ // queue setup
		peer.Queue.Service = queue.NewService(
			config.Queue.Config,
			peer.Database.Queue(),
			peer.Users.Service,
		)

		peer.Queue.PlaceChore = queue.NewChore(
			peer.Log,
			config.Queue.Config,
			peer.Queue.Service,
		)
	}

	{ // admin setup
		peer.Admin.Listener, err = net.Listen("tcp", config.Admins.Server.Address)
		if err != nil {
			return nil, err
		}

		peer.Admin.Endpoint, err = adminserver.NewServer(
			config.Admins.Server,
			logger,
			peer.Admin.Listener,
			peer.Admins.Auth,
			peer.Admins.Service,
			peer.Users.Service,
			peer.Cards.Service,
			config.Cards.PercentageQualities,
			peer.Avatars.Service,
			peer.Marketplace.Service,
			peer.LootBoxes.Service,
			peer.Clubs.Service,
			peer.Queue.Service,
		)
		if err != nil {
			return nil, err
		}
	}

	{ // console setup
		peer.Console.Listener, err = net.Listen("tcp", config.Console.Server.Address)
		if err != nil {
			return nil, err
		}

		peer.Console.Endpoint = consoleserver.NewServer(
			config.Console.Server,
			logger,
			peer.Console.Listener,
			peer.Cards.Service,
			peer.LootBoxes.Service,
			peer.Marketplace.Service,
			peer.Clubs.Service,
			peer.Users.Auth,
			peer.Users.Service,
			peer.Queue.Service,
		)
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
	group.Go(func() error {
		return ignoreCancel(peer.Console.Endpoint.Run(ctx))
	})
	group.Go(func() error {
		return ignoreCancel(peer.Marketplace.ExpirationLotChore.Run(ctx))
	})
	group.Go(func() error {
		return ignoreCancel(peer.Queue.PlaceChore.Run(ctx))
	})

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group

	errlist.Add(peer.Admin.Endpoint.Close())
	errlist.Add(peer.Console.Endpoint.Close())

	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
