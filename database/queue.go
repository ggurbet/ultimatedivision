// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/pagination"
	"ultimatedivision/queue"
)

// ensures that queueDB implements queue.DB.
var _ queue.DB = (*queueDB)(nil)

// ErrQueue indicates that there was an error in the database.
var ErrQueue = errs.Class("queues repository error")

// queueDB provides access to queue database.
//
// architecture: Database
type queueDB struct {
	conn *sql.DB
}

// Create adds place of queue in the database.
func (queueDB *queueDB) Create(ctx context.Context, place queue.Place) error {
	query :=
		`INSERT INTO
			places(user_id, status) 
		VALUES
			($1, $2)`

	_, err := queueDB.conn.ExecContext(ctx, query, place.UserID, place.Status)
	return ErrQueue.Wrap(err)
}

// Get returns place from the database.
func (queueDB *queueDB) Get(ctx context.Context, id uuid.UUID) (queue.Place, error) {
	place := queue.Place{}
	query :=
		`SELECT 
			user_id, status
		FROM 
			places
		WHERE 
			user_id = $1`

	err := queueDB.conn.QueryRowContext(ctx, query, id).Scan(&place.UserID, &place.Status)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return place, queue.ErrNoPlace.Wrap(err)
	case err != nil:
		return place, ErrQueue.Wrap(err)
	default:
		return place, nil
	}
}

// ListPaginated returns places in page from the database.
func (queueDB *queueDB) ListPaginated(ctx context.Context, cursor pagination.Cursor) (queue.Page, error) {
	var placesListPage queue.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query :=
		`SELECT 
			user_id, status 
		FROM 
			places 
		LIMIT 
			$1
		OFFSET 
			$2`

	rows, err := queueDB.conn.QueryContext(ctx, query, cursor.Limit, offset)
	if err != nil {
		return placesListPage, ErrQueue.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	Places := []queue.Place{}
	for rows.Next() {
		Place := queue.Place{}
		if err = rows.Scan(&Place.UserID, &Place.Status); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return placesListPage, queue.ErrNoPlace.Wrap(err)
			}
			return placesListPage, ErrQueue.Wrap(err)
		}
		Places = append(Places, Place)
	}
	if err = rows.Err(); err != nil {
		return placesListPage, ErrQueue.Wrap(err)
	}

	placesListPage, err = queueDB.listPaginated(ctx, cursor, Places)
	return placesListPage, ErrQueue.Wrap(err)
}

// listPaginated returns paginated list of places.
func (queueDB *queueDB) listPaginated(ctx context.Context, cursor pagination.Cursor, queuesList []queue.Place) (queue.Page, error) {
	var placesListPage queue.Page
	offset := (cursor.Page - 1) * cursor.Limit

	totalCount, err := queueDB.totalCount(ctx)
	if err != nil {
		return placesListPage, ErrQueue.Wrap(err)
	}

	pageCount := totalCount / cursor.Limit
	if totalCount%cursor.Limit != 0 {
		pageCount++
	}

	placesListPage = queue.Page{
		Places: queuesList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalCount,
		},
	}
	return placesListPage, nil
}

// totalCount counts all places in the table.
func (queueDB *queueDB) totalCount(ctx context.Context) (int, error) {
	var count int
	query :=
		`SELECT 
			COUNT(*) 
		FROM 
			places`

	err := queueDB.conn.QueryRowContext(ctx, query).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, queue.ErrNoPlace.Wrap(err)
	}
	return count, ErrQueue.Wrap(err)
}

// UpdateStatus updates status of place in the database.
func (queueDB *queueDB) UpdateStatus(ctx context.Context, id uuid.UUID, status queue.Status) error {
	query :=
		`UPDATE
			places 
		SET 
			status = $1 
		WHERE 
			user_id = $2`

	_, err := queueDB.conn.ExecContext(ctx, query, status, id)
	return ErrQueue.Wrap(err)
}

// Delete deletes record place in the database.
func (queueDB *queueDB) Delete(ctx context.Context, id uuid.UUID) error {
	query :=
		`DELETE FROM
			places
		WHERE 
			user_id = $1`

	_, err := queueDB.conn.ExecContext(ctx, query, id)
	return ErrQueue.Wrap(err)
}
