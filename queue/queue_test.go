// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/internal/pagination"
	"ultimatedivision/queue"
	"ultimatedivision/users"
)

func TestQueues(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		LastLogin:    time.Now(),
		Status:       0,
		CreatedAt:    time.Now(),
	}

	user2 := users.User{
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

	queuePlace1 := queue.Place{
		UserID: user1.ID,
		Status: queue.StatusSearches,
	}

	queuePlace2 := queue.Place{
		UserID: user2.ID,
		Status: queue.StatusGames,
	}

	cursor1 := pagination.Cursor{
		Limit: 2,
		Page:  1,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryQueue := db.Queue()
		repositoryUsers := db.Users()
		id := uuid.New()
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryQueue.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, queue.ErrNoPlace.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryQueue.Create(ctx, queuePlace1)
			require.NoError(t, err)

			queueFromDB, err := repositoryQueue.Get(ctx, user1.ID)
			require.NoError(t, err)
			compareQueues(t, queuePlace1, queueFromDB)
		})

		t.Run("list paginated", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			err = repositoryQueue.Create(ctx, queuePlace2)
			require.NoError(t, err)

			queueList, err := repositoryQueue.ListPaginated(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(queueList.Places), 2)
			compareQueues(t, queuePlace1, queueList.Places[0])
			compareQueues(t, queuePlace2, queueList.Places[1])
		})

		t.Run("update status", func(t *testing.T) {
			queuePlace1.Status = queue.StatusGames
			err := repositoryQueue.UpdateStatus(ctx, queuePlace1.UserID, queuePlace1.Status)
			require.NoError(t, err)

			queueList, err := repositoryQueue.ListPaginated(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(queueList.Places), 2)
			compareQueues(t, queuePlace1, queueList.Places[1])
			compareQueues(t, queuePlace2, queueList.Places[0])
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryQueue.Delete(ctx, queuePlace1.UserID)
			require.NoError(t, err)

			queueList, err := repositoryQueue.ListPaginated(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(queueList.Places), 1)
			compareQueues(t, queuePlace2, queueList.Places[0])
		})
	})
}

func compareQueues(t *testing.T, queue1, queue2 queue.Place) {
	assert.Equal(t, queue1.UserID, queue2.UserID)
	assert.Equal(t, queue1.Status, queue2.Status)
}
