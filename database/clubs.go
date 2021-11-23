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

// ensures that clubsDB implements clubs.DB.
var _ clubs.DB = (*clubsDB)(nil)

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
func (clubsDB *clubsDB) Create(ctx context.Context, club clubs.Club) (uuid.UUID, error) {
	tx, err := clubsDB.conn.BeginTx(ctx, nil)
	if err != nil {
		return uuid.Nil, ErrClubs.Wrap(err)
	}

	query := `INSERT INTO clubs(id, owner_id, club_name, status, division_id, created_at)
              VALUES($1,$2,$3,$4,$5,$6)
              RETURNING id`

	var clubID uuid.UUID
	err = clubsDB.conn.QueryRowContext(ctx, query,
		club.ID, club.OwnerID, club.Name, club.Status, club.DivisionID, club.CreatedAt).Scan(&clubID)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return uuid.Nil, ErrClubs.Wrap(err)
		}

		return uuid.Nil, ErrClubs.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return uuid.Nil, ErrClubs.Wrap(err)
	}

	return clubID, ErrClubs.Wrap(err)
}

// List returns all clubs.
func (clubsDB *clubsDB) List(ctx context.Context) ([]clubs.Club, error) {
	query := `SELECT id, owner_id, club_name, status, division_id, created_at
			  FROM clubs`

	rows, err := clubsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allClubs []clubs.Club

	for rows.Next() {
		var club clubs.Club
		err = rows.Scan(&club.ID, &club.OwnerID, &club.Name, &club.Status, &club.DivisionID, &club.CreatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return allClubs, clubs.ErrNoClub.Wrap(err)
			}
			return allClubs, clubs.ErrClubs.Wrap(err)
		}

		allClubs = append(allClubs, club)
	}

	return allClubs, nil
}

// CreateSquad creates squad for clubs in the database.
func (clubsDB *clubsDB) CreateSquad(ctx context.Context, squad clubs.Squad) (uuid.UUID, error) {
	query := `INSERT INTO squads(id, squad_name, club_id, tactic, formation, captain_id)
              VALUES($1,$2,$3,$4,$5,$6)
              RETURNING id`

	var squadID uuid.UUID

	err := clubsDB.conn.QueryRowContext(ctx, query,
		squad.ID, squad.Name, squad.ClubID, squad.Tactic, squad.Formation, squad.CaptainID).Scan(&squadID)

	return squadID, ErrClubs.Wrap(err)
}

// AddSquadCard inserts card to club.
func (clubsDB *clubsDB) AddSquadCard(ctx context.Context, squadCards clubs.SquadCard) error {
	query := `INSERT INTO squad_cards(id, card_id, card_position)
              VALUES($1,$2,$3)`

	_, err := clubsDB.conn.ExecContext(ctx, query,
		squadCards.SquadID, squadCards.CardID, squadCards.Position)

	return ErrClubs.Wrap(err)
}

// DeleteSquadCard deletes card from squad.
func (clubsDB *clubsDB) DeleteSquadCard(ctx context.Context, squadID, cardID uuid.UUID) error {
	query := `DELETE FROM squad_cards
              WHERE card_id = $1 and id = $2`

	result, err := clubsDB.conn.ExecContext(ctx, query, cardID, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return clubs.ErrNoSquadCard.New("squad card does not exist")
	}

	return ErrClubs.Wrap(err)
}

// ListByUserID returns clubs owned by the user.
func (clubsDB *clubsDB) ListByUserID(ctx context.Context, userID uuid.UUID) ([]clubs.Club, error) {
	query := `SELECT id, owner_id, club_name, status, created_at
			  FROM clubs
			  WHERE owner_id = $1`

	rows, err := clubsDB.conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allClubs []clubs.Club

	for rows.Next() {
		var club clubs.Club
		err = rows.Scan(&club.ID, &club.OwnerID, &club.Name, &club.Status, &club.CreatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return allClubs, clubs.ErrNoClub.New("club does not exist")
			}
			return allClubs, clubs.ErrClubs.Wrap(err)
		}

		allClubs = append(allClubs, club)
	}

	return allClubs, nil
}

// Get returns club.
func (clubsDB *clubsDB) Get(ctx context.Context, clubID uuid.UUID) (clubs.Club, error) {
	query := `SELECT id, owner_id, club_name, status, division_id, created_at
			  FROM clubs
			  WHERE id = $1`

	row := clubsDB.conn.QueryRowContext(ctx, query, clubID)

	var club clubs.Club
	err := row.Scan(&club.ID, &club.OwnerID, &club.Name, &club.Status, &club.DivisionID, &club.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return club, clubs.ErrNoClub.Wrap(err)
		}

		return club, clubs.ErrClubs.Wrap(err)
	}

	return club, nil
}

// GetSquadByClubID returns squad from database.
func (clubsDB *clubsDB) GetSquadByClubID(ctx context.Context, clubID uuid.UUID) (clubs.Squad, error) {
	query := `SELECT id, squad_name, club_id, tactic, formation, captain_id
			  FROM squads
			  WHERE club_id = $1`

	row := clubsDB.conn.QueryRowContext(ctx, query, clubID)

	var squad clubs.Squad

	err := row.Scan(&squad.ID, &squad.Name, &squad.ClubID, &squad.Tactic, &squad.Formation, &squad.CaptainID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return squad, clubs.ErrNoSquad.Wrap(err)
		}

		return squad, ErrClubs.Wrap(err)
	}

	return squad, nil
}

