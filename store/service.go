// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package store

import (
	"context"

	"github.com/zeebo/errs"
)

// ErrStore indicated that there was an error in service.
var ErrStore = errs.Class("store service error")

// Service is handling store related logic.
//
// architecture: Service
type Service struct {
	store DB
}

// NewService is a constructor for store service.
func NewService(store DB) *Service {
	return &Service{
		store: store,
	}
}

// Create creates setting of store in database.
func (service *Service) Create(ctx context.Context, setting Setting) error {
	return ErrStore.Wrap(service.store.Create(ctx, setting))
}

// Get returns setting of store by id from database.
func (service *Service) Get(ctx context.Context, id int) (Setting, error) {
	setting, err := service.store.Get(ctx, id)
	return setting, ErrStore.Wrap(err)
}

// List returns settings of store from database.
func (service *Service) List(ctx context.Context) ([]Setting, error) {
	settings, err := service.store.List(ctx)
	return settings, ErrStore.Wrap(err)
}

// Update updates setting of store in database.
func (service *Service) Update(ctx context.Context, setting Setting) error {
	return ErrStore.Wrap(service.store.Update(ctx, setting))
}
