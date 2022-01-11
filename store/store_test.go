// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/store"
)

func TestStore(t *testing.T) {
	setting1 := store.Setting{
		ID:          1,
		CardsAmount: 10,
		IsRenewal:   true,
		DateRenewal: time.Now().UTC(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryStore := db.Store()

		t.Run("Create", func(t *testing.T) {
			err := repositoryStore.Create(ctx, setting1)
			require.NoError(t, err)
		})

		t.Run("Get", func(t *testing.T) {
			settingGet, err := repositoryStore.Get(ctx, setting1.ID)
			require.NoError(t, err)

			compareStoreSlice(t, []store.Setting{settingGet}, []store.Setting{setting1})
		})

		t.Run("List", func(t *testing.T) {
			settingList, err := repositoryStore.List(ctx)
			require.NoError(t, err)

			compareStoreSlice(t, settingList, []store.Setting{setting1})
		})

		t.Run("Update", func(t *testing.T) {
			setting1.CardsAmount = 15
			setting1.IsRenewal = false
			setting1.DateRenewal = time.Now().UTC()

			err := repositoryStore.Update(ctx, setting1)
			require.NoError(t, err)

			settingGet, err := repositoryStore.Get(ctx, setting1.ID)
			require.NoError(t, err)

			compareStoreSlice(t, []store.Setting{settingGet}, []store.Setting{setting1})
		})

	})
}

func compareStoreSlice(t *testing.T, setting1, setting2 []store.Setting) {
	assert.Equal(t, len(setting1), len(setting2))

	for i := 0; i < len(setting1); i++ {
		assert.Equal(t, setting1[i].ID, setting2[i].ID)
		assert.Equal(t, setting1[i].CardsAmount, setting2[i].CardsAmount)
		assert.Equal(t, setting1[i].IsRenewal, setting2[i].IsRenewal)
		assert.WithinDuration(t, setting1[i].DateRenewal, setting2[i].DateRenewal, 1*time.Second)
	}
}
