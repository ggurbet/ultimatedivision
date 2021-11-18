// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/mail"
	"ultimatedivision/pkg/cryptoutils"
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
	rows, err := usersDB.conn.QueryContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, last_login, status, created_at FROM users")
	if err != nil {
		return nil, ErrUsers.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.LastLogin, &user.Status, &user.CreatedAt)
		if err != nil {
			return nil, ErrUsers.Wrap(err)
		}

		data = append(data, user)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrUsers.Wrap(err)
	}

	return data, ErrUsers.Wrap(err)
}

// Get returns user by id from the data base.
func (usersDB *usersDB) Get(ctx context.Context, id uuid.UUID) (users.User, error) {
	var user users.User

	row := usersDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, last_login, status, created_at FROM users WHERE id=$1", id)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, ErrUsers.Wrap(err)
	}

	return user, ErrUsers.Wrap(err)
}

// GetByEmail returns user by email from the data base.
func (usersDB *usersDB) GetByEmail(ctx context.Context, email string) (users.User, error) {
	var user users.User
	emailNormalized := mail.Normalize(email)

	row := usersDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, last_login, status, created_at FROM users WHERE email_normalized=$1", emailNormalized)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}
		return user, ErrUsers.Wrap(err)
	}

	return user, ErrUsers.Wrap(err)
}

// GetByWalletAddress returns user by wallet address from the data base.
func (usersDB *usersDB) GetByWalletAddress(ctx context.Context, walletAddress cryptoutils.Address) (users.User, error) {
	var user users.User

	row := usersDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, last_login, status, created_at FROM users WHERE wallet_address = $1", walletAddress)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}
		return user, ErrUsers.Wrap(err)
	}

	return user, ErrUsers.Wrap(err)
}

// Create creates a user and writes to the database.
func (usersDB *usersDB) Create(ctx context.Context, user users.User) error {
	emailNormalized := mail.Normalize(user.Email)
	query := `INSERT INTO users(
                  id, 
                  email, 
                  email_normalized, 
                  password_hash, 
                  nick_name, 
                  first_name, 
                  last_name,
                  wallet_address,
                  last_login, 
                  status, 
                  created_at) 
                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := usersDB.conn.ExecContext(ctx, query, user.ID, user.Email, emailNormalized, user.PasswordHash,
		user.NickName, user.FirstName, user.LastName, user.Wallet, user.LastLogin, user.Status, user.CreatedAt)

	return ErrUsers.Wrap(err)
}

// Delete deletes a user in the database.
func (usersDB *usersDB) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := usersDB.conn.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}

// Update updates a status in the database.
func (usersDB *usersDB) Update(ctx context.Context, status users.Status, id uuid.UUID) error {
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET status=$1 WHERE id=$2", status, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
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
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET password_hash=$1 WHERE id=$2", passwordHash, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}

// UpdateWalletAddress updates wallet address in the database.
func (usersDB *usersDB) UpdateWalletAddress(ctx context.Context, wallet cryptoutils.Address, id uuid.UUID) error {
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET wallet_address=$1 WHERE id=$2", wallet, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}
