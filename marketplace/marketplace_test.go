// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/marketplace"
	"ultimatedivision/users"
)

func TestMarketplace(t *testing.T) {
	lot1 := marketplace.Lot{
		ID:           uuid.New(),
		ItemID:       uuid.New(),
		Type:         marketplace.TypeCard,
		UserID:       uuid.New(),
		ShopperID:    uuid.New(),
		Status:       marketplace.StatusSoldBuynow,
		StartPrice:   5.0,
		MaxPrice:     30.0,
		CurrentPrice: 30.0,
		StartTime:    time.Now().UTC(),
		EndTime:      time.Now().AddDate(0, 0, 2).UTC(),
		Period:       2,
	}

	lot2 := marketplace.Lot{
		ID:           uuid.New(),
		ItemID:       uuid.New(),
		Type:         marketplace.TypeCard,
		UserID:       uuid.New(),
		Status:       marketplace.StatusActive,
		StartPrice:   5.0,
		CurrentPrice: 25.0,
		StartTime:    time.Now().UTC(),
		EndTime:      time.Now().AddDate(0, 0, 1).UTC(),
		Period:       marketplace.MinPeriod,
	}

	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		LastLogin:    time.Now(),
		Status:       0,
		CreatedAt:    time.Now(),
	}

	user2 := users.User{
		ID:           uuid.New(),
		Email:        "3560876@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "qwerty",
		FirstName:    "Stas",
		LastName:     "Isakov",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryMarketplace := db.Marketplace()
		repositoryUsers := db.Users()
		id := uuid.New()
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryMarketplace.GetLotByID(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, marketplace.ErrNoLot.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			lot1.UserID = user1.ID
			err = repositoryMarketplace.CreateLot(ctx, lot1)
			require.NoError(t, err)

			lotFromDB, err := repositoryMarketplace.GetLotByID(ctx, lot1.ID)
			require.NoError(t, err)
			compareLot(t, lot1, lotFromDB)
		})

		t.Run("list active", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			lot2.UserID = user2.ID
			err = repositoryMarketplace.CreateLot(ctx, lot2)
			require.NoError(t, err)

			activeLots, err := repositoryMarketplace.ListActiveLots(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(activeLots), 1)
			compareLot(t, lot2, activeLots[0])
		})

		t.Run("list expired lot", func(t *testing.T) {
			lot1.EndTime = time.Now().UTC()
			err := repositoryMarketplace.UpdateEndTimeLot(ctx, lot1.ID, lot1.EndTime)
			require.NoError(t, err)

			lot1.Status = marketplace.StatusActive
			err = repositoryMarketplace.UpdateStatusLot(ctx, lot1.ID, marketplace.StatusActive)
			require.NoError(t, err)

			activeLots, err := repositoryMarketplace.ListExpiredLot(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(activeLots), 1)
			compareLot(t, lot1, activeLots[0])
		})

		t.Run("update shopperID of lot", func(t *testing.T) {
			lot1.ShopperID = uuid.New()
			err := repositoryMarketplace.UpdateShopperIDLot(ctx, lot1.ID, lot1.ShopperID)
			require.NoError(t, err)

			lotFromDB, err := repositoryMarketplace.GetLotByID(ctx, lot1.ID)
			require.NoError(t, err)
			compareLot(t, lot1, lotFromDB)
		})

		t.Run("update staus of lot", func(t *testing.T) {
			lot1.Status = marketplace.StatusExpired
			err := repositoryMarketplace.UpdateStatusLot(ctx, lot1.ID, marketplace.StatusExpired)
			require.NoError(t, err)

			lotFromDB, err := repositoryMarketplace.GetLotByID(ctx, lot1.ID)
			require.NoError(t, err)
			compareLot(t, lot1, lotFromDB)
		})

		t.Run("update current price of lot", func(t *testing.T) {
			lot1.CurrentPrice = 25.0
			err := repositoryMarketplace.UpdateCurrentPriceLot(ctx, lot1.ID, 25.0)
			require.NoError(t, err)

			lotFromDB, err := repositoryMarketplace.GetLotByID(ctx, lot1.ID)
			require.NoError(t, err)
			compareLot(t, lot1, lotFromDB)
		})

		t.Run("update end time of lot", func(t *testing.T) {
			lot1.EndTime = time.Now().UTC().Add(time.Hour)
			err := repositoryMarketplace.UpdateEndTimeLot(ctx, lot1.ID, lot1.EndTime)
			require.NoError(t, err)

			lotFromDB, err := repositoryMarketplace.GetLotByID(ctx, lot1.ID)
			require.NoError(t, err)
			compareLot(t, lot1, lotFromDB)
		})

	})
}

func compareLot(t *testing.T, lot1, lot2 marketplace.Lot) {
	assert.Equal(t, lot1.ID, lot2.ID)
	assert.Equal(t, lot1.ItemID, lot2.ItemID)
	assert.Equal(t, lot1.Type, lot2.Type)
	assert.Equal(t, lot1.UserID, lot2.UserID)
	assert.Equal(t, lot1.ShopperID, lot2.ShopperID)
	assert.Equal(t, lot1.Status, lot2.Status)
	assert.Equal(t, lot1.StartPrice, lot2.StartPrice)
	assert.Equal(t, lot1.MaxPrice, lot2.MaxPrice)
	assert.Equal(t, lot1.CurrentPrice, lot2.CurrentPrice)
	assert.WithinDuration(t, lot1.StartTime, lot2.StartTime, 1*time.Second)
	assert.WithinDuration(t, lot1.EndTime, lot2.EndTime, 1*time.Second)
	assert.Equal(t, lot1.Period, lot2.Period)
}
