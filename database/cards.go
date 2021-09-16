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
	"ultimatedivision/internal/pagination"
	"ultimatedivision/marketplace"
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
	allFields = `id, player_name, quality, picture_type, height, weight, skin_color, hair_style, hair_color, dominant_foot, is_tattoos, status,
		type, user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed,
		reaction_speed, agility, stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve,
		volleys, short_passing, long_passing, forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, 
		corners, heading_accuracy, defence, offside_trap, sliding, tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, 
		diving, handling, sweeping, throwing
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
			$50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63)
		`
	_, err = cardsDB.conn.ExecContext(ctx, query,
		card.ID, card.PlayerName, card.Quality, card.PictureType, card.Height, card.Weight, card.SkinColor, card.HairStyle, card.HairColor,
		card.DominantFoot, card.IsTattoos, card.Status, card.Type, card.UserID, card.Tactics, card.Positioning, card.Composure, card.Aggression,
		card.Vision, card.Awareness, card.Crosses, card.Physique, card.Acceleration, card.RunningSpeed, card.ReactionSpeed, card.Agility,
		card.Stamina, card.Strength, card.Jumping, card.Balance, card.Technique, card.Dribbling, card.BallControl, card.WeakFoot, card.SkillMoves,
		card.Finesse, card.Curve, card.Volleys, card.ShortPassing, card.LongPassing, card.ForwardPass, card.Offense, card.FinishingAbility,
		card.ShotPower, card.Accuracy, card.Distance, card.Penalty, card.FreeKicks, card.Corners, card.HeadingAccuracy, card.Defence,
		card.OffsideTrap, card.Sliding, card.Tackles, card.BallFocus, card.Interceptions, card.Vigilance, card.Goalkeeping, card.Reflexes,
		card.Diving, card.Handling, card.Sweeping, card.Throwing,
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
		&card.HairColor, &card.DominantFoot, &card.IsTattoos, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
		&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
		&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
		&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
		&card.ForwardPass, &card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks,
		&card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions,
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
			if errors.Is(err, sql.ErrNoRows) {
				return nil, cards.ErrNoCard.Wrap(err)
			}
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
func (cardsDB *cardsDB) List(ctx context.Context, cursor pagination.Cursor) (cards.Page, error) {
	var cardsListPage cards.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query := fmt.Sprintf(`SELECT %s FROM cards LIMIT %d OFFSET %d`, allFields, cursor.Limit, offset)

	rows, err := cardsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	data := []cards.Card{}
	for rows.Next() {
		card := cards.Card{}
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle,
			&card.HairColor, &card.DominantFoot, &card.IsTattoos, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration,
			&card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique,
			&card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing,
			&card.LongPassing, &card.ForwardPass, &card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance,
			&card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles,
			&card.BallFocus, &card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping,
			&card.Throwing,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return cardsListPage, cards.ErrNoCard.Wrap(err)
			}
			return cardsListPage, ErrCard.Wrap(err)
		}

		accessoryIds, err := listAccessoryIdsByCardID(ctx, cardsDB, card.ID)
		if err != nil {
			return cardsListPage, ErrCard.Wrap(err)
		}
		card.Accessories = accessoryIds

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	cardsListPage, err = cardsDB.listPaginated(ctx, cursor, data)
	return cardsListPage, ErrCard.Wrap(err)
}

// ListByUserID returns all users cards from the database.
func (cardsDB *cardsDB) ListByUserID(ctx context.Context, id uuid.UUID) ([]cards.Card, error) {
	query := `SELECT ` + allFields + ` FROM cards WHERE user_id = $1`

	rows, err := cardsDB.conn.QueryContext(ctx, query, id)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var userCards []cards.Card
	for rows.Next() {
		var card cards.Card
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle,
			&card.HairColor, &card.DominantFoot, &card.IsTattoos, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration,
			&card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique,
			&card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing,
			&card.LongPassing, &card.ForwardPass, &card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance,
			&card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles,
			&card.BallFocus, &card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping,
			&card.Throwing,
		); err != nil {
			return nil, cards.ErrNoCard.Wrap(err)
		}

		accessoryIds, err := listAccessoryIdsByCardID(ctx, cardsDB, card.ID)
		if err != nil {
			return nil, ErrCard.Wrap(err)
		}
		card.Accessories = accessoryIds

		userCards = append(userCards, card)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrCard.Wrap(err)
	}

	return userCards, nil
}

