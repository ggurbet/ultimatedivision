// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"ultimatedivision/admin/admins"
	"ultimatedivision/divisions"
	"ultimatedivision/internal/mail"
	"ultimatedivision/users"
)

// CreateUser creates a user and writes to the database.
func CreateUser(ctx context.Context, db *sql.DB) error {
	testUser := users.User{
		ID:           uuid.New(),
		Email:        "test@test.com",
		PasswordHash: []byte("Qwerty123-"),
		NickName:     "Admin",
		FirstName:    "Test",
		LastName:     "Test",
		Wallet:       "Test",
		LastLogin:    time.Time{},
		Status:       1,
		CreatedAt:    time.Now().UTC(),
	}

	err := testUser.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}

	emailNormalized := mail.Normalize(testUser.Email)
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

	_, err = db.ExecContext(ctx, query, testUser.ID, testUser.Email, emailNormalized, testUser.PasswordHash,
		testUser.NickName, testUser.FirstName, testUser.LastName, testUser.Wallet, testUser.LastLogin, testUser.Status, testUser.CreatedAt)

	return ErrUsers.Wrap(err)
}

// CreateAdmin inserts admin to DB.
func CreateAdmin(ctx context.Context, conn *sql.DB) error {
	testAdmin := admins.Admin{
		ID:           uuid.New(),
		Email:        "test@test.com",
		PasswordHash: []byte("Qwerty123-"),
		CreatedAt:    time.Now().UTC(),
	}
	err := testAdmin.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}
	_, err = conn.ExecContext(ctx,
		`INSERT INTO admins(id,email,password_hash,created_at)
                VALUES($1,$2,$3,$4)`, testAdmin.ID, testAdmin.Email, testAdmin.PasswordHash, testAdmin.CreatedAt)

	return ErrAdmins.Wrap(err)
}

// CreateDivisions creates a division and writes to the database.
func CreateDivisions(ctx context.Context, conn *sql.DB) error {
	division1 := divisions.Division{
		ID:             uuid.New(),
		Name:           1,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}
	division2 := divisions.Division{
		ID:             uuid.New(),
		Name:           2,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}
	division3 := divisions.Division{
		ID:             uuid.New(),
		Name:           3,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}
	division4 := divisions.Division{
		ID:             uuid.New(),
		Name:           4,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}

	divisions := []divisions.Division{division1, division2, division3, division4}

	query := `INSERT INTO divisions(id, name, passing_percent, created_at) 
	VALUES ($1, $2, $3, $4)`

	for _, d := range divisions {
		_, err := conn.ExecContext(ctx, query, d.ID, d.Name, d.PassingPercent, d.CreatedAt)
		if err != nil {
			return ErrDivisions.Wrap(err)
		}
	}

	return nil
}
