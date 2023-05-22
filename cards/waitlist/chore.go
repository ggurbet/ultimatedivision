// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"github.com/BoostyLabs/thelooper"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/users"
)

var (
	// ChoreError represents waitlist chore error type.
	ChoreError = errs.Class("expiration waitlist chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore.
type Chore struct {
	config   Config
	Loop     *thelooper.Loop
	waitList *Service
	nfts     *nfts.Service
	users    *users.Service
	cards    *cards.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, waitList *Service, nfts *nfts.Service, users *users.Service, cards *cards.Service) *Chore {
	return &Chore{
		config:   config,
		Loop:     thelooper.NewLoop(config.WaitListRenewalInterval),
		waitList: waitList,
		nfts:     nfts,
		users:    users,
		cards:    cards,
	}
}
