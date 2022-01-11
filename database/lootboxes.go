// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/store/lootboxes"
)

// ensures that lootboxDB implements lootbox.DB.
var _ lootboxes.DB = (*lootboxesDB)(nil)

// ErrLootBoxes indicates that there was an error in the database.
var ErrLootBoxes = errs.Class("lootboxes repository error")

// lootboxesDB provide access to lootboxes DB.
//
// architecture: Database
type lootboxesDB struct {
	conn *sql.DB
}

// Create creates opened lootbox in db.
func (lootboxesDB *lootboxesDB) Create(ctx context.Context, lootBox lootboxes.LootBox) error {
	tx, err := lootboxesDB.conn.BeginTx(ctx, nil)
	if err != nil {
		return ErrLootBoxes.Wrap(err)
	}
	query := `INSERT INTO lootboxes(user_id, lootbox_id, lootbox_type)
              VALUES($1,$2,$3)`

	_, err = lootboxesDB.conn.ExecContext(ctx, query, lootBox.UserID, lootBox.LootBoxID, lootBox.Type)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return ErrLootBoxes.Wrap(err)
		}
		return ErrLootBoxes.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return ErrLootBoxes.Wrap(err)
	}

	return ErrLootBoxes.Wrap(err)
}

// Delete deletes opened lootbox by user in db.
func (lootboxesDB *lootboxesDB) Delete(ctx context.Context, lootboxID uuid.UUID) error {
	query := `DELETE FROM lootboxes
              WHERE lootbox_id = $1`

	result, err := lootboxesDB.conn.ExecContext(ctx, query, lootboxID)
	if err != nil {
		return ErrLootBoxes.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return lootboxes.ErrNoLootBox.New("invalid query")
	}

	return ErrLootBoxes.Wrap(err)
}

// List returns all loot boxes.
func (lootboxesDB *lootboxesDB) List(ctx context.Context) ([]lootboxes.LootBox, error) {
	query := `SELECT user_id, lootbox_id, lootbox_type
              FROM lootboxes`

	rows, err := lootboxesDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrLootBoxes.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var userLootBoxes []lootboxes.LootBox

	for rows.Next() {
		var userLootBox lootboxes.LootBox

		err = rows.Scan(&userLootBox.UserID, &userLootBox.LootBoxID, &userLootBox.Type)
		if err != nil {
			return nil, ErrLootBoxes.Wrap(err)
		}

		userLootBoxes = append(userLootBoxes, userLootBox)
	}

	return userLootBoxes, nil
}

// Get returns lootbox by user id.
func (lootboxesDB *lootboxesDB) Get(ctx context.Context, lootboxID uuid.UUID) (lootboxes.LootBox, error) {
	query := `SELECT user_id, lootbox_id, lootbox_type
              FROM lootboxes
              WHERE lootbox_id = $1`

	row := lootboxesDB.conn.QueryRowContext(ctx, query, lootboxID)

	var userLootbox lootboxes.LootBox

	err := row.Scan(&userLootbox.UserID, &userLootbox.LootBoxID, &userLootbox.Type)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return userLootbox, lootboxes.ErrNoLootBox.Wrap(err)
		}

		return userLootbox, ErrLootBoxes.Wrap(err)
	}

	return userLootbox, nil
}
