// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/lootboxes"
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

	userLootBox := lootboxes.LootBox{
		UserID:    user1.ID,
		LootBoxID: uuid.New(),
		Type:      lootboxes.RegularBox,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryUsers := db.Users()
		repositoryLootBoxes := db.LootBoxes()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryLootBoxes.Create(ctx, userLootBox)
			require.NoError(t, err)
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryLootBoxes.Delete(ctx, userLootBox.LootBoxID)
			require.NoError(t, err)
		})
	})
}
