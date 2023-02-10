// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"math/big"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/marketplace/bids"
)

// ensures that bidsDB implements bids.DB.
var _ bids.DB = (*bidsDB)(nil)

// ErrBids indicates that there was an error in the database.
var ErrBids = errs.Class("bids repository error")

// bidsDB provides access to bids db.
//
// architecture: Database
type bidsDB struct {
	conn *sql.DB
}

// Create creates bid for lot in the database.
func (bidsDB *bidsDB) Create(ctx context.Context, bid bids.Bid) error {
	query := `INSERT INTO bids(id, lot_id, user_id, amount, created_at)
	          VALUES($1,$2,$3,$4,$5)`
	_, err := bidsDB.conn.ExecContext(ctx, query, bid.ID, bid.LotID, bid.UserID, bid.Amount.String(), bid.CreatedAt)
	return ErrBids.Wrap(err)
}

// GetCurrentBidByLotID returns current bid by lot id from the database.
func (bidsDB *bidsDB) GetCurrentBidByLotID(ctx context.Context, lotID uuid.UUID) (bids.Bid, error) {
	var (
		bid    bids.Bid
		amount string
	)
	query := `SELECT id, lot_id, user_id, amount, created_at
	          FROM bids
	          WHERE lot_id = $1
	          ORDER BY created_at DESC
	          LIMIT 1`

	err := bidsDB.conn.QueryRowContext(ctx, query, lotID).Scan(&bid.ID, &bid.LotID, &bid.UserID, &amount, &bid.CreatedAt)
	if errs.Is(sql.ErrNoRows, err) {
		return bid, bids.ErrNoBid.Wrap(err)
	}
	if _, ok := bid.Amount.SetString(amount, 10); !ok {
		return bid, ErrBids.New("could not parse amount equal %v from db", amount)
	}
	bid.CreatedAt = bid.CreatedAt.UTC()

	return bid, ErrBids.Wrap(err)
}

// ListByLotID returns bids by lot id from the database.
func (bidsDB *bidsDB) ListByLotID(ctx context.Context, lotID uuid.UUID) (_ []bids.Bid, err error) {
	var (
		bidsList []bids.Bid
		amount   string
	)
	query := `SELECT id, lot_id, user_id, amount, created_at
	          FROM bids
	          WHERE lot_id = $1`

	rows, err := bidsDB.conn.QueryContext(ctx, query, lotID)
	if err != nil {
		return bidsList, ErrBids.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var bid bids.Bid
		if err = rows.Scan(
			&bid.ID, &bid.LotID, &bid.UserID, &amount, &bid.CreatedAt); err != nil {
			return nil, ErrBids.Wrap(err)
		}
		if _, ok := bid.Amount.SetString(amount, 10); !ok {
			return nil, ErrBids.New("could not parse amount equal %v from db", amount)
		}

		bidsList = append(bidsList, bid)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrBids.Wrap(err)
	}

	return bidsList, ErrBids.Wrap(err)
}

// GetUserBidsAmountByLotID returns amount of user last bet on certain lot from the database.
func (bidsDB *bidsDB) GetUserBidsAmountByLotID(ctx context.Context, userID, lotID uuid.UUID) (_ []big.Int, err error) {
	var bidsAmount []big.Int
	query := `SELECT amount
	          FROM bids
	          WHERE user_id = $1 AND lot_id=$2`

	rows, err := bidsDB.conn.QueryContext(ctx, query, userID, lotID)
	if err != nil {
		return bidsAmount, ErrBids.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var (
			amount       string
			bigIntAmount big.Int
		)

		if err = rows.Scan(&amount); err != nil {
			return bidsAmount, ErrBids.Wrap(err)
		}
		if _, ok := bigIntAmount.SetString(amount, 10); !ok {
			return nil, ErrBids.New("could not parse amount equal %v from db", amount)
		}
		bidsAmount = append(bidsAmount, bigIntAmount)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrBids.Wrap(err)
	}

	return bidsAmount, ErrBids.Wrap(err)
}

// DeleteByLotID deletes bids by lot id in the database.
func (bidsDB *bidsDB) DeleteByLotID(ctx context.Context, lotID uuid.UUID) error {
	result, err := bidsDB.conn.ExecContext(ctx, "DELETE FROM bids WHERE lot_id=$1", lotID)
	if err != nil {
		return ErrBids.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return bids.ErrNoBid.New("")
	}

	return ErrBids.Wrap(err)
}
