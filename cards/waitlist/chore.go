// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"
	"strconv"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/pkg/jsonrpc"
	"ultimatedivision/users"
)

var (
	// ChoreError represents waitlist chore error type.
	ChoreError = errs.Class("expiration waitlist chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
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

// RunCheckMintEvent runs a task to check the nft assignment.
func (chore *Chore) RunCheckMintEvent(ctx context.Context) (err error) {
	filter := []*jsonrpc.CreateFilter{
		{
			ToBlock: evmsignature.BlockTagLatest,
			Address: chore.config.NFTContract.Address,
			Topics:  []evmsignature.Hex{chore.config.NFTContract.AddressEvent},
		},
	}

	transaction := jsonrpc.NewTransaction(jsonrpc.MethodEthNewFilter, filter, evmsignature.ChainIDRinkeby)
	body, err := jsonrpc.Send(chore.config.AddressNodeServer, transaction)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	addressOfFilter, err := jsonrpc.GetAddressOfFilter(body)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		transaction := jsonrpc.NewTransaction(jsonrpc.MethodEthGetFilterChanges, []evmsignature.Address{addressOfFilter}, evmsignature.ChainIDRinkeby)
		events, err := jsonrpc.GetEvents(chore.config.AddressNodeServer, transaction)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, event := range events {
			fromStr := string(event.Topics[1])
			from, _ := strconv.ParseInt(fromStr[evmsignature.LengthHexPrefix:], 16, 64)

			toStr := string(event.Topics[2])
			toAddress := evmsignature.CreateValidAddress(evmsignature.Hex(toStr))

			tokenIDStr := string(event.Topics[3])
			tokenID, err := strconv.ParseInt(tokenIDStr[evmsignature.LengthHexPrefix:], 16, 64)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			if from == 0 {
				nftWaitList, err := chore.waitList.GetByTokenID(ctx, tokenID)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				nft := nfts.NFT{
					CardID:        nftWaitList.CardID,
					Chain:         evmsignature.ChainPolygon,
					TokenID:       tokenID,
					WalletAddress: toAddress,
				}
				if err = chore.nfts.Create(ctx, nft); err != nil {
					return ChoreError.Wrap(err)
				}
				continue
			}

			nft, err := chore.nfts.Get(ctx, tokenID, evmsignature.ChainPolygon)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			user, err := chore.users.GetByWalletAddress(ctx, toAddress)
			if err != nil {
				if err = chore.nfts.Delete(ctx, nft.CardID); err != nil {
					return ChoreError.Wrap(err)
				}

				if err = chore.cards.UpdateUserID(ctx, nft.CardID, uuid.Nil); err != nil {
					return ChoreError.Wrap(err)
				}
				continue
			}

			if err = chore.nfts.Update(ctx, nft); err != nil {
				return ChoreError.Wrap(err)
			}

			if err = chore.cards.UpdateUserID(ctx, nft.CardID, user.ID); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}
