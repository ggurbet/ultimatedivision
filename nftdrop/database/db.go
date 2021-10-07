// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision/nftdrop"
	"ultimatedivision/nftdrop/whitelist"
)

// ensures that database implements nftdrop.DB.
var _ nftdrop.DB = (*database)(nil)

var (
	// Error is the default nftdrop error class.
	Error = errs.Class("nftdrop db error")
)

// database combines access to different database tables with a record
// of the db driver, db implementation, and db source URL.
//
// architecture: Master Database
type database struct {
	conn *sql.DB
}

// New returns nftdrop.DB postgresql implementation.
func New(databaseURL string) (nftdrop.DB, error) {
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	return &database{conn: conn}, nil
}

// CreateSchema create schema for all tables and databases.
func (db *database) CreateSchema(ctx context.Context) (err error) {
	createTableQuery :=
		`CREATE TABLE IF NOT EXISTS whitelist (
			address  VARCHAR PRIMARY KEY NOT NULL,
			password BYTEA               NOT NULL
		);`

	_, err = db.conn.ExecContext(ctx, createTableQuery)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// Close closes underlying db connection.
func (db *database) Close() error {
	return Error.Wrap(db.conn.Close())
}

// Whitelist provided access to accounts db.
func (db *database) Whitelist() whitelist.DB {
	return &whitelistDB{conn: db.conn}
}