// ListWithFilters returns cards from DB, taking the necessary filters.
func (cardsDB *cardsDB) ListWithFilters(ctx context.Context, filters []cards.Filters, cursor pagination.Cursor) (cards.Page, error) {
	var cardsListPage cards.Page
	whereClause, valuesString := BuildWhereClauseDependsOnCardsFilters(filters)
	valuesInterface := ValidDBParameters(valuesString)
	offset := (cursor.Page - 1) * cursor.Limit
	query := fmt.Sprintf(`
        SELECT
            cards.id, player_name, quality, picture_type, height, weight, skin_color, hair_style, hair_color, dominant_foot, is_tattoos, cards.status, cards.type,
            cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
            stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
            forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
            tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing
        FROM
            cards 
        %s
        LIMIT 
            %d
        OFFSET 
            %d
        `, whereClause, cursor.Limit, offset)

	rows, err := cardsDB.conn.QueryContext(ctx, query, valuesInterface...)
	if err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	data := []cards.Card{}
	for rows.Next() {
		card := cards.Card{}
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle,
			&card.HairColor, &card.DominantFoot, &card.IsTattoos, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
			&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
			&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
			&card.ForwardPass, &card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty,
			&card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus,
			&card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return cardsListPage, cards.ErrNoCard.Wrap(err)
			}
			return cardsListPage, ErrCard.Wrap(err)
		}

		accessoryIds, err := listAccessoryIdsByCardID(ctx, cardsDB, card.ID)
		if err != nil {
			return cardsListPage, ErrCard.Wrap(err)
		}
		card.Accessories = accessoryIds

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	cardsListPage, err = cardsDB.listPaginated(ctx, cursor, data)
	return cardsListPage, ErrCard.Wrap(err)
}

// ListCardIDsWithFiltersWhereActiveLot returns card ids where active lots from DB, taking the necessary filters.
func (cardsDB *cardsDB) ListCardIDsWithFiltersWhereActiveLot(ctx context.Context, filters []cards.Filters) ([]uuid.UUID, error) {
	whereClause, valuesString := BuildWhereClauseDependsOnCardsFilters(filters)
	valuesInterface := ValidDBParameters(valuesString)
	valuesInterface = append(valuesInterface, marketplace.StatusActive)
	query := fmt.Sprintf(`
        SELECT
            cards.id
        FROM
            cards 
		LEFT JOIN lots ON cards.id = lots.item_id
		%s 
		AND 
			lots.status = $%d
        `, whereClause, len(valuesInterface))

	rows, err := cardsDB.conn.QueryContext(ctx, query, valuesInterface...)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var cardIDs []uuid.UUID
	for rows.Next() {
		var cardID uuid.UUID
		if err = rows.Scan(&cardID); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, cards.ErrNoCard.Wrap(err)
			}
			return nil, ErrCard.Wrap(err)
		}

		cardIDs = append(cardIDs, cardID)
	}

	return cardIDs, ErrCard.Wrap(err)
}

// ListByPlayerName returns all cards from DB by player name.
func (cardsDB *cardsDB) ListByPlayerName(ctx context.Context, filter cards.Filters, cursor pagination.Cursor) (cards.Page, error) {
	var cardsListPage cards.Page
	whereClause, valuesString := BuildWhereClauseDependsOnPlayerNameCards(filter)
	valuesInterface := ValidDBParameters(valuesString)
	offset := (cursor.Page - 1) * cursor.Limit
	query := fmt.Sprintf(`SELECT %s FROM cards %s LIMIT %d OFFSET %d`, allFields, whereClause, cursor.Limit, offset)

	rows, err := cardsDB.conn.QueryContext(ctx, query, valuesInterface...)
	if err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []cards.Card
	for rows.Next() {
		card := cards.Card{}
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle,
			&card.HairColor, &card.DominantFoot, &card.IsTattoos, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
			&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
			&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
			&card.ForwardPass, &card.Offense, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty,
			&card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus,
			&card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return cardsListPage, cards.ErrNoCard.Wrap(err)
			}
			return cardsListPage, ErrCard.Wrap(err)
		}

		accessoryIds, err := listAccessoryIdsByCardID(ctx, cardsDB, card.ID)
		if err != nil {
			return cardsListPage, ErrCard.Wrap(err)
		}
		card.Accessories = accessoryIds

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	cardsListPage, err = cardsDB.listPaginated(ctx, cursor, data)
	return cardsListPage, ErrCard.Wrap(err)
}

