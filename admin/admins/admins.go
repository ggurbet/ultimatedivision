// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoAdmin indicates that user does not exist.
var ErrNoAdmin = errs.Class("admin does not exist")

// DB is exposing access to admins db.
//
// architecture: DB
type DB interface {
	// List returns all admins from database.
	List(ctx context.Context) ([]Admin, error)
	// Get return admin by id from database.
	Get(ctx context.Context, id uuid.UUID) (Admin, error)
	// Create creates an admin and write it to database.
	Create(ctx context.Context, admin Admin) error
	// Update updates an admins password.
	Update(ctx context.Context, admin Admin) error
}

// Admin describes admin entity.
type Admin struct {
	ID           uuid.UUID
	Email        string
	PasswordHash []byte
	CreatedAt    time.Time
}
