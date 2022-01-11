// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeebo/errs"

	"ultimatedivision/store"
)

// ensures that storeDB implements store.DB.
var _ store.DB = (*storeDB)(nil)

// ErrStore indicates that there was an error in the database.
var ErrStore = errs.Class("store repository error")

// storeDB provide access to store DB.
//
// architecture: Database
type storeDB struct {
	conn *sql.DB
}

// Create creates setting of store in the database.
func (storeDB *storeDB) Create(ctx context.Context, setting store.Setting) error {
	query := `INSERT INTO store(id, cards_amount, is_renewal, date_renewal) VALUES($1,$2,$3,$4)`

	_, err := storeDB.conn.ExecContext(ctx, query, setting.ID, setting.CardsAmount, setting.IsRenewal, setting.DateRenewal)
	return ErrStore.Wrap(err)
}

// Get returns setting by id from database.
func (storeDB *storeDB) Get(ctx context.Context, id int) (store.Setting, error) {
	var setting store.Setting
	query := `SELECT * FROM store WHERE id = $1`

	row := storeDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&setting.ID, &setting.CardsAmount, &setting.IsRenewal, &setting.DateRenewal)
	if errors.Is(err, sql.ErrNoRows) {
		return setting, store.ErrNoSetting.Wrap(err)
	}

	return setting, ErrStore.Wrap(err)
}

// Update updates setting of store in the database.
func (storeDB *storeDB) Update(ctx context.Context, setting store.Setting) error {
	query := `UPDATE store
	          SET cards_amount = $1, is_renewal = $2, date_renewal = $3
	          WHERE id = $4`

	result, err := storeDB.conn.ExecContext(ctx, query, setting.CardsAmount, setting.IsRenewal, setting.DateRenewal, setting.ID)
	if err != nil {
		return ErrStore.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return store.ErrNoSetting.New("setting does not exist")
	}

	return ErrStore.Wrap(err)
}
