// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencywaitlist

import (
	"math/big"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrNoItem indicates that item of currency wait list does not exist.
var ErrNoItem = errs.Class("item of currency wait list does not exist")

// DB is exposing access to currencywaitlist db.
//
// architecture: DB
type DB interface{}

// Item entity describes item of currency wait list.
type Item struct {
	Wallet    cryptoutils.Address   `json:"wallet"`
	Value     big.Int               `json:"value"`
	Nonce     int64                 `json:"nonce"`
	Signature cryptoutils.Signature `json:"signature"`
}

// Transaction entity describes values for creating transaction to contract.
type Transaction struct {
	Signature cryptoutils.Signature `json:"signature"`
	Contract  cryptoutils.Contract  `json:"contract"`
	Value     big.Int               `json:"value"`
	Nonce     int64                 `json:"nonce"`
}
