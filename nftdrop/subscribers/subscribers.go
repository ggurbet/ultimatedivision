// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package subscribers

import (
	"context"
	"time"

	"github.com/zeebo/errs"
)

// ErrNoSubscriber indicated that subscriber does not exist.
var ErrNoSubscriber = errs.Class("subscriber does not exist")

// DB exposes access to subscribers db.
//
// architecture: DB
type DB interface {
	// Create creates a subscriber and writes to the database.
	Create(ctx context.Context, email Subscriber) error
	// List returns all subscriber from the data base.
	List(ctx context.Context) ([]Subscriber, error)
	// Delete deletes a subscriber in the database.
	Delete(ctx context.Context, email string) error
	// GetByEmail returns subscriber by email from the data base.
	GetByEmail(ctx context.Context, email string) (Subscriber, error)
}

// Subscriber describes subscriber entity.
type Subscriber struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateSubscriberFields for create subscriber.
type CreateSubscriberFields struct {
	Email string `json:"email"`
}
