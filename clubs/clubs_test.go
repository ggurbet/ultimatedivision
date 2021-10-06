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

	testCard1 := cards.Card{
		ID:     uuid.New(),
		UserID: testUser.ID,
	}

	testCard2 := cards.Card{
		ID:     uuid.New(),
		UserID: testUser.ID,
	}

	testSquadCard1 := clubs.SquadCard{
		SquadID:  testSquad.ID,
		CardID:   testCard1.ID,
		Position: clubs.CCAM,
	}

	testSquadCard2 := clubs.SquadCard{
		SquadID:  testSquad.ID,
		CardID:   testCard2.ID,
		Position: clubs.RB,
	}

	updatedSquadCard1 := clubs.SquadCard{
		SquadID:  testSquad.ID,
		CardID:   testCard1.ID,
		Position: clubs.CST,
	}

	updatedSquadCard2 := clubs.SquadCard{
		SquadID:  testSquad.ID,
		CardID:   testCard2.ID,
		Position: clubs.CCD,
	}

	updatedSquad := clubs.Squad{
		ID:        testSquad.ID,
		ClubID:    testClub.ID,
		Formation: clubs.FourFourTwo,
		Tactic:    clubs.Attack,
		CaptainID: testCard1.ID,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryClubs := db.Clubs()
		repositoryCards := db.Cards()
		repositoryUsers := db.Users()

		t.Run("Create club", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, testUser)
			require.NoError(t, err)

			clubsID, err := repositoryClubs.Create(ctx, testClub)
			require.NoError(t, err)
			assert.Equal(t, clubsID, testClub.ID)
		})

		t.Run("Create squad", func(t *testing.T) {
			squadID, err := repositoryClubs.CreateSquad(ctx, testSquad)
			require.NoError(t, err)
			assert.Equal(t, squadID, testSquad.ID)
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

		t.Run("get formation", func(t *testing.T) {
			formation, err := repositoryClubs.GetFormation(ctx, testSquad.ID)
			require.NoError(t, err)
			assert.Equal(t, formation, testSquad.Formation)
		})

		t.Run("Add cards to the squad", func(t *testing.T) {
			err := repositoryCards.Create(ctx, testCard1)
			require.NoError(t, err)

			err = repositoryCards.Create(ctx, testCard2)
			require.NoError(t, err)

			err = repositoryClubs.AddSquadCard(ctx, testSquadCard1)
			require.NoError(t, err)

			err = repositoryClubs.AddSquadCard(ctx, testSquadCard2)
			require.NoError(t, err)
		})

		t.Run("List cards from squad", func(t *testing.T) {
			squadCardsDB, err := repositoryClubs.ListSquadCards(ctx, testSquad.ID)
			require.NoError(t, err)

			comparePlayers(t, squadCardsDB, []clubs.SquadCard{testSquadCard2, testSquadCard1})
		})

		t.Run("Update tactic and formation in squad", func(t *testing.T) {
			err := repositoryClubs.UpdateTacticFormationCaptain(ctx, updatedSquad)
			require.NoError(t, err)
		})

		t.Run("Update card position in squad", func(t *testing.T) {
			err := repositoryClubs.UpdatePosition(ctx, []clubs.SquadCard{updatedSquadCard1, updatedSquadCard2})
			require.NoError(t, err)
		})

		t.Run("Delete card from squad", func(t *testing.T) {
			err := repositoryClubs.DeleteSquadCard(ctx, testSquad.ID, testCard1.ID)
			require.NoError(t, err)
			err = repositoryClubs.DeleteSquadCard(ctx, testSquad.ID, testCard2.ID)
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
