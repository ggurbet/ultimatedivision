// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/divisions"
)

// ensures that divisionsDB implements divisions.DB.
var _ divisions.DB = (*divisionsDB)(nil)

// ErrDivisions indicates that there was an error in the database.
var ErrDivisions = errs.Class("divisions repository error")

// divisionsDB provides access to divisions db.
//
// architecture: Database
type divisionsDB struct {
	conn *sql.DB
}

// Create creates a division and writes to the database.
func (divisionsDB *divisionsDB) Create(ctx context.Context, division divisions.Division) error {
	query := `INSERT INTO divisions(id, name, passing_percent, created_at) 
	VALUES ($1, $2, $3, $4)`

	_, err := divisionsDB.conn.ExecContext(ctx, query, division.ID, division.Name, division.PassingPercent, division.CreatedAt)

	return ErrDivisions.Wrap(err)
}

// List returns all divisions from the data base.
func (divisionsDB *divisionsDB) List(ctx context.Context) ([]divisions.Division, error) {
	query := `SELECT id, name, passing_percent, created_at FROM divisions`

	rows, err := divisionsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrDivisions.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allDivisions []divisions.Division
	for rows.Next() {
		var division divisions.Division
		err := rows.Scan(&division.ID, &division.Name, &division.PassingPercent, &division.CreatedAt)
		if err != nil {
			return nil, ErrDivisions.Wrap(err)
		}

		allDivisions = append(allDivisions, division)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrDivisions.Wrap(err)
	}

	return allDivisions, ErrDivisions.Wrap(err)
}

// Get returns division by id from the data base.
func (divisionsDB *divisionsDB) Get(ctx context.Context, id uuid.UUID) (divisions.Division, error) {
	query := `SELECT id, name, passing_percent, created_at FROM divisions WHERE id=$1`
	var division divisions.Division

	row := divisionsDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&division.ID, &division.Name, &division.PassingPercent, &division.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return division, divisions.ErrNoDivision.Wrap(err)
		}

		return division, ErrDivisions.Wrap(err)
	}

	return division, ErrDivisions.Wrap(err)
}

// GetByName returns division by name from the data base.
func (divisionsDB *divisionsDB) GetByName(ctx context.Context, divisionName int) (divisions.Division, error) {
	query := `SELECT id, name, passing_percent, created_at FROM divisions WHERE name=$1`
	var division divisions.Division

	row := divisionsDB.conn.QueryRowContext(ctx, query, divisionName)

	err := row.Scan(&division.ID, &division.Name, &division.PassingPercent, &division.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return division, divisions.ErrNoDivision.Wrap(err)
		}

		return division, ErrDivisions.Wrap(err)
	}

	return division, ErrDivisions.Wrap(err)
}

// Get returns division by id from the data base.
func (divisionsDB *divisionsDB) GetLastDivision(ctx context.Context) (divisions.Division, error) {
	query := `SELECT * FROM divisions WHERE name=(SELECT MAX(name) FROM divisions)`
	var division divisions.Division

	row := divisionsDB.conn.QueryRowContext(ctx, query)

	err := row.Scan(&division.ID, &division.Name, &division.PassingPercent, &division.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return division, divisions.ErrNoDivision.Wrap(err)
		}

		return division, ErrDivisions.Wrap(err)
	}

	return division, ErrDivisions.Wrap(err)
}

// Delete deletes a division in the database.
func (divisionsDB *divisionsDB) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := divisionsDB.conn.ExecContext(ctx, "DELETE FROM divisions WHERE id=$1", id)
	if err != nil {
		return ErrDivisions.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return divisions.ErrNoDivision.New("division does not exist")
	}

	return ErrDivisions.Wrap(err)
}
