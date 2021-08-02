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
	testUser := users.User{
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

	testClub := clubs.Club{
		ID:        uuid.New(),
		OwnerID:   testUser.ID,
		Name:      testUser.NickName,
		CreatedAt: time.Now().UTC(),
	}

	testSquad := clubs.Squads{
		ID:        uuid.New(),
		Name:      "test squad",
		ClubID:    testClub.ID,
		Tactic:    clubs.Balanced,
		Formation: clubs.FourTwoFour,
	}

	testCard := cards.Card{
		ID:     uuid.New(),
		UserID: testUser.ID,
	}

	testSquadCards := clubs.SquadCards{
		ID:       testSquad.ID,
		CardID:   testCard.ID,
		Position: clubs.CAM,
	}

	updatedSquadCards := []clubs.SquadCards{{
		ID:       testSquad.ID,
		CardID:   testCard.ID,
		Position: clubs.CAM,
		Capitan:  testCard.ID,
	}}

	updatedSquad := clubs.Squads{
		ID:        testSquad.ID,
		ClubID:    testClub.ID,
		Formation: clubs.FourFourTwo,
		Tactic:    clubs.Attack,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryClubs := db.Clubs()
		repositoryCards := db.Cards()
		repositoryUsers := db.Users()

		t.Run("Create club", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, testUser)
			require.NoError(t, err)

			err = repositoryClubs.Create(ctx, testClub)
			require.NoError(t, err)

		})

		t.Run("Create squad", func(t *testing.T) {
			err := repositoryClubs.CreateSquad(ctx, testSquad)
			require.NoError(t, err)
		})

		t.Run("List clubs", func(t *testing.T) {
			clubDB, err := repositoryClubs.List(ctx, testClub.ID)
			require.NoError(t, err)

			compareClubs(t, clubDB, []clubs.Club{testClub})
		})

		t.Run("Get squad", func(t *testing.T) {
			squadDB, err := repositoryClubs.GetSquad(ctx, testClub.ID)
			require.NoError(t, err)

			compareSquads(t, squadDB, testSquad)
		})

		t.Run("Add card to squad", func(t *testing.T) {
			err := repositoryCards.Create(ctx, testCard)
			require.NoError(t, err)

			err = repositoryClubs.Add(ctx, testSquadCards)
			require.NoError(t, err)
		})

		t.Run("Get capitan", func(t *testing.T) {
			capitan, err := repositoryClubs.GetCapitan(ctx, testSquad.ID)
			require.NoError(t, err)

			assert.Equal(t, capitan, uuid.Nil)
		})

		t.Run("Update capitan", func(t *testing.T) {
			err := repositoryClubs.UpdateCapitan(ctx, testCard.ID, testSquad.ID)
			require.NoError(t, err)
		})

		t.Run("List cards from squad", func(t *testing.T) {
			squadCardsDB, err := repositoryClubs.ListSquadCards(ctx, updatedSquad.ID)
			require.NoError(t, err)

			comparePlayers(t, squadCardsDB, updatedSquadCards)
		})

		t.Run("Update tactic and formation in squad", func(t *testing.T) {
			err := repositoryClubs.UpdateTacticFormation(ctx, updatedSquad)
			require.NoError(t, err)
		})

		t.Run("Update card position in squad", func(t *testing.T) {
			err := repositoryClubs.UpdatePosition(ctx, updatedSquad.ID, testCard.ID, clubs.CM)
			require.NoError(t, err)
		})
	})

}

func compareClubs(t *testing.T, clubDB []clubs.Club, clubTest []clubs.Club) {
	assert.Equal(t, len(clubDB), len(clubDB))

	for i := 0; i < len(clubDB); i++ {
		assert.Equal(t, clubDB[i].ID, clubTest[i].ID)
		assert.Equal(t, clubDB[i].Name, clubTest[i].Name)
		assert.Equal(t, clubDB[i].OwnerID, clubTest[i].OwnerID)
		assert.WithinDuration(t, clubDB[i].CreatedAt, clubTest[i].CreatedAt, 1*time.Second)
	}
}

func compareSquads(t *testing.T, squadDB clubs.Squads, squadTest clubs.Squads) {
	assert.Equal(t, squadDB.ID, squadTest.ID)
	assert.Equal(t, squadDB.ClubID, squadTest.ClubID)
	assert.Equal(t, squadDB.Tactic, squadTest.Tactic)
	assert.Equal(t, squadDB.Formation, squadTest.Formation)
}

func comparePlayers(t *testing.T, playersDB []clubs.SquadCards, playersTest []clubs.SquadCards) {
	assert.Equal(t, len(playersDB), len(playersTest))

	for i := 0; i < len(playersTest); i++ {
		assert.Equal(t, playersDB[i].ID, playersTest[i].ID)
		assert.Equal(t, playersDB[i].CardID, playersTest[i].CardID)
		assert.Equal(t, playersDB[i].Capitan, playersTest[i].Capitan)
		assert.Equal(t, playersDB[i].Position, playersTest[i].Position)
	}
}
