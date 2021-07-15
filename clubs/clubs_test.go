// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/users"
)

func TestTeam(t *testing.T) {
	user1 := users.User{
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
		PlayerName:       "test name",
		Quality:          "bronze",
		PictureType:      1,
		Height:           178.8,
		Weight:           72.2,
		SkinColor:        1,
		HairStyle:        1,
		HairColor:        1,
		Accessories:      []int{1, 2},
		DominantFoot:     "left",
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
		Offense:          28,
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

	club := clubs.Club{
		UserID:    user1.ID,
		Formation: 1,
		Tactic:    1,
	}

	capitan := card1.ID

	player := clubs.Player{
		UserID:   user1.ID,
		CardID:   card1.ID,
		Position: clubs.CAM,
		Capitan:  capitan,
	}

	players := []clubs.Player{player}

	updatedClub := clubs.Club{
		UserID:    user1.ID,
		Formation: clubs.FourTwoFour,
		Tactic:    clubs.Regular,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryClub := db.Clubs()
		repositoryCard := db.Cards()
		repositoryUser := db.Users()

		t.Run("Create club", func(t *testing.T) {
			err := repositoryUser.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryClub.Create(ctx, club)
			require.NoError(t, err)
		})

		t.Run("Add card to the club", func(t *testing.T) {
			err := repositoryCard.Create(ctx, card1)
			require.NoError(t, err)

			err = repositoryClub.Add(ctx, user1.ID, card1, capitan, clubs.CAM)
			require.NoError(t, err)
		})

		t.Run("Get club", func(t *testing.T) {
			clubFromDB, err := repositoryClub.GetClub(ctx, user1.ID)
			require.NoError(t, err)

			compareClubs(t, clubFromDB, club)
		})

		t.Run("Get cards from club", func(t *testing.T) {
			playersFromDB, err := repositoryClub.ListCards(ctx, user1.ID)
			require.NoError(t, err)

			comparePlayers(t, playersFromDB, players)
		})

		t.Run("Get capitan", func(t *testing.T) {
			id, err := repositoryClub.GetCapitan(ctx, user1.ID)
			require.NoError(t, err)

			assert.Equal(t, id, capitan)
		})

		t.Run("Update capitan", func(t *testing.T) {
			newCapitan := uuid.New()
			err := repositoryClub.UpdateCapitan(ctx, newCapitan, user1.ID)
			require.NoError(t, err)
		})

		t.Run("Update club", func(t *testing.T) {
			err := repositoryClub.Update(ctx, updatedClub)
			require.NoError(t, err)
		})

	})

}

func compareClubs(t *testing.T, clubDB clubs.Club, clubFake clubs.Club) {
	assert.Equal(t, clubDB.UserID, clubFake.UserID)
	assert.Equal(t, clubDB.Formation, clubFake.Formation)
	assert.Equal(t, clubDB.Tactic, clubFake.Tactic)
}

func comparePlayers(t *testing.T, playersDB []clubs.Player, playersFake []clubs.Player) {
	assert.Equal(t, len(playersDB), len(playersFake))

	for i := 0; i < len(playersFake); i++ {
		assert.Equal(t, playersDB[i].UserID, playersFake[i].UserID)
		assert.Equal(t, playersDB[i].CardID, playersFake[i].CardID)
		assert.Equal(t, playersDB[i].Capitan, playersFake[i].Capitan)
		assert.Equal(t, playersDB[i].Position, playersFake[i].Position)
	}
}
