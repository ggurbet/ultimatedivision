// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoUser indicated that user does not exist.
var ErrNoUser = errs.Class("user does not exist")

// DB exposes access to users db.
//
// architecture: DB
type DB interface {
	// List returns all users from the data base.
	List(ctx context.Context) ([]User, error)
	// Get returns user by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (User, error)
	// GetByEmail returns user by email from the data base.
	GetByEmail(ctx context.Context, email string) (User, error)
	// Create creates a user and writes to the database.
	Create(ctx context.Context, user User) error
}

// Status defines the list of possible user statuses.
type Status int

const (
	// StatusActive indicates that user can login to the account.
	StatusActive Status = 0
	// StatusSuspended indicates that user cannot login to the account.
	StatusSuspended Status = 1
)

// User describes user entity.
type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"passwordHash"`
	NickName     string    `json:"nickName"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	LastLogin    time.Time `json:"lastLogin"`
	Status       Status    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
}
