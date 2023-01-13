// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
)

// ensures that waitlistDB implements waitlist.DB.
var _ waitlist.DB = (*waitlistDB)(nil)

// ErrWaitlist indicates that there was an error in the database.
var ErrWaitlist = errs.Class("ErrWaitlist repository error")

// waitlistDB provides access to waitlist DB.
//
// architecture: Database
type waitlistDB struct {
	conn *sql.DB
}

// Create creates item of wait list in the database.
func (waitlistDB *waitlistDB) Create(ctx context.Context, item waitlist.Item) error {
	query := `INSERT INTO waitlist(token_id, card_id, wallet_address, casper_wallet_address, casper_wallet_hash, value, password, wallet_type) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := waitlistDB.conn.ExecContext(ctx, query, item.TokenID, item.CardID, item.Wallet, item.CasperWallet, item.CasperWalletHash, item.Value.Bytes(), "", item.WalletType)
	return ErrWaitlist.Wrap(err)
}

// GetByTokenID returns item of wait list by token id.
func (waitlistDB *waitlistDB) GetByTokenID(ctx context.Context, tokenNumber int64) (waitlist.Item, error) {
	var value []byte
	query := `SELECT *
	          FROM waitlist
	          WHERE token_number = $1`

	var item waitlist.Item

	err := waitlistDB.conn.QueryRowContext(ctx, query, tokenNumber).Scan(&item.TokenID, &item.TokenNumber, &item.CardID, &item.Wallet, &item.CasperWallet, &item.CasperWalletHash, &value, &item.Password, &item.WalletType)
	if errors.Is(err, sql.ErrNoRows) {
		return item, waitlist.ErrNoItem.Wrap(err)
	}
	item.Value.SetBytes(value)

	return item, ErrWaitlist.Wrap(err)
}

// GetByCardID returns item of wait list by card id.
func (waitlistDB *waitlistDB) GetByCardID(ctx context.Context, cardID uuid.UUID) (waitlist.Item, error) {
	var value []byte
	query := `SELECT *
	          FROM waitlist
	          WHERE card_id = $1`

	var item waitlist.Item

	err := waitlistDB.conn.QueryRowContext(ctx, query, cardID).Scan(&item.TokenID, &item.TokenNumber, &item.CardID, &item.Wallet, &item.CasperWallet, &item.CasperWalletHash, &value, &item.Password, &item.WalletType)
	if errors.Is(err, sql.ErrNoRows) {
		return item, waitlist.ErrNoItem.Wrap(err)
	}
	item.Value.SetBytes(value)

	return item, ErrWaitlist.Wrap(err)
}

// GetLastTokenID returns id of last inserted item of wait list.
func (waitlistDB *waitlistDB) GetLastTokenID(ctx context.Context) (int64, error) {
	query := `SELECT token_number
	          FROM waitlist
	          ORDER BY token_number DESC
	          LIMIT 1`

	var lastToken int64

	err := waitlistDB.conn.QueryRowContext(ctx, query).Scan(&lastToken)
	if errors.Is(err, sql.ErrNoRows) {
		return lastToken, waitlist.ErrNoItem.Wrap(err)
	}

	return lastToken, ErrWaitlist.Wrap(err)
}

// List returns items of wait list from database.
func (waitlistDB *waitlistDB) List(ctx context.Context) ([]waitlist.Item, error) {
	var value []byte
	query := `SELECT * FROM waitlist`

	var waitList []waitlist.Item

	rows, err := waitlistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return waitList, ErrWaitlist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var item waitlist.Item
		err = rows.Scan(&item.TokenID, &item.TokenNumber, &item.CardID, &item.Wallet, &item.CasperWallet, &item.CasperWalletHash, &value, &item.Password, &item.WalletType)
		if err != nil {
			return waitList, ErrWaitlist.Wrap(err)
		}
		item.Value.SetBytes(value)

		waitList = append(waitList, item)
	}
	if err = rows.Err(); err != nil {
		return waitList, ErrWaitlist.Wrap(err)
	}

	return waitList, ErrWaitlist.Wrap(err)
}

// ListWithoutPassword returns items of wait list without password from database.
func (waitlistDB *waitlistDB) ListWithoutPassword(ctx context.Context) ([]waitlist.Item, error) {
	var value []byte
	query :=
		`SELECT *
	     FROM waitlist
	     WHERE password = ''`

	rows, err := waitlistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrWaitlist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var waitListWithoutPassword []waitlist.Item
	for rows.Next() {
		var item waitlist.Item
		if err = rows.Scan(&item.TokenID, &item.TokenNumber, &item.CardID, &item.Wallet, &item.CasperWallet, &item.CasperWalletHash, &value, &item.Password, &item.WalletType); err != nil {
			return nil, ErrWaitlist.Wrap(err)
		}
		item.Value.SetBytes(value)

		waitListWithoutPassword = append(waitListWithoutPassword, item)
	}

	return waitListWithoutPassword, ErrWaitlist.Wrap(rows.Err())
}

// Update updates signature of item by token id.
func (waitlistDB *waitlistDB) Update(ctx context.Context, tokenID uuid.UUID, password evmsignature.Signature) error {
	query := `UPDATE waitlist
	          SET password = $1
	          WHERE token_id = $2`

	result, err := waitlistDB.conn.ExecContext(ctx, query, password, tokenID)
	if err != nil {
		return ErrWaitlist.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return waitlist.ErrNoItem.New("item of wait list does not exist")
	}

	return nil
}

// Delete deletes item of wait list by token id.
func (waitlistDB *waitlistDB) Delete(ctx context.Context, tokenIDs []int64) error {
	query := `DELETE FROM waitlist
	          WHERE token_number = ANY($1)`

	result, err := waitlistDB.conn.ExecContext(ctx, query, pq.Array(tokenIDs))
	if err != nil {
		return ErrWaitlist.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return waitlist.ErrNoItem.New("item of wait list does not exist")
	}

	return ErrWaitlist.Wrap(err)
}
