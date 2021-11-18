// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package divisions_test

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
)

func TestDivisions(t *testing.T) {
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

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repository := db.Divisions()
		id := uuid.New()
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repository.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, divisions.ErrNoDivision.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repository.Create(ctx, division1)
			require.NoError(t, err)

			divisionFromDB, err := repository.Get(ctx, division1.ID)
			require.NoError(t, err)
			compareDivisions(t, division1, divisionFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repository.Create(ctx, division2)
			require.NoError(t, err)

			allDivisions, err := repository.List(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(allDivisions), 2)
			compareDivisions(t, division1, allDivisions[0])
			compareDivisions(t, division2, allDivisions[1])
		})

		t.Run("delete sql no rows", func(t *testing.T) {
			err := repository.Delete(ctx, id)
			require.Error(t, err)
			require.Equal(t, divisions.ErrNoDivision.Has(err), true)
		})

		t.Run("delete", func(t *testing.T) {
			err := repository.Delete(ctx, division1.ID)
			require.NoError(t, err)
		})
	})
}

func compareDivisions(t *testing.T, division1, division2 divisions.Division) {
	assert.Equal(t, division1.ID, division2.ID)
	assert.Equal(t, division1.Name, division2.Name)
	assert.Equal(t, division1.PassingPercent, division2.PassingPercent)
	assert.WithinDuration(t, division1.CreatedAt, division2.CreatedAt, 1*time.Second)
}
