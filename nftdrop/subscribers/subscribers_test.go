// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package subscribers_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/nftdrop"
	"ultimatedivision/nftdrop/database/dbtesting"
	"ultimatedivision/nftdrop/subscribers"
	"ultimatedivision/pkg/pagination"
)

func TestEmails(t *testing.T) {
	subscriber1 := subscribers.Subscriber{
		Email:     "tarkovskynik@gmail.com",
		CreatedAt: time.Now(),
	}

	subscriber2 := subscribers.Subscriber{
		Email:     "3560876@gmail.com",
		CreatedAt: time.Now(),
	}

	cursor := pagination.Cursor{
		Limit: 2,
		Page:  1,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db nftdrop.DB) {
		repository := db.Subscribers()

		t.Run("getByEmail", func(t *testing.T) {
			err := repository.Create(ctx, subscriber1)
			require.NoError(t, err)

			emailFromDB, err := repository.GetByEmail(ctx, subscriber1.Email)
			require.NoError(t, err)
			compareEmails(t, subscriber1, emailFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repository.Create(ctx, subscriber2)
			require.NoError(t, err)

			allUsers, err := repository.List(ctx, cursor)
			assert.NoError(t, err)
			assert.Equal(t, len(allUsers.Subscribers), 2)
			compareEmails(t, subscriber1, allUsers.Subscribers[0])
			compareEmails(t, subscriber2, allUsers.Subscribers[1])
		})

		t.Run("delete", func(t *testing.T) {
			err := repository.Delete(ctx, subscriber1.Email)
			require.NoError(t, err)
		})
	})
}

func compareEmails(t *testing.T, subscriber1, subscriber2 subscribers.Subscriber) {
	assert.Equal(t, subscriber1.Email, subscriber2.Email)
}
