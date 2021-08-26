// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoLot indicated that lot does not exist.
var ErrNoLot = errs.Class("lot does not exist")

// ErrMarketplace indicated that there was an error in service.
var ErrMarketplace = errs.Class("marketplace service error")

// DB is exposing access to lots db.
//
// architecture: DB
type DB interface {
	// CreateLot add lot in the data base.
	CreateLot(ctx context.Context, lot Lot) error
	// GetLotByID returns lot by id from the data base.
	GetLotByID(ctx context.Context, id uuid.UUID) (Lot, error)
	// ListActiveLots returns active lots from the data base.
	ListActiveLots(ctx context.Context) ([]Lot, error)
	// ListExpiredLot returns active lots where end time lower than or equal to time now UTC from the data base.
	ListExpiredLot(ctx context.Context) ([]Lot, error)
	// UpdateShopperIDLot updates shopper id of lot in the database.
	UpdateShopperIDLot(ctx context.Context, id, shopperID uuid.UUID) error
	// UpdateStatusLot updates status of lot in the database.
	UpdateStatusLot(ctx context.Context, id uuid.UUID, status Status) error
	// UpdateCurrentPriceLot updates current price of lot in the database.
	UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice float64) error
	// UpdateEndTimeLot updates end time of lot in the database.
	UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error
}

// Lot describes lot entity.
type Lot struct {
	ID           uuid.UUID `json:"id"`
	ItemID       uuid.UUID `json:"itemId"`
	Type         Type      `json:"type"`
	UserID       uuid.UUID `json:"userId"`
	ShopperID    uuid.UUID `json:"shopperId"`
	Status       Status    `json:"status"`
	StartPrice   float64   `json:"startPrice"`
	MaxPrice     float64   `json:"maxPrice"`
	CurrentPrice float64   `json:"currentPrice"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Period       Period    `json:"period"`
}

// Type defines the list of possible lot types.
type Type string

const (
	// TypeCard indicates that lot type is card.
	TypeCard Type = "card"
)

// Status defines the list of possible lot statuses.
type Status string

const (
	// StatusActive indicates that lot is active that is, the lot is sold at auction.
	StatusActive Status = "active"
	// StatusExpired indicates that the time of lot has expired.
	StatusExpired Status = "expired"
	// StatusSold indicates that the lot was sold after the expiration of the lot at the highest rate.
	StatusSold Status = "sold"
	// StatusSoldBuynow indicates that the lot was purchased at the maximum price.
	StatusSoldBuynow Status = "soldBuynow"
)

// Period defines the list of possible lot periods.
type Period int

const (
	// MinPeriod indicates that lot minimal period time is 1 hour.
	MinPeriod Period = 1
	// MaxPeriod indicates that lot maximal period time is 120 hour.
	MaxPeriod Period = 120
)

// Config defines configuration for marketplace.
type Config struct {
	LotRenewalInterval time.Duration `json:"lotRenewalInterval"`
}

// CreateLot entity that contains the values required to create the lot.
type CreateLot struct {
	ItemID     uuid.UUID `json:"itemId"`
	Type       Type      `json:"type"`
	StartPrice float64   `json:"startPrice"`
	MaxPrice   float64   `json:"maxPrice"`
	Period     Period    `json:"period"`
}

// BetLot entity that contains the values required to place bet the lot.
type BetLot struct {
	ID        uuid.UUID `json:"id"`
	BetAmount float64   `json:"betAmount"`
}

// WinLot entity that contains the values required to win the lot.
type WinLot struct {
	ID        uuid.UUID `json:"id"`
	ItemID    uuid.UUID `json:"itemId"`
	Type      Type      `json:"type"`
	UserID    uuid.UUID `json:"userId"`
	ShopperID uuid.UUID `json:"shopperID"`
	Status    Status    `json:"status"`
	Amount    float64   `json:"amount"`
}

// ValidateCreateLot check is empty fields of create lot entity.
func (createLot CreateLot) ValidateCreateLot() error {
	if createLot.ItemID.String() == "" {
		return ErrMarketplace.New("item id is empty")
	}

	if createLot.StartPrice == 0 {
		return ErrMarketplace.New("start price is empty")
	}

	if createLot.Period == 0 {
		return ErrMarketplace.New("period is empty")
	}

	return nil
}

// ValidateBetLot check is empty fields of bet lot entity.
func (betLot BetLot) ValidateBetLot() error {
	if betLot.ID.String() == "" {
		return ErrMarketplace.New("lot id is empty")
	}

	if betLot.BetAmount == 0 {
		return ErrMarketplace.New("bet amount is empty")
	}

	return nil
}
