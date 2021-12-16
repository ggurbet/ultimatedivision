// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencysigner

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/pkg/sync"
	"ultimatedivision/udts/currencywaitlist"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// ChoreConfig is the global configuration for currencysigner.
type ChoreConfig struct {
	RenewalInterval    time.Duration          `json:"renewalInterval"`
	PrivateKey         cryptoutils.PrivateKey `json:"privateKey"`
	UDTContractAddress cryptoutils.Address    `json:"udtContractAddress"`
}

// Chore requests for unsigned nft tokens and sign all of them .
//
// architecture: Chore
type Chore struct {
	log              logger.Logger
	currencywaitlist *currencywaitlist.Service
	Loop             *sync.Cycle
	config           ChoreConfig
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config ChoreConfig, db currencywaitlist.DB) *Chore {
	return &Chore{
		log:              log,
		Loop:             sync.NewCycle(config.RenewalInterval),
		currencywaitlist: currencywaitlist.NewService(currencywaitlist.Config{}, db, nil, nil),
		config:           config,
	}
}

// Run starts the chore for signing unsigned item of currency waitlist from ultimatedivision.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		unsignedItems, err := chore.currencywaitlist.ListWithoutSignature(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		privateKeyECDSA, err := crypto.HexToECDSA(string(chore.config.PrivateKey))
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, item := range unsignedItems {
			signature, err := cryptoutils.GenerateSignatureWithValueAndNonce(item.WalletAddress, chore.config.UDTContractAddress, &item.Value, item.Nonce, privateKeyECDSA)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			err = chore.currencywaitlist.UpdateSignature(ctx, signature, item.WalletAddress, item.Nonce)
			if err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}

// Close closes the chore for signing unsigned item of currency waitlist from ultimatedivision.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
