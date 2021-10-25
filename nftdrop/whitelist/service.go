// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/pkg/pagination"
)

// ErrWhitelist indicated that there was an error in service.
var ErrWhitelist = errs.Class("whitelist service error")

// Service is handling whitelist related logic.
//
// architecture: Service
type Service struct {
	config    Config
	whitelist DB
}

// NewService is a constructor for whitelist service.
func NewService(config Config, whitelist DB) *Service {
	return &Service{
		config:    config,
		whitelist: whitelist,
	}
}

// Create adds whitelist in the database.
func (service *Service) Create(ctx context.Context, wallet CreateWallet) error {
	var password cryptoutils.Signature

	if wallet.PrivateKey != "" {
		privateKeyECDSA, err := crypto.HexToECDSA(string(wallet.PrivateKey))
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}

		password, err = cryptoutils.GenerateSignature(wallet.Address, service.config.NFTSale, privateKeyECDSA)
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}
	}

	whitelist := Wallet{
		Address:  wallet.Address,
		Password: password,
	}
	return ErrWhitelist.Wrap(service.whitelist.Create(ctx, whitelist))
}

// GetByAddress returns whitelist by address from the database.
func (service *Service) GetByAddress(ctx context.Context, address cryptoutils.Address) (Transaction, error) {
	whitelist, err := service.whitelist.GetByAddress(ctx, address)
	if err != nil {
		return Transaction{}, ErrWhitelist.Wrap(err)
	}
	if whitelist.Password == "" {
		return Transaction{}, ErrWhitelist.New("password is empty")
	}

	transactionValue := Transaction{
		Password: whitelist.Password,
		SmartContractAddress: SmartContractAddress{
			NFT:     service.config.NFT,
			NFTSale: service.config.NFTSale,
		},
	}

	return transactionValue, nil
}

// List returns whitelist page from the database.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	whitelistRecords, err := service.whitelist.List(ctx, cursor)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}

// ListWithoutPassword returns whitelist without password from the database.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Wallet, error) {
	whitelistRecords, err := service.whitelist.ListWithoutPassword(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}

// Update updates whitelist by address.
func (service *Service) Update(ctx context.Context, whitelist Wallet) error {
	return ErrWhitelist.Wrap(service.whitelist.Update(ctx, whitelist))
}

// Delete deletes whitelist.
func (service *Service) Delete(ctx context.Context, address cryptoutils.Address) error {
	return ErrWhitelist.Wrap(service.whitelist.Delete(ctx, address))
}

// SetPassword generates passwords for all whitelist items.
func (service *Service) SetPassword(ctx context.Context, privateKey cryptoutils.PrivateKey) error {
	privateKeyECDSA, err := crypto.HexToECDSA(string(privateKey))
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	whitelist, err := service.ListWithoutPassword(ctx)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	for _, v := range whitelist {
		password, err := cryptoutils.GenerateSignature(v.Address, service.config.NFTSale, privateKeyECDSA)
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}

		whitelist := Wallet{
			Address:  v.Address,
			Password: password,
		}
		if err := service.Update(ctx, whitelist); err != nil {
			return ErrWhitelist.Wrap(err)
		}
	}

	return nil
}
