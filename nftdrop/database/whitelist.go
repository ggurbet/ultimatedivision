// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision/nftdrop/whitelist"
	"ultimatedivision/pkg/cryptoutils"
)

// ensures that whitelistDB implements whitelist.DB.
var _ whitelist.DB = (*whitelistDB)(nil)

// ErrWhitelist indicates that there was an error in the database.
var ErrWhitelist = errs.Class("whitelist repository error")

// whitelistDB provides access to whitelist db.
//
// architecture: Database
type whitelistDB struct {
	conn *sql.DB
}

// Create add wallet in the data base.
func (whitelistDB *whitelistDB) Create(ctx context.Context, wallet whitelist.Wallet) error {
	query :=
		`INSERT INTO
			whitelist(address, password) 
		VALUES 
			($1, $2)`

	_, err := whitelistDB.conn.ExecContext(ctx, query, wallet.Address, wallet.Password)
	return ErrWhitelist.Wrap(err)
}

// GetByAddress returns wallet by address from the data base.
func (whitelistDB *whitelistDB) GetByAddress(ctx context.Context, address cryptoutils.Address) (whitelist.Wallet, error) {
	wallet := whitelist.Wallet{}
	query :=
		`SELECT
			address, password
		FROM 
			whitelist
		WHERE
			address = $1`

	err := whitelistDB.conn.QueryRowContext(ctx, query, address).Scan(&wallet.Address, &wallet.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return wallet, whitelist.ErrNoWhitelist.Wrap(err)
		}

		return wallet, whitelist.ErrNoWhitelist.New("address does not exist")
	}

	return wallet, ErrWhitelist.Wrap(err)
}

// List returns all wallets from the data base.
func (whitelistDB *whitelistDB) List(ctx context.Context) ([]whitelist.Wallet, error) {
	query :=
		`SELECT
			address, password
		FROM 
			whitelist`

	rows, err := whitelistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrWhitelist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrWhitelist.Wrap(rows.Close()))
	}()

	wallets := []whitelist.Wallet{}
	for rows.Next() {
		wallet := whitelist.Wallet{}
		if err = rows.Scan(&wallet.Address, &wallet.Password); err != nil {
			return nil, ErrWhitelist.Wrap(err)
		}
		wallets = append(wallets, wallet)
	}

	return wallets, ErrWhitelist.Wrap(rows.Err())
}

// Delete deletes wallet from the database.
func (whitelistDB *whitelistDB) Delete(ctx context.Context, address cryptoutils.Address) error {
	query := `DELETE FROM whitelist
              WHERE address = $1`

	result, err := whitelistDB.conn.ExecContext(ctx, query, address)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return whitelist.ErrNoWhitelist.New("address does not exist")
	}

	return ErrWhitelist.Wrap(err)
}

// Update updates a wallets password in the data base.
func (whitelistDB *whitelistDB) Update(ctx context.Context, wallet whitelist.Wallet) error {
	query :=
		`UPDATE whitelist 
		 SET password = $1
		 WHERE address = $2`

	result, err := whitelistDB.conn.ExecContext(ctx, query, wallet.Password, wallet.Address)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return whitelist.ErrNoWhitelist.New("wallet does not exist")
	}

	return ErrWhitelist.Wrap(err)
}

// ListWithoutPassword returns all wallets address from the data base.
func (whitelistDB *whitelistDB) ListWithoutPassword(ctx context.Context) ([]whitelist.Wallet, error) {
	query :=
		`SELECT
			address
		FROM 
			whitelist
		WHERE 
			password = ''`

	rows, err := whitelistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrWhitelist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrWhitelist.Wrap(rows.Close()))
	}()

	wallets := []whitelist.Wallet{}
	for rows.Next() {
		wallet := whitelist.Wallet{}
		if err = rows.Scan(&wallet.Address); err != nil {
			return nil, ErrWhitelist.Wrap(err)
		}
		wallets = append(wallets, wallet)
	}

	return wallets, ErrWhitelist.Wrap(rows.Err())
}
