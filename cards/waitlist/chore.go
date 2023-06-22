// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
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
	log      logger.Logger
	config   Config
	loop     *thelooper.Loop
	waitList *Service
	nfts     *nfts.Service
	users    *users.Service
	cards    *cards.Service
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, waitList *Service, nfts *nfts.Service, users *users.Service, cards *cards.Service) *Chore {
	return &Chore{
		log:      log,
		config:   config,
		loop:     thelooper.NewLoop(config.WaitListRenewalInterval),
		waitList: waitList,
		nfts:     nfts,
		users:    users,
		cards:    cards,
	}
}

// RunCasperCheckMintEvent runs a task to check and create the casper nft assignment.
func (chore *Chore) RunCasperCheckMintEvent(ctx context.Context) error {
	err := chore.loop.Run(ctx, func(ctx context.Context) error {
		event, err := chore.waitList.GetNodeEvents(ctx)
		if err != nil {
			chore.log.Error("could not get node events", ChoreError.Wrap(err))
		}
		nftWaitList, err := chore.waitList.GetByTokenID(ctx, event.TokenID)
		if err != nil {
			chore.log.Error("could not get node events", ChoreError.Wrap(err))
		}

		toAddress := common.HexToAddress(nftWaitList.CasperWalletHash)
		nft := nfts.NFT{
			CardID:        nftWaitList.CardID,
			TokenID:       event.TokenID,
			Chain:         evmsignature.ChainEthereum,
			WalletAddress: toAddress,
		}

		if err = chore.nfts.Create(ctx, nft); err != nil {
			chore.log.Error("could not create nft", ChoreError.Wrap(err))
		}

		user, err := chore.users.GetByCasperHash(ctx, nftWaitList.CasperWalletHash)
		if err != nil {
			if err = chore.nfts.Delete(ctx, nft.CardID); err != nil {
				chore.log.Error("could not delete nft events", ChoreError.Wrap(err))
			}
			chore.log.Error("could get user by casper hash", ChoreError.Wrap(err))
		}

		if err = chore.nfts.Update(ctx, nft); err != nil {
			chore.log.Error("could not update nft", ChoreError.Wrap(err))
		}

		if err = chore.cards.UpdateUserID(ctx, nft.CardID, user.ID); err != nil {
			chore.log.Error("could not update user ID by card id", ChoreError.Wrap(err))
		}

		if err = chore.cards.UpdateMintedStatus(ctx, nft.CardID, cards.Minted); err != nil {
			chore.log.Error("could not update minted status to 1", ChoreError.Wrap(err))
		}

		return nil
	})
	if err != nil {
		chore.log.Error("could not get node events", ChoreError.Wrap(err))
	}

	return nil
}
