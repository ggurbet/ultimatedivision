// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"database/sql"

	"github.com/zeebo/errs"

	"ultimatedivision/udts/currencywaitlist"
)

// ensures that currencywaitlistDB implements currencywaitlist.DB.
var _ currencywaitlist.DB = (*currencywaitlistDB)(nil)

// ErrCurrencyWaitlist indicates that there was an error in the database.
var ErrCurrencyWaitlist = errs.Class("ErrCurrencyWaitlist repository error")

// currencywaitlistDB provide access to nfts DB.
//
// architecture: Database
type currencywaitlistDB struct {
	conn *sql.DB
}
