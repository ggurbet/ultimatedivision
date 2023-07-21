// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardgenerator

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"ultimatedivision/cardgenerator/cardavatars"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
)

// Config is the global configuration for cardgenerator.
type Config struct {
	CardAvatars struct {
		cardavatars.Config
	} `json:"cardAvatars"`
}

// Peer is the representation of a cardgenerator.
type Peer struct {
	Config Config
	Log    logger.Logger

	cardsTotal int

	// exposes cards related logic.
	Cards struct {
		Service *cards.Service
	}

	// exposes avatars related logic.
	Avatars struct {
		Service *avatars.Service
	}

	// exposes nfts related logic.
	NFTs struct {
		Service *nfts.Service
	}

	// exposes avatar cards related logic.
	CardAvatars struct {
		Service *cardavatars.Service
	}
}

// New is a constructor for cardgenerator.Peer.
func New(logger logger.Logger, config Config, cardsTotal int) (peer *Peer, err error) {
	peer = &Peer{
		Log:        logger,
		Config:     config,
		cardsTotal: cardsTotal,
	}

	{ // cards setup.
		peer.Cards.Service = cards.NewService(
			nil,
			config.CardAvatars.CardConfig,
		)
	}

	{ // avatars setup.
		peer.Avatars.Service = avatars.NewService(
			nil,
			nil,
			config.CardAvatars.AvatarConfig,
		)
	}

	{ // nfts setup.
		peer.NFTs.Service = nfts.NewService(
			config.CardAvatars.NFTConfig,
			nil,
		)
	}

	{ // avatar cards setup.
		peer.CardAvatars.Service = cardavatars.NewService(
			config.CardAvatars.Config,
			peer.Cards.Service,
			peer.Avatars.Service,
			peer.NFTs.Service,
		)
	}

	return peer, nil
}

// Generate initiates generation of avatar cards.
func (peer *Peer) Generate(ctx context.Context) error {
	for i := 0; i < peer.cardsTotal; i++ {
		allNames := make(map[string]struct{}, peer.cardsTotal)
		for len(allNames) <= peer.cardsTotal {
			if err := peer.CardAvatars.Service.GenerateName(peer.Config.CardAvatars.CardConfig.PathToNamesDataset, allNames); err != nil {
				return err
			}
		}

		var playerName string
		for name := range allNames {
			playerName = name
			delete(allNames, name)
			break
		}

		nft, err := peer.CardAvatars.Service.Generate(ctx, i, playerName)
		if err != nil {
			return err
		}

		file, err := json.MarshalIndent(nft, "", " ")
		if err != nil {
			return err
		}

		if err = os.MkdirAll(peer.Config.CardAvatars.PathToOutputJSONFile, os.ModePerm); err != nil {
			return err
		}

		if err = os.WriteFile(filepath.Join(peer.Config.CardAvatars.PathToOutputJSONFile, strconv.Itoa(i+1)+".json"), file, 0644); err != nil {
			return err
		}
	}

	return nil
}

// TestGenerate initiates generation test version of avatar cards.
func (peer *Peer) TestGenerate(ctx context.Context) error {
	avatars, err := peer.CardAvatars.Service.TestGenerate(ctx, peer.cardsTotal)
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(avatars, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(peer.Config.CardAvatars.PathToOutputJSONFile, "data-that-make-up-avatar.json"), file, 0644)
}
