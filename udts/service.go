// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package udts

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrUDTs indicated that there was an error in service.
var ErrUDTs = errs.Class("UDTs service error")

// Service is handling UDTs related logic.
//
// architecture: Service.
type Service struct {
	udts DB
}

// NewService is a constructor for UDTs service.
func NewService(udts DB) *Service {
	return &Service{
		udts: udts,
	}
}

// Create creates udt in the database.
func (service *Service) Create(ctx context.Context, udt UDT) error {
	return ErrUDTs.Wrap(service.udts.Create(ctx, udt))
}

// GetByUserID returns udt by user id from database.
func (service *Service) GetByUserID(ctx context.Context, userID uuid.UUID) (UDT, error) {
	udt, err := service.udts.GetByUserID(ctx, userID)
	return udt, ErrUDTs.Wrap(err)
}

// List returns udts from database.
func (service *Service) List(ctx context.Context) ([]UDT, error) {
	udts, err := service.udts.List(ctx)
	return udts, ErrUDTs.Wrap(err)
}

// Update updates udt by user's id in the database.
func (service *Service) Update(ctx context.Context, udt UDT) error {
	return ErrUDTs.Wrap(service.udts.Update(ctx, udt))
}

// Delete deletes udt by user's id in the database.
func (service *Service) Delete(ctx context.Context, userID uuid.UUID) error {
	return ErrUDTs.Wrap(service.udts.Delete(ctx, userID))
}
