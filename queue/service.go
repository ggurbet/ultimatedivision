// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/pagination"
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

// Create adds place in database.
func (service *Service) Create(ctx context.Context, place Place) error {
	if _, err := service.users.Get(ctx, place.UserID); err != nil {
		return ErrQueue.Wrap(err)
	}
	if place, err := service.Get(ctx, place.UserID); err == nil {
		return ErrQueue.New("the user " + string(place.Status) + " now")
	}
	return ErrQueue.Wrap(service.queues.Create(ctx, place))
}

// Get returns place from database.
func (service *Service) Get(ctx context.Context, id uuid.UUID) (Place, error) {
	queue, err := service.queues.Get(ctx, id)
	return queue, ErrQueue.Wrap(err)
}

// ListPaginated returns places in page from database.
func (service *Service) ListPaginated(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	queuesListPage, err := service.queues.ListPaginated(ctx, cursor)
	return queuesListPage, ErrQueue.Wrap(err)
}

// UpdateStatus updates place status in database.
func (service *Service) UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error {
	return ErrQueue.Wrap(service.queues.UpdateStatus(ctx, id, status))
}

// Finish finishes place for user in database.
func (service *Service) Finish(ctx context.Context, id uuid.UUID) error {
	return ErrQueue.Wrap(service.queues.Delete(ctx, id))
}
