// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/users"
)

// ErrWaitlist indicated that there was an error in service.
var ErrWaitlist = errs.Class("waitlist service error")

// Service is handling waitList related logic.
//
// architecture: Service
type Service struct {
	waitList DB
	cards    *cards.Service
	avatars  *avatars.Service
	users    *users.Service
	nfts     *nfts.Service
}

// NewService is a constructor for waitlist service.
func NewService(waitList DB, cards *cards.Service, avatars *avatars.Service, users *users.Service, nfts *nfts.Service) *Service {
	return &Service{
		waitList: waitList,
		cards:    cards,
		avatars:  avatars,
		users:    users,
		nfts:     nfts,
	}
}

// Create creates nft for wait list.
func (service *Service) Create(ctx context.Context, createNFT CreateNFT) error {
	card, err := service.cards.Get(ctx, createNFT.CardID)
	if err != nil {
		return ErrWaitlist.Wrap(err)
	}

	if card.UserID != createNFT.UserID {
		return ErrWaitlist.New("it isn't user`s card")
	}

	avatar, err := service.avatars.Get(ctx, createNFT.CardID)
	if err != nil {
		return ErrWaitlist.Wrap(err)
	}

	// TODO: save avatar in file storage

	service.nfts.Generate(ctx, card, avatar.OriginalURL)

	// TODO: save metadata in file storage
	// TODO: add transaction

	if err = service.users.UpdateWalletAddress(ctx, createNFT.WalletAddress, createNFT.UserID); err != nil {
		return ErrWaitlist.Wrap(err)
	}

	return service.waitList.Create(ctx, createNFT.CardID, createNFT.WalletAddress)
}

// List returns all nft for wait list.
func (service *Service) List(ctx context.Context) ([]Item, error) {
	allNFT, err := service.waitList.List(ctx)
	return allNFT, ErrWaitlist.Wrap(err)
}

// Get returns nft for wait list by token id.
func (service *Service) Get(ctx context.Context, tokenID int) (Item, error) {
	nft, err := service.waitList.Get(ctx, tokenID)
	return nft, ErrWaitlist.Wrap(err)
}

// GetLastTokenID returns id of latest nft for wait list.
func (service *Service) GetLastTokenID(ctx context.Context) (int, error) {
	lastID, err := service.waitList.GetLast(ctx)
	return lastID, ErrWaitlist.Wrap(err)
}

// ListWithoutPassword returns nft for wait list without password.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Item, error) {
	nftWithoutPassword, err := service.waitList.ListWithoutPassword(ctx)
	return nftWithoutPassword, ErrWaitlist.Wrap(err)
}

// Delete deletes nft for wait list.
func (service *Service) Delete(ctx context.Context, tokenIDs []int) error {
	return ErrWaitlist.Wrap(service.waitList.Delete(ctx, tokenIDs))
}
