// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencysigner

import (
	"context"
	"errors"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/internal/logger"
	"ultimatedivision/udts/currencywaitlist"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// CurrencyWaitList provides access to currencywaitlist db.
	CurrencyWaitList() currencywaitlist.DB
}

// Config is the global configuration for currencysigner.
type Config struct {
	Chore struct {
		ChoreConfig
	} `json:"chore"`
}

// Peer is the representation of a currencysigner.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB

	Chore *Chore
}

// New is a constructor for currencysigner.Peer.
func New(logger logger.Logger, config Config, database DB) (peer *Peer, err error) {
	peer = &Peer{
		Log:      logger,
		Config:   config,
		Database: database,
	}

	{ // chore setup
		peer.Chore = NewChore(logger, config.Chore.ChoreConfig, peer.Database.CurrencyWaitList())
	}

	return peer, nil
}

// Run runs currencysigner.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return ignoreCancel(peer.Chore.Run(ctx))
	})
	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group

	peer.Chore.Close()

	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
