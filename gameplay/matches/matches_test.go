// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches_test

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
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/users"
)

func TestMatches(t *testing.T) {
	testUser1 := users.User{
		ID:           uuid.New(),
		Email:        "test@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "testNickName",
		FirstName:    "test",
		LastName:     "test",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	testUser2 := users.User{
		ID:           uuid.New(),
		Email:        "test@gmail.com",
		PasswordHash: []byte{2},
		NickName:     "testNickName",
		FirstName:    "test",
		LastName:     "test",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	testCard := cards.Card{
		ID:     uuid.New(),
		UserID: testUser1.ID,
	}

	testClub1 := clubs.Club{
		ID:        uuid.New(),
		OwnerID:   testUser1.ID,
		Name:      testUser1.NickName,
		CreatedAt: time.Now().UTC(),
	}

	testSquad1 := clubs.Squad{
		ID:        uuid.New(),
		Name:      "test squad",
		ClubID:    testClub1.ID,
		Tactic:    clubs.Balanced,
		Formation: clubs.FourTwoFour,
	}

	testClub2 := clubs.Club{
		ID:        uuid.New(),
		OwnerID:   testUser1.ID,
		Name:      testUser1.NickName,
		CreatedAt: time.Now().UTC(),
	}

	testSquad2 := clubs.Squad{
		ID:        uuid.New(),
		Name:      "test squad",
		ClubID:    testClub2.ID,
		Tactic:    clubs.Balanced,
		Formation: clubs.FourTwoFour,
	}

	testMatch := matches.Match{
		ID:       uuid.New(),
		User1ID:  testUser1.ID,
		Squad1ID: testSquad1.ID,
		User2ID:  testUser2.ID,
		Squad2ID: testSquad2.ID,
	}

	testMatchGoal1 := matches.MatchGoals{
		ID:      uuid.New(),
		MatchID: testMatch.ID,
		UserID:  testUser1.ID,
		CardID:  testCard.ID,
		Minute:  25,
	}

	testMatchGoal2 := matches.MatchGoals{
		ID:      uuid.New(),
		MatchID: testMatch.ID,
		UserID:  testUser1.ID,
		CardID:  testCard.ID,
		Minute:  41,
	}

	newCursor := pagination.Cursor{
		Limit: 10,
		Page:  1,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryCards := db.Cards()
		repositoryUsers := db.Users()
		repositoryClubs := db.Clubs()
		repositoryMatches := db.Matches()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, testUser1)
			require.NoError(t, err)

			err = repositoryUsers.Create(ctx, testUser2)
			require.NoError(t, err)

			_, err = repositoryClubs.Create(ctx, testClub1)
			require.NoError(t, err)

			_, err = repositoryClubs.CreateSquad(ctx, testSquad1)
			require.NoError(t, err)

			_, err = repositoryClubs.Create(ctx, testClub2)
			require.NoError(t, err)

			_, err = repositoryClubs.CreateSquad(ctx, testSquad2)
			require.NoError(t, err)

			err = repositoryMatches.Create(ctx, testMatch)
			require.NoError(t, err)
		})

		t.Run("List matches", func(t *testing.T) {
			allMatchesDB, err := repositoryMatches.ListMatches(ctx, newCursor)
			require.NoError(t, err)
			compareMatchesSlice(t, allMatchesDB.Matches, []matches.Match{testMatch})
		})

		t.Run("Get", func(t *testing.T) {
			matchDB, err := repositoryMatches.Get(ctx, testMatch.ID)
			require.NoError(t, err)
			compareMatches(t, matchDB, testMatch)
		})

		t.Run("Add goal in the match", func(t *testing.T) {
			err := repositoryCards.Create(ctx, testCard)
			require.NoError(t, err)

			err = repositoryMatches.AddGoals(ctx, []matches.MatchGoals{testMatchGoal1, testMatchGoal2})
			require.NoError(t, err)
		})

		t.Run("List match goals", func(t *testing.T) {
			matchGoalsDB, err := repositoryMatches.ListMatchGoals(ctx, testMatch.ID)
			require.NoError(t, err)
			compareMatchGoals(t, matchGoalsDB, []matches.MatchGoals{testMatchGoal1, testMatchGoal2})
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryMatches.Delete(ctx, testMatch.ID)
			require.NoError(t, err)
		})
	})
}

func compareMatches(t *testing.T, matchDB, matchTest matches.Match) {
	assert.Equal(t, matchDB.ID, matchTest.ID)
	assert.Equal(t, matchDB.User1ID, matchTest.User1ID)
	assert.Equal(t, matchDB.User2ID, matchTest.User2ID)
}

func compareMatchesSlice(t *testing.T, matchesDB, matchesTest []matches.Match) {
	assert.Equal(t, len(matchesDB), len(matchesTest))

	for i := 0; i < len(matchesDB); i++ {
		assert.Equal(t, matchesDB[i].ID, matchesTest[i].ID)
		assert.Equal(t, matchesDB[i].User1ID, matchesTest[i].User1ID)
		assert.Equal(t, matchesDB[i].User2ID, matchesTest[i].User2ID)
	}
}

func compareMatchGoals(t *testing.T, matchGoalsDB, matchGoalsTest []matches.MatchGoals) {
	assert.Equal(t, len(matchGoalsDB), len(matchGoalsTest))

	for i := 0; i < len(matchGoalsDB); i++ {
		assert.Equal(t, matchGoalsDB[i].ID, matchGoalsTest[i].ID)
		assert.Equal(t, matchGoalsDB[i].MatchID, matchGoalsTest[i].MatchID)
		assert.Equal(t, matchGoalsDB[i].UserID, matchGoalsTest[i].UserID)
		assert.Equal(t, matchGoalsDB[i].CardID, matchGoalsTest[i].CardID)
		assert.Equal(t, matchGoalsDB[i].Minute, matchGoalsTest[i].Minute)
	}
}
