// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/pkg/pagination"
	rand2 "ultimatedivision/pkg/rand"
)

// ErrMatches indicates that there was an error in the service.
var ErrMatches = errs.Class("matches service error")

// Service is handling matches related logic.
//
// architecture: Service
type Service struct {
	matches DB
	config  Config
	clubs   *clubs.Service
	cards   *cards.Service
}

// NewService is a constructor for matches service.
func NewService(matches DB, config Config, clubs *clubs.Service, cards *cards.Service) *Service {
	return &Service{
		matches: matches,
		config:  config,
		clubs:   clubs,
		cards:   cards,
	}
}

// periodBegin defines index of the beginning of period.
const periodBegin = 0

// periodEnd defines index of the ending of period.
const periodEnd = 1

// MinNumberOfMatches defines minimal number of matches to participate in weekly competition.
const MinNumberOfMatches = 3

// maxNumberOfMatches defines maximal number of matches in weekly competition.
const maxNumberOfMatches = 30

// Play initiates match between users, calls methods to generate result.
func (service *Service) Play(ctx context.Context, match Match, squadCards1 []clubs.SquadCard, squadCards2 []clubs.SquadCard) error {
	periods := []int{service.config.Periods.First.Begin, service.config.Periods.First.End,
		service.config.Periods.Second.Begin, service.config.Periods.Second.End,
		service.config.Periods.Third.Begin, service.config.Periods.Third.End,
		service.config.Periods.Fourth.Begin, service.config.Periods.Fourth.End,
		service.config.Periods.Fifth.Begin, service.config.Periods.Fifth.End,
		service.config.Periods.Sixth.Begin, service.config.Periods.Sixth.End,
		service.config.Periods.Seventh.Begin, service.config.Periods.Seventh.End,
		service.config.Periods.Eighth.Begin, service.config.Periods.Eighth.End,
		service.config.Periods.Ninth.Begin, service.config.Periods.Ninth.End,
		service.config.Periods.Tenth.Begin, service.config.Periods.Tenth.End}

	goalProbability := service.config.GoalProbability

	squadPowerAccuracy := service.config.SquadPowerAccuracy

	goalProbabilityByPosition := map[clubs.Position]int{
		clubs.CST:  service.config.GoalProbabilityByPosition.ST,
		clubs.RW:   service.config.GoalProbabilityByPosition.RW,
		clubs.LW:   service.config.GoalProbabilityByPosition.LW,
		clubs.CCAM: service.config.GoalProbabilityByPosition.CAM,
		clubs.CCM:  service.config.GoalProbabilityByPosition.CM,
		clubs.RM:   service.config.GoalProbabilityByPosition.RM,
		clubs.LM:   service.config.GoalProbabilityByPosition.LM,
		clubs.CCDM: service.config.GoalProbabilityByPosition.CDM,
		clubs.CCD:  service.config.GoalProbabilityByPosition.CD,
		clubs.LB:   service.config.GoalProbabilityByPosition.LB,
		clubs.RB:   service.config.GoalProbabilityByPosition.RB,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	goals := make([]MatchGoals, 0, 10)

	for i := 0; i < len(periods); i += 2 {
		randNumber := rand.Intn(100) + 1
		if randNumber > goalProbability {
			continue
		}

		minute := rand2.Minute(periods[i+periodBegin], periods[i+periodEnd])
		userID, cardID, err := service.chooseSquad(ctx, goalProbabilityByPosition,
			squadPowerAccuracy, match.User1ID, match.User2ID, squadCards1, squadCards2)
		if err != nil {
			return ErrMatches.Wrap(err)
		}

		goals = append(goals, MatchGoals{
			ID:      uuid.New(),
			MatchID: match.ID,
			UserID:  userID,
			CardID:  cardID,
			Minute:  minute,
		})
	}

	err := service.matches.AddGoals(ctx, goals)
	if err != nil {
		return ErrMatches.Wrap(err)
	}

	err = service.RankMatch(ctx, match, goals)

	return ErrMatches.Wrap(err)
}

// choseGoalscorer returns id of cards which scored goal.
func chooseGoalscorer(squadCards []clubs.SquadCard, goalByPosition map[clubs.Position]int) uuid.UUID {
	rand.Seed(time.Now().UTC().UnixNano())
	var cardsByPosition []uuid.UUID
	randNumber := rand.Intn(100) + 1

	switch {
	case randNumber > 0 && randNumber <= goalByPosition[clubs.CST]:
		for _, card := range squadCards {
			if card.Position == clubs.CST || card.Position == clubs.LST || card.Position == clubs.RST {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}

		if len(cardsByPosition) > 0 {
			break
		}

		fallthrough
	case randNumber > goalByPosition[clubs.CST] &&
		randNumber < goalByPosition[clubs.CST]+goalByPosition[clubs.RW]:
		for _, card := range squadCards {
			if card.Position == clubs.RW || card.Position == clubs.LW ||
				card.Position == clubs.CCM || card.Position == clubs.CCAM ||
				card.Position == clubs.LCM || card.Position == clubs.LCAM ||
				card.Position == clubs.RCM || card.Position == clubs.RCAM {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}

		if len(cardsByPosition) > 0 {
			break
		}

		fallthrough
	case randNumber > goalByPosition[clubs.CST]+goalByPosition[clubs.RW] &&
		randNumber < 100-goalByPosition[clubs.CCD]:
		for _, card := range squadCards {
			if card.Position == clubs.RM || card.Position == clubs.LM ||
				card.Position == clubs.CCDM ||
				card.Position == clubs.LCDM ||
				card.Position == clubs.RCDM {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}

		if len(cardsByPosition) > 0 {
			break
		}

		fallthrough
	case randNumber >= 100-goalByPosition[clubs.CCD] && randNumber < 100:
		for _, card := range squadCards {
			if card.Position == clubs.CCD || card.Position == clubs.LCD ||
				card.Position == clubs.LB || card.Position == clubs.RCD ||
				card.Position == clubs.RB || card.Position == clubs.RWB ||
				card.Position == clubs.LWB {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}
	}
	if len(cardsByPosition) == 0 {
		return uuid.Nil
	}

	randIndex := rand.Intn(len(cardsByPosition))
	goalscorer := cardsByPosition[randIndex]

	return goalscorer
}

// chooseSquad returns the squad which is stronger in the period.
func (service *Service) chooseSquad(ctx context.Context, goalByPosition map[clubs.Position]int, squadPowerAccuracy int, user1 uuid.UUID, user2 uuid.UUID, squadCards1, squadCards2 []clubs.SquadCard) (uuid.UUID, uuid.UUID, error) {
	squad1Effectiveness, err := service.clubs.CalculateEffectivenessOfSquad(ctx, squadCards1)
	if err != nil {
		return uuid.Nil, uuid.Nil, ErrMatches.Wrap(err)
	}

	squad2Effectiveness, err := service.clubs.CalculateEffectivenessOfSquad(ctx, squadCards2)
	if err != nil {
		return uuid.Nil, uuid.Nil, ErrMatches.Wrap(err)
	}

	randAccuracy1 := float64(rand.Intn(2*squadPowerAccuracy+1)-squadPowerAccuracy) / 100
	randAccuracy2 := float64(rand.Intn(2*squadPowerAccuracy+1)-squadPowerAccuracy) / 100

	squad1Effectiveness += squad1Effectiveness * randAccuracy1
	squad2Effectiveness += squad1Effectiveness * randAccuracy2

	if squad1Effectiveness > squad2Effectiveness {
		return user1, chooseGoalscorer(squadCards1, goalByPosition), nil
	}

	return user2, chooseGoalscorer(squadCards2, goalByPosition), nil
}

// Create creates new match.
func (service *Service) Create(ctx context.Context, squad1ID uuid.UUID, squad2ID uuid.UUID, user1ID, user2ID uuid.UUID, seasonID int) (uuid.UUID, error) {
	squadCards1, err := service.clubs.ListSquadCardIDs(ctx, squad1ID)
	if err != nil {
		return uuid.Nil, ErrMatches.Wrap(err)
	}

	squadCards2, err := service.clubs.ListSquadCardIDs(ctx, squad2ID)
	if err != nil {
		return uuid.Nil, ErrMatches.Wrap(err)
	}

	newMatch := Match{
		ID:       uuid.New(),
		User1ID:  user1ID,
		Squad1ID: squad1ID,
		User2ID:  user2ID,
		Squad2ID: squad2ID,
		SeasonID: seasonID,
	}

	if err = service.matches.Create(ctx, newMatch); err != nil {
		return uuid.Nil, ErrMatches.Wrap(err)
	}

	err = service.Play(ctx, newMatch, squadCards1, squadCards2)

	return newMatch.ID, ErrMatches.Wrap(err)
}

// Get returns match by id.
func (service *Service) Get(ctx context.Context, matchID uuid.UUID) (Match, error) {
	match, err := service.matches.Get(ctx, matchID)

	return match, ErrMatches.Wrap(err)
}

// List returns page of matches.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	allMatches, err := service.matches.ListMatches(ctx, cursor)

	return allMatches, ErrMatches.Wrap(err)
}

// Delete deletes match.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrMatches.Wrap(service.matches.Delete(ctx, id))
}

// ListMatchGoals returns all goals scored in the match.
func (service *Service) ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]MatchGoals, error) {
	matchGoals, err := service.matches.ListMatchGoals(ctx, matchID)

	return matchGoals, ErrMatches.Wrap(err)
}

// GetGameResult returns goals of each user in the match.
func (service *Service) GetGameResult(ctx context.Context, matchID uuid.UUID) (GameResult, error) {
	matchResults, err := service.matches.GetMatchResult(ctx, matchID)
	if err != nil {
		return GameResult{}, ErrMatches.Wrap(err)
	}

	matchGoals, err := service.ListMatchGoals(ctx, matchID)
	if err != nil {
		return GameResult{}, ErrMatches.Wrap(err)
	}

	gameResult := GameResult{
		MatchResults: matchResults,
	}

	if len(gameResult.MatchResults) == 2 {
		for k, result := range gameResult.MatchResults {
			for _, goal := range matchGoals {
				if goal.UserID == result.UserID {
					card, err := service.cards.Get(ctx, goal.CardID)
					if err != nil {
						return gameResult, ErrMatches.Wrap(err)
					}

					gameResult.MatchResults[k].Goalscorers = append(gameResult.MatchResults[k].Goalscorers, Goalscorer{
						Card:   card,
						Minute: goal.Minute,
					})
				}
			}
		}

		return gameResult, nil
	}

	match, err := service.matches.Get(ctx, matchID)
	if err != nil {
		return gameResult, ErrMatches.Wrap(err)
	}

	var newGameResult GameResult
	newGameResult.MatchResults = append(newGameResult.MatchResults, MatchResult{UserID: match.User1ID})
	newGameResult.MatchResults = append(newGameResult.MatchResults, MatchResult{UserID: match.User2ID})

	for k, result := range newGameResult.MatchResults {
		for _, res := range matchResults {
			if result.UserID == res.UserID {
				newGameResult.MatchResults[k].QuantityGoals = res.QuantityGoals
			}
		}

		for _, goal := range matchGoals {
			if goal.UserID == result.UserID {
				card, err := service.cards.Get(ctx, goal.CardID)
				if err != nil {
					return newGameResult, ErrMatches.Wrap(err)
				}

				newGameResult.MatchResults[k].Goalscorers = append(newGameResult.MatchResults[k].Goalscorers, Goalscorer{
					Card:   card,
					Minute: goal.Minute,
				})
			}
		}
	}

	return newGameResult, nil
}

// ListSquadMatches returns all club matches in season.
func (service *Service) ListSquadMatches(ctx context.Context, seasonID int) ([]Match, error) {
	allMatches, err := service.matches.ListSquadMatches(ctx, seasonID)
	return allMatches, ErrMatches.Wrap(err)
}

// RankMatch evaluates how many points each user receive per match.
func (service *Service) RankMatch(ctx context.Context, match Match, matchGoals []MatchGoals) error {
	var (
		user1Goals int
		user2Goals int
	)

	for _, goal := range matchGoals {
		if goal.UserID == match.User1ID {
			user1Goals++
			continue
		}
		user2Goals++
	}

	switch {
	case user1Goals > user2Goals:
		match.User1Points = service.config.NumberOfPointsForWin
		match.User2Points = service.config.NumberOfPointsForLosing
	case user1Goals < user2Goals:
		match.User1Points = service.config.NumberOfPointsForLosing
		match.User2Points = service.config.NumberOfPointsForWin
	case user1Goals == user2Goals:
		match.User1Points = service.config.NumberOfPointsForDraw
		match.User2Points = service.config.NumberOfPointsForDraw
	}

	return ErrMatches.Wrap(service.matches.UpdateMatch(ctx, match))
}

// GetStatistic returns statistic of club in season.
func (service *Service) GetStatistic(ctx context.Context, club clubs.Club, seasonID int) (Statistic, error) {
	var statistic Statistic

	allMatches, err := service.ListSquadMatches(ctx, seasonID)
	if err != nil {
		return statistic, ErrMatches.Wrap(err)
	}

	if len(allMatches) < MinNumberOfMatches {
		return statistic, nil
	}

	if len(allMatches) > maxNumberOfMatches {
		allMatches = allMatches[:maxNumberOfMatches]
	}

	var (
		goalScored    int
		goalsConceded int
	)

	for _, match := range allMatches {
		statistic.MatchPlayed++

		if match.User1ID == club.OwnerID {
			switch {
			case match.User1Points == service.config.NumberOfPointsForWin:
				statistic.Wins++
			case match.User1Points == service.config.NumberOfPointsForDraw:
				statistic.Draws++
			case match.User1Points == service.config.NumberOfPointsForLosing:
				statistic.Losses++
			}
		} else if match.User2ID == club.OwnerID {
			switch {
			case match.User2Points == service.config.NumberOfPointsForWin:
				statistic.Wins++
			case match.User2Points == service.config.NumberOfPointsForDraw:
				statistic.Draws++
			case match.User2Points == service.config.NumberOfPointsForLosing:
				statistic.Losses++
			}
		} else {
			return statistic, err
		}
		matchGoals, err := service.ListMatchGoals(ctx, match.ID)
		if err != nil {
			return statistic, ErrMatches.Wrap(err)
		}
		for _, goal := range matchGoals {
			if goal.UserID == club.OwnerID {
				goalScored++
				continue
			}
			goalsConceded++
		}
	}

	statistic.Points = service.config.NumberOfPointsForWin*statistic.Wins + service.config.NumberOfPointsForDraw*statistic.Draws +
		+service.config.NumberOfPointsForLosing*statistic.Losses

	statistic.GoalDifference = goalScored - goalsConceded
	statistic.Club = club

	return statistic, nil
}
