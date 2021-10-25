// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatars_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/database/dbtesting"
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
		UserID:           user2.ID,
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
	}

	avatar1 := avatars.Avatar{
		CardID:         card1.ID,
		PictureType:    avatars.PictureTypeFirst,
		FaceColor:      1,
		FaceType:       2,
		EyeBrowsType:   1,
		EyeBrowsColor:  2,
		EyeLaserType:   1,
		HairstyleColor: 1,
		HairstyleType:  2,
		Nose:           1,
		Tshirt:         2,
		Beard:          1,
		Lips:           2,
		Tattoo:         1,
	}

	avatar2 := avatars.Avatar{
		CardID:         card2.ID,
		PictureType:    avatars.PictureTypeFirst,
		FaceColor:      1,
		FaceType:       2,
		EyeBrowsType:   1,
		EyeBrowsColor:  2,
		EyeLaserType:   1,
		HairstyleColor: 1,
		HairstyleType:  2,
		Nose:           1,
		Tshirt:         2,
		Beard:          1,
		Lips:           2,
		Tattoo:         1,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryCards := db.Cards()
		repositoryAvatars := db.Avatars()
		repositoryUsers := db.Users()
		id := uuid.New()

		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryAvatars.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, avatars.ErrNoAvatar.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			err = repositoryCards.Create(ctx, card1)
			require.NoError(t, err)

			err = repositoryCards.Create(ctx, card2)
			require.NoError(t, err)

			err = repositoryAvatars.Create(ctx, avatar1)
			require.NoError(t, err)

			err = repositoryAvatars.Create(ctx, avatar2)
			require.NoError(t, err)

			avatarDB1, err := repositoryAvatars.Get(ctx, avatar1.CardID)
			assert.NoError(t, err)

			avatarDB2, err := repositoryAvatars.Get(ctx, avatar2.CardID)
			assert.NoError(t, err)

			compareAvatar(t, avatar1, avatarDB1)
			compareAvatar(t, avatar2, avatarDB2)
		})

	})
}

func compareAvatar(t *testing.T, avatar1, avatar2 avatars.Avatar) {
	assert.Equal(t, avatar1.CardID, avatar2.CardID)
	assert.Equal(t, avatar1.PictureType, avatar2.PictureType)
	assert.Equal(t, avatar1.FaceColor, avatar2.FaceColor)
	assert.Equal(t, avatar1.FaceType, avatar2.FaceType)
	assert.Equal(t, avatar1.EyeBrowsType, avatar2.EyeBrowsType)
	assert.Equal(t, avatar1.EyeBrowsColor, avatar2.EyeBrowsColor)
	assert.Equal(t, avatar1.EyeLaserType, avatar2.EyeLaserType)
	assert.Equal(t, avatar1.HairstyleColor, avatar2.HairstyleColor)
	assert.Equal(t, avatar1.HairstyleType, avatar2.HairstyleType)
	assert.Equal(t, avatar1.Nose, avatar2.Nose)
	assert.Equal(t, avatar1.Tshirt, avatar2.Tshirt)
	assert.Equal(t, avatar1.Beard, avatar2.Beard)
	assert.Equal(t, avatar1.Lips, avatar2.Lips)
	assert.Equal(t, avatar1.Tattoo, avatar2.Tattoo)
}
