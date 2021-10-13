// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardgenerator

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"ultimatedivision/cardgenerator/avatarcards"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/internal/logger"
)

// Config is the global configuration for cardgenerator.
type Config struct {
	AvatarCards struct {
		avatarcards.Config
	} `json:"avatarCards"`
}

// Peer is the representation of a cardgenerator.
type Peer struct {
	Config Config
	Log    logger.Logger

	quantityOfCard int

	// exposes cards related logic.
	Cards struct {
		Service *cards.Service
	}

	// exposes avatars related logic.
	Avatars struct {
		Service *avatars.Service
	}

	// exposes avatar cards related logic.
	AvatarCards struct {
		Service *avatarcards.Service
	}
}

// New is a constructor for cardgenerator.Peer.
func New(logger logger.Logger, config Config, quantityOfCard int) (peer *Peer, err error) {
	peer = &Peer{
		Log:            logger,
		Config:         config,
		quantityOfCard: quantityOfCard,
	}

	{ // Avatars setup
		peer.Avatars.Service = avatars.NewService(
			nil,
			config.AvatarCards.AvatarConfig,
		)
	}

	{ // cards setup
		peer.Cards.Service = cards.NewService(
			nil,
			config.AvatarCards.CardConfig,
			peer.Avatars.Service,
		)
	}

	{ // avatar cards setup
		peer.AvatarCards.Service = avatarcards.NewService(
			config.AvatarCards.Config,
			peer.Cards.Service,
			peer.Avatars.Service,
		)
	}

	return peer, nil
}

// Generate initiates generation of avatar cards.
func (peer *Peer) Generate(ctx context.Context) error {
	cardsWithAvatars, err := peer.AvatarCards.Service.Generate(ctx, peer.quantityOfCard)
	if err != nil {
		return err
	}
	file, err := json.MarshalIndent(cardsWithAvatars, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(peer.Config.AvatarCards.PathToOutputJSONFile, peer.Config.AvatarCards.NameOutputJSONFile+".json"), file, 0644)
}

// TestGenerate initiates generation test version of avatar cards.
func (peer *Peer) TestGenerate(ctx context.Context) error {
	avatars, err := peer.AvatarCards.Service.TestGenerate(ctx, peer.quantityOfCard)
	if err != nil {
		return err
	}
	file, err := json.MarshalIndent(avatars, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(peer.Config.AvatarCards.PathToOutputJSONFile, peer.Config.AvatarCards.NameOutputJSONFile+"_test.json"), file, 0644)
}
