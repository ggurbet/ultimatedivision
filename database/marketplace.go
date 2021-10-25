// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/zeebo/errs"

	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/pagination"
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
			lots.id, item_id, lots.type, lots.user_id, shopper_id, lots.status, start_price, max_price, current_price, start_time, end_time, period,
			cards.id, player_name, quality, height, weight, dominant_foot, is_tattoo, cards.status, cards.type,
			cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
			stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
			forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
			tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing
		FROM 
			lots
		LEFT JOIN 
			cards ON lots.item_id = cards.id
		WHERE 
			lots.id = $1
		`
	err := marketplaceDB.conn.QueryRowContext(ctx, query, id).Scan(
		&lot.ID, &lot.ItemID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status, &lot.StartPrice, &lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
		&lot.Card.ID, &lot.Card.PlayerName, &lot.Card.Quality, &lot.Card.Height, &lot.Card.Weight, &lot.Card.DominantFoot, &lot.Card.IsTattoo, &lot.Card.Status, &lot.Card.Type, &lot.Card.UserID, &lot.Card.Tactics, &lot.Card.Positioning,
		&lot.Card.Composure, &lot.Card.Aggression, &lot.Card.Vision, &lot.Card.Awareness, &lot.Card.Crosses, &lot.Card.Physique, &lot.Card.Acceleration, &lot.Card.RunningSpeed,
		&lot.Card.ReactionSpeed, &lot.Card.Agility, &lot.Card.Stamina, &lot.Card.Strength, &lot.Card.Jumping, &lot.Card.Balance, &lot.Card.Technique, &lot.Card.Dribbling,
		&lot.Card.BallControl, &lot.Card.WeakFoot, &lot.Card.SkillMoves, &lot.Card.Finesse, &lot.Card.Curve, &lot.Card.Volleys, &lot.Card.ShortPassing, &lot.Card.LongPassing,
		&lot.Card.ForwardPass, &lot.Card.Offense, &lot.Card.FinishingAbility, &lot.Card.ShotPower, &lot.Card.Accuracy, &lot.Card.Distance, &lot.Card.Penalty,
		&lot.Card.FreeKicks, &lot.Card.Corners, &lot.Card.HeadingAccuracy, &lot.Card.Defence, &lot.Card.OffsideTrap, &lot.Card.Sliding, &lot.Card.Tackles, &lot.Card.BallFocus,
		&lot.Card.Interceptions, &lot.Card.Vigilance, &lot.Card.Goalkeeping, &lot.Card.Reflexes, &lot.Card.Diving, &lot.Card.Handling, &lot.Card.Sweeping, &lot.Card.Throwing,
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
func (marketplaceDB *marketplaceDB) ListActiveLots(ctx context.Context, cursor pagination.Cursor) (marketplace.Page, error) {
	var lotsListPage marketplace.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query := fmt.Sprintf(
		`SELECT 
			lots.id, item_id, lots.type, lots.user_id, shopper_id, lots.status, start_price, max_price, current_price, start_time, end_time, period,
			cards.id, player_name, quality, height, weight, dominant_foot, is_tattoo, cards.status, cards.type,
			cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
			stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
			forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
			tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing
		FROM 
			lots
		LEFT JOIN 
			cards ON lots.item_id = cards.id
		WHERE
			lots.status = $1
		LIMIT 
			%d 
		OFFSET 
			%d
		`, cursor.Limit, offset)

	rows, err := marketplaceDB.conn.QueryContext(ctx, query, marketplace.StatusActive)
	if err != nil {
		return lotsListPage, ErrMarketplace.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	lots := []marketplace.Lot{}
	for rows.Next() {
		lot := marketplace.Lot{}
		if err = rows.Scan(
			&lot.ID, &lot.ItemID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status, &lot.StartPrice, &lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
			&lot.Card.ID, &lot.Card.PlayerName, &lot.Card.Quality, &lot.Card.Height, &lot.Card.Weight,
			&lot.Card.DominantFoot, &lot.Card.IsTattoo, &lot.Card.Status, &lot.Card.Type, &lot.Card.UserID, &lot.Card.Tactics, &lot.Card.Positioning,
			&lot.Card.Composure, &lot.Card.Aggression, &lot.Card.Vision, &lot.Card.Awareness, &lot.Card.Crosses, &lot.Card.Physique, &lot.Card.Acceleration, &lot.Card.RunningSpeed,
			&lot.Card.ReactionSpeed, &lot.Card.Agility, &lot.Card.Stamina, &lot.Card.Strength, &lot.Card.Jumping, &lot.Card.Balance, &lot.Card.Technique, &lot.Card.Dribbling,
			&lot.Card.BallControl, &lot.Card.WeakFoot, &lot.Card.SkillMoves, &lot.Card.Finesse, &lot.Card.Curve, &lot.Card.Volleys, &lot.Card.ShortPassing, &lot.Card.LongPassing,
			&lot.Card.ForwardPass, &lot.Card.Offense, &lot.Card.FinishingAbility, &lot.Card.ShotPower, &lot.Card.Accuracy, &lot.Card.Distance, &lot.Card.Penalty,
			&lot.Card.FreeKicks, &lot.Card.Corners, &lot.Card.HeadingAccuracy, &lot.Card.Defence, &lot.Card.OffsideTrap, &lot.Card.Sliding, &lot.Card.Tackles, &lot.Card.BallFocus,
			&lot.Card.Interceptions, &lot.Card.Vigilance, &lot.Card.Goalkeeping, &lot.Card.Reflexes, &lot.Card.Diving, &lot.Card.Handling, &lot.Card.Sweeping, &lot.Card.Throwing,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return lotsListPage, marketplace.ErrNoLot.Wrap(err)
			}
			return lotsListPage, ErrMarketplace.Wrap(err)
		}

		lots = append(lots, lot)
	}
	lotsListPage, err = marketplaceDB.listPaginated(ctx, cursor, lots)
	return lotsListPage, ErrMarketplace.Wrap(err)
}

// ListActiveLotsByItemID returns active lots from the data base by item id.
func (marketplaceDB *marketplaceDB) ListActiveLotsByItemID(ctx context.Context, itemIds []uuid.UUID, cursor pagination.Cursor) (marketplace.Page, error) {
	var lotsListPage marketplace.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query := fmt.Sprintf(
		`SELECT 
			lots.id, item_id, lots.type, lots.user_id, shopper_id, lots.status, start_price, max_price, current_price, start_time, end_time, period,
			cards.id, player_name, quality, height, weight, dominant_foot, is_tattoo, cards.status, cards.type,
			cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
			stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
			forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
			tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing
		FROM 
			lots
		LEFT JOIN 
			cards ON lots.item_id = cards.id
		WHERE
			lots.status = $1 AND item_id = ANY($2)
		LIMIT 
			%d 
		OFFSET 
			%d
		`, cursor.Limit, offset)

	rows, err := marketplaceDB.conn.QueryContext(ctx, query, marketplace.StatusActive, pq.Array(itemIds))
	if err != nil {
		return lotsListPage, ErrMarketplace.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	lots := []marketplace.Lot{}
	for rows.Next() {
		lot := marketplace.Lot{}
		if err = rows.Scan(
			&lot.ID, &lot.ItemID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status, &lot.StartPrice, &lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
			&lot.Card.ID, &lot.Card.PlayerName, &lot.Card.Quality, &lot.Card.Height, &lot.Card.Weight, &lot.Card.DominantFoot, &lot.Card.IsTattoo, &lot.Card.Status, &lot.Card.Type, &lot.Card.UserID, &lot.Card.Tactics, &lot.Card.Positioning,
			&lot.Card.Composure, &lot.Card.Aggression, &lot.Card.Vision, &lot.Card.Awareness, &lot.Card.Crosses, &lot.Card.Physique, &lot.Card.Acceleration, &lot.Card.RunningSpeed,
			&lot.Card.ReactionSpeed, &lot.Card.Agility, &lot.Card.Stamina, &lot.Card.Strength, &lot.Card.Jumping, &lot.Card.Balance, &lot.Card.Technique, &lot.Card.Dribbling,
			&lot.Card.BallControl, &lot.Card.WeakFoot, &lot.Card.SkillMoves, &lot.Card.Finesse, &lot.Card.Curve, &lot.Card.Volleys, &lot.Card.ShortPassing, &lot.Card.LongPassing,
			&lot.Card.ForwardPass, &lot.Card.Offense, &lot.Card.FinishingAbility, &lot.Card.ShotPower, &lot.Card.Accuracy, &lot.Card.Distance, &lot.Card.Penalty,
			&lot.Card.FreeKicks, &lot.Card.Corners, &lot.Card.HeadingAccuracy, &lot.Card.Defence, &lot.Card.OffsideTrap, &lot.Card.Sliding, &lot.Card.Tackles, &lot.Card.BallFocus,
			&lot.Card.Interceptions, &lot.Card.Vigilance, &lot.Card.Goalkeeping, &lot.Card.Reflexes, &lot.Card.Diving, &lot.Card.Handling, &lot.Card.Sweeping, &lot.Card.Throwing,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return lotsListPage, marketplace.ErrNoLot.Wrap(err)
			}
			return lotsListPage, ErrMarketplace.Wrap(err)
		}
		lots = append(lots, lot)
	}
	lotsListPage, err = marketplaceDB.listPaginated(ctx, cursor, lots)
	return lotsListPage, ErrMarketplace.Wrap(err)
}

// listPaginated returns paginated list of lots.
func (marketplaceDB *marketplaceDB) listPaginated(ctx context.Context, cursor pagination.Cursor, lotsList []marketplace.Lot) (marketplace.Page, error) {
	var lotsListPage marketplace.Page
	offset := (cursor.Page - 1) * cursor.Limit

	totalActiveCount, err := marketplaceDB.totalActiveCount(ctx)
	if err != nil {
		return lotsListPage, ErrMarketplace.Wrap(err)
	}

	pageCount := totalActiveCount / cursor.Limit
	if totalActiveCount%cursor.Limit != 0 {
		pageCount++
	}

	lotsListPage = marketplace.Page{
		Lots: lotsList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalActiveCount,
		},
	}

	return lotsListPage, nil
}

// totalActiveCount counts active lots in the table.
func (marketplaceDB *marketplaceDB) totalActiveCount(ctx context.Context) (int, error) {
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM lots WHERE lots.status = $1`)
	err := marketplaceDB.conn.QueryRowContext(ctx, query, marketplace.StatusActive).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, marketplace.ErrNoLot.Wrap(err)
	}
	return count, ErrMarketplace.Wrap(err)
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
		return nil, ErrMarketplace.Wrap(err)
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
			if errors.Is(err, sql.ErrNoRows) {
				return nil, marketplace.ErrNoLot.Wrap(err)
			}
			return nil, ErrMarketplace.Wrap(err)
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
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET shopper_id = $1 WHERE id = $2", shopperID, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}

// UpdateStatusLot updates status of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateStatusLot(ctx context.Context, id uuid.UUID, status marketplace.Status) error {
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET status = $1 WHERE id = $2", status, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}

// UpdateCurrentPriceLot updates current price of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice float64) error {
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET current_price = $1 WHERE id = $2", currentPrice, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}

// UpdateEndTimeLot updates end time of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET end_time = $1 WHERE id = $2", endTime, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}
