// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package udts

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoUDT indicates that udt of currency wait list does not exist.
var ErrNoUDT = errs.Class("udt does not exist")

// DB is exposing access to udts db.
//
// architecture: DB.
type DB interface {
	// Create creates udt in the database.
	Create(ctx context.Context, udt UDT) error
	// GetByUserID returns udt by user's id from database.
	GetByUserID(ctx context.Context, userID uuid.UUID) (UDT, error)
	// List returns udts from database.
	List(ctx context.Context) ([]UDT, error)
	// Update updates udt by user's id in the database.
	Update(ctx context.Context, udt UDT) error
	// Delete deletes udt by user's id in the database.
	Delete(ctx context.Context, userID uuid.UUID) error
}

// UDT entity describes how many tokens of udt and what nonce the user has.
type UDT struct {
	UserID uuid.UUID `json:"userId"`
	Nonce  int64     `json:"nonce"`
}
