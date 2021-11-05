// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/matches"
	"ultimatedivision/pkg/pagination"
)

// ensures that matchesDB implements matches.DB.
var _ matches.DB = (*matchesDB)(nil)

// ErrMatches indicates that there was an error in the database.
var ErrMatches = errs.Class("matches repository error")

// matchesDB provide access to matches DB.
//
// architecture: Database
type matchesDB struct {
	conn *sql.DB
}

// Create inserts match in the database.
func (matchesDB *matchesDB) Create(ctx context.Context, match matches.Match) error {
	query := `INSERT INTO matches(id, user1_id, squad1_id, user1_points, user2_id, squad2_id, user2_points)
              VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err := matchesDB.conn.ExecContext(ctx, query, match.ID, match.User1ID,
		match.Squad1ID, match.User1Points, match.User2ID, match.Squad2ID, match.User2Points)

	return ErrMatches.Wrap(err)
}

// Get returns match from the database.
func (matchesDB *matchesDB) Get(ctx context.Context, id uuid.UUID) (matches.Match, error) {
	query := `SELECT id, user1_id, squad1_id, user1_points, user2_id, squad2_id, user2_points
              FROM matches
              WHERE id = $1`

	var match matches.Match

	row := matchesDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&match.ID, &match.User1ID, &match.Squad1ID, &match.User1Points, &match.User2ID, &match.Squad2ID, &match.User2Points)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return match, matches.ErrNoMatch.Wrap(err)
		}

		return match, ErrMatches.Wrap(err)
	}

	return match, ErrMatches.Wrap(err)
}

// ListMatches returns all matches from the database.
func (matchesDB *matchesDB) ListMatches(ctx context.Context, cursor pagination.Cursor) (matches.Page, error) {
	var matchesListPage matches.Page
	offset := (cursor.Page - 1) * cursor.Limit

	query := `SELECT id, user1_id, squad1_id, user1_points, user2_id, squad2_id, user2_points
	          FROM matches
	          LIMIT $1
	          OFFSET $2`

	rows, err := matchesDB.conn.QueryContext(ctx, query, cursor.Limit, offset)
	if err != nil {
		return matchesListPage, ErrMatches.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allMatches []matches.Match

	for rows.Next() {
		var match matches.Match
		err = rows.Scan(&match.ID, &match.User1ID, &match.Squad1ID, &match.User1Points, &match.User2ID, &match.Squad2ID, &match.User2Points)
		if err != nil {
			return matchesListPage, ErrMatches.Wrap(err)
		}

		allMatches = append(allMatches, match)
	}
	if err = rows.Err(); err != nil {
		return matchesListPage, ErrMatches.Wrap(err)
	}

	matchesListPage, err = matchesDB.listPaginated(ctx, cursor, allMatches)

	return matchesListPage, ErrMatches.Wrap(err)
}

// listPaginated returns paginated list of matches.
func (matchesDB *matchesDB) listPaginated(ctx context.Context, cursor pagination.Cursor, matchesList []matches.Match) (matches.Page, error) {
	var matchesPage matches.Page
	offset := (cursor.Page - 1) * cursor.Limit

	totalMatchesCount, err := matchesDB.countMatches(ctx)
	if err != nil {
		return matchesPage, ErrMatches.Wrap(err)
	}

	pageCount := totalMatchesCount / cursor.Limit
	if totalMatchesCount%cursor.Limit != 0 {
		pageCount++
	}

	matchPages := matches.Page{
		Matches: matchesList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalMatchesCount,
		},
	}

	return matchPages, ErrMatches.Wrap(err)
}

// listMatches counts all matches from the database.
func (matchesDB *matchesDB) countMatches(ctx context.Context) (int, error) {
	query := `SELECT count(*) FROM matches`

	var count int

	err := matchesDB.conn.QueryRowContext(ctx, query).Scan(&count)

	return count, ErrMatches.Wrap(err)
}

// Delete deletes match from the database.
func (matchesDB *matchesDB) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM matches
              WHERE id = $1`

	result, err := matchesDB.conn.ExecContext(ctx, query, id)
	if err != nil {
		return ErrMatches.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err != nil {
		return ErrMatches.Wrap(err)
	}
	if rowNum == 0 {
		return matches.ErrNoMatch.New("match does not exist")
	}

	return ErrMatches.Wrap(err)
}

// UpdateMatch updates the number of points that users received for a played match.
func (matchesDB *matchesDB) UpdateMatch(ctx context.Context, match matches.Match) error {
	query := `UPDATE matches
	          SET user1_points = $1, user2_points = $2
	          WHERE id = $3`

	result, err := matchesDB.conn.ExecContext(ctx, query, match.User1Points, match.User2Points, match.ID)
	if err != nil {
		return ErrMatches.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err != nil {
		return ErrMatches.Wrap(err)
	}
	if rowNum == 0 {
		return matches.ErrNoMatch.New("match does not exist")
	}

	return ErrMatches.Wrap(err)
}

// AddGoals adds goals in the match.
func (matchesDB *matchesDB) AddGoals(ctx context.Context, matchGoals []matches.MatchGoals) error {
	query := `INSERT INTO match_results(id, match_id, user_id, card_id, minute)
	          VALUES($1,$2,$3,$4,$5)`

	preparedQuery, err := matchesDB.conn.PrepareContext(ctx, query)
	if err != nil {
		return ErrMatches.Wrap(err)
	}
	defer func() {
		err = preparedQuery.Close()
	}()

	for _, matchGoal := range matchGoals {
		_, err = preparedQuery.ExecContext(ctx, matchGoal.ID, matchGoal.MatchID,
			matchGoal.UserID, matchGoal.CardID, matchGoal.Minute)

		if err != nil {
			return ErrMatches.Wrap(err)
		}
	}

	if err = preparedQuery.Close(); err != nil {
		return ErrMatches.Wrap(err)
	}

	return ErrMatches.Wrap(err)
}

// ListMatchGoals returns all goals from the match from the database.
func (matchesDB *matchesDB) ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]matches.MatchGoals, error) {
	query := `SELECT id, match_id, user_id, card_id, minute
              FROM match_results
              WHERE match_id = $1
              ORDER BY minute`

	rows, err := matchesDB.conn.QueryContext(ctx, query, matchID)
	if err != nil {
		return nil, ErrMatches.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var goals []matches.MatchGoals

	for rows.Next() {
		var goal matches.MatchGoals
		err = rows.Scan(&goal.ID, &goal.MatchID, &goal.UserID, &goal.CardID, &goal.Minute)
		if err != nil {
			return nil, ErrMatches.Wrap(err)
		}

		goals = append(goals, goal)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMatches.Wrap(err)
	}

	return goals, ErrMatches.Wrap(err)
}

// GetMatchResult returns goals of each user in the match from db.
func (matchesDB *matchesDB) GetMatchResult(ctx context.Context, matchID uuid.UUID) ([]matches.MatchResult, error) {
	query := `SELECT user_id, COUNT(user_id)
              FROM match_results
              WHERE match_id = $1
              GROUP BY user_id`

	rows, err := matchesDB.conn.QueryContext(ctx, query, matchID)
	if err != nil {
		return nil, ErrMatches.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var goals []matches.MatchResult

	for rows.Next() {
		var goal matches.MatchResult
		err = rows.Scan(&goal.UserID, &goal.QuantityGoals)
		if err != nil {
			return nil, ErrMatches.Wrap(err)
		}

		goals = append(goals, goal)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMatches.Wrap(err)
	}

	return goals, ErrMatches.Wrap(err)
}
