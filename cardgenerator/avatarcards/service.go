// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
)

// ErrAvatarCard indicated that there was an error in service.
var ErrAvatarCard = errs.Class("avatar card service error")

// Service is handling avatars related logic.
//
// architecture: Service
type Service struct {
	config  Config
	cards   *cards.Service
	avatars *avatars.Service
}

// NewService is a constructor for avatar card service.
func NewService(config Config, cards *cards.Service, avatars *avatars.Service) *Service {
	return &Service{
		config:  config,
		cards:   cards,
		avatars: avatars,
	}
}

// Generate generates avatar cards.
func (service *Service) Generate(ctx context.Context, count int) ([]AvatarCards, error) {
	var (
		err         error
		avatarCards []AvatarCards
	)

	id := uuid.New()
	percentageQualities := []int{
		service.config.PercentageQualities.Wood,
		service.config.PercentageQualities.Silver,
		service.config.PercentageQualities.Gold,
		service.config.PercentageQualities.Diamond,
	}

	for i := 0; i < count; i++ {
		var avatarCard AvatarCards
		var avatar avatars.Avatar
		if avatarCard.Card, err = service.cards.Generate(ctx, id, percentageQualities); err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		if avatar, err = service.avatars.Generate(ctx, avatarCard.Card.ID, avatarCard.Card.IsTattoo, strconv.Itoa(i)); err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}
		avatarCard.OriginalURL = avatar.OriginalURL

		avatarCards = append(avatarCards, avatarCard)
	}

	return avatarCards, nil
}

// TestGenerate generates test version avatar cards.
func (service *Service) TestGenerate(ctx context.Context, count int) ([]avatars.Avatar, error) {
	var (
		err     error
		avatars []avatars.Avatar
	)

	id := uuid.New()
	percentageQualities := []int{
		service.config.PercentageQualities.Wood,
		service.config.PercentageQualities.Silver,
		service.config.PercentageQualities.Gold,
		service.config.PercentageQualities.Diamond,
	}

	for i := 0; i < count; i++ {
		var avatarCard AvatarCards
		if avatarCard.Card, err = service.cards.Generate(ctx, id, percentageQualities); err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		avatar, err := service.avatars.Generate(ctx, avatarCard.Card.ID, avatarCard.Card.IsTattoo, avatarCard.Card.ID.String())
		if err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		avatars = append(avatars, avatar)
	}

	return avatars, nil
}
