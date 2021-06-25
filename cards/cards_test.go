// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/cards"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/users"
)

func TestCards(t *testing.T) {

	card1 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro",
		Quality:          "bronze",
		PictureType:      "test1",
		Height:           178.8,
		Weight:           72.2,
		SkinColor:        1,
		HairStyle:        1,
		HairColor:        1,
		Accessories:      []cards.Accessories{1, 4, 5},
		DominantFoot:     "left",
		UserID:           uuid.New(),
		Positioning:      1,
		Composure:        2,
		Aggression:       3,
		Vision:           4,
		Awareness:        5,
		Crosses:          6,
		Acceleration:     7,
		RunningSpeed:     8,
		ReactionSpeed:    9,
		Agility:          10,
		Stamina:          11,
		Strength:         12,
		Jumping:          13,
		Balance:          14,
		Dribbling:        15,
		BallControl:      16,
		WeakFoot:         17,
		SkillMoves:       18,
		Finesse:          19,
		Curve:            20,
		Volleys:          21,
		ShortPassing:     22,
		LongPassing:      23,
		ForwardPass:      24,
		FinishingAbility: 25,
		ShotPower:        26,
		Accuracy:         27,
		Distance:         28,
		Penalty:          29,
		FreeKicks:        30,
		Corners:          31,
		HeadingAccuracy:  32,
		OffsideTrap:      33,
		Sliding:          34,
		Tackles:          35,
		BallFocus:        36,
		Interceptions:    37,
		Vigilance:        38,
		Reflexes:         39,
		Diving:           40,
		Handling:         41,
		Sweeping:         42,
		Throwing:         43,
	}

	card2 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Vova",
		Quality:          "silver",
		PictureType:      "test2",
		Height:           179.9,
		Weight:           73.3,
		SkinColor:        2,
		HairStyle:        2,
		HairColor:        2,
		Accessories:      []cards.Accessories{2, 4, 5},
		DominantFoot:     "right",
		UserID:           uuid.New(),
		Positioning:      100,
		Composure:        99,
		Aggression:       98,
		Vision:           97,
		Awareness:        96,
		Crosses:          95,
		Acceleration:     94,
		RunningSpeed:     93,
		ReactionSpeed:    92,
		Agility:          91,
		Stamina:          90,
		Strength:         89,
		Jumping:          88,
		Balance:          87,
		Dribbling:        86,
		BallControl:      85,
		WeakFoot:         84,
		SkillMoves:       83,
		Finesse:          82,
		Curve:            81,
		Volleys:          80,
		ShortPassing:     79,
		LongPassing:      78,
		ForwardPass:      77,
		FinishingAbility: 76,
		ShotPower:        75,
		Accuracy:         74,
		Distance:         73,
		Penalty:          72,
		FreeKicks:        71,
		Corners:          70,
		HeadingAccuracy:  69,
		OffsideTrap:      68,
		Sliding:          67,
		Tackles:          66,
		BallFocus:        65,
		Interceptions:    64,
		Vigilance:        63,
		Reflexes:         62,
		Diving:           61,
		Handling:         60,
		Sweeping:         59,
		Throwing:         58,
	}

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

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryCards := db.Cards()
		repositoryUsers := db.Users()
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

		t.Run("list", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			card2.UserID = user2.ID
			err = repositoryCards.Create(ctx, card2)
			require.NoError(t, err)

			allCards, err := repositoryCards.List(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(allCards), 2)
			compareCards(t, card1, allCards[0])
			compareCards(t, card2, allCards[1])
		})
	})
}

func compareCards(t *testing.T, card1, card2 cards.Card) {
	assert.Equal(t, card1.ID, card2.ID)
	assert.Equal(t, card1.PlayerName, card2.PlayerName)
	assert.Equal(t, card1.Quality, card2.Quality)
	assert.Equal(t, card1.Height, card2.Height)
	assert.Equal(t, card1.Weight, card2.Weight)
	assert.Equal(t, card1.SkinColor, card2.SkinColor)
	assert.Equal(t, card1.HairStyle, card2.HairStyle)
	assert.Equal(t, card1.HairColor, card2.HairColor)
	assert.Equal(t, card1.Accessories, card2.Accessories)
	assert.Equal(t, card1.DominantFoot, card2.DominantFoot)
	assert.Equal(t, card1.UserID, card2.UserID)
	assert.Equal(t, card1.Positioning, card2.Positioning)
	assert.Equal(t, card1.Composure, card2.Composure)
	assert.Equal(t, card1.Aggression, card2.Aggression)
	assert.Equal(t, card1.Vision, card2.Vision)
	assert.Equal(t, card1.Awareness, card2.Awareness)
	assert.Equal(t, card1.Crosses, card2.Crosses)
	assert.Equal(t, card1.Acceleration, card2.Acceleration)
	assert.Equal(t, card1.RunningSpeed, card2.RunningSpeed)
	assert.Equal(t, card1.ReactionSpeed, card2.ReactionSpeed)
	assert.Equal(t, card1.Agility, card2.Agility)
	assert.Equal(t, card1.Stamina, card2.Stamina)
	assert.Equal(t, card1.Strength, card2.Strength)
	assert.Equal(t, card1.Jumping, card2.Jumping)
	assert.Equal(t, card1.Balance, card2.Balance)
	assert.Equal(t, card1.Dribbling, card2.Dribbling)
	assert.Equal(t, card1.BallControl, card2.BallControl)
	assert.Equal(t, card1.WeakFoot, card2.WeakFoot)
	assert.Equal(t, card1.SkillMoves, card2.SkillMoves)
	assert.Equal(t, card1.Finesse, card2.Finesse)
	assert.Equal(t, card1.Curve, card2.Curve)
	assert.Equal(t, card1.Volleys, card2.Volleys)
	assert.Equal(t, card1.ShortPassing, card2.ShortPassing)
	assert.Equal(t, card1.LongPassing, card2.LongPassing)
	assert.Equal(t, card1.ForwardPass, card2.ForwardPass)
	assert.Equal(t, card1.FinishingAbility, card2.FinishingAbility)
	assert.Equal(t, card1.ShotPower, card2.ShotPower)
	assert.Equal(t, card1.Accuracy, card2.Accuracy)
	assert.Equal(t, card1.Distance, card2.Distance)
	assert.Equal(t, card1.Penalty, card2.Penalty)
	assert.Equal(t, card1.FreeKicks, card2.FreeKicks)
	assert.Equal(t, card1.Corners, card2.Corners)
	assert.Equal(t, card1.HeadingAccuracy, card2.HeadingAccuracy)
	assert.Equal(t, card1.OffsideTrap, card2.OffsideTrap)
	assert.Equal(t, card1.Sliding, card2.Sliding)
	assert.Equal(t, card1.Tackles, card2.Tackles)
	assert.Equal(t, card1.BallFocus, card2.BallFocus)
	assert.Equal(t, card1.Interceptions, card2.Interceptions)
	assert.Equal(t, card1.Vigilance, card2.Vigilance)
	assert.Equal(t, card1.Reflexes, card2.Reflexes)
	assert.Equal(t, card1.Diving, card2.Diving)
	assert.Equal(t, card1.Handling, card2.Handling)
	assert.Equal(t, card1.Sweeping, card2.Sweeping)
	assert.Equal(t, card1.Throwing, card2.Throwing)
}
