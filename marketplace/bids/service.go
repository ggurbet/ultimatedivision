// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package bids

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/clubs"
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
	clubs       *clubs.Service
	nfts        *nfts.Service
	users       *users.Service
}

// NewService is constructor for Service.
func NewService(bids DB, marketplace *marketplace.Service, cards *cards.Service, clubs *clubs.Service, nfts *nfts.Service, users *users.Service) *Service {
	return &Service{
		bids:        bids,
		marketplace: marketplace,
		cards:       cards,
		clubs:       clubs,
		nfts:        nfts,
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

	if err = service.marketplace.UpdateCurrentPriceLot(ctx, bid.LotID, bid.Amount); err != nil {
		return ErrBids.Wrap(err)
	}
	return nil
}

// GetMakeOfferData returns make offer data by card id from DB.
func (service *Service) GetMakeOfferData(ctx context.Context, cardID, userID uuid.UUID) (nfts.MakeOffer, error) {
	tokenIDWithContractAddress, err := service.marketplace.GetMakeOfferByCardID(ctx, cardID)
	if err != nil {
		return nfts.MakeOffer{}, ErrBids.Wrap(err)
	}

	if err := service.marketplace.UpdateShopperIDLot(ctx, cardID, userID); err != nil {
		log.Error(fmt.Sprintf("could not update update shopper id by card id equal %v in db", cardID), ErrBids.Wrap(err))
	}
	if err = service.marketplace.Delete(ctx, cardID); err != nil {
		log.Error(fmt.Sprintf("could not delete lot by card id equal %v in db", cardID), ErrBids.Wrap(err))
	}
	if err = service.cards.UpdateStatus(ctx, cardID, cards.StatusActive); err != nil {
		log.Error(fmt.Sprintf("could not update card status by card id equal %v in db", cardID), ErrBids.Wrap(err))
	}

	if err = service.bids.DeleteByLotID(ctx, cardID); err != nil {
		log.Error(fmt.Sprintf("could not delete bids by card id equal %v in db", cardID), ErrBids.Wrap(err))
	}

	squadID, err := service.clubs.GetSquadIDByCardID(ctx, cardID)
	if err != nil {
		log.Error(fmt.Sprintf("could not get squad by card id equal %v from db", cardID), ErrBids.Wrap(err))
	}

	if squadID != uuid.Nil {
		if err = service.clubs.DeleteByCardID(ctx, cardID); err != nil {
			log.Error(fmt.Sprintf("could not delete card from club by card id equal %v in db", cardID), ErrBids.Wrap(err))
		}
	}

	user, err := service.users.Get(ctx, userID)
	if err != nil {
		log.Error(fmt.Sprintf("could not get user by user id equal %v from db", userID), ErrBids.Wrap(err))
	}

	if err = service.cards.UpdateUserID(ctx, cardID, userID); err != nil {
		log.Error(fmt.Sprintf("could not get update user id of the card lot id equal %v in db", cardID), ErrBids.Wrap(err))
	}

	card, err := service.cards.Get(ctx, cardID)
	if err != nil {
		log.Error(fmt.Sprintf("could not get card by lot id equal %v from db", cardID), ErrBids.Wrap(err))
	}

	nft, err := service.nfts.GetNFTByCardID(ctx, card.ID)
	if err != nil {
		log.Error(fmt.Sprintf("could not get nft by card id equal %v from db", card.ID), ErrBids.Wrap(err))
	}

	if user.WalletType == users.WalletTypeCasper {
		nft.WalletAddress = common.HexToAddress(user.CasperWallet)
	} else {
		nft.WalletAddress = user.Wallet
	}

	if err = service.nfts.Update(ctx, nft); err != nil {
		log.Error("could not update nft by nft data from db", ErrBids.Wrap(err))
	}

	return tokenIDWithContractAddress, ErrBids.Wrap(err)
}

// GetCurrentBidByLotID returns current bid by lot id from the database.
func (service *Service) GetCurrentBidByLotID(ctx context.Context, lotID uuid.UUID) (Bid, error) {
	currentAmount, err := service.bids.GetCurrentBidByLotID(ctx, lotID)
	return currentAmount, ErrBids.Wrap(err)
}

// DeleteByLotID deletes bids by lot id in the database.
func (service *Service) DeleteByLotID(ctx context.Context, lotID uuid.UUID) error {
	return ErrBids.Wrap(service.bids.DeleteByLotID(ctx, lotID))
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
