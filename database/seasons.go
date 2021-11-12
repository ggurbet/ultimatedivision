// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/seasons"
)

// ensures that seasonsDB implements seasons.DB.
var _ seasons.DB = (*seasonsDB)(nil)

// ErrSeasons indicates that there was an error in the database.
var ErrSeasons = errs.Class("seasons repository error")

// seasonsDB provides access to seasons db.
//
// architecture: Database
type seasonsDB struct {
	conn *sql.DB
}

// Create creates a seasons and writes to the database.
func (seasonsDB *seasonsDB) Create(ctx context.Context, season seasons.Season) error {
	query := `INSERT INTO seasons(division_id, started_at, ended_at) 
	VALUES ($1, $2, $3)`

	_, err := seasonsDB.conn.ExecContext(ctx, query, season.DivisionID, season.StartedAt, season.EndedAt)

	return ErrSeasons.Wrap(err)
}

// EndSeason updates a status in the database when season ended.
func (seasonsDB *seasonsDB) EndSeason(ctx context.Context, id int) error {
	db, err := seasonsDB.conn.ExecContext(ctx, "UPDATE seasons SET ended_at=$1 WHERE id=$2", time.Now().UTC(), id)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	rowNum, err := db.RowsAffected()
	if rowNum == 0 {
		return seasons.ErrNoSeason.New("season does not exist")
	}

	return ErrSeasons.Wrap(err)
}

// List returns all seasons from the data base.
func (seasonsDB *seasonsDB) List(ctx context.Context) ([]seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons`

	rows, err := seasonsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allSeasons []seasons.Season
	for rows.Next() {
		var season seasons.Season
		err := rows.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
		if err != nil {
			return nil, ErrSeasons.Wrap(err)
		}

		allSeasons = append(allSeasons, season)
	}

	return allSeasons, ErrSeasons.Wrap(rows.Err())
}

// Get returns season by id from the data base.
func (seasonsDB *seasonsDB) Get(ctx context.Context, id int) (seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons WHERE id=$1`
	var season seasons.Season

	row := seasonsDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return season, seasons.ErrNoSeason.Wrap(err)
		}

		return season, ErrSeasons.Wrap(err)
	}

	return season, ErrSeasons.Wrap(err)
}

// GetCurrentSeasons returns all current seasons from the data base.
func (seasonsDB *seasonsDB) GetCurrentSeasons(ctx context.Context) ([]seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons WHERE ended_at=$1`

	rows, err := seasonsDB.conn.QueryContext(ctx, query, time.Time{})
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allSeasons []seasons.Season
	for rows.Next() {
		var season seasons.Season
		err := rows.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
		if err != nil {
			return nil, ErrSeasons.Wrap(err)
		}

		allSeasons = append(allSeasons, season)
	}

	return allSeasons, ErrSeasons.Wrap(rows.Err())
}

// Delete deletes a season in the database.
func (seasonsDB *seasonsDB) Delete(ctx context.Context, id int) error {
	result, err := seasonsDB.conn.ExecContext(ctx, "DELETE FROM seasons WHERE id=$1", id)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return seasons.ErrNoSeason.New("season does not exist")
	}

	return ErrSeasons.Wrap(err)
}

// GetSeasonByDivisionID returns season by division id from the data base.
func (seasonsDB *seasonsDB) GetSeasonByDivisionID(ctx context.Context, divisionID uuid.UUID) (seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons WHERE division_id=$1`
	var season seasons.Season

	row := seasonsDB.conn.QueryRowContext(ctx, query, divisionID)

	err := row.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return season, seasons.ErrNoSeason.Wrap(err)
		}

		return season, ErrSeasons.Wrap(err)
	}

	return season, ErrSeasons.Wrap(err)
}
