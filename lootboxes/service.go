// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
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
func (service *Service) Create(ctx context.Context, lootBoxType Type, userID uuid.UUID) (LootBox, error) {
	userLootBox := LootBox{
		UserID:    userID,
		LootBoxID: uuid.New(),
		Type:      lootBoxType,
	}

	return userLootBox, ErrLootBoxes.Wrap(service.lootboxes.Create(ctx, userLootBox))
}

// Open opens lootbox by user.
func (service *Service) Open(ctx context.Context, userID, lootboxID uuid.UUID) ([]cards.Card, error) {
	cardsNum := 0
	probabilities := make([]int, 0, 4)

	userLootBox, err := service.lootboxes.Get(ctx, lootboxID)
	if err != nil {
		return nil, ErrLootBoxes.Wrap(err)
	}
	if userLootBox.Type == RegularBox {
		cardsNum = service.config.RegularBoxConfig.CardsNum
		probabilities = []int{service.config.RegularBoxConfig.Wood, service.config.RegularBoxConfig.Silver, service.config.RegularBoxConfig.Gold, service.config.RegularBoxConfig.Diamond}
	} else if userLootBox.Type == UDReleaseCelebrationBox {
		cardsNum = service.config.UDReleaseCelebrationBoxConfig.CardsNum
		probabilities = []int{service.config.UDReleaseCelebrationBoxConfig.Wood, service.config.UDReleaseCelebrationBoxConfig.Silver, service.config.UDReleaseCelebrationBoxConfig.Gold, service.config.UDReleaseCelebrationBoxConfig.Diamond}
	}

	var lootBoxCards []cards.Card

	for i := 0; i < cardsNum; i++ {
		card, err := service.cards.Create(ctx, userID, probabilities, strconv.Itoa(i))
		if err != nil {
			return lootBoxCards, ErrLootBoxes.Wrap(err)
		}

		lootBoxCards = append(lootBoxCards, card)
	}

	sortLootBoxCards(lootBoxCards)

	err = service.lootboxes.Delete(ctx, lootboxID)

	return lootBoxCards, ErrLootBoxes.Wrap(err)
}

// List returns all loot boxes.
func (service *Service) List(ctx context.Context) ([]LootBox, error) {
	userLootBoxes, err := service.lootboxes.List(ctx)

	return userLootBoxes, ErrLootBoxes.Wrap(err)
}
