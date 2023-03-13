// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards_test

import (
	"context"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/database"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/divisions"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/pkg/sqlsearchoperators"
	"ultimatedivision/users"
)

func TestCards(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		LastLogin:    time.Now(),
		Status:       0,
		CreatedAt:    time.Now(),
	}

	user2 := users.User{
		ID:           uuid.New(),
		Email:        "3560876@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "qwerty",
		FirstName:    "Stas",
		LastName:     "Isakov",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	card1 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro yak muk",
		Quality:          "wood",
		Height:           178.8,
		Weight:           72.2,
		DominantFoot:     "left",
		IsTattoo:         false,
		Status:           cards.StatusActive,
		Type:             cards.TypeWon,
		UserID:           user1.ID,
		Tactics:          1,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offence:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
		IsMinted:         0,
	}

	card2 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Vova",
		Quality:          "gold",
		Height:           179.9,
		Weight:           73.3,
		DominantFoot:     "right",
		IsTattoo:         true,
		Status:           cards.StatusSale,
		Type:             cards.TypeUnordered,
		UserID:           uuid.New(),
		Tactics:          2,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offence:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
		IsMinted:         0,
	}

	division1 := divisions.Division{
		ID:             uuid.New(),
		Name:           10,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}

	testClub := clubs.Club{
		ID:         uuid.New(),
		OwnerID:    user1.ID,
		Name:       "",
		DivisionID: division1.ID,
		CreatedAt:  time.Now(),
	}

	testSquad := clubs.Squad{
		ID:        uuid.New(),
		Name:      "",
		ClubID:    testClub.ID,
		Formation: 1,
		Tactic:    1,
		CaptainID: uuid.Nil,
	}

	testSquadCard := clubs.SquadCard{
		SquadID:  testSquad.ID,
		CardID:   card1.ID,
		Position: 5,
	}

	filter1 := cards.Filters{
		Name:           cards.FilterTactics,
		Value:          "1",
		SearchOperator: sqlsearchoperators.GTE,
	}

	filter2 := cards.Filters{
		Name:           cards.FilterType,
		Value:          string(cards.TypeWon),
		SearchOperator: sqlsearchoperators.EQ,
	}

	filter2a := cards.Filters{
		Name:           cards.FilterQuality,
		Value:          string(cards.QualityGold),
		SearchOperator: sqlsearchoperators.EQ,
	}

	filter2b := cards.Filters{
		Name:           cards.FilterQuality,
		Value:          string(cards.QualityWood),
		SearchOperator: sqlsearchoperators.EQ,
	}

	filter3 := cards.Filters{
		Name:           cards.FilterPlayerName,
		Value:          "yak",
		SearchOperator: sqlsearchoperators.LIKE,
	}

	cursor1 := pagination.Cursor{
		Limit: 2,
		Page:  1,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryCards := db.Cards()
		repositoryUsers := db.Users()
		repositoryClubs := db.Clubs()
		repositoryDivisions := db.Divisions()

		id := uuid.New()

		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryCards.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, cards.ErrNoCard.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			card1.UserID = user1.ID
			err = repositoryCards.Create(ctx, card1)
			require.NoError(t, err)

			cardFromDB, err := repositoryCards.Get(ctx, card1.ID)
			require.NoError(t, err)
			compareCards(t, card1, cardFromDB)
		})

		t.Run("get by player name", func(t *testing.T) {
			cardFromDB, err := repositoryCards.GetByPlayerName(ctx, card1.PlayerName)
			require.NoError(t, err)
			compareCards(t, card1, cardFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			card2.UserID = user2.ID
			err = repositoryCards.Create(ctx, card2)
			require.NoError(t, err)

			allCards, err := repositoryCards.List(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(allCards.Cards), 2)
			compareCards(t, card1, allCards.Cards[0])
			compareCards(t, card2, allCards.Cards[1])
		})

		t.Run("list by user id", func(t *testing.T) {
			userCard, err := repositoryCards.ListByUserID(ctx, user1.ID, cursor1)
			require.NoError(t, err)

			compareCards(t, userCard.Cards[0], card1)
		})

		t.Run("ListByTypeUnordered", func(t *testing.T) {
			userCard, err := repositoryCards.ListByTypeUnordered(ctx)
			require.NoError(t, err)

			compareCards(t, userCard[0], card2)
		})

		t.Run("list with filters", func(t *testing.T) {
			filters := []cards.Filters{}
			filters = append(filters, filter1, filter2)

			for _, v := range filters {
				err := v.Validate()
				assert.NoError(t, err)
			}

			allCards, err := repositoryCards.ListWithFilters(ctx, filters, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(allCards.Cards), 1)
			compareCards(t, card1, allCards.Cards[0])
		})

		t.Run("list by player name", func(t *testing.T) {
			strings.ToValidUTF8(filter3.Value, "")

			_, err := strconv.Atoi(filter3.Value)
			if err == nil {
				assert.NoError(t, err)
			}

			allCards, err := repositoryCards.ListByUserIDAndPlayerName(ctx, card1.UserID, filter3, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(allCards.Cards), 1)
			compareCards(t, card1, allCards.Cards[0])
		})

		t.Run("build where string", func(t *testing.T) {
			filters := []cards.Filters{}
			filters = append(filters, filter1, filter2, filter2a, filter2b)

			for _, v := range filters {
				err := v.Validate()
				assert.NoError(t, err)
			}

			queryString, values := database.BuildWhereClauseDependsOnCardsFilters(filters)

			assert.Equal(t, queryString, ` WHERE (cards.quality = $1 OR cards.quality = $2) AND cards.tactics >= $3 AND cards.type = $4`)
			assert.Equal(t, values, []string{"gold", "wood", "1", "won"})
		})

		t.Run("build where string for player name", func(t *testing.T) {

			strings.ToValidUTF8(filter3.Value, "")

			_, err := strconv.Atoi(filter3.Value)
			if err == nil {
				assert.NoError(t, err)
			}

			queryString, values := database.BuildWhereClauseDependsOnPlayerNameCards(filter3)

			assert.Equal(t, queryString, ` WHERE player_name LIKE $1 OR player_name LIKE $2 OR player_name LIKE $3 OR player_name LIKE $4`)
			assert.Equal(t, values, []string{"yak", "yak %", "% yak", "% yak %"})
		})

		t.Run("update status sql no rows", func(t *testing.T) {
			err := repositoryCards.UpdateStatus(ctx, uuid.New(), cards.StatusActive)
			require.Error(t, err)
			require.Equal(t, cards.ErrNoCard.Has(err), true)
		})

		t.Run("update status", func(t *testing.T) {
			card1.Status = cards.StatusActive
			err := repositoryCards.UpdateStatus(ctx, card1.ID, card1.Status)
			require.NoError(t, err)

			allCards, err := repositoryCards.List(ctx, cursor1)
			require.NoError(t, err)
			require.Equal(t, len(allCards.Cards), 2)
			compareCards(t, card1, allCards.Cards[1])
			compareCards(t, card2, allCards.Cards[0])
		})

		t.Run("update mint status sql no rows", func(t *testing.T) {
			err := repositoryCards.UpdateMintedStatus(ctx, uuid.New(), cards.Minted)
			require.Error(t, err)
			require.Equal(t, cards.ErrNoCard.Has(err), true)
		})

		t.Run("update mint status", func(t *testing.T) {
			card1.IsMinted = cards.Minted
			err := repositoryCards.UpdateMintedStatus(ctx, card1.ID, cards.Minted)
			require.NoError(t, err)

			allCards, err := repositoryCards.List(ctx, cursor1)
			require.NoError(t, err)
			require.Equal(t, len(allCards.Cards), 2)
			compareCards(t, card1, allCards.Cards[1])
			compareCards(t, card2, allCards.Cards[0])
		})

		t.Run("UpdateType", func(t *testing.T) {
			card1.Type = cards.TypeOrdered
			err := repositoryCards.UpdateType(ctx, card1.ID, card1.Type)
			require.NoError(t, err)

			allCards, err := repositoryCards.List(ctx, cursor1)
			require.NoError(t, err)
			require.Equal(t, len(allCards.Cards), 2)
			compareCards(t, card1, allCards.Cards[1])
			compareCards(t, card2, allCards.Cards[0])
		})

		t.Run("update user id sql no rows", func(t *testing.T) {
			err := repositoryCards.UpdateUserID(ctx, uuid.New(), uuid.New())
			require.Error(t, err)
			require.Equal(t, cards.ErrNoCard.Has(err), true)
		})

		t.Run("update user id", func(t *testing.T) {
			card1.UserID = user2.ID
			err := repositoryCards.UpdateUserID(ctx, card1.ID, user2.ID)
			require.NoError(t, err)

			card, err := repositoryCards.Get(ctx, card1.ID)
			assert.NoError(t, err)
			compareCards(t, card1, card)
		})

		t.Run("get cards from squad", func(t *testing.T) {
			err := repositoryDivisions.Create(ctx, division1)
			require.NoError(t, err)

			_, err = repositoryClubs.Create(ctx, testClub)
			require.NoError(t, err)

			_, err = repositoryClubs.CreateSquad(ctx, testSquad)
			require.NoError(t, err)

			err = repositoryClubs.AddSquadCard(ctx, testSquadCard)
			require.NoError(t, err)

			card, err := repositoryCards.GetSquadCards(ctx, testSquad.ID)
			require.NoError(t, err)
			compareCards(t, card[0], card1)
		})

		t.Run("delete sql no rows", func(t *testing.T) {
			err := repositoryCards.Delete(ctx, uuid.New())
			require.Error(t, err)
			require.Equal(t, cards.ErrNoCard.Has(err), true)
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryCards.Delete(ctx, card1.ID)
			require.NoError(t, err)

			allCards, err := repositoryCards.List(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(allCards.Cards), 1)
			compareCards(t, card2, allCards.Cards[0])
		})
	})
}

