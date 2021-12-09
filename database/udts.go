// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/udts"
)

// ensures that udtsDB implements udts.DB.
var _ udts.DB = (*udtsDB)(nil)

// ErrUDTs indicates that there was an error in the database.
var ErrUDTs = errs.Class("ErrUDTs repository error")

// udtsDB provide access to udts DB.
//
// architecture: Database
type udtsDB struct {
	conn *sql.DB
}

// Create creates udt in the database.
func (udtsDB *udtsDB) Create(ctx context.Context, udt udts.UDT) error {
	query := `INSERT INTO udts(user_id, value, nonce) VALUES($1,$2,$3)`
	_, err := udtsDB.conn.ExecContext(ctx, query, udt.UserID, udt.Value.Bytes(), udt.Nonce)
	return ErrUDTs.Wrap(err)
}

// Get returns udt by user's id from database.
func (udtsDB *udtsDB) Get(ctx context.Context, userID uuid.UUID) (udts.UDT, error) {
	var (
		udt   udts.UDT
		value []byte
	)

	query := `SELECT * FROM udts WHERE user_id = $1`
	row := udtsDB.conn.QueryRowContext(ctx, query, userID)

	err := row.Scan(&udt.UserID, &value, &udt.Nonce)
	udt.Value.SetBytes(value)
	if errors.Is(err, sql.ErrNoRows) {
		return udt, udts.ErrNoUDT.Wrap(err)
	}

	return udt, ErrUDTs.Wrap(err)
}

// List returns udts from database.
func (udtsDB *udtsDB) List(ctx context.Context) ([]udts.UDT, error) {
	var (
		udtList []udts.UDT
		value   []byte
	)

	query := `SELECT * FROM udts`

	rows, err := udtsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return udtList, ErrUDTs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var udt udts.UDT

		if err = rows.Scan(&udt.UserID, &value, &udt.Nonce); err != nil {
			return udtList, ErrUDTs.Wrap(err)
		}
		udt.Value.SetBytes(value)
		udtList = append(udtList, udt)
	}

	return udtList, ErrUDTs.Wrap(rows.Err())
}

// Update updates nft by user's id in the database.
func (udtsDB *udtsDB) Update(ctx context.Context, udt udts.UDT) error {
	query := `UPDATE udts
	          SET value = $1, nonce = $2
	          WHERE user_id = $3`

	result, err := udtsDB.conn.ExecContext(ctx, query, udt.Value.Bytes(), udt.Nonce, udt.UserID)
	if err != nil {
		return ErrUDTs.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err != nil {
		return ErrUDTs.Wrap(err)
	}
	if rowNum == 0 {
		return udts.ErrNoUDT.New("udt does not exist")
	}

	return ErrUDTs.Wrap(err)
}

// Delete deletes udt by user's id in the database.
func (udtsDB *udtsDB) Delete(ctx context.Context, userID uuid.UUID) error {
	result, err := udtsDB.conn.ExecContext(ctx, "DELETE FROM udts WHERE user_id = $1", userID)
	if err != nil {
		return ErrUDTs.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return udts.ErrNoUDT.New("udt does not exist")
	}

	return ErrUDTs.Wrap(err)
}
