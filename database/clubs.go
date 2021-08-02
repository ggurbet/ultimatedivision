// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/clubs"
)

// ErrClubs indicates that there was an error in the database.
var ErrClubs = errs.Class("clubs repository error")

// ErrPlayers indicates that there was an error in the database.
var ErrPlayers = errs.Class("players repository error")

// clubsDB provide access to club DB.
//
// architecture: Database
type clubsDB struct {
	conn *sql.DB
}

// Create creates empty club in the db.
func (clubsDB *clubsDB) Create(ctx context.Context, club clubs.Club) error {
	_, err := clubsDB.conn.ExecContext(ctx,
		`INSERT INTO clubs(user_id,formation, tactic)
		VALUES($1,$2,$3)`, club.UserID, club.Formation, club.Tactic)
	return ErrClubs.Wrap(err)
}

// Add inserts card to club.
func (clubsDB clubsDB) Add(ctx context.Context, userID uuid.UUID, card cards.Card, capitan uuid.UUID, position clubs.Position) error {
	query := `INSERT  INTO club_player(user_id, card_id, card_position,capitan)
			  VALUES($1,$2,$3,$4)`

	_, err := clubsDB.conn.ExecContext(ctx, query, userID, card.ID, position, capitan)
	return ErrPlayers.Wrap(err)
}

// GetClub returns club from db.
func (clubsDB clubsDB) GetClub(ctx context.Context, userID uuid.UUID) (clubs.Club, error) {
	query := `SELECT user_id, formation, tactic
			  FROM clubs
			  WHERE user_id = $1`

	row := clubsDB.conn.QueryRowContext(ctx, query, userID)

	var club clubs.Club

	err := row.Scan(&club.UserID, &club.Formation, &club.Tactic)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return club, clubs.ErrNoClub.Wrap(err)
		}

		return club, ErrClubs.Wrap(err)
	}

	return club, nil
}

// ListCards returns all cards from the club.
func (clubsDB clubsDB) ListCards(ctx context.Context, userID uuid.UUID) ([]clubs.Player, error) {
	query := `SELECT user_id, card_id, card_position, capitan 
			  FROM club_player
			  WHERE user_id = $1`

	rows, err := clubsDB.conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, ErrPlayers.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrPlayers.Wrap(rows.Close()))
	}()

	var players []clubs.Player

	for rows.Next() {
		var player clubs.Player
		err = rows.Scan(&player.UserID, &player.CardID, &player.Position, &player.Capitan)
		if err != nil {
			return nil, clubs.ErrNoPlayer.Wrap(err)
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrPlayers.Wrap(err)
	}

	return players, nil
}

// Update updates club in the db.
func (clubsDB clubsDB) Update(ctx context.Context, club clubs.Club) error {
	query := `UPDATE clubs
              SET tactic = $1, formation = $2
			  WHERE user_id = $3;`

	_, err := clubsDB.conn.ExecContext(ctx, query, club.Tactic,
		club.Formation, club.UserID)
	return ErrClubs.Wrap(err)
}

// UpdateCapitan updates capitan in the users team.
func (clubsDB clubsDB) UpdateCapitan(ctx context.Context, capitan uuid.UUID, userID uuid.UUID) error {
	query := `UPDATE club_player
			  SET capitan = $1
			  WHERE user_id = $2`

	_, err := clubsDB.conn.ExecContext(ctx, query, capitan, userID)

	return ErrPlayers.Wrap(err)
}

// GetCapitan returns id of capitan of the users team.
func (clubsDB clubsDB) GetCapitan(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	query := `SELECT user_id, card_id, capitan, card_position
			  FROM club_player
              WHERE user_id = $1`

	var player clubs.Player

	row := clubsDB.conn.QueryRowContext(ctx, query, userID)

	err := row.Scan(&player.UserID, &player.CardID, &player.Capitan, &player.Position)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.New(), clubs.ErrNoPlayer.Wrap(err)
		}

		return uuid.New(), ErrPlayers.Wrap(err)
	}

	return player.Capitan, nil
}
