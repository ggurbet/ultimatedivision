// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencywaitlist_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

func TestCurrencycurrencywaitlist(t *testing.T) {
	var value = new(big.Int)
	value.SetString("5000000000000000000", 10)
	item1 := currencywaitlist.Item{
		WalletAddress:       common.HexToAddress("0x96216849c49358b10257cb55b28ea603c874b05e"),
		CasperWalletAddress: "0202b2a13f20e71016aa8bbca5fbc1cca56af4092c926490e65dcfab2168ab051c42",
		Value:               *value,
		Nonce:               1,
		Signature:           "",
	}

	item2 := currencywaitlist.Item{
		WalletAddress:       common.HexToAddress("0x96216849c49358b10257cb55b28ea603c874b05e"),
		CasperWalletAddress: "0202b2a13f20e71016aa8bbca5fbc1cca56af4092c926490e65dcfab2168ab051c92",
		WalletType:          users.WalletTypeCasper,
		Value:               *value,
		Nonce:               2,
		Signature:           "",
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryCurrencyWaitList := db.CurrencyWaitList()

		t.Run("GetNonce without any entries", func(t *testing.T) {
			nonce, err := repositoryCurrencyWaitList.GetNonce(ctx)
			require.NoError(t, err)

			assert.Equal(t, int64(0), nonce)
		})

		t.Run("Create", func(t *testing.T) {
			err := repositoryCurrencyWaitList.Create(ctx, item1)
			require.NoError(t, err)

			err = repositoryCurrencyWaitList.Create(ctx, item2)
			require.NoError(t, err)
		})

		t.Run("GetByWalletAddressAndNonce", func(t *testing.T) {
			item, err := repositoryCurrencyWaitList.GetByWalletAddressAndNonce(ctx, item1.WalletAddress, item1.Nonce)
			require.NoError(t, err)

			compareItem(t, item, item1)
		})

		t.Run("GetByCasperWalletAddressAndNonce", func(t *testing.T) {
			item, err := repositoryCurrencyWaitList.GetByCasperWalletAddressAndNonce(ctx, item1.CasperWalletAddress, item1.Nonce)
			require.NoError(t, err)

			compareItem(t, item, item1)
		})

		t.Run("GetNonce", func(t *testing.T) {
			nonce, err := repositoryCurrencyWaitList.GetNonce(ctx)
			require.NoError(t, err)

			assert.Equal(t, item2.Nonce, nonce)
		})

		t.Run("List", func(t *testing.T) {
			itemList, err := repositoryCurrencyWaitList.List(ctx)
			require.NoError(t, err)

			compareItemsSlice(t, itemList, []currencywaitlist.Item{item1, item2})
		})

		t.Run("ListWithoutSignature", func(t *testing.T) {
			itemList, err := repositoryCurrencyWaitList.ListWithoutSignature(ctx)
			require.NoError(t, err)

			compareItemsSlice(t, itemList, []currencywaitlist.Item{item1, item2})
		})

		t.Run("UpdateSignature", func(t *testing.T) {
			item1.Signature = evmsignature.Signature("707fb93c61be8d54c6d1fdf4b83c8642831c480194f7cc93ebdd6fe1ac7474ae63efd077cf6398bf00dc0f7ea96be9f3f9a05dfac1382c4d2f1bb11ec46148491b")
			err := repositoryCurrencyWaitList.UpdateSignature(ctx, item1.Signature, item1.WalletAddress, item1.Nonce)
			require.NoError(t, err)

			itemList, err := repositoryCurrencyWaitList.List(ctx)
			require.NoError(t, err)

			compareItemsSlice(t, itemList, []currencywaitlist.Item{item2, item1})
		})

		t.Run("Update", func(t *testing.T) {
			item1.Signature = evmsignature.Signature("")

			var value = new(big.Int)
			value.SetString("25000000000000000000", 10)
			item1.Value = *value

			err := repositoryCurrencyWaitList.Update(ctx, item1)
			require.NoError(t, err)

			itemList, err := repositoryCurrencyWaitList.List(ctx)
			require.NoError(t, err)

			compareItemsSlice(t, itemList, []currencywaitlist.Item{item2, item1})
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryCurrencyWaitList.Delete(ctx, item1.WalletAddress, item1.Nonce)
			require.NoError(t, err)

			itemList, err := repositoryCurrencyWaitList.List(ctx)
			require.NoError(t, err)

			compareItemsSlice(t, itemList, []currencywaitlist.Item{item2})
		})

		t.Run("UpdateNonceByWallet", func(t *testing.T) {
			err := repositoryCurrencyWaitList.UpdateNonceByWallet(ctx, 5, item2.CasperWalletAddress)
			require.NoError(t, err)

			item2.Nonce = 5
			itemList, err := repositoryCurrencyWaitList.List(ctx)
			require.NoError(t, err)

			compareItemsSlice(t, itemList, []currencywaitlist.Item{item2})
			require.NoError(t, err)
		})
	})
}

func compareItemsSlice(t *testing.T, item1, item2 []currencywaitlist.Item) {
	assert.Equal(t, len(item1), len(item2))

	for i := 0; i < len(item1); i++ {
		assert.Equal(t, item1[i].WalletAddress, item2[i].WalletAddress)
		assert.Equal(t, item1[i].Value, item2[i].Value)
		assert.Equal(t, item1[i].Nonce, item2[i].Nonce)
		assert.Equal(t, item1[i].Signature, item2[i].Signature)
	}
}

func compareItem(t *testing.T, item1, item2 currencywaitlist.Item) {
	assert.Equal(t, item1.WalletAddress, item2.WalletAddress)
	assert.Equal(t, item1.Value, item2.Value)
	assert.Equal(t, item1.Nonce, item2.Nonce)
	assert.Equal(t, item1.Signature, item2.Signature)
}