func compareCards(t *testing.T, expected, actual cards.Card) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.PlayerName, actual.PlayerName)
	assert.Equal(t, expected.Quality, actual.Quality)
	assert.Equal(t, expected.Height, actual.Height)
	assert.Equal(t, expected.Weight, actual.Weight)
	assert.Equal(t, expected.DominantFoot, actual.DominantFoot)
	assert.Equal(t, expected.IsTattoo, actual.IsTattoo)
	assert.Equal(t, expected.Status, actual.Status)
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.Positioning, actual.Positioning)
	assert.Equal(t, expected.Composure, actual.Composure)
	assert.Equal(t, expected.Aggression, actual.Aggression)
	assert.Equal(t, expected.Vision, actual.Vision)
	assert.Equal(t, expected.Awareness, actual.Awareness)
	assert.Equal(t, expected.Crosses, actual.Crosses)
	assert.Equal(t, expected.Acceleration, actual.Acceleration)
	assert.Equal(t, expected.RunningSpeed, actual.RunningSpeed)
	assert.Equal(t, expected.ReactionSpeed, actual.ReactionSpeed)
	assert.Equal(t, expected.Agility, actual.Agility)
	assert.Equal(t, expected.Stamina, actual.Stamina)
	assert.Equal(t, expected.Strength, actual.Strength)
	assert.Equal(t, expected.Jumping, actual.Jumping)
	assert.Equal(t, expected.Balance, actual.Balance)
	assert.Equal(t, expected.Dribbling, actual.Dribbling)
	assert.Equal(t, expected.BallControl, actual.BallControl)
	assert.Equal(t, expected.WeakFoot, actual.WeakFoot)
	assert.Equal(t, expected.SkillMoves, actual.SkillMoves)
	assert.Equal(t, expected.Finesse, actual.Finesse)
	assert.Equal(t, expected.Curve, actual.Curve)
	assert.Equal(t, expected.Volleys, actual.Volleys)
	assert.Equal(t, expected.ShortPassing, actual.ShortPassing)
	assert.Equal(t, expected.LongPassing, actual.LongPassing)
	assert.Equal(t, expected.ForwardPass, actual.ForwardPass)
	assert.Equal(t, expected.FinishingAbility, actual.FinishingAbility)
	assert.Equal(t, expected.ShotPower, actual.ShotPower)
	assert.Equal(t, expected.Accuracy, actual.Accuracy)
	assert.Equal(t, expected.Distance, actual.Distance)
	assert.Equal(t, expected.Penalty, actual.Penalty)
	assert.Equal(t, expected.FreeKicks, actual.FreeKicks)
	assert.Equal(t, expected.Corners, actual.Corners)
	assert.Equal(t, expected.HeadingAccuracy, actual.HeadingAccuracy)
	assert.Equal(t, expected.OffsideTrap, actual.OffsideTrap)
	assert.Equal(t, expected.Sliding, actual.Sliding)
	assert.Equal(t, expected.Tackles, actual.Tackles)
	assert.Equal(t, expected.BallFocus, actual.BallFocus)
	assert.Equal(t, expected.Interceptions, actual.Interceptions)
	assert.Equal(t, expected.Vigilance, actual.Vigilance)
	assert.Equal(t, expected.Reflexes, actual.Reflexes)
	assert.Equal(t, expected.Diving, actual.Diving)
	assert.Equal(t, expected.Handling, actual.Handling)
	assert.Equal(t, expected.Sweeping, actual.Sweeping)
	assert.Equal(t, expected.Throwing, actual.Throwing)
	assert.Equal(t, expected.IsMinted, actual.IsMinted)
}
