// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package store

import (
	"context"
	"math/big"
	"math/rand"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/waitlist"
)

// ErrStore indicated that there was an error in service.
var ErrStore = errs.Class("store service error")

// Service is handling store related logic.
//
// architecture: Service.
type Service struct {
	config   Config
	store    DB
	cards    *cards.Service
	waitlist *waitlist.Service
}

// NewService is a constructor for store service.
func NewService(config Config, store DB, cards *cards.Service, waitlist *waitlist.Service) *Service {
	return &Service{
		config:   config,
		store:    store,
		cards:    cards,
		waitlist: waitlist,
	}
}

// Buy returns values required for minting nft with value.
func (service *Service) Buy(ctx context.Context, createNFT waitlist.CreateNFT) (waitlist.Transaction, error) {
	var transaction waitlist.Transaction

	cards, err := service.cards.ListByTypeOrdered(ctx)
	if err != nil {
		return transaction, ErrStore.Wrap(err)
	}
	if len(cards) == 0 {
		return transaction, ErrStore.New("all cards of store are minted")
	}

	randNumberCard := rand.Intn(len(cards)) - 1
	createNFT.CardID = cards[randNumberCard].ID
	// TODO: add price in setting of store.
	createNFT.Value = *big.NewInt(100)

	// TODO: change selector of buy method.
	transaction, err = service.waitlist.Create(ctx, createNFT)

	return transaction, ErrStore.Wrap(err)
}

// Create creates setting of store in database.
func (service *Service) Create(ctx context.Context, setting Setting) error {
	return ErrStore.Wrap(service.store.Create(ctx, setting))
}

// Get returns setting of store by id from database.
func (service *Service) Get(ctx context.Context, id int) (Setting, error) {
	setting, err := service.store.Get(ctx, id)
	return setting, ErrStore.Wrap(err)
}

// List returns settings of store from database.
func (service *Service) List(ctx context.Context) ([]Setting, error) {
	settings, err := service.store.List(ctx)
	return settings, ErrStore.Wrap(err)
}

// Update updates setting of store in database.
func (service *Service) Update(ctx context.Context, setting Setting) error {
	return ErrStore.Wrap(service.store.Update(ctx, setting))
}
