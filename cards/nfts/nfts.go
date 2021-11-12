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
	Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address) error
	// Get returns nft token by card id.
	Get(ctx context.Context, tokenID int) (NFTWaitListItem, error)
	// GetLast returns id of last inserted token.
	GetLast(ctx context.Context) (int, error)
	// List returns all nft token from wait list from database.
	List(ctx context.Context) ([]NFTWaitListItem, error)
	// ListWithoutPassword returns all nft tokens without password from database.
	ListWithoutPassword(ctx context.Context) ([]NFTWaitListItem, error)
	// Delete deletes nft from wait list by id of token.
	Delete(ctx context.Context, tokenIDs []int) error
}

// NFTWaitListItem describes list of nft tokens entity.
type NFTWaitListItem struct {
	TokenID  int                   `json:"tokenId"`
	CardID   uuid.UUID             `json:"cardId"`
	Wallet   cryptoutils.Address   `json:"wallet"`
	Password cryptoutils.Signature `json:"password"`
}

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100

// Config defines values needed by create nft.
type Config struct {
	Description string `json:"description"`
	ExternalURL string `json:"externalUrl"`
}