// GetSquad returns squad from database.
func (clubsDB *clubsDB) GetSquad(ctx context.Context, squadID uuid.UUID) (clubs.Squad, error) {
	query := `SELECT id, squad_name, club_id, tactic, formation, captain_id
			  FROM squads
			  WHERE id = $1`

	row := clubsDB.conn.QueryRowContext(ctx, query, squadID)

	var squad clubs.Squad

	err := row.Scan(&squad.ID, &squad.Name, &squad.ClubID, &squad.Tactic, &squad.Formation, &squad.CaptainID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return squad, clubs.ErrNoSquad.Wrap(err)
		}

		return squad, ErrClubs.Wrap(err)
	}

	return squad, nil
}

// ListSquadCards returns all cards from squad.
func (clubsDB *clubsDB) ListSquadCards(ctx context.Context, squadID uuid.UUID) ([]clubs.SquadCard, error) {
	query := `SELECT id, card_id, card_position 
              FROM squad_cards
              WHERE id = $1
          	  ORDER BY card_position`

	rows, err := clubsDB.conn.QueryContext(ctx, query, squadID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var players []clubs.SquadCard

	for rows.Next() {
		var player clubs.SquadCard
		err = rows.Scan(&player.SquadID, &player.CardID, &player.Position)
		if err != nil {
			return nil, clubs.ErrNoSquadCard.Wrap(err)
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	return players, ErrClubs.Wrap(err)
}

// UpdateTacticCaptain updates tactic and capitan in the squad.
func (clubsDB *clubsDB) UpdateTacticCaptain(ctx context.Context, squad clubs.Squad) error {
	query := `UPDATE squads
			  SET tactic = $1, captain_id = $2
  			  WHERE id = $3`

	result, err := clubsDB.conn.ExecContext(ctx, query, squad.Tactic, squad.CaptainID, squad.ID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return clubs.ErrNoSquad.New("squad does not exist")
	}

	return ErrClubs.Wrap(err)
}

// UpdateFormation updates formation in the squad.
func (clubsDB *clubsDB) UpdateFormation(ctx context.Context, newFormation clubs.Formation, squadID uuid.UUID) error {
	query := `UPDATE squads
			  SET formation = $1
  			  WHERE id = $2`

	result, err := clubsDB.conn.ExecContext(ctx, query, newFormation, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return clubs.ErrNoSquad.New("squad does not exist")
	}

	return ErrSquad.Wrap(err)
}

// GetFormation returns formation of the squad.
func (clubsDB *clubsDB) GetFormation(ctx context.Context, squadID uuid.UUID) (clubs.Formation, error) {
	var formation clubs.Formation
	query := `SELECT formation 
              FROM squads
              WHERE id = $1`

	err := clubsDB.conn.QueryRowContext(ctx, query, squadID).Scan(&formation)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return formation, clubs.ErrNoSquad.Wrap(err)
		}

		return formation, ErrClubs.Wrap(err)
	}

	return formation, ErrClubs.Wrap(err)
}

// UpdateStatuses updates statuses of user clubs.
func (clubsDB *clubsDB) UpdateStatuses(ctx context.Context, allClubs []clubs.Club) error {
	query := `UPDATE clubs
			  SET status = $1
			  WHERE id = $2`

	preparedQuery, err := clubsDB.conn.PrepareContext(ctx, query)
	if err != nil {
		return ErrClubs.Wrap(err)
	}
	defer func() {
		err = preparedQuery.Close()
	}()

	for _, club := range allClubs {
		result, err := preparedQuery.ExecContext(ctx, club.Status, club.ID)
		if err != nil {
			return ErrClubs.Wrap(err)
		}

		rowNum, err := result.RowsAffected()
		if rowNum == 0 {
			return clubs.ErrNoClub.New("club does not exist")
		}

		if err != nil {
			return ErrClubs.Wrap(err)
		}
	}

	return ErrClubs.Wrap(err)
}

// UpdatePositions updates positions of cards in the squad.
func (clubsDB *clubsDB) UpdatePositions(ctx context.Context, squadCards []clubs.SquadCard) error {
	query := `UPDATE squad_cards
			  SET card_position = $1
			  WHERE card_id = $2 and id = $3`

	preparedQuery, err := clubsDB.conn.PrepareContext(ctx, query)
	if err != nil {
		return ErrClubs.Wrap(err)
	}
	defer func() {
		err = preparedQuery.Close()
	}()

	for _, squadCard := range squadCards {
		result, err := preparedQuery.ExecContext(ctx, squadCard.Position, squadCard.CardID, squadCard.SquadID)
		if err != nil {
			return ErrClubs.Wrap(err)
		}

		rowNum, err := result.RowsAffected()
		if rowNum == 0 {
			return clubs.ErrNoSquadCard.New("squad card does not exist")
		}

		if err != nil {
			return ErrClubs.Wrap(err)
		}
	}

	return ErrClubs.Wrap(err)
}

// GetCaptainID returns id of captain of the users team.
func (clubsDB *clubsDB) GetCaptainID(ctx context.Context, squadID uuid.UUID) (uuid.UUID, error) {
	query := `SELECT captain_id
			  FROM squads
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

// UpdateClubToNewDivision updates club to new division.
func (clubsDB *clubsDB) UpdateClubToNewDivision(ctx context.Context, clubID uuid.UUID, newDivisionID uuid.UUID) error {
	query := `UPDATE clubs
			  SET division_id = $1
			  WHERE id = $2`

	result, err := clubsDB.conn.ExecContext(ctx, query, newDivisionID, clubID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return clubs.ErrNoClub.New("club does not exist")
	}

	return ErrSquad.Wrap(err)
}
