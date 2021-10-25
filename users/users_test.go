// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/users"
)

func TestUsers(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		LastLogin:    time.Now().UTC(),
		Status:       0,
		CreatedAt:    time.Now().UTC(),
	}

	user2 := users.User{
		ID:           uuid.New(),
		Email:        "3560876@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "qwerty",
		FirstName:    "Stas",
		LastName:     "Isakov",
		LastLogin:    time.Now().UTC(),
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repository := db.Users()
		id := uuid.New()
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repository.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, users.ErrNoUser.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repository.Create(ctx, user1)
			require.NoError(t, err)

			userFromDB, err := repository.Get(ctx, user1.ID)
			require.NoError(t, err)
			compareUsers(t, user1, userFromDB)
		})

		t.Run("getByEmail", func(t *testing.T) {
			userFromDB, err := repository.GetByEmail(ctx, user1.Email)
			require.NoError(t, err)
			compareUsers(t, user1, userFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repository.Create(ctx, user2)
			require.NoError(t, err)

			allUsers, err := repository.List(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(allUsers), 2)
			compareUsers(t, user1, allUsers[0])
			compareUsers(t, user2, allUsers[1])
		})

		t.Run("update sql no rows", func(t *testing.T) {
			err := repository.Update(ctx, users.StatusSuspended, id)
			require.Error(t, err)
			require.Equal(t, users.ErrNoUser.Has(err), true)
		})

		t.Run("update", func(t *testing.T) {
			err := repository.Update(ctx, users.StatusSuspended, user1.ID)
			require.NoError(t, err)

			userFromDB, err := repository.Get(ctx, user1.ID)
			require.NoError(t, err)
			assert.Equal(t, users.StatusSuspended, userFromDB.Status)
		})

		t.Run("delete sql no rows", func(t *testing.T) {
			err := repository.Delete(ctx, id)
			require.Error(t, err)
			require.Equal(t, users.ErrNoUser.Has(err), true)
		})

		t.Run("delete", func(t *testing.T) {
			err := repository.Delete(ctx, user1.ID)
			require.NoError(t, err)
		})

		t.Run("get users nickname", func(t *testing.T) {
			nickname, err := repository.GetNickNameByID(ctx, user2.ID)
			require.NoError(t, err)

			assert.Equal(t, user2.NickName, nickname)
		})
	})
}

func compareUsers(t *testing.T, user1, user2 users.User) {
	assert.Equal(t, user1.ID, user2.ID)
	assert.Equal(t, user1.Email, user2.Email)
	assert.Equal(t, user1.PasswordHash, user2.PasswordHash)
	assert.Equal(t, user1.NickName, user2.NickName)
	assert.Equal(t, user1.FirstName, user2.FirstName)
	assert.Equal(t, user1.LastName, user2.LastName)
	assert.Equal(t, user1.Status, user2.Status)
	assert.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, 1*time.Second)
	assert.WithinDuration(t, user1.LastLogin, user2.LastLogin, 1*time.Second)
}
