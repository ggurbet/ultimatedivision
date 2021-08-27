// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"
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
	// Update updates a status in the database.
	Update(ctx context.Context, status int, id uuid.UUID) error
	// UpdatePassword updates a password in the database.
	UpdatePassword(ctx context.Context, passwordHash []byte, id uuid.UUID) error
	// Delete deletes a user in the database.
	Delete(ctx context.Context, id uuid.UUID) error
	// GetNickNameByID returns nickname by user id from the database.
	GetNickNameByID(ctx context.Context, id uuid.UUID) (string, error)
}

// Status defines the list of possible user statuses.
type Status int

const (
	// StatusCreated indicates that user email is created.
	StatusCreated Status = 0
	// StatusActive indicates that user can login to the account.
	StatusActive Status = 1
	// StatusSuspended indicates that user cannot login to the account.
	StatusSuspended Status = 2
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

// EncodePass encode the password and generate "hash" to store from users password.
func (user *User) EncodePass() error {
	hash, err := bcrypt.GenerateFromPassword(user.PasswordHash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = hash
	return nil
}

// CreateUserFields for crete user.
type CreateUserFields struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	NickName  string `json:"nickName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Profile for user profile.
type Profile struct {
	Email     string    `json:"email"`
	NickName  string    `json:"nickName"`
	CreatedAt time.Time `json:"registerDate"`
}

// Password for old/new passwords.
type Password struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// IsPasswordValid check the password for all conditions.
func IsPasswordValid(s string) bool {
	var number, upper, special bool
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c) || unicode.IsMark(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		}
	}
	return len(s) >= 8 && letters >= 1 && number && upper && special
}
