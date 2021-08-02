// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ensures that cardsDB implements cards.DB.
var _ cards.DB = (*cardsDB)(nil)

// ErrCard indicates that there was an error in the database.
var ErrCard = errs.Class("cards repository error")

// Action defines the list of possible filter actions.
type Action string

const (
	// EQ - equal to value.
	EQ Action = "="
	// GTE - greater than or equal to value.
	GTE Action = ">="
	// LTE - less than or equal to value.
	LTE Action = "<="
	// LIKE - like to value.
	LIKE Action = "LIKE"
)

// cardsDB provides access to cards db.
//
// architecture: Database
type cardsDB struct {
	conn *sql.DB
}

const (
	allFields = `id, player_name, quality, picture_type, height, weight, skin_color, hair_style, hair_color, dominant_foot, is_tattoos, user_id,
		tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed,
		agility, stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve,
		volleys, short_passing, long_passing, forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty,
		free_kicks, corners, heading_accuracy, defence, offside_trap, sliding, tackles, ball_focus, interceptions, vigilance, goalkeeping,
		reflexes, diving, handling, sweeping, throwing
		`
)

// Create add card in the data base.
func (cardsDB *cardsDB) Create(ctx context.Context, card cards.Card) error {
	tx, err := cardsDB.conn.BeginTx(ctx, nil)
	if err != nil {
		return ErrCard.Wrap(err)
	}

	query :=
		`INSERT INTO
			cards(` + allFields + `) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25,
			$26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49,
			$50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61)
		`
	_, err = cardsDB.conn.ExecContext(ctx, query,
		card.ID, card.PlayerName, card.Quality, card.PictureType, card.Height, card.Weight, card.SkinColor, card.HairStyle, card.HairColor,
		card.DominantFoot, card.IsTattoos, card.UserID, card.Tactics, card.Positioning, card.Composure, card.Aggression, card.Vision, card.Awareness,
		card.Crosses, card.Physique, card.Acceleration, card.RunningSpeed, card.ReactionSpeed, card.Agility, card.Stamina, card.Strength,
		card.Jumping, card.Balance, card.Technique, card.Dribbling, card.BallControl, card.WeakFoot, card.SkillMoves, card.Finesse,
		card.Curve, card.Volleys, card.ShortPassing, card.LongPassing, card.ForwardPass, card.Offense, card.FinishingAbility, card.ShotPower,
		card.Accuracy, card.Distance, card.Penalty, card.FreeKicks, card.Corners, card.HeadingAccuracy, card.Defence, card.OffsideTrap,
		card.Sliding, card.Tackles, card.BallFocus, card.Interceptions, card.Vigilance, card.Goalkeeping, card.Reflexes, card.Diving,
		card.Handling, card.Sweeping, card.Throwing,
	)
	if err != nil {
		// TODO: add defer for Rollback()
		err = tx.Rollback()
		if err != nil {
			return ErrCard.Wrap(err)
		}
		return ErrCard.Wrap(err)
	}

	err = createCardsAccessories(ctx, cardsDB, card)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return ErrCard.Wrap(err)
		}
		return ErrCard.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return ErrCard.Wrap(err)
	}

	return nil
}

// createCardsAccessories add cards - accessories relation in the database.
func createCardsAccessories(ctx context.Context, cardsDB *cardsDB, card cards.Card) error {
	query :=
		`INSERT INTO
            cards_accessories (card_id, accessory_id) 
        VALUES
        `
	query, values := buildStringForManyRecordsValue(query, card.ID, card.Accessories)
	if _, err := cardsDB.conn.ExecContext(ctx, query, values...); err != nil {
		return ErrCard.Wrap(err)
	}

	return nil
}

// buildStringForManyRecordsValue build string for many records value.
func buildStringForManyRecordsValue(query string, cardID uuid.UUID, accessories []int) (string, []interface{}) {
	values := []interface{}{}
	countAccessories := len(accessories)
	for i, accessory := range accessories {
		values = append(values, cardID, accessory)

		n := i * countAccessories
		query += `(`
		for j := 0; j < countAccessories; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}

	return query[:len(query)-1], values
}

