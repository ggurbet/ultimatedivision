// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/store/lootboxes"
	"ultimatedivision/users"
)

func TestLootBox(t *testing.T) {
	user1 := users.User{
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

	userLootBox1 := lootboxes.LootBox{
		UserID:    user1.ID,
		LootBoxID: uuid.New(),
		Type:      lootboxes.RegularBox,
	}

	userLootBox2 := lootboxes.LootBox{
		UserID:    user1.ID,
		LootBoxID: uuid.New(),
		Type:      lootboxes.UDReleaseCelebrationBox,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryUsers := db.Users()
		repositoryLootBoxes := db.LootBoxes()
		id := uuid.New()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryLootBoxes.Create(ctx, userLootBox1)
			require.NoError(t, err)
		})

		t.Run("List", func(t *testing.T) {
			err := repositoryLootBoxes.Create(ctx, userLootBox2)
			require.NoError(t, err)

			userLootBoxes, err := repositoryLootBoxes.List(ctx)
			require.NoError(t, err)

			compareLootBoxesSlice(t, userLootBoxes, []lootboxes.LootBox{userLootBox1, userLootBox2})
		})

		t.Run("Get", func(t *testing.T) {
			lootBoxDB, err := repositoryLootBoxes.Get(ctx, userLootBox1.LootBoxID)
			require.NoError(t, err)

			compareLootBoxes(t, lootBoxDB, userLootBox1)
		})

		t.Run("Delete sql no rows", func(t *testing.T) {
			err := repositoryLootBoxes.Delete(ctx, id)
			require.Error(t, err)
			assert.Equal(t, lootboxes.ErrNoLootBox.Has(err), true)
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryLootBoxes.Delete(ctx, userLootBox1.LootBoxID)
			require.NoError(t, err)
		})
	})
}

func compareLootBoxesSlice(t *testing.T, userLootBoxesDB, userLootBoxesTest []lootboxes.LootBox) {
	assert.Equal(t, len(userLootBoxesDB), len(userLootBoxesTest))

	for i := 0; i < len(userLootBoxesDB); i++ {
		assert.Equal(t, userLootBoxesDB[i].UserID, userLootBoxesTest[i].UserID)
		assert.Equal(t, userLootBoxesDB[i].LootBoxID, userLootBoxesTest[i].LootBoxID)
		assert.Equal(t, userLootBoxesDB[i].Type, userLootBoxesTest[i].Type)
	}
}

func compareLootBoxes(t *testing.T, userLootBoxDB, userLootBoxTest lootboxes.LootBox) {
	assert.Equal(t, userLootBoxDB.UserID, userLootBoxTest.UserID)
	assert.Equal(t, userLootBoxDB.LootBoxID, userLootBoxTest.LootBoxID)
	assert.Equal(t, userLootBoxDB.Type, userLootBoxTest.Type)
}
