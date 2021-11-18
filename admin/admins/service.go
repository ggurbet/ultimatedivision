// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrAdmins indicates that there was an error in the service.
var ErrAdmins = errs.Class("admins service error")

// Service is handling admins related logic.
//
// architecture: Service
type Service struct {
	admins DB
}

// NewService is constructor for Service.
func NewService(admins DB) *Service {
	return &Service{
		admins: admins,
	}
}

// List returns all admins from DB.
func (service *Service) List(ctx context.Context) ([]Admin, error) {
	allAdmins, err := service.admins.List(ctx)
	return allAdmins, ErrAdmins.Wrap(err)
}

// Get returns admin from DB.
func (service *Service) Get(ctx context.Context, id uuid.UUID) (Admin, error) {
	admin, err := service.admins.Get(ctx, id)
	return admin, ErrAdmins.Wrap(err)
}

// Create insert admin to DB.
func (service *Service) Create(ctx context.Context, email string, password []byte) error {
	admin := Admin{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: password,
		CreatedAt:    time.Now().UTC(),
	}

	err := admin.EncodePass()
	if err != nil {
		return ErrAdmins.Wrap(err)
	}

	return ErrAdmins.Wrap(service.admins.Create(ctx, admin))
}

// Update updates admin from DB.
func (service *Service) Update(ctx context.Context, id uuid.UUID, newPassword []byte) error {
	admin, err := service.Get(ctx, id)
	if err != nil {
		return ErrAdmins.Wrap(err)
	}

	err = admin.EncodePass()
	if err != nil {
		return ErrAdmins.Wrap(err)
	}

	return ErrAdmins.Wrap(service.admins.Update(ctx, admin))
}
