// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/lootboxes"
)

// ensures that cardsDB implements cards.DB.
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
	query := `INSERT INTO lootboxes(lootbox_id, user_id, lootbox_name)
              VALUES($1,$2,$3)`

	_, err = lootboxesDB.conn.ExecContext(ctx, query, lootBox.UserID, lootBox.UserID, lootBox.Type)

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

	_, err := lootboxesDB.conn.ExecContext(ctx, query, lootboxID)

	return ErrLootBoxes.Wrap(err)
}
