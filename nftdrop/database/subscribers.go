// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/mail"
	"ultimatedivision/nftdrop/subscribers"
	"ultimatedivision/pkg/pagination"
)

// ErrSubscribers indicates that there was an error in the database.
var ErrSubscribers = errs.Class("subscribers repository error")

// subscribersDB provides access to subscribers db.
//
// architecture: Database
type subscribersDB struct {
	conn *sql.DB
}

// List returns all subscribers from the data base.
func (subscribersDB *subscribersDB) List(ctx context.Context, cursor pagination.Cursor) (subscribers.Page, error) {
	var subscribersPage subscribers.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query := `SELECT email, created_at FROM subscribers LIMIT $1 OFFSET $2`

	rows, err := subscribersDB.conn.QueryContext(ctx, query, cursor.Limit, offset)
	if err != nil {
		return subscribersPage, ErrSubscribers.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var dataSubscribers []subscribers.Subscriber
	for rows.Next() {
		var subscriber subscribers.Subscriber
		err := rows.Scan(&subscriber.Email, &subscriber.CreatedAt)
		if err != nil {
			return subscribersPage, subscribers.ErrSubscribers.Wrap(err)
		}

		dataSubscribers = append(dataSubscribers, subscriber)
	}
	if err = rows.Err(); err != nil {
		return subscribersPage, ErrSubscribers.Wrap(err)
	}

	subscribersPage, err = subscribersDB.listPaginated(ctx, cursor, dataSubscribers)

	return subscribersPage, ErrSubscribers.Wrap(err)
}

// GetByEmail returns subscriber by email from the data base.
func (subscribersDB *subscribersDB) GetByEmail(ctx context.Context, email string) (subscribers.Subscriber, error) {
	var subscriber subscribers.Subscriber
	emailNormalized := mail.Normalize(email)

	row := subscribersDB.conn.QueryRowContext(ctx, "SELECT email, created_at FROM subscribers WHERE email_normalized=$1", emailNormalized)

	err := row.Scan(&subscriber.Email, &subscriber.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return subscriber, subscribers.ErrNoSubscriber.Wrap(err)
		}

		return subscriber, ErrSubscribers.Wrap(err)
	}

	return subscriber, nil
}

// Create creates a subscriber and writes to the database.
func (subscribersDB *subscribersDB) Create(ctx context.Context, subscriber subscribers.Subscriber) error {
	emailNormalized := mail.Normalize(subscriber.Email)
	query := `INSERT INTO subscribers(
                  email, 
                  email_normalized, 
                  created_at) 
                  VALUES ($1, $2, $3)`

	_, err := subscribersDB.conn.ExecContext(ctx, query, subscriber.Email, emailNormalized, subscriber.CreatedAt)

	return subscribers.ErrSubscribersDB.Wrap(err)
}

// Delete deletes a subscriber in the database.
func (subscribersDB *subscribersDB) Delete(ctx context.Context, email string) error {
	_, err := subscribersDB.conn.ExecContext(ctx, "DELETE FROM subscribers WHERE email=$1", email)

	return ErrSubscribers.Wrap(err)
}

// listPaginated returns paginated list of subscribers.
func (subscribersDB *subscribersDB) listPaginated(ctx context.Context, cursor pagination.Cursor, subscribersList []subscribers.Subscriber) (subscribers.Page, error) {
	var subscribersPage subscribers.Page
	offset := (cursor.Page - 1) * cursor.Limit

	totalCount, err := subscribersDB.totalCount(ctx)
	if err != nil {
		return subscribersPage, ErrSubscribers.Wrap(err)
	}

	pageCount := totalCount / cursor.Limit
	if totalCount%cursor.Limit != 0 {
		pageCount++
	}

	subscribersPage = subscribers.Page{
		Subscribers: subscribersList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalCount,
		},
	}

	return subscribersPage, nil
}

// totalCount counts all the subscribers in the table.
func (subscribersDB *subscribersDB) totalCount(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM subscribers`
	err := subscribersDB.conn.QueryRowContext(ctx, query).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, subscribers.ErrNoSubscriber.Wrap(err)
	}
	return count, ErrSubscribers.Wrap(err)
}
