// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/internal/logger"
)

// ErrLootBoxes indicates that there was an error in the service.
var ErrLootBoxes = errs.Class("lootboxes service error")

// Service is handling lootboxes related logic.
//
// architecture: Service
type Service struct {
	log       logger.Logger
	config    Config
	lootboxes DB
	cards     *cards.Service
	avatars   *avatars.Service
}

// NewService is a constructor for lootboxes service.
func NewService(log logger.Logger, config Config, lootboxes DB, cards *cards.Service, avatars *avatars.Service) *Service {
	return &Service{
		log:       log,
		config:    config,
		lootboxes: lootboxes,
		cards:     cards,
		avatars:   avatars,
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
		card, err := service.cards.Create(ctx, userID, probabilities)
		if err != nil {
			return lootBoxCards, ErrLootBoxes.Wrap(err)
		}

		lootBoxCards = append(lootBoxCards, card)
	}

	sortLootBoxCards(lootBoxCards)

	var wg sync.WaitGroup
	wg.Add(len(lootBoxCards))
	for _, card := range lootBoxCards {
		go service.GenerateAvatarForLootboxCards(ctx, card, &wg)
	}
	wg.Wait()

	err = service.lootboxes.Delete(ctx, lootboxID)

	return lootBoxCards, ErrLootBoxes.Wrap(err)
}

// GenerateAvatarForLootboxCards generates and saves avatar for card.
func (service *Service) GenerateAvatarForLootboxCards(ctx context.Context, card cards.Card, wg *sync.WaitGroup) {
	defer wg.Done()

	avatar, err := service.avatars.Generate(ctx, card, card.ID.String())
	if err != nil {
		service.log.Error("could not create card", err)
		return
	}

	if err := service.avatars.Create(ctx, avatar); err != nil {
		service.log.Error("could not create card", err)
		return
	}
}

// List returns all loot boxes.
func (service *Service) List(ctx context.Context) ([]LootBox, error) {
	userLootBoxes, err := service.lootboxes.List(ctx)

	return userLootBoxes, ErrLootBoxes.Wrap(err)
}
