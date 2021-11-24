// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrNoNFTs indicated that nfts does not exist.
var ErrNoNFTs = errs.Class("nfts does not exist")

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create creates nft token in the database.
	Create(ctx context.Context, nft NFT) error
	// Get returns nft by token id and chain from database.
	Get(ctx context.Context, tokenID int64, chain cryptoutils.Chain) (NFT, error)
	// List returns all nft token from database.
	List(ctx context.Context) ([]NFT, error)
	// Update updates users wallet address for nft token in the database.
	Update(ctx context.Context, nft NFT) error
	// Delete deletes nft token in the database.
	Delete(ctx context.Context, cardID uuid.UUID) error
}

// NFT entity describes values released nft token.
type NFT struct {
	CardID        uuid.UUID           `json:"cardId"`
	TokenID       int64               `json:"tokenId"`
	Chain         cryptoutils.Chain   `json:"chain"`
	WalletAddress cryptoutils.Address `json:"walletAddress"`
}

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100

// Config defines values needed by create nft.
type Config struct {
	Description        string               `json:"description"`
	ExternalURL        string               `json:"externalUrl"`
	NFTRenewalInterval time.Duration        `json:"nftRenewalInterval"`
	Contract           cryptoutils.Contract `json:"contract"`
	AddressNodeServer  string               `json:"addressNodeServer"`
}
