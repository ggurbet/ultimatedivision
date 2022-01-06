// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencywaitlist

import (
	"context"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/udts"
	"ultimatedivision/users"
)

// ErrCurrencyWaitlist indicated that there was an error in service.
var ErrCurrencyWaitlist = errs.Class("currency waitlist service error")

// Service is handling currency wait list related logic.
//
// architecture: Service
type Service struct {
	config           Config
	currencyWaitList DB
	users            *users.Service
	udts             *udts.Service
}

// NewService is a constructor for currencyWaitList service.
func NewService(config Config, currencyWaitList DB, users *users.Service, udts *udts.Service) *Service {
	return &Service{
		config:           config,
		currencyWaitList: currencyWaitList,
		users:            users,
		udts:             udts,
	}
}

// Create creates item of currency wait list.
func (service *Service) Create(ctx context.Context, userID uuid.UUID, value big.Int, nonce int64) (Transaction, error) {
	var transaction Transaction

	user, err := service.users.Get(ctx, userID)
	if err != nil {
		return transaction, ErrCurrencyWaitlist.Wrap(err)
	}

	item := Item{
		WalletAddress: user.Wallet,
		Value:         value,
		Nonce:         nonce,
		Signature:     "",
	}

	// TODO: catch dublicale error from db
	if _, err = service.currencyWaitList.GetByWalletAddressAndNonce(ctx, item.WalletAddress, item.Nonce); err != nil {
		if ErrNoItem.Has(err) {
			if err = service.currencyWaitList.Create(ctx, item); err != nil {
				return transaction, ErrCurrencyWaitlist.Wrap(err)
			}
		}
	}

	if err = service.Update(ctx, item); err != nil {
		return transaction, ErrCurrencyWaitlist.Wrap(err)
	}

	for range time.NewTicker(time.Millisecond * service.config.IntervalSignatureCheck).C {
		if item, err := service.GetByWalletAddressAndNonce(ctx, item.WalletAddress, item.Nonce); item.Signature != "" && err == nil {
			transaction = Transaction{
				Signature:   item.Signature,
				UDTContract: service.config.UDTContract,
				Value:       item.Value.String(),
			}
			break
		}
	}

	return transaction, err
}

// GetByWalletAddressAndNonce returns item of currency wait list by wallet address and nonce.
func (service *Service) GetByWalletAddressAndNonce(ctx context.Context, walletAddress evmsignature.Address, nonce int64) (Item, error) {
	item, err := service.currencyWaitList.GetByWalletAddressAndNonce(ctx, walletAddress, nonce)
	return item, ErrCurrencyWaitlist.Wrap(err)
}

// List returns items of currency wait list.
func (service *Service) List(ctx context.Context) ([]Item, error) {
	items, err := service.currencyWaitList.List(ctx)
	return items, ErrCurrencyWaitlist.Wrap(err)
}

// ListWithoutSignature returns items of currency waitlist without signature from database.
func (service *Service) ListWithoutSignature(ctx context.Context) ([]Item, error) {
	items, err := service.currencyWaitList.ListWithoutSignature(ctx)
	return items, ErrCurrencyWaitlist.Wrap(err)
}

// UpdateSignature updates signature of item by wallet address and nonce.
func (service *Service) UpdateSignature(ctx context.Context, signature evmsignature.Signature, walletAddress evmsignature.Address, nonce int64) error {
	return ErrCurrencyWaitlist.Wrap(service.currencyWaitList.UpdateSignature(ctx, signature, walletAddress, nonce))
}

// Update updates item by wallet address and nonce.
func (service *Service) Update(ctx context.Context, item Item) error {
	return ErrCurrencyWaitlist.Wrap(service.currencyWaitList.Update(ctx, item))
}

// Delete deletes item of currency wait list by wallet address and nonce.
func (service *Service) Delete(ctx context.Context, walletAddress evmsignature.Address, nonce int64) error {
	return ErrCurrencyWaitlist.Wrap(service.currencyWaitList.Delete(ctx, walletAddress, nonce))
}
