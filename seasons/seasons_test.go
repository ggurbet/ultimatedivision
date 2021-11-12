// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/divisions"
	"ultimatedivision/seasons"
)

func TestSeasons(t *testing.T) {
	division1 := divisions.Division{
		ID:             uuid.New(),
		Name:           "10",
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}
	division2 := divisions.Division{
		ID:             uuid.New(),
		Name:           "9",
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

		t.Run("endSeason", func(t *testing.T) {
			err := repository.EndSeason(ctx, season1.ID)
			require.NoError(t, err)
		})

		t.Run("delete sql no rows", func(t *testing.T) {
			err := repository.Delete(ctx, id)
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
