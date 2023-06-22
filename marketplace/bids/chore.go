// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package bids

import (
	"context"
	"fmt"

	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/cards/waitlist"
	"ultimatedivision/clubs"
	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
	"ultimatedivision/users"
)

var (
	// ChoreError represents bid chore error type.
	ChoreError = errs.Class("expiration bid chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log         logger.Logger
	config      Config
	loop        *thelooper.Loop
	bids        *Service
	clubs       *clubs.Service
	marketplace *marketplace.Service
	users       *users.Service
	cards       *cards.Service
	nfts        *nfts.Service
	waitlist    *waitlist.Service
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, bids *Service, clubs *clubs.Service, marketplace *marketplace.Service, users *users.Service,
	cards *cards.Service, nfts *nfts.Service, waitlist *waitlist.Service) *Chore {
	return &Chore{
		log:         log,
		config:      config,
		loop:        thelooper.NewLoop(config.ExpiredLotRenewalInterval),
		bids:        bids,
		clubs:       clubs,
		marketplace: marketplace,
		users:       users,
		cards:       cards,
		nfts:        nfts,
		waitlist:    waitlist,
	}
}

// Run starts the chore for re-check the expiration time of the lot.
func (chore *Chore) Run(ctx context.Context) error {
	err := chore.loop.Run(ctx, func(ctx context.Context) error {
		expiredLots, err := chore.marketplace.ListExpiredLots(ctx)
		if err != nil {
			return nil
		}
		for _, lot := range expiredLots {
			currentBid, err := chore.bids.GetCurrentBidByLotID(ctx, lot.CardID)
			if err != nil {
				if !ErrNoBid.Has(err) {
					chore.log.Error(fmt.Sprintf("could not get current bid by lot id equal %v from db", lot.CardID), ChoreError.Wrap(err))
				}
			}

			// TODO: transaction required.
			if err = chore.marketplace.Delete(ctx, lot.CardID); err != nil {
				chore.log.Error(fmt.Sprintf("could not delete lot by card id equal %v in db", lot.CardID), ChoreError.Wrap(err))
			}
			if err = chore.cards.UpdateStatus(ctx, lot.CardID, cards.StatusActive); err != nil {
				chore.log.Error(fmt.Sprintf("could not update card status by card id equal %v in db", lot.CardID), ChoreError.Wrap(err))
			}

			if err = chore.bids.DeleteByLotID(ctx, lot.CardID); err != nil {
				chore.log.Error(fmt.Sprintf("could not delete bids by card id equal %v in db", lot.CardID), ChoreError.Wrap(err))
			}

			squadID, err := chore.clubs.GetSquadIDByCardID(ctx, lot.CardID)
			if err != nil {
				chore.log.Error(fmt.Sprintf("could not get squad by card id equal %v from db", lot.CardID), ChoreError.Wrap(err))
			}

			if squadID != uuid.Nil {
				if err = chore.clubs.DeleteByCardID(ctx, lot.CardID); err != nil {
					chore.log.Error(fmt.Sprintf("could not delete card from club by card id equal %v in db", lot.CardID), ChoreError.Wrap(err))
				}
			}

			user, err := chore.users.Get(ctx, currentBid.UserID)
			if err != nil {
				chore.log.Error(fmt.Sprintf("could not get user by user id equal %v from db", currentBid.UserID), ChoreError.Wrap(err))
			}

			if err = chore.cards.UpdateUserID(ctx, currentBid.LotID, currentBid.UserID); err != nil {
				chore.log.Error(fmt.Sprintf("could not get update user id of the card lot id equal %v in db", currentBid.LotID), ChoreError.Wrap(err))
			}

			card, err := chore.cards.Get(ctx, currentBid.LotID)
			if err != nil {
				chore.log.Error(fmt.Sprintf("could not get card by lot id equal %v from db", currentBid.LotID), ChoreError.Wrap(err))
			}

			nft, err := chore.nfts.GetNFTByCardID(ctx, card.ID)
			if err != nil {
				chore.log.Error(fmt.Sprintf("could not get nft by card id equal %v from db", card.ID), ChoreError.Wrap(err))
			}

			if user.WalletType == users.WalletTypeCasper {
				nft.WalletAddress = common.HexToAddress(user.CasperWallet)
			} else {
				nft.WalletAddress = user.Wallet
			}

			if err = chore.nfts.Update(ctx, nft); err != nil {
				chore.log.Error("could not update nft by nft data from db", ChoreError.Wrap(err))
			}
		}

		return nil
	})
	if err != nil {
		chore.log.Error("could not check expired lots", ChoreError.Wrap(err))
	}

	return nil
}

// Close closes the chore for re-check the expiration time of the lot.
func (chore *Chore) Close() {
	chore.loop.Close()
}
