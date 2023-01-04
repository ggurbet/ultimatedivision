// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons_test

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
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/divisions"
	"ultimatedivision/seasons"
	"ultimatedivision/users"
)

func TestSeasons(t *testing.T) {
	division1 := divisions.Division{
		ID:             uuid.New(),
		Name:           10,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}
	division2 := divisions.Division{
		ID:             uuid.New(),
		Name:           9,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}

	season1 := seasons.Season{
		ID:         1,
		DivisionID: division1.ID,
		StartedAt:  time.Now().UTC(),
		EndedAt:    time.Time{},
	}
	season2 := seasons.Season{
		ID:         2,
		DivisionID: division2.ID,
		StartedAt:  time.Now().UTC(),
		EndedAt:    time.Time{},
	}

	season3 := seasons.Season{
		ID:         3,
		DivisionID: division1.ID,
		StartedAt:  time.Now().UTC(),
		EndedAt:    time.Time{},
	}
	season4 := seasons.Season{
		ID:         4,
		DivisionID: division2.ID,
		StartedAt:  time.Now().UTC(),
		EndedAt:    time.Time{},
	}

	user := users.User{
		ID:             uuid.New(),
		Email:          "oleksii@gmail.com",
		PasswordHash:   []byte{0},
		NickName:       "Free",
		FirstName:      "Oleksii",
		LastName:       "Prysiazhniuk",
		Wallet:         common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b7"),
		CasperWallet:   "01a4db357602c3d45a2b7b68110e66440ac2a2e792cebffbce83eaefb73e65aef1",
		CasperWalletID: "4bfcd0ebd44c3de9d1e6556336cbb73259649b7d6b344bc1499d40652fd5781a",
		WalletType:     users.WalletTypeETH,
		LastLogin:      time.Now().UTC(),
		Status:         0,
		CreatedAt:      time.Now().UTC(),
	}

	value := *big.NewInt(100)

	reward := seasons.Reward{
		UserID:              user.ID,
		SeasonID:            season1.ID,
		WalletAddress:       common.Address{},
		CasperWalletAddress: user.CasperWallet,
		WalletType:          user.WalletType,
		Value:               value,
		Status:              1,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repository := db.Seasons()
		repositoryDivision := db.Divisions()

		id := 3
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repository.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, seasons.ErrNoSeason.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryDivision.Create(ctx, division1)
			require.NoError(t, err)
			err = repository.Create(ctx, season1)
			require.NoError(t, err)

			seasonFromDB, err := repository.Get(ctx, season1.ID)
			require.NoError(t, err)
			compareSeasons(t, season1, seasonFromDB)
		})

		t.Run("Create Reward", func(t *testing.T) {
			err := repository.CreateReward(ctx, reward)
			require.NoError(t, err)
		})

		t.Run("list", func(t *testing.T) {
			err := repositoryDivision.Create(ctx, division2)
			require.NoError(t, err)
			err = repository.Create(ctx, season2)
			require.NoError(t, err)

			allSeasons, err := repository.List(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(allSeasons), 2)
			compareSeasons(t, season1, allSeasons[0])
			compareSeasons(t, season2, allSeasons[1])
		})

		t.Run("get current season by division id", func(t *testing.T) {
			err := repository.EndSeason(ctx, season1.ID)
			require.NoError(t, err)
			err = repository.EndSeason(ctx, season2.ID)
			require.NoError(t, err)
			err = repository.Create(ctx, season3)
			require.NoError(t, err)
			err = repository.Create(ctx, season4)
			require.NoError(t, err)

			seasonFromDB, err := repository.GetSeasonByDivisionID(ctx, division1.ID)
			require.NoError(t, err)
			compareSeasons(t, season3, seasonFromDB)
		})

		t.Run("endSeason", func(t *testing.T) {
			err := repository.EndSeason(ctx, season4.ID)
			require.NoError(t, err)
		})

		t.Run("delete sql no rows", func(t *testing.T) {
			err := repository.Delete(ctx, 5)
			require.Error(t, err)
			require.Equal(t, seasons.ErrNoSeason.Has(err), true)
		})

		t.Run("delete", func(t *testing.T) {
			err := repository.Delete(ctx, season1.ID)
			require.NoError(t, err)
		})
	})
}

func compareSeasons(t *testing.T, season1, season2 seasons.Season) {
	assert.Equal(t, season1.ID, season2.ID)
	assert.Equal(t, season1.DivisionID, season2.DivisionID)
	assert.WithinDuration(t, season1.StartedAt, season2.StartedAt, 1*time.Second)
}