// Get returns card by id from the data base.
func (cardsDB *cardsDB) Get(ctx context.Context, id uuid.UUID) (cards.Card, error) {
	card := cards.Card{}
	query :=
		`SELECT 
            ` + allFields + `
        FROM 
            cards
        WHERE 
            id = $1
        `
	err := cardsDB.conn.QueryRowContext(ctx, query, id).Scan(
		&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle,
		&card.HairColor, &card.DominantFoot, &card.IsTattoos, &card.UserID, &card.Tactics, &card.Positioning, &card.Composure, &card.Aggression,
		&card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed, &card.ReactionSpeed,
		&card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling, &card.BallControl,
		&card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing, &card.ForwardPass,
		&card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks, &card.Corners,
		&card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions,
		&card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return card, cards.ErrNoCard.Wrap(err)
	case err != nil:
		return card, ErrCard.Wrap(err)
	default:
		accessoryIds, err := listAccessoryIdsByCardID(ctx, cardsDB, id)
		if err != nil {
			return card, ErrCard.Wrap(err)
		}
		card.Accessories = accessoryIds
		return card, nil
	}
}

// listAccessoryIdsByCardID returns all accessories for card by id from the database.
func listAccessoryIdsByCardID(ctx context.Context, cardsDB *cardsDB, cardID uuid.UUID) ([]int, error) {
	query :=
		`SELECT 
            accessory_id
        FROM 
            cards_accessories
        WHERE
            card_id = $1
        `
	rows, err := cardsDB.conn.QueryContext(ctx, query, cardID)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []int
	for rows.Next() {
		var cardID int
		if err = rows.Scan(&cardID); err != nil {
			return nil, ErrCard.Wrap(err)
		}
		data = append(data, cardID)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrCard.Wrap(err)
	}

	return data, nil
}

// List returns all cards from the data base.
func (cardsDB *cardsDB) List(ctx context.Context) ([]cards.Card, error) {
	query :=
		`SELECT 
            ` + allFields + ` 
        FROM 
            cards
        `

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
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle,
			&card.HairColor, &card.DominantFoot, &card.IsTattoos, &card.UserID, &card.Tactics, &card.Positioning, &card.Composure, &card.Aggression,
			&card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed, &card.ReactionSpeed,
			&card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling, &card.BallControl,
			&card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing, &card.ForwardPass,
			&card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks, &card.Corners,
			&card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions,
			&card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
		); err != nil {
			return nil, cards.ErrNoCard.Wrap(err)
		}

		accessoryIds, err := listAccessoryIdsByCardID(ctx, cardsDB, card.ID)
		if err != nil {
			return nil, ErrCard.Wrap(err)
		}
		card.Accessories = accessoryIds

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrCard.Wrap(err)
	}

	return data, nil
}

// ListWithFilters returns all cards from DB, taking the necessary filters.
func (cardsDB *cardsDB) ListWithFilters(ctx context.Context, filters []cards.Filters) ([]cards.Card, error) {
	whereClause, valuesString := BuildWhereClauseDependsOnCardsFilters(filters)
	valuesInterface := ValidDBParameters(valuesString)
	query := fmt.Sprintf("SELECT %s FROM cards %s", allFields, whereClause)

	rows, err := cardsDB.conn.QueryContext(ctx, query, valuesInterface...)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	data := []cards.Card{}
	for rows.Next() {
		card := cards.Card{}
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle,
			&card.HairColor, &card.DominantFoot, &card.UserID, &card.Tactics, &card.Positioning, &card.Composure, &card.Aggression,
			&card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed, &card.ReactionSpeed,
			&card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling, &card.BallControl,
			&card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing, &card.ForwardPass,
			&card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks, &card.Corners,
			&card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions,
			&card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
		); err != nil {
			return nil, cards.ErrNoCard.Wrap(err)
		}

		accessoryIds, err := listAccessoryIdsByCardID(ctx, cardsDB, card.ID)
		if err != nil {
			return nil, ErrCard.Wrap(err)
		}
		card.Accessories = accessoryIds

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrCard.Wrap(err)
	}

	return data, nil
}

