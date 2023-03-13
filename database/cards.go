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
	_ "github.com/lib/pq" // using postgres driver.
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/pagination"
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
	allFields = `id, player_name, quality, height, weight, dominant_foot, is_tattoo, status, type, user_id, tactics, positioning, composure, 
		aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility, stamina, strength, jumping, 
		balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing, forward_pass, 
		offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, 
		sliding, tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing, is_minted`
)

// Create adds card in the data base.
func (cardsDB *cardsDB) Create(ctx context.Context, card cards.Card) error {
	query :=
		`INSERT INTO
			cards(` + allFields + `) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25,
			$26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49,
			$50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60)`

	_, err := cardsDB.conn.ExecContext(ctx, query,
		card.ID, card.PlayerName, card.Quality, card.Height, card.Weight,
		card.DominantFoot, card.IsTattoo, card.Status, card.Type, card.UserID, card.Tactics, card.Positioning, card.Composure, card.Aggression,
		card.Vision, card.Awareness, card.Crosses, card.Physique, card.Acceleration, card.RunningSpeed, card.ReactionSpeed, card.Agility,
		card.Stamina, card.Strength, card.Jumping, card.Balance, card.Technique, card.Dribbling, card.BallControl, card.WeakFoot, card.SkillMoves,
		card.Finesse, card.Curve, card.Volleys, card.ShortPassing, card.LongPassing, card.ForwardPass, card.Offence, card.FinishingAbility,
		card.ShotPower, card.Accuracy, card.Distance, card.Penalty, card.FreeKicks, card.Corners, card.HeadingAccuracy, card.Defence,
		card.OffsideTrap, card.Sliding, card.Tackles, card.BallFocus, card.Interceptions, card.Vigilance, card.Goalkeeping, card.Reflexes,
		card.Diving, card.Handling, card.Sweeping, card.Throwing, card.IsMinted,
	)

	return ErrCard.Wrap(err)
}

// Get returns card by id from the data base.
func (cardsDB *cardsDB) Get(ctx context.Context, id uuid.UUID) (cards.Card, error) {
	card := cards.Card{}
	query :=
		`SELECT * FROM  
            cards
        WHERE 
            id = $1`

	err := cardsDB.conn.QueryRowContext(ctx, query, id).Scan(
		&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight, &card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
		&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
		&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
		&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
		&card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks,
		&card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions,
		&card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing, &card.IsMinted,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return card, cards.ErrNoCard.Wrap(err)
	}

	return card, ErrCard.Wrap(err)
}

// GetByPlayerName returns card by player name from DB.
func (cardsDB *cardsDB) GetByPlayerName(ctx context.Context, playerName string) (cards.Card, error) {
	card := cards.Card{}
	query :=
		`SELECT * FROM 
            cards
        WHERE 
            player_name = $1`

	err := cardsDB.conn.QueryRowContext(ctx, query, playerName).Scan(
		&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight, &card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
		&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
		&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
		&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
		&card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks,
		&card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions,
		&card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing, &card.IsMinted,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return card, cards.ErrNoCard.Wrap(err)
	}

	return card, ErrCard.Wrap(err)
}

// List returns all cards from the data base.
func (cardsDB *cardsDB) List(ctx context.Context, cursor pagination.Cursor) (cards.Page, error) {
	var cardsListPage cards.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query :=
		`SELECT * FROM
			cards 
		LIMIT 
			$1
		OFFSET 
			$2`

	rows, err := cardsDB.conn.QueryContext(ctx, query, cursor.Limit, offset)
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
			&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight, &card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration,
			&card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique,
			&card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing,
			&card.LongPassing, &card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance,
			&card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles,
			&card.BallFocus, &card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping,
			&card.Throwing, &card.IsMinted,
		); err != nil {
			return cardsListPage, ErrCard.Wrap(err)
		}

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	totalCount, err := cardsDB.totalCount(ctx)
	if err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	cardsListPage, err = cardsDB.listPaginated(ctx, cursor, data, totalCount)
	return cardsListPage, ErrCard.Wrap(err)
}

// ListByUserID returns all users cards from the database.
func (cardsDB *cardsDB) ListByUserID(ctx context.Context, id uuid.UUID, cursor pagination.Cursor) (cards.Page, error) {
	var userCardsPage cards.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query :=
		`SELECT * FROM  
			cards 
		WHERE 
			user_id = $1
		LIMIT 
			$2
		OFFSET 
			$3`

	rows, err := cardsDB.conn.QueryContext(ctx, query, id, cursor.Limit, offset)
	if err != nil {
		return userCardsPage, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var userCards []cards.Card
	for rows.Next() {
		var card cards.Card
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight, &card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration,
			&card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique,
			&card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing,
			&card.LongPassing, &card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance,
			&card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles,
			&card.BallFocus, &card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping,
			&card.Throwing, &card.IsMinted,
		); err != nil {
			return userCardsPage, ErrCard.Wrap(err)
		}

		userCards = append(userCards, card)
	}
	if err = rows.Err(); err != nil {
		return userCardsPage, ErrCard.Wrap(err)
	}

	totalCount, err := cardsDB.totalCountWithFilters(ctx, "WHERE user_id = $1", []interface{}{id})
	if err != nil {
		return userCardsPage, ErrCard.Wrap(err)
	}

	userCardsPage, err = cardsDB.listPaginated(ctx, cursor, userCards, totalCount)
	return userCardsPage, ErrCard.Wrap(err)
}

