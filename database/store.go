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
	query := `INSERT INTO store VALUES($1,$2,$3,$4,$5)`

	_, err := storeDB.conn.ExecContext(ctx, query, setting.ID, setting.CardsAmount, setting.IsRenewal, setting.HourRenewal, setting.Price.Bytes())
	return ErrStore.Wrap(err)
}

// Get returns setting by id from database.
func (storeDB *storeDB) Get(ctx context.Context, id int) (store.Setting, error) {
	var (
		setting store.Setting
		price   []byte
	)
	query := `SELECT * FROM store WHERE id = $1`

	row := storeDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&setting.ID, &setting.CardsAmount, &setting.IsRenewal, &setting.HourRenewal, &price)
	if errors.Is(err, sql.ErrNoRows) {
		return setting, store.ErrNoSetting.Wrap(err)
	}
	setting.Price.SetBytes(price)

	return setting, ErrStore.Wrap(err)
}

// List returns settings of store from database.
func (storeDB *storeDB) List(ctx context.Context) ([]store.Setting, error) {
	var (
		settings []store.Setting
		price    []byte
	)
	query := `SELECT * FROM store`

	rows, err := storeDB.conn.QueryContext(ctx, query)
	if err != nil {
		return settings, ErrNFTs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var setting store.Setting

		if err = rows.Scan(&setting.ID, &setting.CardsAmount, &setting.IsRenewal, &setting.HourRenewal, &price); err != nil {
			return settings, ErrStore.Wrap(err)
		}
		setting.Price.SetBytes(price)

		settings = append(settings, setting)
	}

	return settings, ErrStore.Wrap(rows.Err())
}

// Update updates setting of store in the database.
func (storeDB *storeDB) Update(ctx context.Context, setting store.Setting) error {
	query := `UPDATE store
	          SET cards_amount = $1, is_renewal = $2, hour_renewal = $3, price = $4
	          WHERE id = $5`

	result, err := storeDB.conn.ExecContext(ctx, query, setting.CardsAmount, setting.IsRenewal, setting.HourRenewal, setting.Price.Bytes(), setting.ID)
	if err != nil {
		return ErrStore.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return store.ErrNoSetting.New("setting does not exist")
	}

	return ErrStore.Wrap(err)
}
