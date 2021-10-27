// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"context"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/pkg/fileutils"
)

// ErrCardWithLinkToAvatar indicated that there was an error in service.
var ErrCardWithLinkToAvatar = errs.Class("card with link to avatar service error")

// Service is handling cards with link to avatars related logic.
//
// architecture: Service
type Service struct {
	config  Config
	cards   *cards.Service
	avatars *avatars.Service
}

// NewService is a constructor for card with link to avatar service.
func NewService(config Config, cards *cards.Service, avatars *avatars.Service) *Service {
	return &Service{
		config:  config,
		cards:   cards,
		avatars: avatars,
	}
}

// Generate generates cards with avatar link.
func (service *Service) Generate(ctx context.Context, count int) ([]CardWithLinkToAvatar, error) {
	var (
		err                   error
		cardsWithLinkToAvatar []CardWithLinkToAvatar
	)

	id := uuid.New()
	percentageQualities := []int{
		service.config.PercentageQualities.Wood,
		service.config.PercentageQualities.Silver,
		service.config.PercentageQualities.Gold,
		service.config.PercentageQualities.Diamond,
	}

	allNames := make(map[string]struct{}, count)

	for i := 0; i < count; i++ {
		var cardWithAvatar CardWithLinkToAvatar
		var avatar avatars.Avatar
		if cardWithAvatar.Card, err = service.cards.Generate(ctx, id, percentageQualities); err != nil {
			return nil, ErrCardWithLinkToAvatar.Wrap(err)
		}

		if avatar, err = service.avatars.Generate(ctx, cardWithAvatar.Card, strconv.Itoa(i)); err != nil {
			return nil, ErrCardWithLinkToAvatar.Wrap(err)
		}
		cardWithAvatar.OriginalURL = avatar.OriginalURL

		cardsWithLinkToAvatar = append(cardsWithLinkToAvatar, cardWithAvatar)
	}

	for len(allNames) < count {
		err = generateName(service.config.PathToNamesDataset, allNames)
		if err != nil {
			return nil, ErrCardWithLinkToAvatar.Wrap(err)
		}
	}

	for i := 0; i < count; i++ {
		for name := range allNames {
			cardsWithLinkToAvatar[i].PlayerName = name
			delete(allNames, name)
			break
		}
	}

	return cardsWithLinkToAvatar, nil
}

// generateName generates name of card.
func generateName(path string, names map[string]struct{}) error {
	file, err := os.Open(path)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, file.Close())
	}()

	rand.Seed(time.Now().UTC().UnixNano())

	totalCount, err := fileutils.CountLines(file)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}

	randomNum := rand.Intn(totalCount) + 1

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}

	name, err := fileutils.ReadLine(file, randomNum)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}

	names[name] = struct{}{}

	return ErrCardWithLinkToAvatar.Wrap(err)
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
		var avatarCard CardWithLinkToAvatar
		if avatarCard.Card, err = service.cards.Generate(ctx, id, percentageQualities); err != nil {
			return nil, ErrCardWithLinkToAvatar.Wrap(err)
		}

		avatar, err := service.avatars.Generate(ctx, avatarCard.Card, avatarCard.Card.ID.String())
		if err != nil {
			return nil, ErrCardWithLinkToAvatar.Wrap(err)
		}

		avatars = append(avatars, avatar)
	}

	return avatars, nil
}
