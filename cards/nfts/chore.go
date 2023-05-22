// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"github.com/BoostyLabs/thelooper"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/users"
)

var (
	// ChoreError represents nfts chore error type.
	ChoreError = errs.Class("expiration nfts chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	config Config
	Loop   *thelooper.Loop
	nfts   *Service
	users  *users.Service
	cards  *cards.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, nfts *Service, users *users.Service, cards *cards.Service) *Chore {
	return &Chore{
		config: config,
		Loop:   thelooper.NewLoop(config.NFTRenewalInterval),
		nfts:   nfts,
		users:  users,
		cards:  cards,
	}
}
