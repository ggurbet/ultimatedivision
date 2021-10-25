// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/admin/admins"
)

// ErrAdmins indicates that there was an error in the database.
var ErrAdmins = errs.Class("admins repository error")

// adminsDB provide access to admin DB.
//
// architecture: Database
type adminsDB struct {
	conn *sql.DB
}

// List returns all admins from db.
func (adminsDB *adminsDB) List(ctx context.Context) ([]admins.Admin, error) {
	rows, err := adminsDB.conn.QueryContext(ctx, "SELECT id, email, password_hash, created_at FROM admins")
	if err != nil {
		return nil, ErrAdmins.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []admins.Admin
	for rows.Next() {
		var admin admins.Admin
		err = rows.Scan(&admin.ID, &admin.Email, &admin.PasswordHash, &admin.CreatedAt)
		if err != nil {
			return nil, ErrAdmins.Wrap(err)
		}

		data = append(data, admin)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrAdmins.Wrap(err)
	}
	return data, nil
}

// Get returns admin from db by id.
func (adminsDB *adminsDB) Get(ctx context.Context, id uuid.UUID) (admins.Admin, error) {
	var admin admins.Admin

	row := adminsDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, created_at FROM admins WHERE id=$1", id)

	err := row.Scan(&admin.ID, &admin.Email, &admin.PasswordHash, &admin.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return admin, admins.ErrNoAdmin.Wrap(err)
		}

		return admin, ErrAdmins.Wrap(err)
	}
	return admin, nil
}

// GetByEmail is a method for querying admin from the database by email.
func (adminsDB *adminsDB) GetByEmail(ctx context.Context, email string) (admins.Admin, error) {
	query := `SELECT * FROM admins WHERE email = $1;`

	var admin admins.Admin

	row := adminsDB.conn.QueryRowContext(ctx, query, email)
	if err := row.Scan(&admin.ID, &admin.Email, &admin.PasswordHash, &admin.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return admin, admins.ErrNoAdmin.Wrap(err)
		}

		return admins.Admin{}, Error.Wrap(err)
	}

	return admin, nil
}

// Create inserts admin to DB.
func (adminsDB *adminsDB) Create(ctx context.Context, admin admins.Admin) error {
	_, err := adminsDB.conn.ExecContext(ctx,
		`INSERT INTO admins(id,email,password_hash,created_at)
                VALUES($1,$2,$3,$4)`, admin.ID, admin.Email, admin.PasswordHash, admin.CreatedAt)

	return ErrAdmins.Wrap(err)
}

// Update updates admin in DB.
func (adminsDB *adminsDB) Update(ctx context.Context, admin admins.Admin) error {
	result, err := adminsDB.conn.ExecContext(ctx,
		"UPDATE admins SET password_hash = $1 WHERE id = $2",
		admin.PasswordHash, admin.ID)
	if err != nil {
		return ErrAdmins.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return admins.ErrNoAdmin.New("admin does not exist")
	}

	return ErrAdmins.Wrap(err)
}
