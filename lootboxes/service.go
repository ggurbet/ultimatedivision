// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/auth"
	"ultimatedivision/users/userauth"
)

// ErrLootBoxes indicates that there was an error in the service.
var ErrLootBoxes = errs.Class("lootboxes service error")

// Service is handling lootboxes related logic.
//
// architecture: Service
type Service struct {
	config    Config
	lootboxes DB
	cards     *cards.Service
}

// NewService is a constructor for lootboxes service.
func NewService(config Config, lootboxes DB, cards *cards.Service) *Service {
	return &Service{
		config:    config,
		lootboxes: lootboxes,
		cards:     cards,
	}
}

// Create creates LootBox.
func (service *Service) Create(ctx context.Context, userLootBox LootBox) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return userauth.ErrUnauthenticated.Wrap(err)
	}

	userLootBox.UserID = claims.ID

	return ErrLootBoxes.Wrap(service.lootboxes.Create(ctx, userLootBox))
}

// Open opens lootbox by user.
func (service *Service) Open(ctx context.Context, userLootBox LootBox) ([]cards.Card, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, userauth.ErrUnauthenticated.Wrap(err)
	}

	userLootBox.UserID = claims.ID

	cardsNum := 0
	probabilities := make([]int, 0, 4)

	if userLootBox.Type == RegularBox {
		cardsNum = service.config.RegularBoxConfig.CardsNum
		probabilities = []int{service.config.RegularBoxConfig.Wood, service.config.RegularBoxConfig.Silver, service.config.RegularBoxConfig.Gold, service.config.RegularBoxConfig.Diamond}
	} else if userLootBox.Type == UDReleaseCelebrationBox {
		cardsNum = service.config.UDReleaseCelebrationBoxConfig.CardsNum
		probabilities = []int{service.config.UDReleaseCelebrationBoxConfig.Wood, service.config.UDReleaseCelebrationBoxConfig.Silver, service.config.UDReleaseCelebrationBoxConfig.Gold, service.config.UDReleaseCelebrationBoxConfig.Diamond}
	}

	var lootBoxCards []cards.Card

	for i := 0; i < cardsNum; i++ {
		card, err := service.cards.Create(ctx, userLootBox.UserID, probabilities)
		if err != nil {
			return lootBoxCards, ErrLootBoxes.Wrap(err)
		}

		lootBoxCards = append(lootBoxCards, card)
	}

	err = service.lootboxes.Delete(ctx, userLootBox)

	return lootBoxCards, ErrLootBoxes.Wrap(err)
}
