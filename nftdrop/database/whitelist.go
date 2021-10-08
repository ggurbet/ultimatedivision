// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision/nftdrop/whitelist"
)

// ensures that whitelistDB implements whitelist.DB.
var _ whitelist.DB = (*whitelistDB)(nil)

// ErrWhitelist indicates that there was an error in the database.
var ErrWhitelist = errs.Class("whitelist repository error")

// whitelistDB provides access to whitelist db.
//
// architecture: Database
type whitelistDB struct {
	conn *sql.DB
}

// Create add record whitelist in the data base.
func (whitelistDB *whitelistDB) Create(ctx context.Context, whitelist whitelist.Whitelist) error {
	query :=
		`INSERT INTO
			whitelist(address, password) 
		VALUES 
			($1, $2)`

	_, err := whitelistDB.conn.ExecContext(ctx, query, whitelist.Address, whitelist.Password)
	return ErrWhitelist.Wrap(err)
}

// GetByAddress returns record whitelist by address from the data base.
func (whitelistDB *whitelistDB) GetByAddress(ctx context.Context, address whitelist.Address) (whitelist.Whitelist, error) {
	whitelistRecord := whitelist.Whitelist{}
	query :=
		`SELECT
			address, password
		FROM 
			whitelist
		WHERE
			address = $1`

	err := whitelistDB.conn.QueryRowContext(ctx, query, address).Scan(&whitelistRecord.Address, &whitelistRecord.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return whitelistRecord, whitelist.ErrNoWhitelist.Wrap(err)
	}

	return whitelistRecord, ErrWhitelist.Wrap(err)
}

// List returns all whitelist from the data base.
func (whitelistDB *whitelistDB) List(ctx context.Context) ([]whitelist.Whitelist, error) {
	query :=
		`SELECT
			address, password
		FROM 
			whitelist`

	rows, err := whitelistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrWhitelist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrWhitelist.Wrap(rows.Close()))
	}()

	whitelistRecords := []whitelist.Whitelist{}
	for rows.Next() {
		whitelistRecord := whitelist.Whitelist{}
		if err = rows.Scan(&whitelistRecord.Address, &whitelistRecord.Password); err != nil {
			return nil, ErrWhitelist.Wrap(err)
		}
		whitelistRecords = append(whitelistRecords, whitelistRecord)
	}

	return whitelistRecords, ErrWhitelist.Wrap(rows.Err())
}

// Delete deletes whitelist from the database.
func (whitelistDB *whitelistDB) Delete(ctx context.Context, address whitelist.Address) error {
	query := `DELETE FROM whitelist
              WHERE address = $1`

	_, err := whitelistDB.conn.ExecContext(ctx, query, address)

	return ErrWhitelist.Wrap(err)
}
