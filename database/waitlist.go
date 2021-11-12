// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
	"ultimatedivision/pkg/cryptoutils"
)

// ensures that waitlistDB implements waitlist.DB.
var _ waitlist.DB = (*waitlistDB)(nil)

// ErrWaitlist indicates that there was an error in the database.
var ErrWaitlist = errs.Class("ErrWaitlist repository error")

// waitlistDB provide access to nfts DB.
//
// architecture: Database
type waitlistDB struct {
	conn *sql.DB
}

// Create creates nft for wait list in the database.
func (waitlistDB *waitlistDB) Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address) error {
	query := `INSERT INTO waitlist(card_id, wallet_address, password)
	          VALUES($1,$2,$3)`

	_, err := waitlistDB.conn.ExecContext(ctx, query, cardID, wallet, "")
	return ErrWaitlist.Wrap(err)
}

// List returns all nft for wait list from wait list from database.
func (waitlistDB *waitlistDB) List(ctx context.Context) ([]waitlist.Item, error) {
	query := `SELECT *
	          FROM waitlist`

	var WaitList []waitlist.Item

	rows, err := waitlistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return WaitList, ErrWaitlist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var nft waitlist.Item
		err = rows.Scan(&nft.TokenID, &nft.CardID, &nft.Wallet, &nft.Password)
		if err != nil {
			return WaitList, ErrWaitlist.Wrap(err)
		}

		WaitList = append(WaitList, nft)
	}
	if err = rows.Err(); err != nil {
		return WaitList, ErrWaitlist.Wrap(err)
	}

	return WaitList, ErrWaitlist.Wrap(err)
}

// ListWithoutPassword returns all nft for wait list without password from database.
func (waitlistDB *waitlistDB) ListWithoutPassword(ctx context.Context) ([]waitlist.Item, error) {
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

	var WaitListWithoutPassword []waitlist.Item
	for rows.Next() {
		var nft waitlist.Item
		if err = rows.Scan(&nft.TokenID, &nft.CardID, &nft.Wallet, &nft.Password); err != nil {
			return nil, ErrWaitlist.Wrap(err)
		}
		WaitListWithoutPassword = append(WaitListWithoutPassword, nft)
	}

	return WaitListWithoutPassword, ErrWaitlist.Wrap(rows.Err())
}

// Get returns nft for wait list by card id.
func (waitlistDB *waitlistDB) Get(ctx context.Context, tokenID int) (waitlist.Item, error) {
	query := `SELECT *
	          FROM waitlist
	          WHERE token_id = $1`

	var WaitList waitlist.Item

	err := waitlistDB.conn.QueryRowContext(ctx, query, tokenID).Scan(&WaitList.TokenID, &WaitList.CardID, &WaitList.Wallet, &WaitList.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return WaitList, waitlist.ErrNoItem.Wrap(err)
	}

	return WaitList, ErrWaitlist.Wrap(err)
}

// GetLast returns id of last inserted nft for wait list.
func (waitlistDB *waitlistDB) GetLast(ctx context.Context) (int, error) {
	query := `SELECT token_id
	          FROM waitlist
	          ORDER BY token_id DESC
	          LIMIT 1`

	var lastToken int

	err := waitlistDB.conn.QueryRowContext(ctx, query).Scan(&lastToken)
	if errors.Is(err, sql.ErrNoRows) {
		return lastToken, waitlist.ErrNoItem.Wrap(err)
	}

	return lastToken, ErrWaitlist.Wrap(err)
}

// Delete deletes nft from wait list by id of token.
func (waitlistDB *waitlistDB) Delete(ctx context.Context, tokenIDs []int) error {
	query := `DELETE FROM waitlist
	          WHERE token_id = $1`

	preparedQuery, err := waitlistDB.conn.PrepareContext(ctx, query)
	if err != nil {
		return ErrWaitlist.Wrap(err)
	}
	defer func() {
		err = preparedQuery.Close()
	}()

	for _, tokenID := range tokenIDs {
		result, err := waitlistDB.conn.ExecContext(ctx, query, tokenID)
		if err != nil {
			return ErrWaitlist.Wrap(err)
		}

		rowNum, err := result.RowsAffected()
		if err != nil {
			return ErrWaitlist.Wrap(err)
		}
		if rowNum == 0 {
			return waitlist.ErrNoItem.New("nft token does not exist")
		}
	}

	return nil
}
