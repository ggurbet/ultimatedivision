// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package bids

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoBid indicates that bid does not exist.
var ErrNoBid = errs.Class("bid does not exist")

// DB is exposing access to bids db.
//
// architecture: DB
type DB interface {
	// Create creates bid for lot in the database.
	Create(ctx context.Context, bid Bid) error
	// GetCurrentBidByLotID returns current bid by lot id from the database.
	GetCurrentBidByLotID(ctx context.Context, lotID uuid.UUID) (Bid, error)
	// ListByLotID returns bids by lot id from the database.
	ListByLotID(ctx context.Context, lotID uuid.UUID) ([]Bid, error)
	// ListByUserID returns bids by user id from the database.
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]Bid, error)
	// GetUserBidsAmountByLotID returns amount of user last bet on certain lot form the database.
	GetUserBidsAmountByLotID(ctx context.Context, userID, lotID uuid.UUID) ([]big.Int, error)
	// DeleteByLotID deletes bids by lot id in the database.
	DeleteByLotID(ctx context.Context, lotID uuid.UUID) error
}

// Bid describes bids placed on a specific lot.
type Bid struct {
	ID        uuid.UUID `json:"id"`
	LotID     uuid.UUID `json:"lotId"`
	UserID    uuid.UUID `json:"userId"`
	UserName  string    `json:"userName"`
	Amount    big.Int   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

// Compare compares two bids.
func (b Bid) Compare(bidToCompare Bid) bool {
	return b.ID == bidToCompare.ID && b.LotID == bidToCompare.LotID && b.UserID == bidToCompare.UserID &&
		b.UserName == bidToCompare.UserName && b.Amount.Cmp(&bidToCompare.Amount) == 0 && b.CreatedAt == bidToCompare.CreatedAt
}
