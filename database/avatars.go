// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // using postgres driver.
	"github.com/zeebo/errs"

	"ultimatedivision/cards/avatars"
)

// ensures that avatarsDB implements avatars.DB.
var _ avatars.DB = (*avatarsDB)(nil)

// ErrAvatar indicates that there was an error in the database.
var ErrAvatar = errs.Class("avatars repository error")

// avatarsDB provides access to avatars db.
//
// architecture: Database
type avatarsDB struct {
	conn *sql.DB
}

const (
	allFieldsOfAvatar = `card_id, picture_type, face_color, face_type, eyebrows_type, eyebrows_color,
		eyelaser_type, hairstyle_color, hairstyle_type, nose, tshirt, beard, lips, tattoo, original_url, preview_url`
)

// Create adds avatar in the data base.
func (avatarsDB *avatarsDB) Create(ctx context.Context, avatar avatars.Avatar) error {
	query :=
		`INSERT INTO
			avatars(` + allFieldsOfAvatar + `) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		`
	_, err := avatarsDB.conn.ExecContext(ctx, query,
		avatar.CardID, avatar.PictureType, avatar.FaceColor, avatar.FaceType, avatar.EyeBrowsType, avatar.EyeBrowsColor, avatar.HairstyleColor,
		avatar.EyeLaserType, avatar.HairstyleType, avatar.Nose, avatar.Tshirt, avatar.Beard, avatar.Lips, avatar.Tattoo, avatar.OriginalURL, avatar.PreviewURL)

	return ErrAvatar.Wrap(err)
}

// Get returns avatar by id from the data base.
func (avatarsDB *avatarsDB) Get(ctx context.Context, cardID uuid.UUID) (avatars.Avatar, error) {
	avatar := avatars.Avatar{}
	query :=
		`SELECT
            ` + allFieldsOfAvatar + `
        FROM 
            avatars
        WHERE
            card_id = $1
        `
	err := avatarsDB.conn.QueryRowContext(ctx, query, cardID).Scan(
		&avatar.CardID, &avatar.PictureType, &avatar.FaceColor, &avatar.FaceType, &avatar.EyeBrowsType, &avatar.EyeBrowsColor, &avatar.HairstyleColor,
		&avatar.EyeLaserType, &avatar.HairstyleType, &avatar.Nose, &avatar.Tshirt, &avatar.Beard, &avatar.Lips, &avatar.Tattoo, &avatar.OriginalURL, &avatar.PreviewURL)
	if errors.Is(err, sql.ErrNoRows) {
		return avatar, avatars.ErrNoAvatar.Wrap(err)
	}

	return avatar, ErrAvatar.Wrap(err)
}
