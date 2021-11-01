// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package subscribers

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/pagination"
)

// ErrSubscribers indicates that there was an error in the service.
var ErrSubscribers = errs.Class("subscribers service error")

// ErrSubscribersDB indicates that there was an error in the database.
var ErrSubscribersDB = errs.Class("subscribers repository error")

// Service is handling subscribers related logic.
//
// architecture: Service
type Service struct {
	subscribers DB
	config      Config
}

// NewService is a constructor for subscribers service.
func NewService(subscribers DB, config Config) *Service {
	return &Service{
		subscribers: subscribers,
		config:      config,
	}
}

// GetByEmail returns subscriber by email from DB.
func (service *Service) GetByEmail(ctx context.Context, email string) (Subscriber, error) {
	subscriber, err := service.subscribers.GetByEmail(ctx, email)
	return subscriber, ErrSubscribers.Wrap(err)
}

// List returns all subscribers from DB.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	subscribers, err := service.subscribers.List(ctx, cursor)
	return subscribers, ErrSubscribers.Wrap(err)
}

// Create creates a subscriber.
func (service *Service) Create(ctx context.Context, email string) error {
	subscriber := Subscriber{
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}

	return ErrSubscribers.Wrap(service.subscribers.Create(ctx, subscriber))
}

// Delete deletes a subscriber.
func (service *Service) Delete(ctx context.Context, email string) error {
	return ErrSubscribers.Wrap(service.subscribers.Delete(ctx, email))
}
