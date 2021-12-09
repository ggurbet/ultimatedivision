// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/queue"
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
	queueHub.hub.Queue = append(queueHub.hub.Queue, client)
}

// Get returns client from the hub of queue.
func (queueHub *queueHub) Get(userID uuid.UUID) (queue.Client, error) {
	for _, client := range queueHub.hub.Queue {
		if client.UserID == userID {
			return client, nil
		}
	}
	// TODO: change error
	return queue.Client{}, queue.ErrNoClient.New("not found user's values")
}

// List returns clients from the hub of queue.
func (queueHub *queueHub) List() []queue.Client {
	return queueHub.hub.Queue
}

// ListNotPlayingUsers returns clients who don't play game from database.
func (queueHub *queueHub) ListNotPlayingUsers() []queue.Client {
	var listNotPlayingUsers []queue.Client
	for _, client := range queueHub.hub.Queue {
		if client.IsPlaying == false {
			listNotPlayingUsers = append(listNotPlayingUsers, client)
		}
	}
	return listNotPlayingUsers
}

// UpdateIsPlaying updates is playing status of client in database.
func (queueHub *queueHub) UpdateIsPlaying(userID uuid.UUID, isPlaying bool) error {
	for k, client := range queueHub.hub.Queue {
		if client.UserID == userID {
			queueHub.hub.Queue[k].IsPlaying = isPlaying
			return nil
		}
	}
	return queue.ErrNoClient.New("not found user's values")
}

// Delete deletes record client in the hub of queue.
func (queueHub *queueHub) Delete(userID uuid.UUID) error {
	for k, client := range queueHub.hub.Queue {
		if client.UserID == userID {
			if k+1 == len(queueHub.hub.Queue) {
				queueHub.hub.Queue = queueHub.hub.Queue[:k]
				return nil
			}
			queueHub.hub.Queue = append(queueHub.hub.Queue[:k], queueHub.hub.Queue[k+1:]...)
			return nil
		}
	}
	return queue.ErrNoClient.New("not found user's values")
}
