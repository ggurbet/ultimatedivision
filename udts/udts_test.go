// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package udts_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/udts"
	"ultimatedivision/users"
)

func TestUDTs(t *testing.T) {
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

	user2 := users.User{
		ID:           uuid.New(),
		Email:        "356087612@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "qwerty",
		FirstName:    "Stas12",
		LastName:     "Isakov12",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	udt1 := udts.UDT{
		UserID: user1.ID,
		Value:  *big.NewInt(100000000000000),
		Nonce:  0,
	}

	udt2 := udts.UDT{
		UserID: user2.ID,
		Value:  *big.NewInt(100000000000000),
		Nonce:  0,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryUsers := db.Users()
		repositoryUDTs := db.UDTs()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			err = repositoryUDTs.Create(ctx, udt1)
			require.NoError(t, err)

			err = repositoryUDTs.Create(ctx, udt2)
			require.NoError(t, err)
		})

		t.Run("Get", func(t *testing.T) {
			udtGet, err := repositoryUDTs.Get(ctx, udt1.UserID)
			require.NoError(t, err)

			compareUDTsSlice(t, []udts.UDT{udtGet}, []udts.UDT{udt1})
		})

		t.Run("List", func(t *testing.T) {
			udtList, err := repositoryUDTs.List(ctx)
			require.NoError(t, err)

			compareUDTsSlice(t, udtList, []udts.UDT{udt1, udt2})
		})

		t.Run("Update", func(t *testing.T) {
			udt1.Value = *big.NewInt(200000000000000)
			udt1.Nonce = 1
			err := repositoryUDTs.Update(ctx, udt1)
			require.NoError(t, err)

			udtGet, err := repositoryUDTs.Get(ctx, udt1.UserID)
			require.NoError(t, err)

			compareUDTsSlice(t, []udts.UDT{udtGet}, []udts.UDT{udt1})
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryUDTs.Delete(ctx, udt1.UserID)
			require.NoError(t, err)

			udtList, err := repositoryUDTs.List(ctx)
			require.NoError(t, err)

			compareUDTsSlice(t, udtList, []udts.UDT{udt2})
		})
	})
}

func compareUDTsSlice(t *testing.T, udt1, udt2 []udts.UDT) {
	assert.Equal(t, len(udt1), len(udt2))

	for i := 0; i < len(udt1); i++ {
		assert.Equal(t, udt1[i].UserID, udt2[i].UserID)
		assert.Equal(t, udt1[i].Value, udt2[i].Value)
		assert.Equal(t, udt1[i].Nonce, udt2[i].Nonce)
	}
}
