// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/marketplace"
)

// ensures that marketplaceDB implements marketplace.DB.
var _ marketplace.DB = (*marketplaceDB)(nil)

// ErrMarketplace indicates that there was an error in the database.
var ErrMarketplace = errs.Class("marketplace repository error")

// marketplaceDB provides access to marketplace db.
//
// architecture: Database
type marketplaceDB struct {
	conn *sql.DB
}

const (
	allLotOfFields = `id, item_id, type, user_id, shopper_id, status, start_price, max_price, current_price, start_time, end_time, period`
)

// CreateLot creates lot in the db.
func (marketplaceDB *marketplaceDB) CreateLot(ctx context.Context, lot marketplace.Lot) error {
	query :=
		`INSERT INTO 
			lots(` + allLotOfFields + `)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		`

	_, err := marketplaceDB.conn.ExecContext(ctx, query,
		lot.ID, lot.ItemID, lot.Type, lot.UserID, lot.ShopperID, lot.Status,
		lot.StartPrice, lot.MaxPrice, lot.CurrentPrice, lot.StartTime, lot.EndTime, lot.Period)

	return ErrMarketplace.Wrap(err)
}

// GetLotByID returns lot by id from the data base.
func (marketplaceDB *marketplaceDB) GetLotByID(ctx context.Context, id uuid.UUID) (marketplace.Lot, error) {
	lot := marketplace.Lot{}
	query :=
		`SELECT 
			` + allLotOfFields + `
		FROM 
			lots
		WHERE 
			id = $1
		`
	err := marketplaceDB.conn.QueryRowContext(ctx, query, id).Scan(
		&lot.ID, &lot.ItemID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status,
		&lot.StartPrice, &lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return lot, marketplace.ErrNoLot.Wrap(err)
	case err != nil:
		return lot, ErrMarketplace.Wrap(err)
	default:
		return lot, nil
	}
}

// ListActiveLots returns active lots from the data base.
func (marketplaceDB *marketplaceDB) ListActiveLots(ctx context.Context) ([]marketplace.Lot, error) {
	query :=
		`SELECT 
			` + allLotOfFields + ` 
		FROM 
			lots
		WHERE
			status = $1
		`

	rows, err := marketplaceDB.conn.QueryContext(ctx, query, marketplace.StatusActive)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	lots := []marketplace.Lot{}
	for rows.Next() {
		lot := marketplace.Lot{}
		if err = rows.Scan(
			&lot.ID, &lot.ItemID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status,
			&lot.StartPrice, &lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
		); err != nil {
			return nil, marketplace.ErrNoLot.Wrap(err)
		}

		lots = append(lots, lot)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	return lots, nil
}

// ListExpiredLot returns active lots where end time lower than or equal to time now UTC from the data base.
func (marketplaceDB *marketplaceDB) ListExpiredLot(ctx context.Context) ([]marketplace.Lot, error) {
	query :=
		`SELECT 
			` + allLotOfFields + ` 
		FROM 
			lots
		WHERE
			status = $1
		AND
			end_time <= $2
		`

	rows, err := marketplaceDB.conn.QueryContext(ctx, query, marketplace.StatusActive, time.Now().UTC())
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	lots := []marketplace.Lot{}
	for rows.Next() {
		lot := marketplace.Lot{}
		if err = rows.Scan(
			&lot.ID, &lot.ItemID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status,
			&lot.StartPrice, &lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
		); err != nil {
			return nil, marketplace.ErrNoLot.Wrap(err)
		}

		lots = append(lots, lot)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	return lots, nil
}

// UpdateShopperIDLot updates shopper id of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateShopperIDLot(ctx context.Context, id, shopperID uuid.UUID) error {
	_, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET shopper_id = $1 WHERE id = $2", shopperID, id)
	return ErrMarketplace.Wrap(err)
}

// UpdateStatusLot updates status of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateStatusLot(ctx context.Context, id uuid.UUID, status marketplace.Status) error {
	_, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET status = $1 WHERE id = $2", status, id)
	return ErrMarketplace.Wrap(err)
}

// UpdateCurrentPriceLot updates current price of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice float64) error {
	_, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET current_price = $1 WHERE id = $2", currentPrice, id)
	return ErrMarketplace.Wrap(err)
}

// UpdateEndTimeLot updates end time of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	_, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET end_time = $1 WHERE id = $2", endTime, id)
	return ErrMarketplace.Wrap(err)
}
