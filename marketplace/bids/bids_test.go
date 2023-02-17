// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package bids_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/cards"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/marketplace"
	"ultimatedivision/marketplace/bids"
	"ultimatedivision/users"
)

func TestBids(t *testing.T) {
	user1 := users.User{
		ID:               uuid.New(),
		Email:            "tarkovskynik@gmail.com",
		PasswordHash:     []byte{0},
		NickName:         "Nik",
		FirstName:        "Nikita",
		LastName:         "Tarkovskyi",
		Wallet:           common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b7"),
		CasperWallet:     "01a4db357602c3d45a2b7b68110e66440ac2a2e792cebffbce83eaefb73e65aef1",
		CasperWalletHash: "4bfcd0ebd44c3de9d1e6556336cbb73259649b7d6b344bc1499d40652fd5781a",
		WalletType:       users.WalletTypeETH,
		LastLogin:        time.Now().UTC(),
		Status:           0,
		CreatedAt:        time.Now().UTC(),
	}

	user2 := users.User{
		ID:               uuid.New(),
		Email:            "3560876@gmail.com",
		PasswordHash:     []byte{1},
		NickName:         "qwerty",
		FirstName:        "Stas",
		LastName:         "Isakov",
		Wallet:           common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b8"),
		CasperWallet:     "01a4db357602c3d45a2b7b68110e66440ac2a2e792cebffbce83eaefb73e65aef1",
		CasperWalletHash: "4bfcd0ebd44c3de9d1e6556336cbb73259649b7d6b344bc1499d40652fd5781a",
		WalletType:       users.WalletTypeVelas,
		LastLogin:        time.Now().UTC(),
		Status:           1,
		CreatedAt:        time.Now().UTC(),
	}

	card1 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro yak muk",
		Quality:          "wood",
		Height:           178.8,
		Weight:           72.2,
		DominantFoot:     "left",
		IsTattoo:         false,
		Status:           cards.StatusActive,
		Type:             cards.TypeWon,
		UserID:           user1.ID,
		Tactics:          1,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offence:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
	}

	card2 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Vova",
		Quality:          "gold",
		Height:           179.9,
		Weight:           73.3,
		DominantFoot:     "right",
		IsTattoo:         true,
		Status:           cards.StatusSale,
		Type:             cards.TypeUnordered,
		UserID:           uuid.New(),
		Tactics:          2,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offence:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
	}

	lot1 := marketplace.Lot{
		CardID:       card1.ID,
		Type:         marketplace.TypeCard,
		UserID:       uuid.New(),
		ShopperID:    uuid.New(),
		Status:       marketplace.StatusSoldBuynow,
		StartPrice:   *big.NewInt(500000000000000),
		MaxPrice:     *big.NewInt(3000000000000000),
		CurrentPrice: *big.NewInt(3000000000000000),
		StartTime:    time.Now().UTC(),
		EndTime:      time.Now().AddDate(0, 0, 2).UTC(),
		Period:       2,
	}

	_ = marketplace.Lot{
		CardID:       card2.ID,
		Type:         marketplace.TypeCard,
		UserID:       uuid.New(),
		Status:       marketplace.StatusActive,
		StartPrice:   *big.NewInt(500000000000000),
		CurrentPrice: *big.NewInt(2500000000000000),
		StartTime:    time.Now().UTC(),
		EndTime:      time.Now().AddDate(0, 0, 1).UTC(),
		Period:       marketplace.MinPeriod,
	}

	bid1 := bids.Bid{
		ID:        uuid.New(),
		LotID:     lot1.CardID,
		UserID:    user1.ID,
		Amount:    *big.NewInt(2100000000000000000),
		CreatedAt: time.Now().UTC().Add(5 * time.Minute).Round(time.Second),
	}

	bid2 := bids.Bid{
		ID:        uuid.New(),
		LotID:     lot1.CardID,
		UserID:    user2.ID,
		Amount:    *big.NewInt(2200000000000000000),
		CreatedAt: time.Now().UTC().Add(7 * time.Minute).Round(time.Second),
	}

	bid3 := bids.Bid{
		ID:        uuid.New(),
		LotID:     lot1.CardID,
		UserID:    user1.ID,
		Amount:    *big.NewInt(2100000000000000000),
		CreatedAt: time.Now().Add(5 * time.Minute).UTC(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		userRepository := db.Users()
		cardsRepository := db.Cards()
		bidsRepository := db.Bids()

		t.Run("Create", func(t *testing.T) {
			err := userRepository.Create(ctx, user1)
			require.NoError(t, err)
			err = userRepository.Create(ctx, user2)
			require.NoError(t, err)
			require.NoError(t, err)
			err = cardsRepository.Create(ctx, card1)
			require.NoError(t, err)
			err = cardsRepository.Create(ctx, card2)
			require.NoError(t, err)
			require.NoError(t, err)
			err = bidsRepository.Create(ctx, bid1)
			require.NoError(t, err)
			err = bidsRepository.Create(ctx, bid2)
			require.NoError(t, err)
		})

		t.Run("GetCurrentBidByLotID", func(t *testing.T) {
			currentBid, err := bidsRepository.GetCurrentBidByLotID(ctx, bid2.LotID)
			require.NoError(t, err)
			assert.Equal(t, currentBid, bid2)
		})

		t.Run("ListByLotID", func(t *testing.T) {
			bids, err := bidsRepository.ListByLotID(ctx, lot1.CardID)
			require.NoError(t, err)
			require.Equal(t, len(bids), 2)
			compareBids(t, bids[0], bid1)
			compareBids(t, bids[1], bid2)
		})

		t.Run("ListByUserID", func(t *testing.T) {
			bids, err := bidsRepository.ListByUserID(ctx, user1.ID)
			require.NoError(t, err)
			require.Equal(t, len(bids), 1)
			compareBids(t, bids[0], bid1)
		})

		t.Run("list user bids amount", func(t *testing.T) {
			err := bidsRepository.Create(ctx, bid3)
			require.NoError(t, err)

			amount, err := bidsRepository.GetUserBidsAmountByLotID(ctx, bid1.UserID, bid1.LotID)
			require.NoError(t, err)
			require.Equal(t, len(amount), 2)
		})

		t.Run("DeleteByLotID", func(t *testing.T) {
			err := bidsRepository.DeleteByLotID(ctx, lot1.CardID)
			require.NoError(t, err)
		})

		t.Run("Negative DeleteByLotID", func(t *testing.T) {
			err := bidsRepository.DeleteByLotID(ctx, lot1.CardID)
			require.Error(t, err)

			assert.True(t, bids.ErrNoBid.Has(err))
		})
	})
}

func compareBids(t *testing.T, bid1, bid2 bids.Bid) {
	assert.Equal(t, bid1.ID, bid2.ID)
	assert.Equal(t, bid1.UserID, bid2.UserID)
	assert.Equal(t, bid1.LotID, bid2.LotID)
	assert.Equal(t, bid1.Amount, bid2.Amount)
	assert.WithinDuration(t, bid1.CreatedAt, bid2.CreatedAt, 1*time.Second)
}
