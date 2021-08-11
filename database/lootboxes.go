// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/zeebo/errs"

	"ultimatedivision/lootboxes"
)

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
	query := `INSERT INTO lootboxes(lootbox_id, user_id, lootbox_name)
              VALUES($1,$2,$3)`

	_, err = lootboxesDB.conn.ExecContext(ctx, query, lootBox.UserID, lootBox.UserID, lootBox.Name)

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
func (lootboxesDB *lootboxesDB) Delete(ctx context.Context, lootBox lootboxes.LootBox) error {
	query := `DELETE FROM lootboxes
              WHERE user_id = $1 and lootbox_id = $2`

	_, err := lootboxesDB.conn.ExecContext(ctx, query, lootBox.UserID, lootBox.LootBoxID)

	return ErrLootBoxes.Wrap(err)
}
