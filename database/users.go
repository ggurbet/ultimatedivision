// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/users"
)

// ErrUsers indicates that there was an error in the database.
var ErrUsers = errs.Class("users repository error")

// usersDB provides access to users db.
//
// architecture: Database
type usersDB struct {
	conn *sql.DB
}

// List returns all users from the data base.
func (usersDB *usersDB) List(ctx context.Context) ([]users.User, error) {
	rows, err := usersDB.conn.QueryContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, last_login, status, created_at FROM users")
	if err != nil {
		return nil, ErrUsers.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrUsers.Wrap(rows.Close()))
	}()

	var data []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
		if err != nil {
			return nil, users.ErrNoUser.Wrap(err)
		}

		data = append(data, user)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrUsers.Wrap(err)
	}

	return data, nil
}

// Get returns user by id from the data base.
func (usersDB *usersDB) Get(ctx context.Context, id uuid.UUID) (users.User, error) {
	var user users.User

	row := usersDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, last_login, status, created_at FROM users WHERE id=$1", id)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, ErrUsers.Wrap(err)
	}

	return user, nil
}

// GetByEmail returns user by email from the data base.
func (usersDB *usersDB) GetByEmail(ctx context.Context, email string) (users.User, error) {
	var user users.User
	emailNormalized := normalizeEmail(email)

	row := usersDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, last_login, status, created_at FROM users WHERE email_normalized=$1", emailNormalized)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, ErrUsers.Wrap(err)
	}

	return user, nil
}

// Create creates a user and writes to the database.
func (usersDB *usersDB) Create(ctx context.Context, user users.User) error {
	emailNormalized := normalizeEmail(user.Email)
	query := `INSERT INTO users(
                  id, 
                  email, 
                  email_normalized, 
                  password_hash, 
                  nick_name, 
                  first_name, 
                  last_name, 
                  last_login, 
                  status, 
                  created_at) 
                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := usersDB.conn.QueryContext(ctx, query, user.ID, user.Email, emailNormalized, user.PasswordHash,
		user.NickName, user.FirstName, user.LastName, user.LastLogin, user.Status, user.CreatedAt)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return nil
}

// Delete deletes a user in the database.
func (usersDB *usersDB) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := usersDB.conn.QueryContext(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return nil
}

// Update updates a status in the database.
func (usersDB *usersDB) Update(ctx context.Context, status int, id uuid.UUID) error {
	_, err := usersDB.conn.QueryContext(ctx, "UPDATE users SET status=$1 WHERE id=$2", status, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return nil
}

// GetNickNameByID returns users nickname by user id.
func (usersDB *usersDB) GetNickNameByID(ctx context.Context, id uuid.UUID) (string, error) {
	query := `SELECT nick_name
              FROM users
              WHERE id = $1`
	row := usersDB.conn.QueryRowContext(ctx, query, id)

	var nickname string

	err := row.Scan(&nickname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nickname, users.ErrNoUser.Wrap(err)
		}

		return nickname, ErrUsers.Wrap(err)
	}

	return nickname, nil
}

// UpdatePassword updates a password in the database.
func (usersDB *usersDB) UpdatePassword(ctx context.Context, passwordHash []byte, id uuid.UUID) error {
	_, err := usersDB.conn.QueryContext(ctx, "UPDATE users SET password_hash=$1 WHERE id=$2", passwordHash, id)
	return ErrUsers.Wrap(err)
}
