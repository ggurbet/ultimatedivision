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

	testSquad := clubs.Squad{
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

	testSquadCards := clubs.SquadCard{
		SquadID:  testSquad.ID,
		CardID:   testCard.ID,
		Position: clubs.CAM,
	}

	updatedSquadCards := []clubs.SquadCard{{
		SquadID:  testSquad.ID,
		CardID:   testCard.ID,
		Position: clubs.CAM,
	}}

	updatedSquad := clubs.Squad{
		ID:        testSquad.ID,
		ClubID:    testClub.ID,
		Formation: clubs.FourFourTwo,
		Tactic:    clubs.Attack,
		CaptainID: testCard.ID,
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

		t.Run("Get club", func(t *testing.T) {
			clubDB, err := repositoryClubs.GetByUserID(ctx, testUser.ID)
			require.NoError(t, err)

			compareClubs(t, clubDB, testClub)
		})

		t.Run("Get squad", func(t *testing.T) {
			squadDB, err := repositoryClubs.GetSquad(ctx, testClub.ID)
			require.NoError(t, err)

			compareSquads(t, squadDB, testSquad)
		})

		t.Run("AddSquadCard card to squad", func(t *testing.T) {
			err := repositoryCards.Create(ctx, testCard)
			require.NoError(t, err)

			err = repositoryClubs.AddSquadCard(ctx, testSquadCards)
			require.NoError(t, err)
		})

		t.Run("Get capitan", func(t *testing.T) {
			capitan, err := repositoryClubs.GetCaptainID(ctx, testSquad.ID)
			require.NoError(t, err)

			assert.Equal(t, capitan, uuid.Nil)
		})

		t.Run("List cards from squad", func(t *testing.T) {
			squadCardsDB, err := repositoryClubs.ListSquadCards(ctx, updatedSquad.ID)
			require.NoError(t, err)

			comparePlayers(t, squadCardsDB, updatedSquadCards)
		})

		t.Run("Update tactic and formation in squad", func(t *testing.T) {
			err := repositoryClubs.UpdateTacticFormationCaptain(ctx, updatedSquad)
			require.NoError(t, err)
		})

		t.Run("Update card position in squad", func(t *testing.T) {
			err := repositoryClubs.UpdatePosition(ctx, updatedSquad.ID, testCard.ID, clubs.CM)
			require.NoError(t, err)
		})

		t.Run("Delete card from squad", func(t *testing.T) {
			err := repositoryClubs.DeleteSquadCard(ctx, updatedSquad.ID, testCard.ID)
			require.NoError(t, err)
		})
	})
}

func compareClubs(t *testing.T, clubDB clubs.Club, clubTest clubs.Club) {
	assert.Equal(t, clubDB.ID, clubTest.ID)
	assert.Equal(t, clubDB.OwnerID, clubTest.OwnerID)
	assert.Equal(t, clubDB.Name, clubTest.Name)
	assert.WithinDuration(t, clubDB.CreatedAt, clubTest.CreatedAt, 1*time.Second)
}

func compareSquads(t *testing.T, squadDB clubs.Squad, squadTest clubs.Squad) {
	assert.Equal(t, squadDB.ID, squadTest.ID)
	assert.Equal(t, squadDB.ClubID, squadTest.ClubID)
	assert.Equal(t, squadDB.Tactic, squadTest.Tactic)
	assert.Equal(t, squadDB.Formation, squadTest.Formation)
	assert.Equal(t, squadDB.CaptainID, squadDB.CaptainID)
}

func comparePlayers(t *testing.T, playersDB []clubs.SquadCard, playersTest []clubs.SquadCard) {
	assert.Equal(t, len(playersDB), len(playersTest))

	for i := 0; i < len(playersTest); i++ {
		assert.Equal(t, playersDB[i].SquadID, playersTest[i].SquadID)
		assert.Equal(t, playersDB[i].CardID, playersTest[i].CardID)
		assert.Equal(t, playersDB[i].Position, playersTest[i].Position)
	}
}
