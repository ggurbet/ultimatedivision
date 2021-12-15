// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/udts/currencywaitlist"
)

// ensures that currencywaitlistDB implements currencywaitlist.DB.
var _ currencywaitlist.DB = (*currencywaitlistDB)(nil)

// ErrCurrencyWaitlist indicates that there was an error in the database.
var ErrCurrencyWaitlist = errs.Class("ErrCurrencyWaitlist repository error")

// currencywaitlistDB provide access to currency_waitlist DB.
//
// architecture: Database
type currencywaitlistDB struct {
	conn *sql.DB
}

// Create creates item of currency waitlist in the database.
func (currencywaitlistDB *currencywaitlistDB) Create(ctx context.Context, item currencywaitlist.Item) error {
	query := `INSERT INTO currency_waitlist(wallet_address, value, nonce, signature)
	          VALUES($1,$2,$3,$4)`

	_, err := currencywaitlistDB.conn.ExecContext(ctx, query, item.WalletAddress, item.Value.Bytes(), item.Nonce, item.Signature)
	return ErrCurrencyWaitlist.Wrap(err)
}

// List returns items of currency waitlist from database.
func (currencywaitlistDB *currencywaitlistDB) List(ctx context.Context) ([]currencywaitlist.Item, error) {
	var (
		itemList []currencywaitlist.Item
		value    []byte
	)
	query := `SELECT * FROM currency_waitlist`

	rows, err := currencywaitlistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return itemList, ErrCurrencyWaitlist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var item currencywaitlist.Item

		if err = rows.Scan(&item.WalletAddress, &value, &item.Nonce, &item.Signature); err != nil {
			return itemList, ErrCurrencyWaitlist.Wrap(err)
		}
		item.Value.SetBytes(value)
		itemList = append(itemList, item)
	}

	return itemList, ErrCurrencyWaitlist.Wrap(rows.Err())
}

// UpdateSignature updates signature of item by wallet address and nonce in the database.
func (currencywaitlistDB *currencywaitlistDB) UpdateSignature(ctx context.Context, signature cryptoutils.Signature, walletAddress cryptoutils.Address, nonce int64) error {
	query := `UPDATE currency_waitlist
	          SET signature = $1
	          WHERE wallet_address = $2 and nonce = $3`

	result, err := currencywaitlistDB.conn.ExecContext(ctx, query, signature, walletAddress, nonce)
	if err != nil {
		return ErrCurrencyWaitlist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return currencywaitlist.ErrNoItem.New("item does not exist")
	}

	return ErrCurrencyWaitlist.Wrap(err)
}

// Delete deletes item of currency waitlist by wallet address and nonce in the database.
func (currencywaitlistDB *currencywaitlistDB) Delete(ctx context.Context, walletAddress cryptoutils.Address, nonce int64) error {
	result, err := currencywaitlistDB.conn.ExecContext(ctx, "DELETE FROM currency_waitlist WHERE wallet_address = $1 and nonce = $2", walletAddress, nonce)
	if err != nil {
		return ErrCurrencyWaitlist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return currencywaitlist.ErrNoItem.New("item does not exist")
	}

	return ErrCurrencyWaitlist.Wrap(err)
}
