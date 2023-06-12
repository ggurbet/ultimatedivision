// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
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
	nfts        *nfts.Service
}

// NewService is a constructor for marketplace service.
func NewService(config Config, marketplace DB, users *users.Service, cards *cards.Service, nfts *nfts.Service) *Service {
	return &Service{
		config:      config,
		marketplace: marketplace,
		users:       users,
		cards:       cards,
		nfts:        nfts,
	}
}

// CreateLot add lot in DB.
func (service *Service) CreateLot(ctx context.Context, createLot CreateLot) error {
	// TODO: add transaction.
	card, err := service.cards.Get(ctx, createLot.CardID)
	if err == nil {
		if card.UserID != createLot.UserID {
			return ErrMarketplace.New("it is not the user's card")
		}

		if card.Status == cards.StatusSale {
			return ErrMarketplace.New("the card is already on sale")
		}

		if err := service.cards.UpdateStatus(ctx, createLot.CardID, cards.StatusSale); err != nil {
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
		CardID:     card.ID,
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

// GetLotEndTimeByID returns lot end time by id from DB.
func (service *Service) GetLotEndTimeByID(ctx context.Context, id uuid.UUID) (bool, error) {
	endTime, err := service.marketplace.GetLotEndTimeByID(ctx, id)
	if time.Now().UTC().After(endTime) {
		return true, ErrMarketplace.Wrap(err)
	}

	return false, ErrMarketplace.Wrap(err)
}

// GetCurrentPriceByCardID returns current price by card id from the data base.
func (service *Service) GetCurrentPriceByCardID(ctx context.Context, cardID uuid.UUID) (big.Int, error) {
	currentPrice, err := service.marketplace.GetCurrentPriceByCardID(ctx, cardID)
	return currentPrice, ErrMarketplace.Wrap(err)
}

// GetNFTTokenIDbyCardID returns nft token id by card id from database.
func (service *Service) GetNFTTokenIDbyCardID(ctx context.Context, cardID uuid.UUID) (uuid.UUID, error) {
	tokenID, err := service.nfts.GetNFTTokenIDbyCardID(ctx, cardID)
	return tokenID, ErrMarketplace.Wrap(err)
}

// GetNFTByCardID returns nft by card id from DB.
func (service *Service) GetNFTByCardID(ctx context.Context, id uuid.UUID) (nfts.NFT, error) {
	nft, err := service.nfts.GetNFTByCardID(ctx, id)
	return nft, ErrMarketplace.Wrap(err)
}

// IsMinted returns 1 if minted or 0 if not minted.
func (service *Service) IsMinted(ctx context.Context, id uuid.UUID) (int, error) {
	return service.nfts.IsMinted(ctx, id)
}

// GetNFTDataByCardID returns nft data by card id from DB.
func (service *Service) GetNFTDataByCardID(ctx context.Context, cardID uuid.UUID) (nfts.TokenIDWithContractAddress, error) {
	var lotData nfts.TokenIDWithContractAddress

	tokenID, err := service.nfts.GetNFTTokenIDbyCardID(ctx, cardID)
	if err != nil {
		return lotData, ErrMarketplace.Wrap(err)
	}
	lotData.TokenID = tokenID
	lotData.Address = service.config.MarketContractAddress
	lotData.AddressNodeServer = service.config.RPCNodeAddress
	lotData.ContractHash = fmt.Sprintf("%s%s", service.config.NFTContractPrefix, service.config.NFTContractAddress)

	return lotData, ErrMarketplace.Wrap(err)
}

// GetApproveByCardID returns nft data by card id from DB.
func (service *Service) GetApproveByCardID(ctx context.Context, cardID string) (nfts.TokenIDWithApproveData, error) {
	var approveData nfts.TokenIDWithApproveData
	var err error

	if cardID != "" {
		cardIDUuid, err := uuid.Parse(cardID)
		if err != nil {
			return approveData, ErrMarketplace.Wrap(err)
		}
		tokenID, err := service.nfts.GetNFTTokenIDbyCardID(ctx, cardIDUuid)
		if err != nil {
			return approveData, ErrMarketplace.Wrap(err)
		}
		approveData.TokenID = tokenID.String()
	}

	approveData.AddressNodeServer = service.config.RPCNodeAddress
	approveData.NFTContractAddress = service.config.NFTContractAddress
	approveData.TokenRewardContractAddress = service.config.TokenContractAddress
	approveData.Amount = service.config.TokenAmountForApproving
	approveData.ApproveNFTSpender = fmt.Sprintf("%s%s", service.config.NFTApprovePrefix, service.config.MarketContractPackageAddress)
	approveData.ApproveTokensSpender = service.config.MarketContractPackageAddress

	return approveData, ErrMarketplace.Wrap(err)
}

// GetMakeOfferByCardID returns make offer data by card id from DB.
func (service *Service) GetMakeOfferByCardID(ctx context.Context, cardID uuid.UUID) (nfts.MakeOffer, error) {
	tokenID, err := service.nfts.GetNFTTokenIDbyCardID(ctx, cardID)
	if err != nil {
		return nfts.MakeOffer{}, ErrMarketplace.Wrap(err)
	}

	return nfts.MakeOffer{
		TokenID:           tokenID,
		Address:           service.config.MarketContractAddress,
		AddressNodeServer: service.config.RPCNodeAddress,
		ContractHash:      fmt.Sprintf("%s%s", service.config.NFTContractPrefix, service.config.NFTContractAddress),
		TokenContractHash: service.config.TokenContractAddress,
	}, ErrMarketplace.Wrap(err)
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

// ListExpiredLots returns all expired lots form the database.
func (service *Service) ListExpiredLots(ctx context.Context) ([]Lot, error) {
	lots, err := service.marketplace.ListExpiredLot(ctx)
	return lots, ErrMarketplace.Wrap(err)
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
	lotsPage, err = service.marketplace.ListActiveLotsByCardID(ctx, cardIDs, cursor)
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
	lotsPage, err = service.marketplace.ListActiveLotsByCardID(ctx, cardIDs, cursor)
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

	lot, err := service.GetLotByID(ctx, betLot.CardID)
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

	if err := service.UpdateShopperIDLot(ctx, betLot.CardID, betLot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if (betLot.BetAmount.Cmp(&lot.MaxPrice) == 1 || betLot.BetAmount.Cmp(&lot.MaxPrice) == 0) && lot.MaxPrice.BitLen() != 0 {
		if err = service.UpdateCurrentPriceLot(ctx, betLot.CardID, lot.MaxPrice); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		winLot := WinLot{
			CardID:    betLot.CardID,
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
		if err = service.UpdateCurrentPriceLot(ctx, betLot.CardID, betLot.BetAmount); err != nil {
			return ErrMarketplace.Wrap(err)
		}
		if lot.EndTime.Sub(time.Now().UTC()) < time.Minute {
			if err = service.UpdateEndTimeLot(ctx, betLot.CardID, time.Now().UTC().Add(time.Minute)); err != nil {
				return ErrMarketplace.Wrap(err)
			}
		}
	}

	return nil
}

// WinLot changes owner of the item and transfers money.
func (service *Service) WinLot(ctx context.Context, winLot WinLot) error {
	if err := service.UpdateStatusLot(ctx, winLot.CardID, winLot.Status); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	// TODO: transfer money to the old cardholder from new user. If userID == shopperID not transfer mb.

	if winLot.Type == TypeCard {
		if err := service.cards.UpdateStatus(ctx, winLot.CardID, cards.StatusActive); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		if winLot.UserID != winLot.ShopperID {
			if err := service.cards.UpdateUserID(ctx, winLot.CardID, winLot.ShopperID); err != nil {
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

// Delete deletes lot in the database.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return ErrMarketplace.Wrap(service.marketplace.Delete(ctx, cardID))
}
