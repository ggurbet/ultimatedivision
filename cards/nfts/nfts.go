// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrNoNFT indicates that nft token does not exist.
var ErrNoNFT = errs.Class("nft does not exist")

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create creates nft token in the database.
	Create(ctx context.Context, nft NFT) error
	// List returns all nft token from database.
	List(ctx context.Context) ([]NFT, error)
}

// NFT entity describes values released nft token.
type NFT struct {
	TokenID       int                 `json:"tokenId"`
	Сhain         Сhain               `json:"chain"`
	CardID        uuid.UUID           `json:"cardId"`
	WalletAddress cryptoutils.Address `json:"walletAddress"`
}

// Сhain defines the list of possible chains in blockchain.
type Сhain string

const (
	// СhainEthereum indicates that chain is ethereum.
	СhainEthereum Сhain = "ethereum"
	// СhainPolygon indicates that chain is polygon.
	СhainPolygon Сhain = "polygon"
)

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100

// Config defines values needed by create nft.
type Config struct {
	Description string `json:"description"`
	ExternalURL string `json:"externalUrl"`
}
