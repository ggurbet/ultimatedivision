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

// RegularIsEthereumAddress indicated that expression is regular expression for ethereum address.
const RegularIsEthereumAddress = "^0x[0-9a-fA-F]{40}$"

// DB is exposing access to whitelist db.
//
// architecture: DB
type DB interface {
	// Create adds whitelist in the database.
	Create(ctx context.Context, whitelist Whitelist) error
	// GetByAddress returns whitelist by address from the database.
	GetByAddress(ctx context.Context, address Address) (Whitelist, error)
	// List returns all whitelist from the database.
	List(ctx context.Context) ([]Whitelist, error)
	// Delete deletes whitelist from the database.
	Delete(ctx context.Context, address Address) error
}

// Whitelist describes whitelist entity.
type Whitelist struct {
	Address  Address `json:"address"`
	Password []byte  `json:"password"`
}

// Address defines address of user's wallet.
type Address string

// ValidateAddress checks if the address is valid.
func (address Address) ValidateAddress() bool {
	return common.IsHexAddress(string(address))
}
