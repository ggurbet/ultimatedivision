// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BoostyLabs/evmsignature"
	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision/nftdrop/whitelist"
	"ultimatedivision/pkg/pagination"
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
func (whitelistDB *whitelistDB) GetByAddress(ctx context.Context, address evmsignature.Address) (whitelist.Wallet, error) {
	wallet := whitelist.Wallet{}
	query :=
		`SELECT
			address, password
		FROM 
			whitelist
		WHERE
			address = $1`

	err := whitelistDB.conn.QueryRowContext(ctx, query, address).Scan(&wallet.Address, &wallet.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return wallet, whitelist.ErrNoWallet.New("address does not exist")
	}

	return wallet, ErrWhitelist.Wrap(err)
}

// List returns whitelist page from the database.
func (whitelistDB *whitelistDB) List(ctx context.Context, cursor pagination.Cursor) (whitelist.Page, error) {
	var whitelistPage whitelist.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query := `SELECT address, password
	          FROM whitelist
	          LIMIT $1
	          OFFSET $2`

	rows, err := whitelistDB.conn.QueryContext(ctx, query, cursor.Limit, offset)
	if err != nil {
		return whitelistPage, ErrWhitelist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrWhitelist.Wrap(rows.Close()))
	}()

	var wallets []whitelist.Wallet
	for rows.Next() {
		wallet := whitelist.Wallet{}
		if err = rows.Scan(&wallet.Address, &wallet.Password); err != nil {
			return whitelistPage, ErrWhitelist.Wrap(err)
		}
		wallets = append(wallets, wallet)
	}
	if err = rows.Err(); err != nil {
		return whitelistPage, ErrWhitelist.Wrap(err)
	}

	whitelistPage, err = whitelistDB.listPaginated(ctx, cursor, wallets)
	return whitelistPage, ErrWhitelist.Wrap(err)
}

// listPaginated returns paginated list of whitelist wallets.
func (whitelistDB *whitelistDB) listPaginated(ctx context.Context, cursor pagination.Cursor, walletsList []whitelist.Wallet) (whitelist.Page, error) {
	var walletsListPage whitelist.Page
	offset := (cursor.Page - 1) * cursor.Limit

	totalCount, err := whitelistDB.totalCount(ctx)
	if err != nil {
		return walletsListPage, ErrWhitelist.Wrap(err)
	}

	pageCount := totalCount / cursor.Limit
	if totalCount%cursor.Limit != 0 {
		pageCount++
	}

	walletsListPage = whitelist.Page{
		Wallets: walletsList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalCount,
		},
	}

	return walletsListPage, nil
}

// totalCount counts all the wallets in the table.
func (whitelistDB *whitelistDB) totalCount(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM whitelist`
	err := whitelistDB.conn.QueryRowContext(ctx, query).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, whitelist.ErrNoWallet.Wrap(err)
	}
	return count, ErrWhitelist.Wrap(err)
}

// Delete deletes wallet from the database.
func (whitelistDB *whitelistDB) Delete(ctx context.Context, address evmsignature.Address) error {
	query := `DELETE FROM whitelist
              WHERE address = $1`

	result, err := whitelistDB.conn.ExecContext(ctx, query, address)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return whitelist.ErrNoWallet.New("address does not exist")
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
		return whitelist.ErrNoWallet.New("wallet does not exist")
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
