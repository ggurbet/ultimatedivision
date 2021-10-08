// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/nftdrop"
	"ultimatedivision/nftdrop/database/dbtesting"
	"ultimatedivision/nftdrop/whitelist"
)

func TestWhitelists(t *testing.T) {
	whitelist1 := whitelist.Whitelist{
		Address:  "address1",
		Password: []byte{},
	}

	whitelist2 := whitelist.Whitelist{
		Address:  "address2",
		Password: []byte{},
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db nftdrop.DB) {
		repositoryWhitelist := db.Whitelist()

		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryWhitelist.GetByAddress(ctx, "address0")
			require.Error(t, err)
			assert.Equal(t, true, whitelist.ErrNoWhitelist.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryWhitelist.Create(ctx, whitelist1)
			require.NoError(t, err)

			whitelistFromDB, err := repositoryWhitelist.GetByAddress(ctx, whitelist1.Address)
			require.NoError(t, err)
			compareWhitelists(t, whitelist1, whitelistFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repositoryWhitelist.Create(ctx, whitelist2)
			require.NoError(t, err)

			whitelistRecordsFromDB, err := repositoryWhitelist.List(ctx)
			require.NoError(t, err)
			compareWhitelists(t, whitelist1, whitelistRecordsFromDB[0])
			compareWhitelists(t, whitelist2, whitelistRecordsFromDB[1])
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryWhitelist.Delete(ctx, whitelist1.Address)
			require.NoError(t, err)

			err = repositoryWhitelist.Delete(ctx, whitelist2.Address)
			require.NoError(t, err)
		})
	})
}

func compareWhitelists(t *testing.T, whitelist1, whitelist2 whitelist.Whitelist) {
	assert.Equal(t, whitelist1.Address, whitelist2.Address)
	assert.Equal(t, whitelist1.Password, whitelist2.Password)
}
