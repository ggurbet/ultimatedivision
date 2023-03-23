// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/gameengine"
)

// ensures that gameDB implements game.DB.
var _ gameengine.DB = (*gameengineDB)(nil)

// ErrGames indicates that there was an error in the database.
var ErrGames = errs.Class("games repository error")

// gameengineDB provide access to games DB.
//
// architecture: Database
type gameengineDB struct {
	conn *sql.DB
}

// Create creates game in db.
func (gameengineDB *gameengineDB) Create(ctx context.Context, matchID uuid.UUID, gameInformationInJSON string) error {
	tx, err := gameengineDB.conn.BeginTx(ctx, nil)
	if err != nil {
		return ErrGames.Wrap(err)
	}
	query := `INSERT INTO games(match_id, game_info)
              VALUES($1,$2)`

	_, err = gameengineDB.conn.ExecContext(ctx, query, matchID, gameInformationInJSON)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return ErrGames.Wrap(err)
		}
		return ErrGames.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return ErrGames.Wrap(err)
	}

	return ErrGames.Wrap(err)
}

// Get returns game information in JSON by match id.
func (gameengineDB *gameengineDB) Get(ctx context.Context, matchID uuid.UUID) (string, error) {
	query := `SELECT game_info
              FROM games
              WHERE match_id = $1`

	row := gameengineDB.conn.QueryRowContext(ctx, query, matchID)

	var gameDataInJSON string

	err := row.Scan(&gameDataInJSON)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return gameDataInJSON, gameengine.ErrNoGames.Wrap(err)
		}

		return gameDataInJSON, ErrGames.Wrap(err)
	}

	return gameDataInJSON, nil
}

// Update updates game info in the database by match id.
func (gameengineDB *gameengineDB) Update(ctx context.Context, matchID uuid.UUID, gameInformationInJSON string) error {
	result, err := gameengineDB.conn.ExecContext(ctx, "UPDATE games SET game_info=$1 WHERE match_id=$2", gameInformationInJSON, matchID)
	if err != nil {
		return ErrGames.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return gameengine.ErrNoGames.New("")
	}

	return ErrGames.Wrap(err)
}

// Delete deletes game information in JSON.
func (gameengineDB *gameengineDB) Delete(ctx context.Context, matchID uuid.UUID) error {
	query := `DELETE FROM games
              WHERE match_id = $1`

	result, err := gameengineDB.conn.ExecContext(ctx, query, matchID)
	if err != nil {
		return ErrGames.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return gameengine.ErrNoGames.New("invalid query")
	}

	return ErrGames.Wrap(err)
}
