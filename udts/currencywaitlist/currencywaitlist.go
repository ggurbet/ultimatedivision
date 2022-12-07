// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencywaitlist

import (
	"context"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

	"ultimatedivision/users"
)

// ErrNoItem indicates that item of currency wait list does not exist.
var ErrNoItem = errs.Class("item of currency wait list does not exist")

// DB is exposing access to currencywaitlist db.
//
// architecture: DB.
type DB interface {
	// Create creates item of currency waitlist in the database.
	Create(ctx context.Context, item Item) error
	// GetByWalletAddressAndNonce returns item of currency wait list by wallet address and nonce.
	GetByWalletAddressAndNonce(ctx context.Context, walletAddress common.Address, nonce int64) (Item, error)
	// GetByCasperWalletAddressAndNonce returns item of currency wait list by casper wallet address and nonce.
	GetByCasperWalletAddressAndNonce(ctx context.Context, casperWallet string, nonce int64) (Item, error)
	// List returns items of currency waitlist from database.
	List(ctx context.Context) ([]Item, error)
	// GetNonce returns number of nonce from database.
	GetNonce(ctx context.Context) (int64, error)
	// GetNonceByWallet returns number of nonce by wallet from database.
	GetNonceByWallet(ctx context.Context, wallet string) (int64, error)
	// ListWithoutSignature returns items of currency waitlist without signature from database.
	ListWithoutSignature(ctx context.Context) ([]Item, error)
	// Update updates item by wallet address and nonce in the database.
	Update(ctx context.Context, item Item) error
	// UpdateNonceByWallet updates nonce by wallet address in the database.
	UpdateNonceByWallet(ctx context.Context, nonce int64, casperWallet string) error
	// UpdateSignature updates signature of item by wallet address and nonce in the database.
	UpdateSignature(ctx context.Context, signature evmsignature.Signature, walletAddress common.Address, nonce int64) error
	// UpdateCasperSignature updates casper signature of item by wallet address and nonce in the database.
	UpdateCasperSignature(ctx context.Context, signature evmsignature.Signature, casperWallet string, nonce int64) error
	// Delete deletes item of currency waitlist by wallet address and nonce in the database.
	Delete(ctx context.Context, walletAddress common.Address, nonce int64) error
}

// Item entity describes item of currency wait list.
type Item struct {
	WalletAddress       common.Address         `json:"walletAddress"`
	CasperWalletAddress string                 `json:"casperWalletAddress"`
	WalletType          users.WalletType       `json:"walleType"`
	Value               big.Int                `json:"value"`
	Nonce               int64                  `json:"nonce"`
	Signature           evmsignature.Signature `json:"signature"`
}

// Transaction entity describes values for creating transaction to contract.
type Transaction struct {
	Signature   evmsignature.Signature `json:"signature"`
	UDTContract evmsignature.Contract  `json:"udtContract"`
	Value       string                 `json:"value"`
}

// CasperTransaction entity describes values for creating transaction to contract.
type CasperTransaction struct {
	Signature           evmsignature.Signature `json:"signature"`
	CasperTokenContract evmsignature.Contract  `json:"casperTokenContract"`
	Value               string                 `json:"value"`
	Nonce               int64                  `json:"nonce"`
}

// Config defines values needed by mint udt tokens in blockchain.
type Config struct {
	IntervalSignatureCheck time.Duration         `json:"intervalSignatureCheck"`
	UDTContract            evmsignature.Contract `json:"udtContract"`
	CasperTokenContract    evmsignature.Contract `json:"casperTokenContract"`
}
