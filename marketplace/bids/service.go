// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package bids

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/marketplace"
	"ultimatedivision/users"
)

var (
	// ErrBids indicates that there was an error in the service.
	ErrBids = errs.Class("bids service error")
	// ErrSmallAmountOfBid indicates that bid amount is small.
	ErrSmallAmountOfBid = errs.New("the amount of the bet is too small")
)

// Service is handling bids related logic.
//
// architecture: Service
type Service struct {
	bids        DB
	marketplace *marketplace.Service
	cards       *cards.Service
	users       *users.Service
}

// NewService is constructor for Service.
func NewService(bids DB, marketplace *marketplace.Service, cards *cards.Service, users *users.Service) *Service {
	return &Service{
		bids:        bids,
		marketplace: marketplace,
		cards:       cards,
		users:       users,
	}
}

// Create creates bid for lot in the database.
func (service *Service) Create(ctx context.Context, bid Bid) error {
	lot, err := service.marketplace.GetLotByID(ctx, bid.LotID)
	if err != nil {
		return ErrBids.Wrap(err)
	}
	if _, err := service.cards.Get(ctx, lot.CardID); err != nil {
		return ErrBids.Wrap(err)
	}

	currentBid, err := service.GetCurrentBidByLotID(ctx, lot.CardID)
	if err != nil && !ErrNoBid.Has(err) {
		return ErrBids.Wrap(err)
	}
	if ErrNoBid.Has(err) {
		if bid.Amount.Cmp(&bid.Amount) < 0 {
			return ErrSmallAmountOfBid
		}
	}

	if currentBid.Amount.Cmp(&bid.Amount) >= 0 {
		return ErrSmallAmountOfBid
	}

	// TODO: check if user have this amount of the money.
	_, err = service.GetBidsAmountOnSpecificLot(ctx, bid.UserID, bid.LotID)
	if err != nil {
		return ErrBids.Wrap(err)
	}

	bid = Bid{
		ID:        uuid.New(),
		LotID:     bid.LotID,
		UserID:    bid.UserID,
		Amount:    bid.Amount,
		CreatedAt: time.Now().UTC(),
	}
	if err = service.bids.Create(ctx, bid); err != nil {
		return ErrBids.Wrap(err)
	}

	return nil
}

// GetCurrentBidByLotID returns current bid by lot id from the database.
func (service *Service) GetCurrentBidByLotID(ctx context.Context, lotID uuid.UUID) (Bid, error) {
	currentAmount, err := service.bids.GetCurrentBidByLotID(ctx, lotID)
	return currentAmount, ErrBids.Wrap(err)
}

// GetBidsAmountOnSpecificLot returns all amount of user bets on specific lot.
func (service *Service) GetBidsAmountOnSpecificLot(ctx context.Context, userID, lotID uuid.UUID) ([]big.Int, error) {
	amounts, err := service.bids.GetUserBidsAmountByLotID(ctx, userID, lotID)

	return amounts, ErrBids.Wrap(err)
}

// ListByUserID returns bids by user id from the database.
func (service *Service) ListByUserID(ctx context.Context, userID uuid.UUID) ([]Bid, error) {
	bids, err := service.bids.ListByUserID(ctx, userID)
	return bids, ErrBids.Wrap(err)
}

// ListByLotID returns bids by lot id from the database.
func (service *Service) ListByLotID(ctx context.Context, lotID uuid.UUID) ([]Bid, error) {
	bids, err := service.bids.ListByLotID(ctx, lotID)
	if err != nil {
		return bids, ErrBids.Wrap(err)
	}

	for i := 0; i < len(bids); i++ {
		bids[i].UserName, err = service.users.GetNickNameByID(ctx, bids[i].UserID)
		if err != nil {
			return bids, ErrBids.Wrap(err)
		}
	}

	return bids, nil
}
