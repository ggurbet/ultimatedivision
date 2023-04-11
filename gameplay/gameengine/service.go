// Copyright (C) 2021 - 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/clubs"
	"ultimatedivision/gameplay/matches"
)

// ErrGameEngine indicates that there was an error in the service.
var ErrGameEngine = errs.Class("game engine service error")

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
	games   DB
	clubs   *clubs.Service
	avatars *avatars.Service
	cards   *cards.Service
	matches *matches.Service
	config  Config
}

// NewService is a constructor for game engine service.
func NewService(games DB, clubs *clubs.Service, avatars *avatars.Service, cards *cards.Service, matches *matches.Service, config Config) *Service {
	return &Service{
		games:   games,
		clubs:   clubs,
		avatars: avatars,
		cards:   cards,
		matches: matches,
		config:  config,
	}
}

const (
	minPlace = 0
	maxPlace = 83
)

// GetCardMoves get all card possible moves.
func (service *Service) GetCardMoves(cardPlace int, isThreeSteps bool) ([]int, error) {
	bottom := []int{6, 13, 20, 27, 34, 41, 48, 55, 62, 69, 76, 83}
	bottom1 := []int{82, 75, 68, 61, 54, 47, 40, 33, 26, 19, 12, 5}
	bottom2 := []int{81, 74, 67, 60, 53, 46, 39, 32, 25, 18, 11, 4}

	top := []int{77, 70, 63, 56, 49, 42, 35, 28, 21, 14, 7, 0}
	top1 := []int{71, 64, 57, 50, 43, 36, 29, 22, 15}
	top2 := []int{72, 65, 58, 51, 44, 37, 30, 23, 16, 9, 2, 79}

	exceptions := []int{1, 5, 78, 82}

	if cardPlace < minPlace || cardPlace > maxPlace {
		return []int{}, ErrGameEngine.New("player place can not be more 83 or les than 0, player place is %d", cardPlace)
	}

	var stepInWidth []int
	var moves []int

	if isThreeSteps == true {

		switch {
		case contains(top, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace, cardPlace+1, cardPlace+2, cardPlace+3)

		case contains(top1, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		case contains(top2, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2, cardPlace+3)

		case contains(bottom, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace)

		case contains(bottom1, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1)

		case contains(bottom2, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		case contains(exceptions, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2, cardPlace+3)

		case cardPlace == 8:
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2, cardPlace+3)

		case cardPlace == 12:
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1)

		}

		for _, w := range stepInWidth {
			min := w - 14
			max := w + 14
			min21 := w - 21
			max21 := w + 21
			moves = append(moves, min, min+7, max-7, max, w, min21, max21)
		}

	} else {
		switch {
		case contains(top, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace, cardPlace+1, cardPlace+2)

		case contains(bottom, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace)

		case contains(exceptions, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1)

		case cardPlace == 8:
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		case cardPlace == 12:
			stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1)

		default:
		}
		stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		for _, w := range stepInWidth {
			min := w - 14
			max := w + 14
			moves = append(moves, min, min+7, max-7, max, w)
		}
	}

	sort.Ints(moves)
	moves = removeMin(moves, minPlace)
	moves = removeMax(moves, maxPlace)

	return moves, nil
}

