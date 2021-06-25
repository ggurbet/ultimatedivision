// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ensures that cardsDB implements cards.DB.
var _ cards.DB = (*cardsDB)(nil)

// ErrCard indicates that there was an error in the database.
var ErrCard = errs.Class("cards repository error")

// cardsDB provides access to cards db.
//
// architecture: Database
type cardsDB struct {
	conn *sql.DB
}

const (
	allFields = `id, player_name, quality, picture_type, height, weight, skin_color, hair_style, hair_color, accessories, dominant_foot, user_id,
		positioning, composure, aggression, vision, awareness, crosses, acceleration, running_speed, reaction_speed, agility, stamina, strength, 
		jumping, balance, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing, forward_pass,
		finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, offside_trap, sliding, tackles,
		ball_focus, interceptions, vigilance, reflexes, diving, handling, sweeping, throwing
		`
)

// Create add card in the data base.
func (cardsDB *cardsDB) Create(ctx context.Context, card cards.Card) error {
	query :=
		`INSERT INTO
			cards(` + allFields + `) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27,$28,
			$29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)
		`
	_, err := cardsDB.conn.ExecContext(ctx, query,
		card.ID, card.PlayerName, card.Quality, card.PictureType, card.Height, card.Weight, card.SkinColor, card.HairStyle, card.HairColor,
		pq.Array(card.Accessories), card.DominantFoot, card.UserID, card.Positioning, card.Composure, card.Aggression, card.Vision, card.Awareness,
		card.Crosses, card.Acceleration, card.RunningSpeed, card.ReactionSpeed, card.Agility, card.Stamina, card.Strength, card.Jumping, card.Balance,
		card.Dribbling, card.BallControl, card.WeakFoot, card.SkillMoves, card.Finesse, card.Curve, card.Volleys, card.ShortPassing, card.LongPassing,
		card.ForwardPass, card.FinishingAbility, card.ShotPower, card.Accuracy, card.Distance, card.Penalty, card.FreeKicks, card.Corners,
		card.HeadingAccuracy, card.OffsideTrap, card.Sliding, card.Tackles, card.BallFocus, card.Interceptions, card.Vigilance, card.Reflexes,
		card.Diving, card.Handling, card.Sweeping, card.Throwing,
	)

	if err != nil {
		return ErrCard.Wrap(err)
	}

	return nil
}

// Get returns card by id from the data base.
func (cardsDB *cardsDB) Get(ctx context.Context, id uuid.UUID) (cards.Card, error) {
	fmt.Println(id)

	card := cards.Card{}
	var accessoriesArray pq.Int64Array
	query := "SELECT " + allFields + " FROM cards WHERE id=$1"
	err := cardsDB.conn.QueryRowContext(ctx, query, id).Scan(
		&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle, &card.HairColor,
		&accessoriesArray, &card.DominantFoot, &card.UserID, &card.Positioning, &card.Composure, &card.Aggression, &card.Vision, &card.Awareness,
		&card.Crosses, &card.Acceleration, &card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping,
		&card.Balance, &card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys,
		&card.ShortPassing, &card.LongPassing, &card.ForwardPass, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance,
		&card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus,
		&card.Interceptions, &card.Vigilance, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
	)

	for _, v := range accessoriesArray {
		card.Accessories = append(card.Accessories, cards.Accessories(v))
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return card, cards.ErrNoCard.Wrap(err)
	case err != nil:
		return card, ErrCard.Wrap(err)
	default:
		return card, nil
	}
}

// List returns all cards from the data base.
func (cardsDB *cardsDB) List(ctx context.Context) ([]cards.Card, error) {
	query := "SELECT " + allFields + " FROM cards"
	rows, err := cardsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	data := []cards.Card{}
	for rows.Next() {
		card := cards.Card{}
		var accessoriesArray pq.Int64Array
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle, &card.HairColor,
			&accessoriesArray, &card.DominantFoot, &card.UserID, &card.Positioning, &card.Composure, &card.Aggression, &card.Vision, &card.Awareness,
			&card.Crosses, &card.Acceleration, &card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping,
			&card.Balance, &card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys,
			&card.ShortPassing, &card.LongPassing, &card.ForwardPass, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance,
			&card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus,
			&card.Interceptions, &card.Vigilance, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
		); err != nil {
			return nil, err
		}

		for _, v := range accessoriesArray {
			card.Accessories = append(card.Accessories, cards.Accessories(v))
		}
		data = append(data, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