// ValidDBParameters build valid parameter with string to sinterface.
func ValidDBParameters(stringSlice []string) []interface{} {
	interfaceSlice := make([]interface{}, 0, len(stringSlice))
	for _, v := range stringSlice {
		interfaceSlice = append(interfaceSlice, v)
	}
	return interfaceSlice
}

// BuildWhereClauseDependsOnCardsFilters build string for WHERE.
func BuildWhereClauseDependsOnCardsFilters(filters []cards.Filters) (string, []string) {
	var query string
	var values []string
	var valuesAND []string
	var valuesOR []string
	var whereAND []string
	var whereOR []string

	for _, v := range filters {
		if _, found := v[cards.Tactics]; found == true {
			valuesAND = append(valuesAND, v[cards.Tactics])
			whereAND = append(whereAND, fmt.Sprintf(`%s %s %s`, cards.Tactics, EQ, "$"+strconv.Itoa(len(valuesAND))))
		}

		if _, found := v[cards.MinPhysique]; found == true {
			valuesAND = append(valuesAND, v[cards.MinPhysique])
			whereAND = append(whereAND, fmt.Sprintf(`%s %s %s`, cards.Physique, GTE, "$"+strconv.Itoa(len(valuesAND))))
		}

		if _, found := v[cards.MaxPhysique]; found == true {
			valuesAND = append(valuesAND, v[cards.MaxPhysique])
			whereAND = append(whereAND, fmt.Sprintf(`%s %s %s`, cards.Physique, LTE, "$"+strconv.Itoa(len(valuesAND))))
		}
	}
	if len(whereAND) > 0 {
		query = (" WHERE " + strings.Join(whereAND, " AND "))
		values = append(values, valuesAND...)
	}

	for _, v := range filters {
		if _, found := v[cards.PlayerName]; found == true {
			valuesOR = append(valuesOR, v[cards.PlayerName])
			whereOR = append(whereOR, fmt.Sprintf(`%s %s %s`, cards.PlayerName, LIKE, "$"+strconv.Itoa(len(valuesAND)+len(valuesOR))))
			valuesOR = append(valuesOR, v[cards.PlayerName]+" %")
			whereOR = append(whereOR, fmt.Sprintf(`%s %s %s`, cards.PlayerName, LIKE, "$"+strconv.Itoa(len(valuesAND)+len(valuesOR))))
			valuesOR = append(valuesOR, "% "+v[cards.PlayerName])
			whereOR = append(whereOR, fmt.Sprintf(`%s %s %s`, cards.PlayerName, LIKE, "$"+strconv.Itoa(len(valuesAND)+len(valuesOR))))
			valuesOR = append(valuesOR, "% "+v[cards.PlayerName]+" %")
			whereOR = append(whereOR, fmt.Sprintf(`%s %s %s`, cards.PlayerName, LIKE, "$"+strconv.Itoa(len(valuesAND)+len(valuesOR))))
		}
	}
	if len(whereAND) > 0 && len(whereOR) > 0 {
		query += (" AND (" + strings.Join(whereOR, " OR ") + ")")
		values = append(values, valuesOR...)
	} else if len(whereOR) > 0 {
		query += (" WHERE (" + strings.Join(whereOR, " OR ") + ")")
		values = append(values, valuesOR...)
	}

	return query, values
}

// Delete deletes record card in the data base.
func (cardsDB *cardsDB) Delete(ctx context.Context, id uuid.UUID) error {
	query :=
		`DELETE FROM
            cards
        WHERE 
            id = $1
        `
	_, err := cardsDB.conn.ExecContext(ctx, query, id)
	if err != nil {
		return ErrCard.Wrap(err)
	}

	return nil
}