func removeMin(l []int, min int) []int {
	for i, v := range l {
		if v >= min {
			return l[i:]
		}
	}
	return l
}
func removeMax(l []int, max int) []int {
	for i, v := range l {
		if v > max {
			return l[:i]
		}
	}
	return l
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func removeIntersections(moves, playerPositions []int) (movesWithoutIntersections []int) {
	for _, v := range moves {
		if !contains(playerPositions, v) {
			movesWithoutIntersections = append(movesWithoutIntersections, v)
		}
	}
	return movesWithoutIntersections
}

// Move update card moves and get possible moves cells.
func (service *Service) Move(ctx context.Context, matchID uuid.UUID, card CardIDWithPosition) (CardAvailableAction, error) {
	var cardAvailableAction CardAvailableAction

	gameInfoJSON, err := service.games.Get(ctx, matchID)
	if err != nil {
		return CardAvailableAction{}, ErrGameEngine.Wrap(err)
	}

	var game Game
	game.MatchID = matchID

	err = json.Unmarshal([]byte(gameInfoJSON), &game.GameInfo)
	if err != nil {
		return CardAvailableAction{}, ErrGameEngine.Wrap(err)
	}

	var moves []int
	var allPositionsInUse []int

	allPositionsInUse = append(allPositionsInUse, card.Position)

	for _, cardWithPosition := range game.GameInfo {
		allPositionsInUse = append(allPositionsInUse, cardWithPosition.Position)
	}

	// check whether the position we want to go to is occupied.
	if contains(allPositionsInUse, card.Position) {
		return CardAvailableAction{}, ErrGameEngine.New("Can not move to position, already in use")
	}

	// check, Update and get all possible moves.
	for i, cardData := range game.GameInfo {
		if cardData.CardID == card.CardID {
			game.GameInfo[i].Position = card.Position

			newGameInfoJSON, err := json.Marshal(game.GameInfo)
			if err != nil {
				return CardAvailableAction{}, ErrGameEngine.Wrap(err)
			}

			err = service.games.Update(ctx, matchID, string(newGameInfoJSON))
			if err != nil {
				return CardAvailableAction{}, ErrGameEngine.Wrap(err)
			}
			isThreeSteps := true
			moves, err = service.GetCardMoves(card.Position, isThreeSteps)
			if err != nil {
				return CardAvailableAction{}, ErrGameEngine.Wrap(err)
			}
		}
	}

	// remove already occupied positions.
	moves = removeIntersections(moves, allPositionsInUse)

	cardAvailableAction = CardAvailableAction{
		Action:        ActionMove,
		CardID:        card.CardID,
		FieldPosition: moves,
	}

	return cardAvailableAction, nil
}

// GameInformation creates a player by user.
func (service *Service) GameInformation(ctx context.Context, player1SquadID, player2SquadID uuid.UUID) (MatchRepresentation, error) {
	var cardsWithPositionPlayer1 []CardWithPosition
	var cardsWithPositionPlayer2 []CardWithPosition
	var cardsAvailableAction []CardAvailableAction

	squadCardsPlayer1, err := service.clubs.ListCards(ctx, player1SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	squadCardsPlayer2, err := service.clubs.ListCards(ctx, player2SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	squadPlayer1, err := service.clubs.GetSquad(ctx, player1SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	squadPlayer2, err := service.clubs.GetSquad(ctx, player2SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	clubPlayer1, err := service.clubs.Get(ctx, squadPlayer1.ClubID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	clubPlayer2, err := service.clubs.Get(ctx, squadPlayer2.ClubID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	var matchInfo []CardIDWithPosition

	for _, sqCard := range squadCardsPlayer1 {
		avatar, err := service.avatars.Get(ctx, sqCard.Card.ID)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardWithPositionPlayer := CardWithPosition{
			Card:          sqCard.Card,
			Avatar:        avatar,
			FieldPosition: service.squadPositionToFieldPositionLeftSide(sqCard.Position),
		}

		cardInfo := CardIDWithPosition{
			CardID:   sqCard.Card.ID,
			Position: cardWithPositionPlayer.FieldPosition,
		}

		matchInfo = append(matchInfo, cardInfo)
		isThreeSteps := true
		fieldPosition, err := service.GetCardMoves(cardWithPositionPlayer.FieldPosition, isThreeSteps)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardAvailableAction := CardAvailableAction{
			Action:        ActionMove,
			CardID:        sqCard.Card.ID,
			FieldPosition: fieldPosition,
		}

		cardsWithPositionPlayer1 = append(cardsWithPositionPlayer1, cardWithPositionPlayer)
		cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction)
	}

	for _, sqCard := range squadCardsPlayer2 {
		avatar, err := service.avatars.Get(ctx, sqCard.Card.ID)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardWithPositionPlayer := CardWithPosition{
			Card:          sqCard.Card,
			Avatar:        avatar,
			FieldPosition: service.squadPositionToFieldPositionRightSide(sqCard.Position),
		}
		isThreeSteps := true
		fieldPosition, err := service.GetCardMoves(cardWithPositionPlayer.FieldPosition, isThreeSteps)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardAvailableAction := CardAvailableAction{
			Action:        ActionMove,
			CardID:        sqCard.Card.ID,
			FieldPosition: fieldPosition,
		}

		cardInfo := CardIDWithPosition{
			CardID:   sqCard.Card.ID,
			Position: cardWithPositionPlayer.FieldPosition,
		}

		matchInfo = append(matchInfo, cardInfo)

		cardsWithPositionPlayer2 = append(cardsWithPositionPlayer2, cardWithPositionPlayer)
		cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction)
	}

	matchID, err := service.matches.CreateMatchID(ctx, player1SquadID, player2SquadID, clubPlayer1.OwnerID, clubPlayer2.OwnerID, 1)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	gameInfo, err := json.Marshal(matchInfo)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	err = service.games.Create(ctx, matchID, string(gameInfo))
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	return MatchRepresentation{
		MatchID:                matchID,
		User1CardsWithPosition: cardsWithPositionPlayer1,
		User2CardsWithPosition: cardsWithPositionPlayer2,
		BallPosition:           0,
		CardAvailableAction:    cardsAvailableAction,
		User1ClubInformation:   clubPlayer1,
		User2ClubInformation:   clubPlayer2,
		User1SquadInformation:  squadPlayer1,
		User2SquadInformation:  squadPlayer2,
		Rounds:                 service.config.Rounds,
	}, nil
}

func (service *Service) squadPositionToFieldPositionLeftSide(squadPosition clubs.Position) int {
	switch squadPosition {
	case clubs.GK:
		return service.config.LeftSide.Goalkeeper
	case clubs.LB:
		return service.config.LeftSide.LeftBack
	case clubs.RB:
		return service.config.LeftSide.RightBack
	case clubs.LM:
		return service.config.LeftSide.LeftMid
	case clubs.RM:
		return service.config.LeftSide.RightMid
	case clubs.LCD:
		return service.config.LeftSide.CenterBackLeft
	case clubs.RCD:
		return service.config.LeftSide.CenterBackRight
	case clubs.LCM:
		return service.config.LeftSide.CenterMidLeft
	case clubs.RCM:
		return service.config.LeftSide.CenterMidRight
	case clubs.LST:
		return service.config.LeftSide.ForwardLeft
	case clubs.RST:
		return service.config.LeftSide.ForwardRight
	}

	return 0
}

func (service *Service) squadPositionToFieldPositionRightSide(squadPosition clubs.Position) int {
	switch squadPosition {
	case clubs.GK:
		return service.config.RightSide.Goalkeeper
	case clubs.LB:
		return service.config.RightSide.LeftBack
	case clubs.RB:
		return service.config.RightSide.RightBack
	case clubs.LM:
		return service.config.RightSide.LeftMid
	case clubs.RM:
		return service.config.RightSide.RightMid
	case clubs.LCD:
		return service.config.RightSide.CenterBackLeft
	case clubs.RCD:
		return service.config.RightSide.CenterBackRight
	case clubs.LCM:
		return service.config.RightSide.CenterMidLeft
	case clubs.RCM:
		return service.config.RightSide.CenterMidRight
	case clubs.LST:
		return service.config.RightSide.ForwardLeft
	case clubs.RST:
		return service.config.RightSide.ForwardRight
	}
	return 0
}
