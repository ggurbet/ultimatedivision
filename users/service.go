// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"

	"github.com/google/uuid"
)

// Service is handling users related logic.
//
// architecture: Service
type Service struct {
	users DB
}

// NewService is a constructor for users service.
func NewService(users DB) *Service {
	return &Service{
		users: users,
	}
}

// Get returns user from DB.
func (service *Service) Get(ctx context.Context, userID uuid.UUID) (User, error) {
	return service.users.Get(ctx, userID)
}

// GetByEmail returns user by email from DB.
func (service *Service) GetByEmail(ctx context.Context, email string) (User, error) {
	return service.users.GetByEmail(ctx, email)
}

// List returns all users from DB.
func (service *Service) List(ctx context.Context) ([]User, error) {
	return service.users.List(ctx)
}

// Create creates a user and returns user email.
func (service *Service) Create(ctx context.Context, user User) error {
	return service.users.Create(ctx, user)
}
