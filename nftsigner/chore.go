// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"
	"math/big"
	"time"
	"ultimatedivision/pkg/signer"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
	"ultimatedivision/users"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// Address defines address type.
type Address string

// ChoreConfig is the global configuration for nftsigner.
type ChoreConfig struct {
	RenewalInterval            time.Duration           `json:"renewalInterval"`
	PrivateKey                 evmsignature.PrivateKey `json:"privateKey"`
	NFTCreateContractAddress   common.Address          `json:"nftCreateContractAddress"`
	VelasSmartContractAddress  common.Address          `json:"velasSmartContractAddress"`
	CasperSmartContractAddress string                  `json:"casperSmartContractAddress"`
	CasperTokenContract        string                  `json:"casperTokenContract"`
}

// Chore requests for unsigned nft tokens and sign all of them .
//
// architecture: Chore
type Chore struct {
	loop   *thelooper.Loop
	config ChoreConfig
	nfts   *waitlist.Service
}

// NewChore instantiates Chore.
func NewChore(config ChoreConfig, db waitlist.DB) *Chore {
	return &Chore{
		loop:   thelooper.NewLoop(config.RenewalInterval),
		config: config,
		nfts:   waitlist.NewService(waitlist.Config{}, db, nil, nil, nil, nil),
	}
}

// Run starts the chore for signing unsigned nft token from ultimatedivision.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.loop.Run(ctx, func(ctx context.Context) error {
		unsignedNFTs, err := chore.nfts.ListWithoutPassword(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}
		privateKeyECDSA, err := crypto.HexToECDSA(string(chore.config.PrivateKey))
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, token := range unsignedNFTs {
			var (
				signature           evmsignature.Signature
				smartContract       common.Address
				casperTokenContract string
				casperContract      string
			)

			switch token.WalletType {
			case users.WalletTypeETH:
				smartContract = chore.config.NFTCreateContractAddress
			case users.WalletTypeVelas:
				smartContract = chore.config.VelasSmartContractAddress
			case users.WalletTypeCasper:
				casperContract = chore.config.CasperSmartContractAddress
				casperTokenContract = chore.config.CasperTokenContract
			}

			if token.Value.Cmp(big.NewInt(0)) <= 0 {
				if casperContract != "" {
					signature, err = signer.GenerateCasperSignatureWithValue(signer.Address(token.CasperWallet),
						signer.Address(casperContract), token.TokenID, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				} else {
					signature, err = signer.GenerateSignatureWithValue(signer.Address(token.Wallet.String()),
						signer.Address(smartContract.String()), token.TokenID, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				}
			} else {
				if casperContract != "" {
					signature, err = signer.GenerateCasperSignatureWithValueAndNonce(signer.Address(token.CasperWallet),
						signer.Address(casperTokenContract), &token.Value, token.TokenNumber, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				} else {
					signature, err = signer.GenerateSignatureWithValueAndNonce(signer.Address(token.Wallet.String()),
						signer.Address(smartContract.String()), &token.Value, token.TokenNumber, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				}
			}

			if err = chore.nfts.Update(ctx, token.TokenID, signature); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}
