// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/users"
)

// ErrQueue indicated that there was an error in service.
var ErrQueue = errs.Class("queue service error")

// Service is handling queues related logic.
//
// architecture: Service
type Service struct {
	config Config
	queues DB
	users  *users.Service
}

// NewService is a constructor for queues service.
func NewService(config Config, queues DB, users *users.Service) *Service {
	return &Service{
		config: config,
		queues: queues,
		users:  users,
	}
}

// Create adds client's queue in database.
func (service *Service) Create(ctx context.Context, client Client) error {
	if _, err := service.users.Get(ctx, client.UserID); err != nil {
		return ErrQueue.Wrap(err)
	}
	service.queues.Create(client)
	return nil
}

// Get returns client from database.
func (service *Service) Get(userID uuid.UUID) (Client, error) {
	queue, err := service.queues.Get(userID)
	return queue, ErrQueue.Wrap(err)
}

// List returns clients from database.
func (service *Service) List() []Client {
	return service.queues.List()
}

// Finish finishes client's queue in database.
func (service *Service) Finish(userID uuid.UUID) {
	service.queues.Delete(userID)
}
