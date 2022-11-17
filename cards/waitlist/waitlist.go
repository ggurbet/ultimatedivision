// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/remotefilestorage/storj"
	"ultimatedivision/users"
)

// ErrNoItem indicates that item for wait list does not exist.
var ErrNoItem = errs.Class("item for wait list does not exist")

// DB is exposing access to waitlist db.
//
// architecture: DB
type DB interface {
	// Create creates nft for wait list in the database.
	Create(ctx context.Context, item Item) error
	// Get returns nft for wait list by token id.
	GetByTokenID(ctx context.Context, tokenID int64) (Item, error)
	// GetByCardID returns nft for wait list by card id.
	GetByCardID(ctx context.Context, cardID uuid.UUID) (Item, error)
	// GetLastTokenID returns id of last inserted token.
	GetLastTokenID(ctx context.Context) (int64, error)
	// List returns all nft tokens from wait list from database.
	List(ctx context.Context) ([]Item, error)
	// ListWithoutPassword returns nfts for wait list without password from database.
	ListWithoutPassword(ctx context.Context) ([]Item, error)
	// Delete deletes nft from wait list by id of token.
	Delete(ctx context.Context, tokenIDs []int64) error
	// Update updates signature to nft token.
	Update(ctx context.Context, tokenID int64, password evmsignature.Signature) error
}

// Item entity describes item fot wait list nfts.
type Item struct {
	TokenID      int64                  `json:"tokenId"`
	CardID       uuid.UUID              `json:"cardId"`
	Wallet       common.Address         `json:"wallet"`
	CasperWallet string                 `json:"casperWallet"`
	WalletType   users.WalletType       `json:"walletType"`
	Value        big.Int                `json:"value"`
	Password     evmsignature.Signature `json:"password"`
}

// CreateNFT describes body of request for creating nft token.
type CreateNFT struct {
	CardID        uuid.UUID      `json:"cardId"`
	CasperWallet  string         `json:"casperWallet"`
	WalletAddress common.Address `json:"walletAddress"`
	UserID        uuid.UUID      `json:"userId"`
	Value         big.Int        `json:"value"`
}

// Transaction entity describes values required to sent transaction.
type Transaction struct {
	Password                evmsignature.Signature  `json:"password"`
	NFTCreateContract       NFTCreateContract       `json:"nftCreateContract"`
	NFTCreateCasperContract NFTCreateCasperContract `json:"nftCreateCasperContract"`
	TokenID                 int64                   `json:"tokenId"`
	Value                   big.Int                 `json:"value"`
	WalletType              users.WalletType        `json:"walletType"`
	RPCNodeAddress          string                  `json:"rpcNodeAddress"`
}

// Config defines values needed by check mint nft in blockchain.
type Config struct {
	WaitListRenewalInterval time.Duration `json:"waitListRenewalInterval"`
	WaitListCheckSignature  time.Duration `json:"waitListCheckSignature"`
	NFTContract             struct {
		Address      common.Address   `json:"address"`
		AddressEvent evmsignature.Hex `json:"addressEvent"`
	} `json:"nftContract"`
	NFTCreateContract       NFTCreateContract       `json:"nftCreateContract"`
	NFTCreateVelasContract  NFTCreateVelasContract  `json:"nftCreateVelasContract"`
	NFTCreateCasperContract NFTCreateCasperContract `json:"nftCreateCasperContract"`
	AddressNodeServer       string                  `json:"addressNodeServer"`
	FileStorage             storj.Config            `json:"fileStorage"`
	Bucket                  string                  `json:"bucket"`
	URLToAvatar             string                  `json:"urlToAvatar"`
	RPCNodeAddress          string                  `json:"rpcNodeAddress"`
}

// NFTCreateContract describes the meaning of the contract.
type NFTCreateContract struct {
	Address                           common.Address   `json:"address"`
	MintWithSignatureSelector         evmsignature.Hex `json:"mintWithSignatureSelector"`
	MintWithSignatureAndValueSelector evmsignature.Hex `json:"mintWithSignatureAndValueSelector"`
	ChainID                           int              `json:"chainId"`
}

// NFTCreateVelasContract describes the meaning of the contract.
type NFTCreateVelasContract struct {
	Address                           common.Address   `json:"address"`
	MintWithSignatureSelector         evmsignature.Hex `json:"mintWithSignatureSelector"`
	MintWithSignatureAndValueSelector evmsignature.Hex `json:"mintWithSignatureAndValueSelector"`
	ChainID                           int              `json:"chainId"`
}

// NFTCreateCasperContract describes the meaning of the contract.
type NFTCreateCasperContract struct {
	Address                           string           `json:"address"`
	MintWithSignatureSelector         evmsignature.Hex `json:"mintWithSignatureSelector"`
	MintWithSignatureAndValueSelector evmsignature.Hex `json:"mintWithSignatureAndValueSelector"`
	ChainID                           int              `json:"chainId"`
}
