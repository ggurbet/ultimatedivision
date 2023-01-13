// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

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
	query := `INSERT INTO currency_waitlist(wallet_address, casper_wallet_address, casper_wallet_hash, wallet_type, value, nonce, signature)
	          VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := currencywaitlistDB.conn.ExecContext(ctx, query, item.WalletAddress, item.CasperWalletAddress, item.CasperWalletHash, item.WalletType, item.Value.Bytes(), item.Nonce, item.Signature)
	return ErrCurrencyWaitlist.Wrap(err)
}

// GetByWalletAddressAndNonce returns item of currency wait list by wallet address and nonce.
func (currencywaitlistDB *currencywaitlistDB) GetByWalletAddressAndNonce(ctx context.Context, walletAddress common.Address, nonce int64) (currencywaitlist.Item, error) {
	var (
		item  currencywaitlist.Item
		value []byte
	)
	query := `SELECT *
	          FROM currency_waitlist
	          WHERE wallet_address = $1 and nonce = $2`

	err := currencywaitlistDB.conn.QueryRowContext(ctx, query, walletAddress, nonce).Scan(&item.WalletAddress, &item.CasperWalletAddress, &item.CasperWalletHash, &item.WalletType, &value, &item.Nonce, &item.Signature)
	if errors.Is(err, sql.ErrNoRows) {
		return item, currencywaitlist.ErrNoItem.Wrap(err)
	}
	item.Value.SetBytes(value)

	return item, ErrCurrencyWaitlist.Wrap(err)
}

// GetByCasperWalletAddressAndNonce returns item of currency wait list by casper wallet address and nonce.
func (currencywaitlistDB *currencywaitlistDB) GetByCasperWalletAddressAndNonce(ctx context.Context, casperWallet string, nonce int64) (currencywaitlist.Item, error) {
	var (
		item  currencywaitlist.Item
		value []byte
	)
	query := `SELECT *
	          FROM currency_waitlist
	          WHERE casper_wallet_address = $1 and nonce = $2`

	err := currencywaitlistDB.conn.QueryRowContext(ctx, query, casperWallet, nonce).Scan(&item.WalletAddress, &item.CasperWalletAddress, &item.CasperWalletHash, &item.WalletType, &value, &item.Nonce, &item.Signature)
	if errors.Is(err, sql.ErrNoRows) {
		return item, currencywaitlist.ErrNoItem.Wrap(err)
	}
	item.Value.SetBytes(value)

	return item, ErrCurrencyWaitlist.Wrap(err)
}

// GetNonce returns number of nonce. from database.
func (currencywaitlistDB *currencywaitlistDB) GetNonce(ctx context.Context) (int64, error) {
	var nonce int64
	query := `SELECT coalesce(MAX(DISTINCT nonce),0) FROM currency_waitlist`

	err := currencywaitlistDB.conn.QueryRowContext(ctx, query).Scan(&nonce)
	if errors.Is(err, sql.ErrNoRows) {
		return nonce, currencywaitlist.ErrNoItem.Wrap(err)
	}

	return nonce, ErrCurrencyWaitlist.Wrap(err)
}

// GetNonceByWallet returns number of nonce by wallet from database.
func (currencywaitlistDB *currencywaitlistDB) GetNonceByWallet(ctx context.Context, wallet string) (int64, error) {
	var nonce int64
	query := `SELECT coalesce(MAX(DISTINCT nonce),0) FROM currency_waitlist WHERE casper_wallet_address = $1`

	err := currencywaitlistDB.conn.QueryRowContext(ctx, query, wallet).Scan(&nonce)
	if errors.Is(err, sql.ErrNoRows) {
		return nonce, currencywaitlist.ErrNoItem.Wrap(err)
	}

	return nonce, ErrCurrencyWaitlist.Wrap(err)
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

		if err = rows.Scan(&item.WalletAddress, &item.CasperWalletAddress, &item.CasperWalletHash, &item.WalletType, &value, &item.Nonce, &item.Signature); err != nil {
			return itemList, ErrCurrencyWaitlist.Wrap(err)
		}
		item.Value.SetBytes(value)
		itemList = append(itemList, item)
	}

	return itemList, ErrCurrencyWaitlist.Wrap(rows.Err())
}

// ListWithoutSignature returns items of currency waitlist without signature from database.
func (currencywaitlistDB *currencywaitlistDB) ListWithoutSignature(ctx context.Context) ([]currencywaitlist.Item, error) {
	var (
		itemList []currencywaitlist.Item
		value    []byte
	)
	query := `SELECT * FROM currency_waitlist WHERE signature = ''`

	rows, err := currencywaitlistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return itemList, ErrCurrencyWaitlist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var item currencywaitlist.Item

		if err = rows.Scan(&item.WalletAddress, &item.CasperWalletAddress, &item.CasperWalletHash, &item.WalletType, &value, &item.Nonce, &item.Signature); err != nil {
			return itemList, ErrCurrencyWaitlist.Wrap(err)
		}
		item.Value.SetBytes(value)
		itemList = append(itemList, item)
	}

	return itemList, ErrCurrencyWaitlist.Wrap(rows.Err())
}

// Update updates item by wallet address and nonce in the database.
func (currencywaitlistDB *currencywaitlistDB) Update(ctx context.Context, item currencywaitlist.Item) error {
	query := `UPDATE currency_waitlist
	          SET value = $1, signature = $2
	          WHERE wallet_address = $3 and nonce = $4`

	result, err := currencywaitlistDB.conn.ExecContext(ctx, query, item.Value.Bytes(), item.Signature, item.WalletAddress, item.Nonce)
	if err != nil {
		return ErrCurrencyWaitlist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return currencywaitlist.ErrNoItem.New("item does not exist")
	}

	return ErrCurrencyWaitlist.Wrap(err)
}

// UpdateNonceByWallet updates item by wallet address and nonce in the database.
func (currencywaitlistDB *currencywaitlistDB) UpdateNonceByWallet(ctx context.Context, nonce int64, casperWallet string) error {
	query := `UPDATE currency_waitlist
	          SET nonce = $1
	          WHERE casper_wallet_address = $2`

	result, err := currencywaitlistDB.conn.ExecContext(ctx, query, nonce, casperWallet)
	if err != nil {
		return ErrCurrencyWaitlist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return currencywaitlist.ErrNoItem.New("item does not exist")
	}

	return ErrCurrencyWaitlist.Wrap(err)
}

// UpdateSignature updates signature of item by wallet address and nonce in the database.
func (currencywaitlistDB *currencywaitlistDB) UpdateSignature(ctx context.Context, signature evmsignature.Signature, walletAddress common.Address, nonce int64) error {
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

// UpdateCasperSignature updates signature of item by casper wallet address and nonce in the database.
func (currencywaitlistDB *currencywaitlistDB) UpdateCasperSignature(ctx context.Context, signature evmsignature.Signature, casperWallet string, nonce int64) error {
	query := `UPDATE currency_waitlist
	          SET signature = $1
	          WHERE casper_wallet_address = $2 and nonce = $3`

	result, err := currencywaitlistDB.conn.ExecContext(ctx, query, signature, casperWallet, nonce)
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
func (currencywaitlistDB *currencywaitlistDB) Delete(ctx context.Context, walletAddress common.Address, nonce int64) error {
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