// ListCardIDsByPlayerNameWhereActiveLot returns card ids where active lot from DB by player name.
func (cardsDB *cardsDB) ListCardIDsByPlayerNameWhereActiveLot(ctx context.Context, filter cards.Filters) ([]uuid.UUID, error) {
	whereClause, valuesString := BuildWhereClauseDependsOnPlayerNameCards(filter)
	valuesInterface := ValidDBParameters(valuesString)
	valuesInterface = append(valuesInterface, marketplace.StatusActive)
	query := fmt.Sprintf(`
        SELECT
            cards.id
        FROM
            cards 
		LEFT JOIN lots ON cards.id = lots.item_id
		%s 
		AND 
			lots.status = $%d
        `, whereClause, len(valuesInterface))

	rows, err := cardsDB.conn.QueryContext(ctx, query, valuesInterface...)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var cardIDs []uuid.UUID
	for rows.Next() {
		var cardID uuid.UUID
		if err = rows.Scan(&cardID); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, cards.ErrNoCard.Wrap(err)
			}
			return nil, ErrCard.Wrap(err)
		}
		cardIDs = append(cardIDs, cardID)
	}
	return cardIDs, ErrCard.Wrap(err)
}

// listPaginated returns paginated list of cards.
func (cardsDB *cardsDB) listPaginated(ctx context.Context, cursor pagination.Cursor, cardsList []cards.Card) (cards.Page, error) {
	var cardsListPage cards.Page
	offset := (cursor.Page - 1) * cursor.Limit

	totalCount, err := cardsDB.totalCount(ctx)
	if err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	pageCount := totalCount / cursor.Limit
	if totalCount%cursor.Limit != 0 {
		pageCount++
	}

	cardsListPage = cards.Page{
		Cards: cardsList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalCount,
		},
	}

	return cardsListPage, nil
}

// totalCount counts all the cards in the table.
func (cardsDB *cardsDB) totalCount(ctx context.Context) (int, error) {
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM cards`)
	err := cardsDB.conn.QueryRowContext(ctx, query).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, cards.ErrNoCard.Wrap(err)
	}
	return count, ErrCard.Wrap(err)
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
	var where []string
	var leftJoin string

	for _, filter := range filters {
		if filter.Name != cards.FilterPrice {
			values = append(values, filter.Value)
			where = append(where, fmt.Sprintf(`cards.%s %s %s`, filter.Name, filter.SearchOperator, "$"+strconv.Itoa(len(values))))
			continue
		}

		for _, v := range filters {
			if v.Name == cards.FilterType && v.Value == string(cards.TypeBought) {
				leftJoin = " LEFT JOIN lots ON cards.id = lots.item_id "
				values = append(values, filter.Value)
				where = append(where, fmt.Sprintf(`
					CASE WHEN
						lots.current_price = 0
					THEN
						lots.start_price
					ELSE
						lots.current_price
					END
						%s %s`,
					filter.SearchOperator,
					"$"+strconv.Itoa(len(values))))
			}
		}
	}

	if leftJoin != "" {
		query += leftJoin
	}
	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}
	return query, values
}

// BuildWhereClauseDependsOnPlayerNameCards build WHERE string for player name.
func BuildWhereClauseDependsOnPlayerNameCards(filter cards.Filters) (string, []string) {
	var query string
	var values []string
	var where []string

	values = append(values, filter.Value)
	where = append(where, fmt.Sprintf(`%s %s %s`, filter.Name, filter.SearchOperator, "$"+strconv.Itoa(len(values))))
	values = append(values, filter.Value+" %")
	where = append(where, fmt.Sprintf(`%s %s %s`, filter.Name, filter.SearchOperator, "$"+strconv.Itoa(len(values))))
	values = append(values, "% "+filter.Value)
	where = append(where, fmt.Sprintf(`%s %s %s`, filter.Name, filter.SearchOperator, "$"+strconv.Itoa(len(values))))
	values = append(values, "% "+filter.Value+" %")
	where = append(where, fmt.Sprintf(`%s %s %s`, filter.Name, filter.SearchOperator, "$"+strconv.Itoa(len(values))))

	query = (" WHERE " + strings.Join(where, " OR "))
	return query, values
}

// UpdateStatus updates status card in the database.
func (cardsDB *cardsDB) UpdateStatus(ctx context.Context, id uuid.UUID, status cards.Status) error {
	_, err := cardsDB.conn.ExecContext(ctx, "UPDATE cards SET status=$1 WHERE id=$2", status, id)
	return ErrCard.Wrap(err)
}

// UpdateUserID updates user id card in the database.
func (cardsDB *cardsDB) UpdateUserID(ctx context.Context, id, userID uuid.UUID) error {
	_, err := cardsDB.conn.ExecContext(ctx, "UPDATE cards SET user_id=$1 WHERE id=$2", userID, id)
	return ErrCard.Wrap(err)
}

// Delete deletes record card in the database.
func (cardsDB *cardsDB) Delete(ctx context.Context, id uuid.UUID) error {
	query :=
		`DELETE FROM
            cards
        WHERE 
            id = $1`

	_, err := cardsDB.conn.ExecContext(ctx, query, id)
	return ErrCard.Wrap(err)
}
