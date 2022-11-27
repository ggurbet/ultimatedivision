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
	"ultimatedivision/gameplay/queue"
	"ultimatedivision/users"
)

func TestQueue(t *testing.T) {
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

	queueClient1 := queue.Client{UserID: user1.ID, Connection: nil}
	queueClient2 := queue.Client{UserID: user2.ID, Connection: nil}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryQueue := db.Queue()
		repositoryUsers := db.Users()
		userID := uuid.New()

		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryQueue.Get(userID)
			require.Error(t, err)
			assert.Equal(t, true, queue.ErrNoClient.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			repositoryQueue.Create(queueClient1)

			queueFromDB, err := repositoryQueue.Get(user1.ID)
			require.NoError(t, err)
			compareQueues(t, queueClient1, queueFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			repositoryQueue.Create(queueClient2)

			queueList := repositoryQueue.List()
			assert.Equal(t, len(queueList), 2)
			compareQueues(t, queueClient1, queueList[0])
			compareQueues(t, queueClient2, queueList[1])
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryQueue.Delete(queueClient1.UserID)
			require.NoError(t, err)

			queueList := repositoryQueue.List()
			assert.Equal(t, len(queueList), 1)
			compareQueues(t, queueClient2, queueList[0])
		})
	})
}

func compareQueues(t *testing.T, queue1, queue2 queue.Client) {
	assert.Equal(t, queue1.UserID, queue2.UserID)
}

func TestDivideClients(t *testing.T) {
	type testDividing struct {
		incomeData []queue.Client
		result     [][]queue.Client
	}

	client1 := queue.Client{
		UserID:     uuid.New(),
		Connection: nil,
		SquadID:    uuid.New(),
		IsPlaying:  true,
		CreatedAt:  time.Now().UTC(),
	}

	client2 := queue.Client{
		UserID:     uuid.New(),
		Connection: nil,
		SquadID:    uuid.New(),
		IsPlaying:  false,
		CreatedAt:  time.Now().UTC(),
	}

	testCases := []testDividing{
		{
			incomeData: nil,
			result:     nil,
		},
		{
			incomeData: []queue.Client{client1, client2},
			result:     [][]queue.Client{{client1, client2}},
		},
	}

	for _, testCase := range testCases {
		result := queue.DivideClients(testCase.incomeData)

		if testCase.result == nil && result == nil {
			continue
		}

		compareClients(t, result, testCase.result)
	}
}

func compareClients(t *testing.T, result, expectedResult [][]queue.Client) {
	assert.Equal(t, len(result), len(expectedResult))

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			assert.Equal(t, result[i][j].Connection, expectedResult[i][j].Connection)
			assert.Equal(t, result[i][j].UserID, expectedResult[i][j].UserID)
			assert.Equal(t, result[i][j].SquadID, expectedResult[i][j].SquadID)
			assert.Equal(t, result[i][j].IsPlaying, expectedResult[i][j].IsPlaying)
		}
	}
}
