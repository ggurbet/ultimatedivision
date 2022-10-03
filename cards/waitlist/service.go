// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/remotefilestorage/storj"
	"ultimatedivision/pkg/imageprocessing"
	"ultimatedivision/users"
)

// ErrWaitlist indicated that there was an error in service.
var ErrWaitlist = errs.Class("waitlist service error")

// Service is handling waitList related logic.
//
// architecture: Service
type Service struct {
	config   Config
	waitList DB
	cards    *cards.Service
	avatars  *avatars.Service
	users    *users.Service
	nfts     *nfts.Service
}

// NewService is a constructor for waitlist service.
func NewService(config Config, waitList DB, cards *cards.Service, avatars *avatars.Service, users *users.Service, nfts *nfts.Service) *Service {
	return &Service{
		config:   config,
		waitList: waitList,
		cards:    cards,
		avatars:  avatars,
		users:    users,
		nfts:     nfts,
	}
}

// Create creates nft for wait list.
func (service *Service) Create(ctx context.Context, createNFT CreateNFT) (Transaction, error) {
	var transaction Transaction

	user, err := service.users.Get(ctx, createNFT.UserID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if len(createNFT.WalletAddress.String()) == 0 {
		createNFT.WalletAddress = user.Wallet
	}

	card, err := service.cards.Get(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if createNFT.Value.Cmp(big.NewInt(0)) <= 0 {
		if card.UserID != createNFT.UserID {
			return transaction, ErrWaitlist.New("this card does not belongs to user")
		}
	}

	if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
		transaction = Transaction{
			Password:          item.Password,
			NFTCreateContract: service.config.NFTCreateContract,
			TokenID:           item.TokenID,
			Value:             item.Value,
			WalletType:        item.WalletType,
		}
		return transaction, nil
	}

	image, err := service.avatars.GetImage(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	client, err := storj.NewClient(service.config.FileStorage)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	// TODO: add transaction and mb lock db.
	lastTokenID, err := service.GetLastTokenID(ctx)
	if err != nil {
		if !ErrNoItem.Has(err) {
			return transaction, ErrWaitlist.Wrap(err)
		}
	}

	nextTokenID := lastTokenID + 1

	if err = client.Upload(ctx, service.config.Bucket, fmt.Sprintf("%d.%s", nextTokenID, imageprocessing.TypeFilePNG), image); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	nft := service.nfts.Generate(ctx, card, fmt.Sprintf(service.config.URLToAvatar, nextTokenID))
	fileMetadata, err := json.MarshalIndent(nft, "", " ")
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if err = client.Upload(ctx, service.config.Bucket, fmt.Sprintf("%d.%s", nextTokenID, imageprocessing.TypeFileJSON), fileMetadata); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if err = service.users.UpdateWalletAddress(ctx, createNFT.WalletAddress, createNFT.UserID, user.WalletType); err != nil {
		if !users.ErrWalletAddressAlreadyInUse.Has(err) {
			return transaction, ErrWaitlist.Wrap(err)
		}
	}

	item := Item{
		CardID:     createNFT.CardID,
		Wallet:     createNFT.WalletAddress,
		WalletType: user.WalletType,
		Value:      createNFT.Value,
	}
	if err = service.waitList.Create(ctx, item); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	for range time.NewTicker(time.Millisecond * service.config.WaitListCheckSignature).C {
		if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
			transaction = Transaction{
				Password:          item.Password,
				NFTCreateContract: service.config.NFTCreateContract,
				TokenID:           item.TokenID,
				Value:             item.Value,
				WalletType:        item.WalletType,
			}
			break
		}
	}

	return transaction, err
}

// GetByTokenID returns nft for wait list by token id.
func (service *Service) GetByTokenID(ctx context.Context, tokenID int64) (Item, error) {
	nft, err := service.waitList.GetByTokenID(ctx, tokenID)
	return nft, ErrWaitlist.Wrap(err)
}

// GetByCardID returns nft for wait list by card id.
func (service *Service) GetByCardID(ctx context.Context, cardID uuid.UUID) (Item, error) {
	nft, err := service.waitList.GetByCardID(ctx, cardID)
	return nft, ErrWaitlist.Wrap(err)
}

// GetLastTokenID returns id of latest nft for wait list.
func (service *Service) GetLastTokenID(ctx context.Context) (int64, error) {
	lastID, err := service.waitList.GetLastTokenID(ctx)
	return lastID, ErrWaitlist.Wrap(err)
}

// List returns all nft for wait list.
func (service *Service) List(ctx context.Context) ([]Item, error) {
	allNFT, err := service.waitList.List(ctx)
	return allNFT, ErrWaitlist.Wrap(err)
}

// ListWithoutPassword returns nft for wait list without password.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Item, error) {
	nftWithoutPassword, err := service.waitList.ListWithoutPassword(ctx)
	return nftWithoutPassword, ErrWaitlist.Wrap(err)
}

// Update updates signature to nft token.
func (service *Service) Update(ctx context.Context, tokenID int64, password evmsignature.Signature) error {
	return ErrWaitlist.Wrap(service.waitList.Update(ctx, tokenID, password))
}

// Delete deletes nft for wait list.
func (service *Service) Delete(ctx context.Context, tokenIDs []int64) error {
	return ErrWaitlist.Wrap(service.waitList.Delete(ctx, tokenIDs))
}
