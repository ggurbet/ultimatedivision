// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision"
	"ultimatedivision/users"
)

// ensures that database implements ultimatedivision.DB.
var _ ultimatedivision.DB = (*database)(nil)

var (
	// Error is the default ultimatedivision error class.
	Error = errs.Class("ultimatedivision db error")
)

// database combines access to different database tables with a record
// of the db driver, db implementation, and db source URL.
//
// architecture: Master Database
type database struct {
	conn *sql.DB
}

// New returns ultimatedivision.DB postgresql implementation.
func New(databaseURL string) (ultimatedivision.DB, error) {
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	return &database{conn: conn}, nil
}

// CreateSchema create schema for all tables and databases.
func (db *database) CreateSchema(ctx context.Context) (err error) {
	createTableQuery :=
		`CREATE TABLE IF NOT EXISTS users (
            id         BYTEA     PRIMARY KEY 	NOT NULL,
            email      VARCHAR                  NOT NULL,
            password   BYTEA                    NOT NULL,
            nick_name  VARCHAR                  NOT NULL,
            first_name VARCHAR                  NOT NULL,
            last_name  VARCHAR                  NOT NULL,
            last_login TIMESTAMP WITH TIME ZONE NOT NULL,
            status     INTEGER                  NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL
		);
		`

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

// usersDB provided access to accounts db.
func (db *database) Users() users.DB {
	return &usersDB{conn: db.conn}
}
