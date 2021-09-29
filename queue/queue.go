// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/pagination"
)

// ErrNoPlace indicated that place does not exist.
var ErrNoPlace = errs.Class("place does not exist")

// DB is exposing access to queues database.
//
// architecture: DB
type DB interface {
	// Create adds place in database.
	Create(ctx context.Context, place Place) error
	// Get returns place from database.
	Get(ctx context.Context, id uuid.UUID) (Place, error)
	// ListPaginated returns page of places from database.
	ListPaginated(ctx context.Context, cursor pagination.Cursor) (Page, error)
	// UpdateStatus updates status place in database.
	UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error
	// Delete deletes place record in database.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Place entity describes place of the queue.
type Place struct {
	UserID uuid.UUID `json:"userId"`
	Status Status    `json:"status"`
}

// Status defines list of possible place statuses.
type Status string

const (
	// StatusSearches indicates that user in place searches game.
	StatusSearches Status = "searches"
	// StatusPlays indicates that in place plays game.
	StatusPlays Status = "plays"
)

// Config defines configuration for places.
type Config struct {
	PlaceRenewalInterval time.Duration     `json:"placeRenewalInterval"`
	Cursor               pagination.Cursor `json:"cursor"`
}

// Page holds place page entity which is used to show listed page of places.
type Page struct {
	Places []Place         `json:"place"`
	Page   pagination.Page `json:"page"`
}
