// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/queue"
)

// ensures that queueHub implements queue.DB.
var _ queue.DB = (*queueHub)(nil)

// ErrQueue indicates that there was an error in the hub.
var ErrQueue = errs.Class("queues repository error")

// queueHub provides access to queue hub.
//
// architecture: Hub
type queueHub struct {
	hub *Hub
}

// Create adds client in the hub of queue.
func (queueHub *queueHub) Create(client queue.Client) {
	queueHub.hub.Queue[client.UserID] = client.Conn
}

// Get returns client from the hub of queue.
func (queueHub *queueHub) Get(userID uuid.UUID) (queue.Client, error) {
	var client queue.Client
	if _, ok := queueHub.hub.Queue[userID]; !ok {
		return client, queue.ErrNoClient.New("not found user's websocket connection")
	}
	client = queue.Client{
		UserID: userID,
		Conn:   queueHub.hub.Queue[userID],
	}
	return client, nil
}

// List returns clients from the hub of queue.
func (queueHub *queueHub) List() []queue.Client {
	clients := []queue.Client{}
	for userID, conn := range queueHub.hub.Queue {
		client := queue.Client{
			UserID: userID,
			Conn:   conn,
		}
		clients = append(clients, client)
	}
	return clients
}

// Delete deletes record client in the hub of queue.
func (queueHub *queueHub) Delete(userID uuid.UUID) {
	delete(queueHub.hub.Queue, userID)
}
