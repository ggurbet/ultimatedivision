// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/admin/admins"
	"ultimatedivision/database/dbtesting"
)

func TestAdmin(t *testing.T) {
	admin1 := admins.Admin{
		ID:           uuid.New(),
		Email:        "admin1@gmail.com",
		PasswordHash: []byte{0},
		CreatedAt:    time.Now(),
	}

	admin2 := admins.Admin{
		ID:           uuid.New(),
		Email:        "admin2@gmail.com",
		PasswordHash: []byte{1},
		CreatedAt:    time.Now(),
	}

	updatedAdmin := admins.Admin{
		ID:           admin1.ID,
		Email:        admin1.Email,
		PasswordHash: []byte{3},
		CreatedAt:    admin1.CreatedAt,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repository := db.Admins()
		id := uuid.New()

		t.Run("Get sql no rows", func(t *testing.T) {
			_, err := repository.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, admins.ErrNoAdmin.Has(err))
		})

		t.Run("Get", func(t *testing.T) {
			err := repository.Create(ctx, admin1)
			require.NoError(t, err)

			adminFromDB, err := repository.Get(ctx, admin1.ID)
			require.NoError(t, err)
			compareAdmins(t, adminFromDB, admin1)
		})

		t.Run("List", func(t *testing.T) {
			err := repository.Create(ctx, admin2)
			require.NoError(t, err)

			allAdmins, err := repository.List(ctx)
			require.NoError(t, err)
			compareAdmins(t, allAdmins[0], admin1)
			compareAdmins(t, allAdmins[1], admin2)
		})

		t.Run("Update", func(t *testing.T) {
			err := repository.Update(ctx, updatedAdmin)
			require.NoError(t, err)

			adminFromDB, err := repository.Get(ctx, admin1.ID)
			require.NoError(t, err)

			compareAdmins(t, adminFromDB, updatedAdmin)
		})
	})
}

func compareAdmins(t *testing.T, adminFromDB admins.Admin, testAdmin admins.Admin) {
	assert.Equal(t, adminFromDB.ID, testAdmin.ID)
	assert.Equal(t, adminFromDB.Email, testAdmin.Email)
	assert.Equal(t, adminFromDB.PasswordHash, testAdmin.PasswordHash)
	assert.Equal(t, adminFromDB.CreatedAt, testAdmin.CreatedAt)
}
