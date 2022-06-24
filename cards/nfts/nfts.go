// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoNFT indicated that nft does not exist.
var ErrNoNFT = errs.Class("nft does not exist")

// DB is exposing access to nfts db.
//
// architecture: DB
type DB interface {
	// Create creates nft token in the database.
	Create(ctx context.Context, nft NFT) error
	// Get returns nft by token id and chain from database.
	Get(ctx context.Context, tokenID int64, chain evmsignature.Chain) (NFT, error)
	// List returns all nft token from database.
	List(ctx context.Context) ([]NFT, error)
	// Update updates users wallet address for nft token in the database.
	Update(ctx context.Context, nft NFT) error
	// Delete deletes nft token in the database.
	Delete(ctx context.Context, cardID uuid.UUID) error
}

// NFT entity describes values released nft token.
type NFT struct {
	CardID        uuid.UUID          `json:"cardId"`
	TokenID       int64              `json:"tokenId"`
	Chain         evmsignature.Chain `json:"chain"`
	WalletAddress common.Address     `json:"walletAddress"`
}

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100

// Config defines values needed by create nft.
type Config struct {
	Description        string        `json:"description"`
	ExternalURL        string        `json:"externalUrl"`
	NFTRenewalInterval time.Duration `json:"nftRenewalInterval"`
	NFTContract        struct {
		Address         common.Address   `json:"address"`
		OwnerOfSelector evmsignature.Hex `json:"ownerOfSelector"`
	} `json:"nftContract"`
	AddressNodeServer string `json:"addressNodeServer"`
}
