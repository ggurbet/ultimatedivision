// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
)

// ErrClubs indicates that there was an error in the database.
var ErrClubs = errs.Class("clubs repository error")

// ErrSquad indicates that there was an error in the database.
var ErrSquad = errs.Class("squad repository error")

// clubsDB provide access to club DB.
//
// architecture: Database
type clubsDB struct {
	conn *sql.DB
}

// Create creates empty club in the db.
func (clubsDB *clubsDB) Create(ctx context.Context, club clubs.Club) error {
	query := `INSERT INTO clubs(id, owner_id, club_name, created_at)
			   VALUES($1,$2,$3,$4)`

	_, err := clubsDB.conn.ExecContext(ctx, query,
		club.ID, club.OwnerID, club.Name, club.CreatedAt)

	return ErrClubs.Wrap(err)
}

func (clubsDB *clubsDB) CreateSquad(ctx context.Context, squad clubs.Squads) error {
	query := `INSERT INTO squads(id, squad_name, club_id, tactic, formation)
			   VALUES($1,$2,$3,$4,$5)`

	_, err := clubsDB.conn.ExecContext(ctx, query,
		squad.ID, squad.Name, squad.ClubID, squad.Tactic, squad.Formation)

	return ErrClubs.Wrap(err)
}

// Add inserts card to club.
func (clubsDB *clubsDB) Add(ctx context.Context, squadCards clubs.SquadCards) error {
	query := `INSERT INTO squad_cards(id, card_id, card_position, capitan)
			  VALUES($1,$2,$3,$4)`

	_, err := clubsDB.conn.ExecContext(ctx, query,
		squadCards.ID, squadCards.CardID, squadCards.Position, squadCards.Capitan)

	return ErrSquad.Wrap(err)
}

// List returns all the clubs owned by the user.
func (clubsDB *clubsDB) List(ctx context.Context, userID uuid.UUID) ([]clubs.Club, error) {
	query := `SELECT id, owner_id, club_name, created_at
			  FROM clubs
			  WHERE owner_id = $1`

	rows, err := clubsDB.conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrClubs.Wrap(rows.Close()))
	}()

	var userClubs []clubs.Club

	for rows.Next() {
		var club clubs.Club

		err = rows.Scan(&club.ID, &club.OwnerID, &club.Name, &club.CreatedAt)
		if err != nil {
			return nil, clubs.ErrNoClub.Wrap(err)
		}

		userClubs = append(userClubs, club)
	}

	return userClubs, nil
}

// GetSquad returns squad from database.
func (clubsDB *clubsDB) GetSquad(ctx context.Context, clubID uuid.UUID) (clubs.Squads, error) {
	query := `SELECT id, squad_name, club_id, tactic, formation 
			  FROM squads
			  WHERE club_id = $1`

	row := clubsDB.conn.QueryRowContext(ctx, query, clubID)

	var squad clubs.Squads

	err := row.Scan(&squad.ID, &squad.Name, &squad.ClubID, &squad.Tactic, &squad.Formation)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return squad, clubs.ErrNoSquad.Wrap(err)
		}

		return squad, ErrClubs.Wrap(err)
	}

	return squad, nil
}

// ListSquadCards returns all cards from squad.
func (clubsDB *clubsDB) ListSquadCards(ctx context.Context, squadID uuid.UUID) ([]clubs.SquadCards, error) {
	query := `SELECT id, card_id, card_position, capitan 
			  FROM squad_cards
			  WHERE id = $1`

	rows, err := clubsDB.conn.QueryContext(ctx, query, squadID)
	if err != nil {
		return nil, ErrSquad.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrSquad.Wrap(rows.Close()))
	}()

	var players []clubs.SquadCards

	for rows.Next() {
		var player clubs.SquadCards
		err = rows.Scan(&player.ID, &player.CardID, &player.Position, &player.Capitan)
		if err != nil {
			return nil, clubs.ErrNoSquad.Wrap(err)
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrSquad.Wrap(err)
	}

	return players, nil
}

// UpdateTacticFormation updates tactic and formation for squad.
func (clubsDB *clubsDB) UpdateTacticFormation(ctx context.Context, squad clubs.Squads) error {
	query := `UPDATE squads
			  SET tactic = $1, formation = $2
  			  WHERE id = $3`

	_, err := clubsDB.conn.ExecContext(ctx, query, squad.Tactic, squad.Formation, squad.ID)

	return ErrSquad.Wrap(err)
}

// UpdateCapitan updates capitan in the users team.
func (clubsDB *clubsDB) UpdateCapitan(ctx context.Context, capitan uuid.UUID, squadID uuid.UUID) error {
	query := `UPDATE squad_cards
			  SET capitan = $1
			  WHERE id = $2`

	_, err := clubsDB.conn.ExecContext(ctx, query, capitan, squadID)

	return ErrSquad.Wrap(err)
}

// UpdatePosition updates position of card in the squad.
func (clubsDB *clubsDB) UpdatePosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, newPosition clubs.Position) error {
	query := `UPDATE squad_cards
			  SET card_position = $1
			  WHERE card_id = $2 AND id = $3`

	_, err := clubsDB.conn.ExecContext(ctx, query, newPosition, cardID, squadID)

	return ErrSquad.Wrap(err)
}

// GetCapitan returns id of capitan of the users team.
func (clubsDB *clubsDB) GetCapitan(ctx context.Context, squadID uuid.UUID) (uuid.UUID, error) {
	query := `SELECT capitan
			  FROM squad_cards
              WHERE id = $1`

	var id uuid.UUID

	row := clubsDB.conn.QueryRowContext(ctx, query, squadID)

	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.New(), clubs.ErrNoSquad.Wrap(err)
		}

		return uuid.New(), ErrSquad.Wrap(err)
	}

	return id, nil
}
