// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrUsers indicates that there was an error in the service.
var ErrUsers = errs.Class("users service error")

// ErrUnauthenticated should be returned when user performs unauthenticated action.
var ErrUnauthenticated = errs.Class("user unauthenticated error")

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
	user, err := service.users.Get(ctx, userID)
	return user, ErrUsers.Wrap(err)
}

// GetByEmail returns user by email from DB.
func (service *Service) GetByEmail(ctx context.Context, email string) (User, error) {
	user, err := service.users.GetByEmail(ctx, email)
	return user, ErrUsers.Wrap(err)
}

// GetByWalletAddress returns user by wallet address from the data base.
func (service *Service) GetByWalletAddress(ctx context.Context, walletAddress cryptoutils.Address) (User, error) {
	user, err := service.users.GetByWalletAddress(ctx, walletAddress)
	return user, ErrUsers.Wrap(err)
}

// List returns all users from DB.
func (service *Service) List(ctx context.Context) ([]User, error) {
	users, err := service.users.List(ctx)
	return users, ErrUsers.Wrap(err)
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
		Status:       StatusCreated,
		CreatedAt:    time.Now().UTC(),
	}
	err := user.EncodePass()
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return ErrUsers.Wrap(service.users.Create(ctx, user))
}

// Delete deletes a user.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrUsers.Wrap(service.users.Delete(ctx, id))
}

// Update updates a users status.
func (service *Service) Update(ctx context.Context, status Status, id uuid.UUID) error {
	return ErrUsers.Wrap(service.users.Update(ctx, status, id))
}

// GetProfile returns user profile.
func (service *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*Profile, error) {
	user, err := service.users.Get(ctx, userID)
	if err != nil {
		return nil, ErrUsers.Wrap(err)
	}

	return &Profile{
		Email:     user.Email,
		NickName:  user.NickName,
		CreatedAt: user.CreatedAt,
		LastLogin: user.LastLogin,
	}, nil
}

// GetNickNameByID returns nickname of user.
func (service *Service) GetNickNameByID(ctx context.Context, id uuid.UUID) (string, error) {
	nickname, err := service.users.GetNickNameByID(ctx, id)

	return nickname, ErrUsers.Wrap(err)
}

// UpdateWalletAddress updates wallet address.
func (service *Service) UpdateWalletAddress(ctx context.Context, wallet cryptoutils.Address, id uuid.UUID) error {
	wallet = cryptoutils.Address(strings.ToLower(string(wallet)))
	return ErrUsers.Wrap(service.users.UpdateWalletAddress(ctx, wallet, id))
}
