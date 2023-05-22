// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"
	"strconv"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/pkg/jsonrpc"
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

// RunNFTSynchronization runs the check of re-own of nft.
func (chore *Chore) RunNFTSynchronization(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		nfts, err := chore.nfts.List(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, nft := range nfts {
			tokenID, err := strconv.ParseInt(nft.TokenID.String(), 10, 64)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			data := evmsignature.Data{
				AddressContractMethod: chore.config.NFTContract.OwnerOfSelector,
				TokenID:               tokenID,
			}

			dataHex := evmsignature.NewDataHex(data)
			params := jsonrpc.Parameter{
				To:   chore.config.NFTContract.Address,
				Data: dataHex,
			}

			transaction := jsonrpc.NewTransaction(jsonrpc.MethodEthCall, []interface{}{&params, evmsignature.BlockTagLatest}, evmsignature.ChainIDRinkeby)
			body, err := jsonrpc.Send(chore.config.AddressNodeServer, transaction)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			ownersWalletAddress, err := jsonrpc.GetOwnersWalletAddress(body)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			if ownersWalletAddress == nft.WalletAddress {
				continue
			}

			nft := NFT{
				Chain:         evmsignature.ChainPolygon,
				TokenID:       nft.TokenID,
				WalletAddress: ownersWalletAddress,
			}
			if err = chore.nfts.Update(ctx, nft); err != nil {
				return ChoreError.Wrap(err)
			}

			user, err := chore.users.GetByWalletAddress(ctx, ownersWalletAddress, users.WalletTypeETH)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			if err = chore.cards.UpdateUserID(ctx, nft.CardID, user.ID); err != nil {
				return ChoreError.Wrap(err)
			}

		}

		return ChoreError.Wrap(err)
	})
}
