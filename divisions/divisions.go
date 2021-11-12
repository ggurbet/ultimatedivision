// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package divisions

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoDivision indicated that division does not exist.
var ErrNoDivision = errs.Class("division does not exist")

// DB exposes access to divisions db.
//
// architecture: DB
type DB interface {
	// Create creates a division and writes to the database.
	Create(ctx context.Context, division Division) error
	// List returns all divisions from the data base.
	List(ctx context.Context) ([]Division, error)
	// Get returns division by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Division, error)
	// GetLastDivision returns last division from the data base.
	GetLastDivision(ctx context.Context) (Division, error)
	// Delete deletes a division in the database.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Division describes divisions entity.
type Division struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	PassingPercent int       `json:"passingPercent"`
	CreatedAt      time.Time `json:"createdAt"`
}

// Config defines configuration for divisions.
type Config struct {
	PassingPercent int `json:"passingPercent"`
}
