// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrUsers indicates that there was an error in the service.
var ErrUsers = errs.Class("users service error")

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
func (service *Service) Create(ctx context.Context, email, password, nickName, firstName, lastName string) error {
	user := User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: []byte(password),
		NickName:     nickName,
		FirstName:    firstName,
		LastName:     lastName,
		LastLogin:    time.Time{},
		Status:       StatusActive,
		CreatedAt:    time.Now(),
	}
	err := user.EncodePass()
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return service.users.Create(ctx, user)
}

// Delete deletes a user.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return service.users.Delete(ctx, id)
}

// Update updates a users status.
func (service *Service) Update(ctx context.Context, status int, id uuid.UUID) error {
	return service.users.Update(ctx, status, id)
}
