// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"time"

	"github.com/google/uuid"

	"ultimatedivision/cards"
	"ultimatedivision/users"
)

// Service is handling marketplace related logic.
//
// architecture: Service
type Service struct {
	marketplace DB
	users       *users.Service
	cards       *cards.Service
}

// NewService is a constructor for marketplace service.
func NewService(marketplace DB, users *users.Service, cards *cards.Service) *Service {
	return &Service{
		marketplace: marketplace,
		users:       users,
		cards:       cards,
	}
}

// CreateLot add lot in DB.
func (service *Service) CreateLot(ctx context.Context, userID uuid.UUID, createLot CreateLot) error {
	// TODO: add transaction
	card, err := service.cards.Get(ctx, createLot.ItemID)
	if err == nil {
		if card.UserID != userID {
			return ErrMarketplace.New("it is not the user's card")
		}

		if card.Status == cards.StatusSale {
			return ErrMarketplace.New("the card is already on sale")
		}

		if err := service.cards.UpdateStatus(ctx, createLot.ItemID, cards.StatusSale); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		createLot.Type = TypeCard
	}
	// TODO: check other items

	if createLot.Type == "" {
		return ErrMarketplace.New("not found item by id")
	}

	if _, err := service.users.Get(ctx, userID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if createLot.MaxPrice != 0 && createLot.MaxPrice < createLot.StartPrice {
		return ErrMarketplace.New("max price less start price")
	}

	if createLot.Period < MinPeriod && createLot.Period < MaxPeriod {
		return ErrMarketplace.New("period exceed the range from 1 to 120 hours")
	}

	lot := Lot{
		ID:         uuid.New(),
		ItemID:     createLot.ItemID,
		Type:       createLot.Type,
		UserID:     userID,
		Status:     StatusActive,
		StartPrice: createLot.StartPrice,
		MaxPrice:   createLot.MaxPrice,
		StartTime:  time.Now().UTC(),
		EndTime:    time.Now().UTC().Add(time.Duration(createLot.Period) * time.Hour),
		Period:     createLot.Period,
	}

	return ErrMarketplace.Wrap(service.marketplace.CreateLot(ctx, lot))
}

// GetLotByID returns lot by id from DB.
func (service *Service) GetLotByID(ctx context.Context, id uuid.UUID) (Lot, error) {
	lot, err := service.marketplace.GetLotByID(ctx, id)

	return lot, ErrMarketplace.Wrap(err)
}

// ListActiveLots returns active lots from DB.
func (service *Service) ListActiveLots(ctx context.Context) ([]Lot, error) {
	lots, err := service.marketplace.ListActiveLots(ctx)

	return lots, ErrMarketplace.Wrap(err)
}

// ListExpiredLot returns active lots from DB.
func (service *Service) ListExpiredLot(ctx context.Context) ([]Lot, error) {
	lots, err := service.marketplace.ListExpiredLot(ctx)
	return lots, ErrMarketplace.Wrap(err)
}

// PlaceBetLot checks the amount of money and makes a bet.
func (service *Service) PlaceBetLot(ctx context.Context, userID uuid.UUID, betLot BetLot) error {
	if _, err := service.users.Get(ctx, userID); err != nil {
		return ErrMarketplace.Wrap(err)
	}
	// TODO: check if the user has the required amount of money.

	lot, err := service.GetLotByID(ctx, betLot.ID)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}
	if lot.Status == StatusSold || lot.Status == StatusSoldBuynow {
		return ErrMarketplace.New("the lot is already on sale")
	}
	if lot.Status == StatusExpired {
		return ErrMarketplace.New("the lot is already on expired")
	}

	if betLot.BetAmount < lot.StartPrice || betLot.BetAmount <= lot.CurrentPrice {
		return ErrMarketplace.New("not enough money")
	}

	/** TODO: the transaction may be required for all operations,
	so that an error in the middle does not lead to an unwanted result in the database. **/

	// TODO: update status to `hold` for new user's money.
	// TODO: unhold old user's money if exist.

	if err = service.UpdateShopperIDLot(ctx, betLot.ID, userID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if betLot.BetAmount >= lot.MaxPrice && lot.MaxPrice != 0 {
		if err = service.UpdateCurrentPriceLot(ctx, betLot.ID, lot.MaxPrice); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		winLot := WinLot{
			ID:        betLot.ID,
			ItemID:    lot.ItemID,
			Type:      TypeCard,
			UserID:    lot.UserID,
			ShopperID: userID,
			Status:    StatusSoldBuynow,
			Amount:    lot.MaxPrice,
		}

		if err = service.WinLot(ctx, winLot); err != nil {
			return ErrMarketplace.Wrap(err)
		}

	} else {
		if err = service.UpdateCurrentPriceLot(ctx, betLot.ID, betLot.BetAmount); err != nil {
			return ErrMarketplace.Wrap(err)
		}
		if lot.EndTime.Sub(time.Now().UTC()) < time.Minute {
			if err = service.UpdateEndTimeLot(ctx, betLot.ID, time.Now().UTC().Add(time.Minute)); err != nil {
				return ErrMarketplace.Wrap(err)
			}
		}
	}

	return nil
}

// WinLot changes owner of the item and transfers money.
func (service *Service) WinLot(ctx context.Context, winLot WinLot) error {
	if err := service.UpdateStatusLot(ctx, winLot.ID, winLot.Status); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	// TODO: transfer money to the old cardholder from new user. If userID == shopperID not transfer mb

	if winLot.Type == TypeCard {
		if err := service.cards.UpdateStatus(ctx, winLot.ItemID, cards.StatusActive); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		if winLot.UserID != winLot.ShopperID {
			if err := service.cards.UpdateUserID(ctx, winLot.ItemID, winLot.ShopperID); err != nil {
				return ErrMarketplace.Wrap(err)
			}
		}
	}
	// TODO: check other items

	return nil
}

// UpdateShopperIDLot updates shopper id of lot.
func (service *Service) UpdateShopperIDLot(ctx context.Context, id, shopperID uuid.UUID) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateShopperIDLot(ctx, id, shopperID))
}

// UpdateStatusLot updates status of lot.
func (service *Service) UpdateStatusLot(ctx context.Context, id uuid.UUID, status Status) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateStatusLot(ctx, id, status))
}

// UpdateCurrentPriceLot updates current price of lot.
func (service *Service) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice float64) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateCurrentPriceLot(ctx, id, currentPrice))
}

// UpdateEndTimeLot updates end time of lot.
func (service *Service) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateEndTimeLot(ctx, id, endTime))
}
