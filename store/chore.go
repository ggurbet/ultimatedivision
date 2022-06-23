// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package store

import (
	"context"
	"math/big"
	"time"

	"github.com/BoostyLabs/thelooper"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
)

var (
	// ChoreError represents store chore error type.
	ChoreError = errs.Class("expiration store chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore.
type Chore struct {
	Loop    *thelooper.Loop
	config  Config
	store   *Service
	cards   *cards.Service
	avatars *avatars.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, store *Service, cards *cards.Service, avatars *avatars.Service) *Chore {
	return &Chore{
		Loop:    thelooper.NewLoop(config.StoreRenewalInterval),
		config:  config,
		store:   store,
		cards:   cards,
		avatars: avatars,
	}
}

// Run runs the renewal of cards in store.
func (chore *Chore) Run(ctx context.Context) error {
	if _, err := chore.store.Get(ctx, ActiveSetting); err != nil {
		if !ErrNoSetting.Has(err) {
			return ChoreError.Wrap(err)
		}

		setting := Setting{
			ID:          ActiveSetting,
			CardsAmount: 10,
			IsRenewal:   true,
			HourRenewal: 0,
			Price:       *big.NewInt(100),
		}
		if err = chore.store.Create(ctx, setting); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	t := time.Now().UTC()
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		setting, err := chore.store.Get(ctx, ActiveSetting)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		timeRenewal := time.Date(t.Year(), t.Month(), t.Day(), setting.HourRenewal, 0, 0, 0, time.UTC)
		duration := time.Until(timeRenewal)
		if duration > 0 {
			return nil
		}
		t = t.Add(24 * time.Hour)

		if !setting.IsRenewal {
			return nil
		}

		cardsList, err := chore.cards.ListByTypeNoOrdered(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		cardsAmount := setting.CardsAmount - len(cardsList)
		percentageQualities := []int{
			chore.config.PercentageQualities.Wood,
			chore.config.PercentageQualities.Silver,
			chore.config.PercentageQualities.Gold,
			chore.config.PercentageQualities.Diamond,
		}

		for i := 0; i < cardsAmount; i++ {
			card, err := chore.cards.Create(ctx, uuid.Nil, percentageQualities, cards.TypeUnordered)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			_, err = chore.avatars.Generate(ctx, card, card.ID.String())
			if err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}

// Close closes the chore for renewal cards in store.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
