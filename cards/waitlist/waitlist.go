// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrNoItem indicates that item for wait list does not exist.
var ErrNoItem = errs.Class("item for wait list does not exist")

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create creates nft for wait list in the database.
	Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address) error
	// Get returns nft for wait list by card id.
	Get(ctx context.Context, tokenID int) (Item, error)
	// GetLast returns id of last inserted token.
	GetLast(ctx context.Context) (int, error)
	// List returns all nft tokens from wait list from database.
	List(ctx context.Context) ([]Item, error)
	// ListWithoutPassword returns nfts for wait list without password from database.
	ListWithoutPassword(ctx context.Context) ([]Item, error)
	// Delete deletes nft from wait list by id of token.
	Delete(ctx context.Context, tokenIDs []int) error
	// Update updates signature to nft token.
	Update(ctx context.Context, tokenID int, password cryptoutils.Signature) error
}

// Item entity describes item fot wait list nfts.
type Item struct {
	TokenID  int                   `json:"tokenId"`
	CardID   uuid.UUID             `json:"cardId"`
	Wallet   cryptoutils.Address   `json:"wallet"`
	Password cryptoutils.Signature `json:"password"`
}

// CreateNFT describes body of request for creating nft token.
type CreateNFT struct {
	CardID        uuid.UUID           `json:"cardId"`
	WalletAddress cryptoutils.Address `json:"walletAddress"`
	UserID        uuid.UUID           `json:"userId"`
}
