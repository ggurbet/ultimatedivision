// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"strings"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrUsers indicates that there was an error in the service.
var ErrUsers = errs.Class("users service error")

// ErrUnauthenticated should be returned when user performs unauthenticated action.
var ErrUnauthenticated = errs.Class("user unauthenticated error")

// ErrWalletAddressAlreadyInUse should be returned when users wallet address is already in use.
var ErrWalletAddressAlreadyInUse = errs.Class("wallet address is already in use")

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
func (service *Service) GetByWalletAddress(ctx context.Context, walletAddress evmsignature.Address) (User, error) {
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
		Wallet:    user.Wallet,
	}, nil
}

// GetNickNameByID returns nickname of user.
func (service *Service) GetNickNameByID(ctx context.Context, id uuid.UUID) (string, error) {
	nickname, err := service.users.GetNickNameByID(ctx, id)

	return nickname, ErrUsers.Wrap(err)
}

// UpdateWalletAddress updates wallet address.
func (service *Service) UpdateWalletAddress(ctx context.Context, wallet evmsignature.Address, id uuid.UUID) error {
	wallet = evmsignature.Address(strings.ToLower(string(wallet)))

	_, err := service.GetByWalletAddress(ctx, wallet)
	if err == nil {
		return ErrWalletAddressAlreadyInUse.New("wallet address already in use")
	}

	return ErrUsers.Wrap(service.users.UpdateWalletAddress(ctx, wallet, id))
}

// ChangeWalletAddress changes wallet address.
func (service *Service) ChangeWalletAddress(ctx context.Context, wallet evmsignature.Address, id uuid.UUID) error {
	wallet = evmsignature.Address(strings.ToLower(string(wallet)))

	user, err := service.GetByWalletAddress(ctx, wallet)
	if err != nil {
		return ErrUsers.Wrap(err)
	}
	if user.ID == id {
		return ErrUsers.New("this address is used by you")
	}
	emptyWallet := evmsignature.Address("")
	err = service.users.UpdateWalletAddress(ctx, emptyWallet, user.ID)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return ErrUsers.Wrap(service.users.UpdateWalletAddress(ctx, wallet, id))
}
