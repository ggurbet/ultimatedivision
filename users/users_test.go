// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users_test

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"ultimatedivision"
	"ultimatedivision/console/emails"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/pkg/velas"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
)

func TestUsers(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		Wallet:       "0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b7",
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
		Wallet:       "0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b8",
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

		t.Run("get users nickname", func(t *testing.T) {
			nickname, err := repository.GetNickNameByID(ctx, user2.ID)
			require.NoError(t, err)

			assert.Equal(t, user2.NickName, nickname)
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

		t.Run("update nonce", func(t *testing.T) {
			nonce := make([]byte, 32)
			_, err := rand.Read(nonce)
			require.NoError(t, err)

			err = repository.UpdateNonce(ctx, user1.ID, nonce)
			require.NoError(t, err)

			userFromDB, err := repository.GetByWalletAddress(ctx, user1.Wallet, users.Wallet)
			require.NoError(t, err)
			assert.Equal(t, nonce, userFromDB.Nonce)
		})

		t.Run("update wallet address sql no rows", func(t *testing.T) {
			err := repository.UpdateWalletAddress(ctx, evmsignature.Address("wallet_address"), uuid.New())
			require.Error(t, err)
			assert.Equal(t, true, users.ErrNoUser.Has(err))
		})

		t.Run("update wallet address", func(t *testing.T) {
			err := repository.UpdateWalletAddress(ctx, "wallet_address", user1.ID)
			require.NoError(t, err)
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
	})
}

func TestUserService(t *testing.T) {
	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		var log logger.Logger
		userService := users.NewService(db.Users())
		authService := userauth.NewService(db.Users(), auth.TokenSigner{}, &emails.Service{}, log, &velas.Service{})
		wallet := evmsignature.Address("0x2346b33F2E379dDA22b2563B009382a0Fc9aA926")
		signature, err := hexutil.Decode("0x3b8566bdf04b3e11e825196e6d25c436814c43e40a44db89b17450d286b2e531640309222d2a30a9235f9521b9cba34ca6d601d9e04ec0f75466c33fe2c720b71b")
		require.NoError(t, err)

		testUser1 := users.User{
			Email:        "testUser@gmail.com",
			PasswordHash: []byte("my-pass"),
			NickName:     "Nik",
			FirstName:    "Nikita",
			LastName:     "Test",
		}

		t.Run("create", func(t *testing.T) {
			err := userService.Create(ctx, testUser1.Email, "my-pass", testUser1.NickName, testUser1.FirstName, testUser1.LastName)
			require.NoError(t, err)
		})

		t.Run("check is password correct", func(t *testing.T) {
			allUsers, err := userService.List(ctx)
			require.NoError(t, err)
			require.Equal(t, len(allUsers), 1)

			err = bcrypt.CompareHashAndPassword(allUsers[0].PasswordHash, testUser1.PasswordHash)
			require.NoError(t, err)
		})

		t.Run("Register user", func(t *testing.T) {
			err := authService.RegisterWithMetamask(ctx, signature)
			require.NoError(t, err)
		})

		t.Run("get nonce", func(t *testing.T) {
			user, err := db.Users().GetByWalletAddress(ctx, "0x2346b33F2E379dDA22b2563B009382a0Fc9aA926", users.Wallet)
			require.NoError(t, err)
			userNonce := hexutil.Encode(user.Nonce)

			nonce, err := authService.Nonce(ctx, wallet, users.Wallet)
			require.NoError(t, err)
			require.Equal(t, nonce, userNonce)
		})

		t.Run("login", func(t *testing.T) {
			nonce := "0xfd71049a838a6f7983dae92608f85d56f42318d5d481697c714c634be4875315"
			sign := "0x64bfdf19719d55fde70a9f9e3b1087d2429b54ed712142855d9b98d0db7643f840ca438eb9cc6363fbc772a6c11030a5ae30d9b78ab6f19fed742910caa29ab31c"
			userNonce, err := hexutil.Decode(nonce)
			require.NoError(t, err)
			signature, err := hexutil.Decode(sign)
			require.NoError(t, err)

			user := users.User{
				ID:           uuid.New(),
				Wallet:       evmsignature.Address("0x2346b33F2E379dDA22b2563B009382a0Fc9aA926"),
				Nonce:        userNonce,
				Email:        "",
				PasswordHash: nil,
				LastLogin:    time.Time{},
				Status:       1,
				CreatedAt:    time.Time{},
			}

			oldUser, err := db.Users().GetByWalletAddress(ctx, "0x2346b33F2E379dDA22b2563B009382a0Fc9aA926", users.Wallet)
			require.NoError(t, err)

			err = db.Users().Delete(ctx, oldUser.ID)
			require.NoError(t, err)

			err = db.Users().Create(ctx, user)
			require.NoError(t, err)

			_, err = authService.LoginWithMetamask(ctx, nonce, signature)
			require.NoError(t, err)
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
