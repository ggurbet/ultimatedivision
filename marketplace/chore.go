// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"

	"github.com/BoostyLabs/thelooper"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

var (
	// ChoreError represents lot chore error type.
	ChoreError = errs.Class("expiration lot chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	Loop        *thelooper.Loop
	marketplace *Service
}

// NewChore instantiates Chore.
func NewChore(config Config, marketplace *Service) *Chore {
	return &Chore{
		marketplace: marketplace,
		Loop:        thelooper.NewLoop(config.LotRenewalInterval),
	}
}

// Run starts the chore for re-check the expiration time of the lot.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		lots, err := chore.marketplace.ListExpiredLot(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		// TODO: the transaction may be required for all operations.
		for _, lot := range lots {
			if lot.CurrentPrice.BitLen() != 0 {
				// TODO: unhold old user's money.

				winLot := WinLot{
					CardID:    lot.CardID,
					Type:      lot.Type,
					UserID:    lot.UserID,
					ShopperID: lot.ShopperID,
					Status:    StatusSold,
					Amount:    lot.MaxPrice,
				}

				err := chore.marketplace.WinLot(ctx, winLot)
				if err != nil {
					return ChoreError.Wrap(err)
				}
				continue
			}

			err := chore.marketplace.UpdateStatusLot(ctx, lot.CardID, StatusExpired)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			if lot.Type == TypeCard {
				if err := chore.marketplace.cards.UpdateStatus(ctx, lot.CardID, cards.StatusActive); err != nil {
					return ErrMarketplace.Wrap(err)
				}
			}
			// TODO: check other items.

		}
		return ChoreError.Wrap(err)
	})
}

// Close closes the chore for re-check the expiration time of the lot.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
