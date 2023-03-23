package gameengine_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/clubs"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/gameengine"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/seasons"
	"ultimatedivision/users"
)

func TestMatches(t *testing.T) {
	matchID := uuid.New()
	card1 := gameengine.CardIDWithPosition{
		CardID:   uuid.New(),
		Position: 1,
	}
	card2 := gameengine.CardIDWithPosition{
		CardID:   uuid.New(),
		Position: 2,
	}
	card3 := gameengine.CardIDWithPosition{
		CardID:   uuid.New(),
		Position: 3,
	}

	testGame := gameengine.Game{
		MatchID: matchID,
		GameInfo: []gameengine.CardIDWithPosition{
			card1,
			card2,
			card3,
		},
	}

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

	division1 := divisions.Division{
		ID:             uuid.New(),
		Name:           10,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}

	season1 := seasons.Season{
		ID:         1,
		DivisionID: division1.ID,
		StartedAt:  time.Now().UTC(),
		EndedAt:    time.Time{},
	}

	testClub1 := clubs.Club{
		ID:         uuid.New(),
		OwnerID:    testUser1.ID,
		Name:       testUser1.NickName,
		DivisionID: division1.ID,
		CreatedAt:  time.Now().UTC(),
	}

	testSquad1 := clubs.Squad{
		ID:        uuid.New(),
		Name:      "test squad",
		ClubID:    testClub1.ID,
		Tactic:    clubs.Balanced,
		Formation: clubs.FourTwoFour,
	}

	testClub2 := clubs.Club{
		ID:         uuid.New(),
		OwnerID:    testUser2.ID,
		Name:       testUser2.NickName,
		DivisionID: division1.ID,
		CreatedAt:  time.Now().UTC(),
	}

	testSquad2 := clubs.Squad{
		ID:        uuid.New(),
		Name:      "test squad",
		ClubID:    testClub2.ID,
		Tactic:    clubs.Balanced,
		Formation: clubs.FourTwoFour,
	}

	testMatch := matches.Match{
		ID:       matchID,
		User1ID:  testUser1.ID,
		Squad1ID: testSquad1.ID,
		User2ID:  testUser2.ID,
		Squad2ID: testSquad2.ID,
		SeasonID: season1.ID,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryGames := db.Games()

		repositoryUsers := db.Users()
		repositoryClubs := db.Clubs()
		repositoryMatches := db.Matches()
		repositoryDivisions := db.Divisions()
		repositorySeasons := db.Seasons()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, testUser1)
			require.NoError(t, err)

			err = repositoryUsers.Create(ctx, testUser2)
			require.NoError(t, err)

			err = repositoryDivisions.Create(ctx, division1)
			require.NoError(t, err)

			err = repositorySeasons.Create(ctx, season1)
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

		t.Run("Create", func(t *testing.T) {
			gameInfo, err := json.Marshal(testGame.GameInfo)
			require.NoError(t, err)

			err = repositoryGames.Create(ctx, matchID, string(gameInfo))
			require.NoError(t, err)
		})

		t.Run("Get", func(t *testing.T) {
			gameInfo, err := repositoryGames.Get(ctx, matchID)
			require.NoError(t, err)

			var game gameengine.Game
			game.MatchID = matchID

			err = json.Unmarshal([]byte(gameInfo), &game.GameInfo)
			require.NoError(t, err)
			compareGame(t, testGame, game)
		})

		t.Run("Update", func(t *testing.T) {
			testGameNew := gameengine.Game{
				MatchID: matchID,
				GameInfo: []gameengine.CardIDWithPosition{
					{
						CardID:   card1.CardID,
						Position: 10,
					},
					card2,
					card3,
				},
			}

			gameInfoJSON, err := json.Marshal(testGameNew.GameInfo)
			require.NoError(t, err)

			err = repositoryGames.Update(ctx, matchID, string(gameInfoJSON))
			require.NoError(t, err)

			gameInfo, err := repositoryGames.Get(ctx, matchID)
			require.NoError(t, err)

			var game gameengine.Game
			game.MatchID = matchID

			err = json.Unmarshal([]byte(gameInfo), &game.GameInfo)
			require.NoError(t, err)
			compareGame(t, testGameNew, game)
		})

		t.Run("Delete sql no rows", func(t *testing.T) {
			err := repositoryGames.Delete(ctx, uuid.New())
			require.Error(t, err)
			assert.Equal(t, true, gameengine.ErrNoGames.Has(err))
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryGames.Delete(ctx, matchID)
			require.NoError(t, err)

			_, err = repositoryGames.Get(ctx, matchID)
			require.Error(t, err)
		})

	})
}

func compareGame(t *testing.T, expectedGame, actualGame gameengine.Game) {
	assert.Equal(t, expectedGame.MatchID, actualGame.MatchID)
	for i, card := range actualGame.GameInfo {
		assert.Equal(t, expectedGame.GameInfo[i].CardID, card.CardID)
		assert.Equal(t, expectedGame.GameInfo[i].Position, card.Position)
	}
}
