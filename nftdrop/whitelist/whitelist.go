// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrNoWhitelist indicated that whitelist does not exist.
var ErrNoWhitelist = errs.Class("whitelist does not exist")

// DB is exposing access to whitelist db.
//
// architecture: DB
type DB interface {
	// Create adds wallet in the database.
	Create(ctx context.Context, wallet Wallet) error
	// GetByAddress returns wallet by address from the database.
	GetByAddress(ctx context.Context, address cryptoutils.Address) (Wallet, error)
	// List returns all wallets from the database.
	List(ctx context.Context) ([]Wallet, error)
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
	SmartContractAddress `json:"smartContractAddress"`
}

// SmartContractAddress entity describes smart contract addresses.
type SmartContractAddress struct {
	NFT     cryptoutils.Address `json:"nft"`
	NFTSale cryptoutils.Address `json:"nftSale"`
}

// Transaction entity describes password wallet and smart contract addresses.
type Transaction struct {
	Password             cryptoutils.Signature `json:"password"`
	SmartContractAddress `json:"smartContractAddress"`
}
