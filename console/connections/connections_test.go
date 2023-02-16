// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package connections_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/console/connections"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/users"
)

func TestConnections(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		LastLogin:    time.Now().UTC(),
		Status:       0,
		CreatedAt:    time.Now().UTC(),
	}

	connection := websocket.Conn{}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryConnections := db.Connections()

		t.Run("connection does not exist", func(t *testing.T) {
			_, err := repositoryConnections.Get(user1.ID)
			require.Error(t, err)
			assert.Equal(t, true, connections.ErrNoConnection.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryConnections.Create(user1.ID, &connection)
			require.NoError(t, err)

			_, err = repositoryConnections.Get(user1.ID)
			require.NoError(t, err)
		})

		t.Run("list", func(t *testing.T) {
			allConnections := repositoryConnections.List()
			require.Len(t, allConnections, 1)
		})

		t.Run("delete connection does not exist", func(t *testing.T) {
			err := repositoryConnections.Delete(uuid.UUID{})
			require.Error(t, err)
			require.Equal(t, connections.ErrNoConnection.Has(err), true)
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryConnections.Delete(user1.ID)
			require.NoError(t, err)
		})
	})
}
