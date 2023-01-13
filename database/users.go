// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/mail"
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

// GetVelasData get json string by user id from the database.
func (usersDB usersDB) GetVelasData(ctx context.Context, userID uuid.UUID) (users.VelasData, error) {
	var user users.VelasData

	row := usersDB.conn.QueryRowContext(ctx, "SELECT user_id, response FROM velas_register_data WHERE user_id=$1", userID)

	err := row.Scan(&user.ID, &user.Response)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, ErrUsers.Wrap(err)
		}

		return user, ErrUsers.Wrap(err)
	}

	return user, ErrUsers.Wrap(err)
}

// List returns all users from the database.
func (usersDB *usersDB) List(ctx context.Context) ([]users.User, error) {
	rows, err := usersDB.conn.QueryContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, casper_wallet_address, casper_wallet_hash, wallet_type, nonce, public_key, private_key, last_login, status, created_at FROM users")
	if err != nil {
		return nil, ErrUsers.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.CasperWallet, &user.CasperWalletHash, &user.WalletType, &user.Nonce, &user.PublicKey, &user.PrivateKey, &user.LastLogin, &user.Status, &user.CreatedAt)
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

// Get returns user by id from the database.
func (usersDB *usersDB) Get(ctx context.Context, id uuid.UUID) (users.User, error) {
	var user users.User

	row := usersDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, casper_wallet_address, casper_wallet_hash, wallet_type, nonce, public_key, private_key, last_login, status, created_at FROM users WHERE id=$1", id)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.CasperWallet, &user.CasperWalletHash, &user.WalletType, &user.Nonce, &user.PublicKey, &user.PrivateKey, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, ErrUsers.Wrap(err)
	}

	return user, ErrUsers.Wrap(err)
}

// GetByEmail returns user by email from the database.
func (usersDB *usersDB) GetByEmail(ctx context.Context, email string) (users.User, error) {
	var user users.User
	emailNormalized := mail.Normalize(email)

	row := usersDB.conn.QueryRowContext(ctx, "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, casper_wallet_address, casper_wallet_hash, wallet_type, nonce, public_key, private_key, last_login, status, created_at FROM users WHERE email_normalized=$1", emailNormalized)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.CasperWallet, &user.CasperWalletHash, &user.WalletType, &user.Nonce, &user.PublicKey, &user.PrivateKey, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}
		return user, ErrUsers.Wrap(err)
	}

	return user, ErrUsers.Wrap(err)
}

// GetByWalletAddress returns user by wallet address from the database.
func (usersDB *usersDB) GetByWalletAddress(ctx context.Context, walletAddress string, walletType users.WalletType) (users.User, error) {
	var user users.User
	var row *sql.Row

	query := "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, casper_wallet_address, casper_wallet_hash, wallet_type, nonce, public_key, private_key, last_login, status, created_at FROM users WHERE "

	switch walletType {
	case users.WalletTypeCasper:
		row = usersDB.conn.QueryRowContext(ctx, query+"casper_wallet_address=$1", walletAddress)
	case users.WalletTypeETH, users.WalletTypeVelas:
		row = usersDB.conn.QueryRowContext(ctx, query+"wallet_address=$1", common.HexToAddress(walletAddress))
	}

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.CasperWallet, &user.CasperWalletHash, &user.WalletType, &user.Nonce, &user.PublicKey, &user.PrivateKey, &user.LastLogin, &user.Status, &user.CreatedAt)
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
                  casper_wallet_address,
                  casper_wallet_hash,
                  wallet_type,
                  nonce,
                  public_key,
                  private_key,
                  last_login, 
                  status, 
                  created_at) 
                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`

	_, err := usersDB.conn.ExecContext(ctx, query, user.ID, user.Email, emailNormalized, user.PasswordHash,
		user.NickName, user.FirstName, user.LastName, user.Wallet, user.CasperWallet, user.CasperWalletHash, user.WalletType, user.Nonce, user.PublicKey, user.PrivateKey, user.LastLogin, user.Status, user.CreatedAt)

	return ErrUsers.Wrap(err)
}

// SetVelasData save json to db while register velas user.
func (usersDB *usersDB) SetVelasData(ctx context.Context, velasData users.VelasData) error {
	query := `INSERT INTO velas_register_data(
                 user_id,
                 response
                )
                 VALUES ($1, $2)`

	_, err := usersDB.conn.ExecContext(ctx, query, velasData.ID, velasData.Response)

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
func (usersDB *usersDB) UpdateWalletAddress(ctx context.Context, wallet common.Address, walletType users.WalletType, id uuid.UUID) error {
	var query string

	switch walletType {
	case users.WalletTypeETH:
		query = "UPDATE users SET wallet_address=$1, wallet_type=$2 WHERE id=$3"
	case users.WalletTypeVelas:
		query = "UPDATE users SET wallet_address=$1, wallet_type=$2 WHERE id=$3"
	case users.WalletTypeCasper:
		query = "UPDATE users SET casper_wallet_address=$1, wallet_type=$2 WHERE id=$3"
	}

	result, err := usersDB.conn.ExecContext(ctx, query, wallet, walletType, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}

// UpdateCasperWalletAddress updates Casper wallet address in the database.
func (usersDB *usersDB) UpdateCasperWalletAddress(ctx context.Context, wallet string, walletType users.WalletType, id uuid.UUID) error {
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET casper_wallet_address=$1, wallet_type=$2 WHERE id=$3", wallet, walletType, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}

// UpdateNonce updates nonce by user.
func (usersDB *usersDB) UpdateNonce(ctx context.Context, id uuid.UUID, nonce []byte) error {
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET nonce=$1 WHERE id=$2", nonce, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}

// UpdateLastLogin updates last login time.
func (usersDB *usersDB) UpdateLastLogin(ctx context.Context, id uuid.UUID) error {
	lastLogin := time.Now().UTC()
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET last_login=$1 WHERE id=$2", lastLogin, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}

// UpdateEmail updates an email address in the database.
func (usersDB *usersDB) UpdateEmail(ctx context.Context, id uuid.UUID, newEmail string) error {
	emailNormalized := mail.Normalize(newEmail)
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET email=$1, email_normalized=$2 WHERE id=$3",
		newEmail, emailNormalized, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}

// GetByPublicKey returns user by public key from the database.
func (usersDB *usersDB) GetByPublicKey(ctx context.Context, publicKey string) (users.User, error) {
	var user users.User
	var row *sql.Row

	query := "SELECT id, email, password_hash, nick_name, first_name, last_name, wallet_address, casper_wallet_address, casper_wallet_hash, wallet_type, nonce, public_key, private_key, last_login, status, created_at FROM users WHERE public_key=$1"

	row = usersDB.conn.QueryRowContext(ctx, query, publicKey)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.Wallet, &user.CasperWallet, &user.CasperWalletHash, &user.WalletType, &user.Nonce, &user.PublicKey, &user.PrivateKey, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, users.ErrNoUser.Wrap(err)
		}
		return user, ErrUsers.Wrap(err)
	}

	return user, nil
}

// UpdatePublicPrivateKey updates public and private key by user.
func (usersDB *usersDB) UpdatePublicPrivateKey(ctx context.Context, id uuid.UUID, publicKey, privateKey string) error {
	result, err := usersDB.conn.ExecContext(ctx, "UPDATE users SET public_key=$1, private_key=$2 WHERE id=$3", publicKey, privateKey, id)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return users.ErrNoUser.New("user does not exist")
	}

	return ErrUsers.Wrap(err)
}
