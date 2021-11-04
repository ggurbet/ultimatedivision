// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/pkg/pagination"
)

// ErrNoWallet indicated that wallet in whitelist does not exist.
var ErrNoWallet = errs.Class("wallet does not exist")

// DB is exposing access to whitelist db.
//
// architecture: DB
type DB interface {
	// Create adds wallet in the database.
	Create(ctx context.Context, wallet Wallet) error
	// GetByAddress returns whitelist by address from the database.
	GetByAddress(ctx context.Context, address cryptoutils.Address) (Wallet, error)
	// List returns whitelist page from the database.
	List(ctx context.Context, cursor pagination.Cursor) (Page, error)
	// ListWithoutPassword returns wallet without password from the database.
	ListWithoutPassword(ctx context.Context) ([]Wallet, error)
	// Update updates wallet by address.
	Update(ctx context.Context, wallet Wallet) error
	// Delete deletes wallet from the database.
	Delete(ctx context.Context, address cryptoutils.Address) error
}

// Wallet describes whitelist entity.
type Wallet struct {
	Address  cryptoutils.Address   `json:"address"`
	Password cryptoutils.Signature `json:"password"`
}

// CreateWallet entity describes request values for create whitelist.
type CreateWallet struct {
	Address    cryptoutils.Address    `json:"address"`
	PrivateKey cryptoutils.PrivateKey `json:"privateKey"`
}

// Config defines configuration for queue.
type Config struct {
	pagination.Cursor `json:"cursor"`
	NFTSaleContract   cryptoutils.Address `json:"nftSaleContract"`
}

// Transaction entity describes password wallet and smart contracts address.
type Transaction struct {
	Password        cryptoutils.Signature `json:"password"`
	NFTSaleContract cryptoutils.Address   `json:"nftSaleContract"`
}

// Page holds wallets page entity which is used to show listed page of wallets.
type Page struct {
	Wallets []Wallet        `json:"wallets"`
	Page    pagination.Page `json:"page"`
}
