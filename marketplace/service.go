// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/users"
)

// ErrMarketplace indicated that there was an error in service.
var ErrMarketplace = errs.Class("marketplace service error")

// Service is handling marketplace related logic.
//
// architecture: Service.
type Service struct {
	config      Config
	marketplace DB
	users       *users.Service
	cards       *cards.Service
}

// NewService is a constructor for marketplace service.
func NewService(config Config, marketplace DB, users *users.Service, cards *cards.Service) *Service {
	return &Service{
		config:      config,
		marketplace: marketplace,
		users:       users,
		cards:       cards,
	}
}

// CreateLot add lot in DB.
func (service *Service) CreateLot(ctx context.Context, createLot CreateLot) error {
	// TODO: add transaction.
	card, err := service.cards.Get(ctx, createLot.ItemID)
	if err == nil {
		if card.UserID != createLot.UserID {
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
	// TODO: check other items.

	if createLot.Type == "" {
		return ErrMarketplace.New("not found item by id")
	}

	if _, err := service.users.Get(ctx, createLot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if createLot.MaxPrice.BitLen() != 0 && createLot.MaxPrice.Cmp(&createLot.StartPrice) == -1 {
		return ErrMarketplace.New("max price less start price")
	}

	if createLot.Period < MinPeriod && createLot.Period < MaxPeriod {
		return ErrMarketplace.New("period exceed the range from 1 to 120 hours")
	}

	lot := Lot{
		ID:         uuid.New(),
		ItemID:     createLot.ItemID,
		Type:       createLot.Type,
		UserID:     createLot.UserID,
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
func (service *Service) ListActiveLots(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}
	lotsPage, err := service.marketplace.ListActiveLots(ctx, cursor)
	return lotsPage, ErrMarketplace.Wrap(err)
}

// ListActiveLotsWithFilters returns active lots from DB, taking the necessary filters.
func (service *Service) ListActiveLotsWithFilters(ctx context.Context, filters []cards.Filters, cursor pagination.Cursor) (Page, error) {
	var lotsPage Page
	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return lotsPage, ErrMarketplace.Wrap(err)
		}
	}

	cardIDs, err := service.cards.ListCardIDsWithFiltersWhereActiveLot(ctx, filters)
	if err != nil {
		return lotsPage, ErrMarketplace.Wrap(err)
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}
	lotsPage, err = service.marketplace.ListActiveLotsByItemID(ctx, cardIDs, cursor)
	return lotsPage, ErrMarketplace.Wrap(err)
}

// ListActiveLotsByPlayerName returns active lots from DB by player name card.
func (service *Service) ListActiveLotsByPlayerName(ctx context.Context, filter cards.Filters, cursor pagination.Cursor) (Page, error) {
	var lotsPage Page
	strings.ToValidUTF8(filter.Value, "")

	// TODO: add best check.
	_, err := strconv.Atoi(filter.Value)
	if err == nil {
		return lotsPage, ErrMarketplace.Wrap(cards.ErrInvalidFilter.New("%s %s", filter.Value, err))
	}

	cardIDs, err := service.cards.ListCardIDsByPlayerNameWhereActiveLot(ctx, filter)
	if err != nil {
		return lotsPage, ErrMarketplace.Wrap(err)
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}
	lotsPage, err = service.marketplace.ListActiveLotsByItemID(ctx, cardIDs, cursor)
	return lotsPage, ErrMarketplace.Wrap(err)
}

// ListExpiredLot returns not active lots from DB.
func (service *Service) ListExpiredLot(ctx context.Context) ([]Lot, error) {
	lots, err := service.marketplace.ListExpiredLot(ctx)
	return lots, ErrMarketplace.Wrap(err)
}

// PlaceBetLot checks the amount of money and makes a bet.
func (service *Service) PlaceBetLot(ctx context.Context, betLot BetLot) error {
	if _, err := service.users.Get(ctx, betLot.UserID); err != nil {
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

	if betLot.BetAmount.Cmp(&lot.StartPrice) == -1 || betLot.BetAmount.Cmp(&lot.CurrentPrice) == -1 || betLot.BetAmount.Cmp(&lot.CurrentPrice) == 0 {
		return ErrMarketplace.New("not enough money")
	}

	/** TODO: the transaction may be required for all operations,
	  so that an error in the middle does not lead to an unwanted result in the database. **/

	// TODO: update status to `hold` for new user's money.
	// TODO: unhold old user's money if exist.

	if err := service.UpdateShopperIDLot(ctx, betLot.ID, betLot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if (betLot.BetAmount.Cmp(&lot.MaxPrice) == 1 || betLot.BetAmount.Cmp(&lot.MaxPrice) == 0) && lot.MaxPrice.BitLen() != 0 {
		if err = service.UpdateCurrentPriceLot(ctx, betLot.ID, lot.MaxPrice); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		winLot := WinLot{
			ID:        betLot.ID,
			ItemID:    lot.ItemID,
			Type:      TypeCard,
			UserID:    lot.UserID,
			ShopperID: betLot.UserID,
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

	// TODO: transfer money to the old cardholder from new user. If userID == shopperID not transfer mb.

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
	// TODO: check other items.

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
func (service *Service) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice big.Int) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateCurrentPriceLot(ctx, id, currentPrice))
}

// UpdateEndTimeLot updates end time of lot.
func (service *Service) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateEndTimeLot(ctx, id, endTime))
}
