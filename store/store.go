// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package store

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ErrNoSetting indicated that setting does not exist.
var ErrNoSetting = errs.Class("setting does not exist")

// DB is exposing access to store db.
//
// architecture: DB.
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
	ID          int  `json:"id"`
	CardsAmount int  `json:"cardsAmount"`
	IsRenewal   bool `json:"isRenewal"`
	HourRenewal int  `json:"dateRenewal"`
}

// Config defines values needed by create cards.
type Config struct {
	StoreRenewalInterval time.Duration             `json:"storeRenewalInterval"`
	PercentageQualities  cards.PercentageQualities `json:"percentageQualities"`
}

// ActiveSetting indicates that the number is active setting.
const ActiveSetting int = 1

// HourOfDay defines hours of day.
type HourOfDay int

const (
	// HourOfDayMin indicates that the minimum hour of the day is 0.
	HourOfDayMin int = 0
	// HourOfDayMax indicates that the maximum hour of the day is 23.
	HourOfDayMax int = 23
)
