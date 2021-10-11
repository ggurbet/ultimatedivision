// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"
)

// ErrNoWhitelist indicated that whitelist does not exist.
var ErrNoWhitelist = errs.Class("whitelist does not exist")

// DB is exposing access to whitelist db.
//
// architecture: DB
type DB interface {
	// Create adds whitelist in the database.
	Create(ctx context.Context, wallet Wallet) error
	// GetByAddress returns whitelist by address from the database.
	GetByAddress(ctx context.Context, address Hex) (Wallet, error)
	// List returns all whitelist from the database.
	List(ctx context.Context) ([]Wallet, error)
	// ListWithoutPassword returns whitelist without password from the database.
	ListWithoutPassword(ctx context.Context) ([]Wallet, error)
	// Update updates whitelist by address.
	Update(ctx context.Context, wallet Wallet) error
	// Delete deletes whitelist from the database.
	Delete(ctx context.Context, address Hex) error
}

// Wallet describes whitelist entity.
type Wallet struct {
	Address  Hex    `json:"address"`
	Password []byte `json:"password"`
}

// Hex defines hex type.
type Hex string

// IsValidAddress checks if the address is valid.
func (hex Hex) IsValidAddress() bool {
	return common.IsHexAddress(string(hex))
}

// IsHex validates whether each byte is valid hexadecimal string.
func (hex Hex) IsHex() bool {
	if len(string(hex))%2 != 0 {
		return false
	}
	for _, c := range []byte(string(hex)) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// CreateWallet entity describes request values for create whitelist.
type CreateWallet struct {
	Address    Hex `json:"address"`
	PrivateKey Hex `json:"privateKey"`
}

// Config defines configuration for queue.
type Config struct {
	SmartContract struct {
		Address string  `json:"address"`
		Price   float64 `json:"price"`
	} `json:"smartContract"`
}

// SmartContractWithWhiteList entity describes whitelist and smart contract.
type SmartContractWithWhiteList struct {
	Wallet
	Address string  `json:"address"`
	Price   float64 `json:"price"`
}
