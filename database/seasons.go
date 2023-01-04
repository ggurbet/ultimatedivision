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

	"ultimatedivision/seasons"
)

// ensures that seasonsDB implements seasons.DB.
var _ seasons.DB = (*seasonsDB)(nil)

// ErrSeasons indicates that there was an error in the database.
var ErrSeasons = errs.Class("seasons repository error")

// seasonsDB provides access to seasons db.
//
// architecture: Database
type seasonsDB struct {
	conn *sql.DB
}

// Create creates a seasons and writes to the database.
func (seasonsDB *seasonsDB) Create(ctx context.Context, season seasons.Season) error {
	query := `INSERT INTO seasons(division_id, started_at, ended_at) 
	VALUES ($1, $2, $3)`

	_, err := seasonsDB.conn.ExecContext(ctx, query, season.DivisionID, season.StartedAt, season.EndedAt)

	return ErrSeasons.Wrap(err)
}

// CreateReward creates a season reward and writes to the database.
func (seasonsDB *seasonsDB) CreateReward(ctx context.Context, reward seasons.Reward) error {
	query := `INSERT INTO season_rewards(id, season_id, user_id, wallet_address, casper_wallet_address, wallet_type, value, status)
	          VALUES($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := seasonsDB.conn.ExecContext(ctx, query, &reward.ID, &reward.SeasonID, &reward.UserID, &reward.WalletAddress, &reward.CasperWalletAddress, &reward.WalletType, reward.Value.Bytes(), &reward.Status)

	return ErrSeasons.Wrap(err)
}

// EndSeason updates a status in the database when season ended.
func (seasonsDB *seasonsDB) EndSeason(ctx context.Context, id int) error {
	db, err := seasonsDB.conn.ExecContext(ctx, "UPDATE seasons SET ended_at=$1 WHERE id=$2", time.Now().UTC(), id)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	rowNum, err := db.RowsAffected()
	if rowNum == 0 {
		return seasons.ErrNoSeason.New("season does not exist")
	}

	return ErrSeasons.Wrap(err)
}

// List returns all seasons from the database.
func (seasonsDB *seasonsDB) List(ctx context.Context) ([]seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons`

	rows, err := seasonsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allSeasons []seasons.Season
	for rows.Next() {
		var season seasons.Season
		err := rows.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
		if err != nil {
			return nil, ErrSeasons.Wrap(err)
		}

		allSeasons = append(allSeasons, season)
	}

	return allSeasons, ErrSeasons.Wrap(rows.Err())
}

// ListOfUnpaidRewardsByUserID returns all unpaid season rewards from the database by user id.
func (seasonsDB *seasonsDB) ListOfUnpaidRewardsByUserID(ctx context.Context, userID uuid.UUID) ([]seasons.Reward, error) {
	query := `SELECT * FROM season_rewards WHERE user_id=$1 and status=$2`

	rows, err := seasonsDB.conn.QueryContext(ctx, query, userID, seasons.StatusUnPaid)
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allRewards []seasons.Reward
	for rows.Next() {
		var reward seasons.Reward
		var value []byte
		err := rows.Scan(&reward.ID, &reward.SeasonID, &reward.UserID, &reward.WalletAddress, &reward.CasperWalletAddress, &reward.WalletType, &value, &reward.Status)
		if err != nil {
			return nil, ErrSeasons.Wrap(err)
		}

		reward.Value.SetBytes(value)

		allRewards = append(allRewards, reward)
	}

	return allRewards, ErrSeasons.Wrap(rows.Err())
}

// Get returns season by id from the database.
func (seasonsDB *seasonsDB) Get(ctx context.Context, id int) (seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons WHERE id=$1`
	var season seasons.Season

	row := seasonsDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return season, seasons.ErrNoSeason.Wrap(err)
		}

		return season, ErrSeasons.Wrap(err)
	}

	return season, ErrSeasons.Wrap(err)
}

// GetRewardByUserID returns user reward by id from the database.
func (seasonsDB *seasonsDB) GetRewardByUserID(ctx context.Context, userID uuid.UUID) (seasons.Reward, error) {
	query := `SELECT * FROM season_rewards WHERE user_id=$1`
	var reward seasons.Reward

	row := seasonsDB.conn.QueryRowContext(ctx, query, userID)

	err := row.Scan(&reward.ID, &reward.SeasonID, &reward.UserID, &reward.WalletAddress, &reward.CasperWalletAddress, &reward.WalletType, &reward.Value, &reward.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return reward, seasons.ErrNoSeason.Wrap(err)
		}

		return reward, ErrSeasons.Wrap(err)
	}

	return reward, ErrSeasons.Wrap(err)
}

// GetCurrentSeasons returns all current seasons from the database.
func (seasonsDB *seasonsDB) GetCurrentSeasons(ctx context.Context) ([]seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons WHERE ended_at=$1`

	rows, err := seasonsDB.conn.QueryContext(ctx, query, time.Time{})
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allSeasons []seasons.Season
	for rows.Next() {
		var season seasons.Season
		err := rows.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
		if err != nil {
			return nil, ErrSeasons.Wrap(err)
		}

		allSeasons = append(allSeasons, season)
	}

	return allSeasons, ErrSeasons.Wrap(rows.Err())
}

// Delete deletes a season in the database.
func (seasonsDB *seasonsDB) Delete(ctx context.Context, id int) error {
	result, err := seasonsDB.conn.ExecContext(ctx, "DELETE FROM seasons WHERE id=$1", id)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return seasons.ErrNoSeason.New("season does not exist")
	}

	return ErrSeasons.Wrap(err)
}

// GetSeasonByDivisionID returns season by division id from the database.
func (seasonsDB *seasonsDB) GetSeasonByDivisionID(ctx context.Context, divisionID uuid.UUID) (seasons.Season, error) {
	query := `SELECT id, division_id, started_at, ended_at FROM seasons WHERE division_id=$1 AND ended_at=$2`
	var season seasons.Season

	row := seasonsDB.conn.QueryRowContext(ctx, query, divisionID, time.Time{})

	err := row.Scan(&season.ID, &season.DivisionID, &season.StartedAt, &season.EndedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return season, seasons.ErrNoSeason.Wrap(err)
		}

		return season, ErrSeasons.Wrap(err)
	}

	return season, ErrSeasons.Wrap(err)
}