// ListByTypeUnordered returns cards where type is unordered from the database.
func (cardsDB *cardsDB) ListByTypeUnordered(ctx context.Context) ([]cards.Card, error) {
	query := `SELECT * FROM cards WHERE type = $1`

	rows, err := cardsDB.conn.QueryContext(ctx, query, cards.TypeUnordered)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	cardsList := []cards.Card{}
	for rows.Next() {
		card := cards.Card{}
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight, &card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration,
			&card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique,
			&card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing,
			&card.LongPassing, &card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance,
			&card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles,
			&card.BallFocus, &card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping,
			&card.Throwing, &card.IsMinted,
		); err != nil {
			return nil, ErrCard.Wrap(err)
		}

		cardsList = append(cardsList, card)
	}

	return cardsList, ErrCard.Wrap(rows.Err())
}

// ListWithFilters returns cards from DB, taking the necessary filters.
func (cardsDB *cardsDB) ListWithFilters(ctx context.Context, filters []cards.Filters, cursor pagination.Cursor) (cards.Page, error) {
	var cardsListPage cards.Page
	whereClause, valuesString := BuildWhereClauseDependsOnCardsFilters(filters)
	valuesInterface := ValidDBParameters(valuesString)
	offset := (cursor.Page - 1) * cursor.Limit
	query := fmt.Sprintf(`
        SELECT
            cards.id, player_name, quality, height, weight, dominant_foot, is_tattoo, cards.status, cards.type,
            cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
            stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
            forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
            tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing, is_minted
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
			&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight,
			&card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
			&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
			&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
			&card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty,
			&card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus,
			&card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
			&card.IsMinted,
		); err != nil {
			return cardsListPage, ErrCard.Wrap(err)
		}

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	totalCount, err := cardsDB.totalCountWithFilters(ctx, whereClause, valuesInterface)
	if err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	cardsListPage, err = cardsDB.listPaginated(ctx, cursor, data, totalCount)
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
			return nil, ErrCard.Wrap(err)
		}

		cardIDs = append(cardIDs, cardID)
	}

	return cardIDs, ErrCard.Wrap(err)
}

// ListByUserIDAndPlayerName returns cards from DB by user id and player name.
func (cardsDB *cardsDB) ListByUserIDAndPlayerName(ctx context.Context, userID uuid.UUID, filter cards.Filters, cursor pagination.Cursor) (cards.Page, error) {
	var cardsListPage cards.Page
	whereClause, valuesString := BuildWhereClauseDependsOnPlayerNameCards(filter)
	whereClause += fmt.Sprintf(" AND user_id = $%d", len(valuesString)+1)
	valuesString = append(valuesString, userID.String())
	valuesInterface := ValidDBParameters(valuesString)
	offset := (cursor.Page - 1) * cursor.Limit
	query := fmt.Sprintf(`SELECT * FROM cards %s LIMIT %d OFFSET %d`, whereClause, cursor.Limit, offset)

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
			&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight,
			&card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
			&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
			&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
			&card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty,
			&card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus,
			&card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
			&card.IsMinted,
		); err != nil {
			return cardsListPage, ErrCard.Wrap(err)
		}

		data = append(data, card)
	}
	if err = rows.Err(); err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	totalCount, err := cardsDB.totalCountWithFilters(ctx, whereClause, valuesInterface)
	if err != nil {
		return cardsListPage, ErrCard.Wrap(err)
	}

	cardsListPage, err = cardsDB.listPaginated(ctx, cursor, data, totalCount)
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
			return nil, ErrCard.Wrap(err)
		}
		cardIDs = append(cardIDs, cardID)
	}
	return cardIDs, ErrCard.Wrap(err)
}

// listPaginated returns paginated list of cards.
func (cardsDB *cardsDB) listPaginated(ctx context.Context, cursor pagination.Cursor, cardsList []cards.Card, totalCount int) (cards.Page, error) {
	var cardsListPage cards.Page
	offset := (cursor.Page - 1) * cursor.Limit
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
	query := fmt.Sprintf("SELECT COUNT(*) FROM cards")
	err := cardsDB.conn.QueryRowContext(ctx, query).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, cards.ErrNoCard.Wrap(err)
	}
	return count, ErrCard.Wrap(err)
}

// totalCountWithFilters counts cards with filtes in the table.
func (cardsDB *cardsDB) totalCountWithFilters(ctx context.Context, whereClause string, valuesInterface []interface{}) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM cards %s", whereClause)
	err := cardsDB.conn.QueryRowContext(ctx, query, valuesInterface...).Scan(&count)
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
	var (
		query        string
		values       []string
		where        []string
		whereOR      []string
		leftJoin     string
		countQuality int
	)

	for _, filter := range filters {
		if filter.Name == cards.FilterQuality {
			countQuality++
		}
	}

	if countQuality > 1 {
		var newFilters []cards.Filters
		for _, filter := range filters {
			if filter.Name == cards.FilterQuality {
				values = append(values, filter.Value)
				whereOR = append(whereOR, fmt.Sprintf(`cards.%s %s %s`, filter.Name, filter.SearchOperator, "$"+strconv.Itoa(len(values))))
				continue
			}
			newFilters = append(newFilters, filter)
		}
		filters = newFilters
	}

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

	if len(whereOR) > 0 {
		query += " WHERE (" + strings.Join(whereOR, " OR ") + ") "
		if len(where) > 0 {
			query += "AND " + strings.Join(where, " AND ")
		}
		return query, values
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
	result, err := cardsDB.conn.ExecContext(ctx, "UPDATE cards SET status=$1 WHERE id=$2", status, id)
	if err != nil {
		return ErrCard.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return cards.ErrNoCard.New("")
	}

	return ErrCard.Wrap(err)
}

// UpdateMintedStatus updates status card in the database.
func (cardsDB *cardsDB) UpdateMintedStatus(ctx context.Context, id uuid.UUID, status int) error {
	result, err := cardsDB.conn.ExecContext(ctx, "UPDATE cards SET is_minted=$1 WHERE id=$2", status, id)
	if err != nil {
		return ErrCard.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return cards.ErrNoCard.New("")
	}

	return ErrCard.Wrap(err)
}

// UpdateType updates type of card in the database.
func (cardsDB *cardsDB) UpdateType(ctx context.Context, id uuid.UUID, typeCard cards.Type) error {
	result, err := cardsDB.conn.ExecContext(ctx, "UPDATE cards SET type=$1 WHERE id=$2", typeCard, id)
	if err != nil {
		return ErrCard.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return cards.ErrNoCard.New("")
	}

	return ErrCard.Wrap(err)
}

// UpdateUserID updates user id card in the database.
func (cardsDB *cardsDB) UpdateUserID(ctx context.Context, id, userID uuid.UUID) error {
	result, err := cardsDB.conn.ExecContext(ctx, "UPDATE cards SET user_id=$1 WHERE id=$2", userID, id)
	if err != nil {
		return ErrCard.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return cards.ErrNoCard.New("")
	}

	return ErrCard.Wrap(err)
}

// Delete deletes record card in the database.
func (cardsDB *cardsDB) Delete(ctx context.Context, id uuid.UUID) error {
	query :=
		`DELETE FROM
            cards
        WHERE 
            id = $1`

	result, err := cardsDB.conn.ExecContext(ctx, query, id)
	if err != nil {
		return ErrCard.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return cards.ErrNoCard.New("")
	}

	return ErrCard.Wrap(err)
}

// GetSquadCards returns all cards with characteristics from the squad from the database.
func (cardsDB *cardsDB) GetSquadCards(ctx context.Context, id uuid.UUID) ([]cards.Card, error) {
	var cardsFromSquad []cards.Card
	query := `SELECT * FROM 
            cards
        WHERE id IN (SELECT card_id
                     FROM squad_cards
                     WHERE id = $1)
        `

	rows, err := cardsDB.conn.QueryContext(ctx, query, id)
	if err != nil {
		return cardsFromSquad, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		card := cards.Card{}
		if err = rows.Scan(
			&card.ID, &card.PlayerName, &card.Quality, &card.Height, &card.Weight,
			&card.DominantFoot, &card.IsTattoo, &card.Status, &card.Type, &card.UserID, &card.Tactics, &card.Positioning,
			&card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses, &card.Physique, &card.Acceleration, &card.RunningSpeed,
			&card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance, &card.Technique, &card.Dribbling,
			&card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing,
			&card.ForwardPass, &card.Offence, &card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty,
			&card.FreeKicks, &card.Corners, &card.HeadingAccuracy, &card.Defence, &card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus,
			&card.Interceptions, &card.Vigilance, &card.Goalkeeping, &card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
			&card.IsMinted,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return cardsFromSquad, cards.ErrNoCard.Wrap(err)
			}
			return cardsFromSquad, ErrCard.Wrap(err)
		}

		cardsFromSquad = append(cardsFromSquad, card)
	}
	if err = rows.Err(); err != nil {
		return cardsFromSquad, ErrCard.Wrap(err)
	}

	return cardsFromSquad, nil
}
