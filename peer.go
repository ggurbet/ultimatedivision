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
	"ultimatedivision/cards/nfts"
	"ultimatedivision/cards/waitlist"
	"ultimatedivision/clubs"
	"ultimatedivision/console/consoleserver"
	"ultimatedivision/console/emails"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/auth"
	mail2 "ultimatedivision/pkg/mail"
	"ultimatedivision/queue"
	"ultimatedivision/seasons"
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

	// WaitList provides access to waitlist db.
	WaitList() waitlist.DB

	// NFTs provides access to nfts db.
	NFTs() nfts.DB

	// Clubs provides access to clubs db.
	Clubs() clubs.DB

	// LootBoxes provides access to lootboxes db.
	LootBoxes() lootboxes.DB

	// Marketplace provides access to marketplace db.
	Marketplace() marketplace.DB

	// Matches provides access to matches db.
	Matches() matches.DB

	// Queue provides access to queue db.
	Queue() queue.DB

	// Divisions provides access to divisions db.
	Divisions() divisions.DB

	// Seasons provides access to seasons db.
	Seasons() seasons.DB

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

	NFTs struct {
		nfts.Config
	} `json:"nfts"`

	WaitList struct {
		waitlist.Config
	} `json:"waitList"`

	LootBoxes struct {
		Config lootboxes.Config `json:"lootBoxes"`
	} `json:"lootBoxes"`

	Marketplace struct {
		marketplace.Config
	} `json:"marketplace"`

	Queue struct {
		queue.Config
	} `json:"queue"`

	Divisions struct {
		divisions.Config
	} `json:"divisions"`

	Seasons struct {
		seasons.Config
	} `json:"seasons"`

	Matches struct {
		matches.Config
	} `json:"matches"`
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

	// exposes waitlist related logic.
	WaitList struct {
		Service       *waitlist.Service
		WaitListChore *waitlist.Chore
	}

	// exposes nfts related logic.
	NFTs struct {
		Service  *nfts.Service
		NFTChore *nfts.Chore
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

	// exposes matches related logic.
	Matches struct {
		Service *matches.Service
	}

	// exposes queue related logic.
	Queue struct {
		Service    *queue.Service
		PlaceChore *queue.Chore
	}

	// exposes divisions related logic.
	Divisions struct {
		Service *divisions.Service
	}

	// exposes divisions related logic.
	Seasons struct {
		Service           *seasons.Service
		ExpirationSeasons *seasons.Chore
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

	{ // cards setup
		peer.Cards.Service = cards.NewService(
			peer.Database.Cards(),
			config.Cards.Config,
		)
	}

	{ // avatars setup
		peer.Avatars.Service = avatars.NewService(
			peer.Database.Avatars(),
			config.Avatars.Config,
		)
	}

	{ // nfts setup
		peer.NFTs.Service = nfts.NewService(
			config.NFTs.Config,
			peer.Database.NFTs(),
		)
		peer.NFTs.NFTChore = nfts.NewChore(
			config.NFTs.Config,
			peer.Log,
			peer.NFTs.Service,
			peer.Users.Service,
			peer.Cards.Service,
		)
	}

	{ // waitlist setup
		peer.WaitList.Service = waitlist.NewService(
			config.WaitList.Config,
			peer.Database.WaitList(),
			peer.Cards.Service,
			peer.Avatars.Service,
			peer.Users.Service,
			peer.NFTs.Service,
		)

		peer.WaitList.WaitListChore = waitlist.NewChore(
			config.WaitList.Config,
			peer.Log,
			peer.WaitList.Service,
			peer.NFTs.Service,
			peer.Users.Service,
			peer.Cards.Service,
		)
	}

	{ // clubs setup
		peer.Clubs.Service = clubs.NewService(
			peer.Database.Clubs(),
			peer.Users.Service,
			peer.Cards.Service,
			peer.Database.Divisions(),
		)
	}

	{ // lootboxes setup
		peer.LootBoxes.Service = lootboxes.NewService(
			config.LootBoxes.Config,
			peer.Database.LootBoxes(),
			peer.Cards.Service,
			peer.Avatars.Service,
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

	{ // divisions setup
		peer.Divisions.Service = divisions.NewService(
			peer.Database.Divisions(),
			config.Divisions.Config)
	}

	{ // matches setup
		peer.Matches.Service = matches.NewService(
			peer.Database.Matches(),
			config.Matches.Config,
			peer.Clubs.Service,
			peer.Cards.Service,
		)
	}

	{ // seasons setup
		peer.Seasons.Service = seasons.NewService(
			peer.Database.Seasons(),
			config.Seasons.Config,
			peer.Divisions.Service,
			peer.Matches.Service,
			peer.Clubs.Service,
		)

		peer.Seasons.ExpirationSeasons = seasons.NewChore(
			config.Seasons.Config,
			peer.Seasons.Service,
		)
	}

	{ // queue setup
		peer.Queue.Service = queue.NewService(
			config.Queue.Config,
			peer.Database.Queue(),
			peer.Users.Service,
			peer.Clubs.Service,
		)

		peer.Queue.PlaceChore = queue.NewChore(
			peer.Log,
			config.Queue.Config,
			peer.Queue.Service,
			peer.Matches.Service,
			peer.Seasons.Service,
			peer.Clubs.Service,
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
			peer.Divisions.Service,
			peer.Matches.Service,
			peer.Seasons.Service,
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
			peer.Seasons.Service,
			peer.WaitList.Service,
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
	group.Go(func() error {
		return ignoreCancel(peer.Seasons.ExpirationSeasons.Run(ctx))
	})
	group.Go(func() error {
		return ignoreCancel(peer.NFTs.NFTChore.RunNFTSynchronization(ctx))
	})
	group.Go(func() error {
		return ignoreCancel(peer.WaitList.WaitListChore.RunCheckMintEvent(ctx))
	})

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group

	errlist.Add(peer.Admin.Endpoint.Close())
	errlist.Add(peer.Console.Endpoint.Close())
	peer.Marketplace.ExpirationLotChore.Close()
	peer.Queue.PlaceChore.Close()
	peer.Seasons.ExpirationSeasons.Close()

	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
