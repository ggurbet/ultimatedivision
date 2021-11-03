// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package divisions

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrDivisions indicates that there was an error in the service.
var ErrDivisions = errs.Class("divisions service error")

// Service is handling divisions related logic.
//
// architecture: Service
type Service struct {
	divisions DB
	config    Config
}

// NewService is a constructor for divisions service.
func NewService(divisions DB, config Config) *Service {
	return &Service{
		divisions: divisions,
		config:    config,
	}
}

// Create creates a division.
func (service *Service) Create(ctx context.Context, name string) error {
	division := Division{
		ID:             uuid.New(),
		Name:           name,
		PassingPercent: service.config.PassingPercent,
		CreatedAt:      time.Now().UTC(),
	}

	return ErrDivisions.Wrap(service.divisions.Create(ctx, division))
}

// List returns all divisions from DB.
func (service *Service) List(ctx context.Context) ([]Division, error) {
	divisions, err := service.divisions.List(ctx)
	return divisions, ErrDivisions.Wrap(err)
}

// Get returns division from DB.
func (service *Service) Get(ctx context.Context, divisionID uuid.UUID) (Division, error) {
	division, err := service.divisions.Get(ctx, divisionID)
	return division, ErrDivisions.Wrap(err)
}

// Delete deletes a division.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrDivisions.Wrap(service.divisions.Delete(ctx, id))
}
