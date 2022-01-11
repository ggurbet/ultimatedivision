// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package store

import (
	"context"
	"time"

	"github.com/zeebo/errs"
)

// ErrNoSetting indicated that setting does not exist.
var ErrNoSetting = errs.Class("setting does not exist")

// DB is exposing access to store db.
//
// architecture: DB
type DB interface {
	// Create creates setting of store in the database.
	Create(ctx context.Context, setting Setting) error
	// Get returns setting of store by id from database.
	Get(ctx context.Context, id int) (Setting, error)
	// List returns settings of store from database.
	List(ctx context.Context) ([]Setting, error)
	// Update updates setting of store in the database.
	Update(ctx context.Context, setting Setting) error
}

// Setting entity describes the values required to configure the store.
type Setting struct {
	ID          int       `json:"id"`
	CardsAmount int       `json:"cardsAmount"`
	IsRenewal   bool      `json:"isRenewal"`
	DateRenewal time.Time `json:"dateRenewal"`
}
