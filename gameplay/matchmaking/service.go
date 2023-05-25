// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matchmaking

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/console/connections"
	"ultimatedivision/gameplay/gameengine"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/gameplay/queue"
)

// ErrMatchmaking indicates that there was an error in the service.
var ErrMatchmaking = errs.Class("matchmaking service error")

// Service is handling matchmaking related logic.
//
// architecture: Service
type Service struct {
	players     DB
	connections *connections.Service
	gameEngine  *gameengine.Service
	queue       *queue.Chore
}

// NewService is a constructor for matchmaking service.
func NewService(players DB, connections *connections.Service, gameEngine *gameengine.Service, queue *queue.Chore) *Service {
	return &Service{
		players:     players,
		connections: connections,
		gameEngine:  gameEngine,
		queue:       queue,
	}
}

// Create creates a player by user.
func (service *Service) Create(ctx context.Context, userID uuid.UUID) error {
	type request struct {
		Action  queue.Action `json:"action"`
		SquadID uuid.UUID    `json:"squadId"`
	}

	var req request

	conn, err := service.connections.Get(userID)
	if err != nil {
		return ErrMatchmaking.Wrap(err)
	}

	if err = conn.ReadJSON(&req); err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			err = conn.ReadJSON(&req)
			if err != nil {
				return ErrMatchmaking.Wrap(err)
			}
		} else {
			return ErrMatchmaking.Wrap(err)
		}
	}

	player := Player{
		UserID:  userID,
		SquadID: req.SquadID,
		Conn:    conn,
		Waiting: true,
	}

	fmt.Println("action1 ------>>>", req.Action)

	if req.Action == queue.ActionStartSearch {
		if err = service.players.Create(player); err != nil {
			return ErrMatchmaking.Wrap(err)
		}

		resp := queue.Response{
			Status:  http.StatusOK,
			Message: "you added",
		}
		if err = conn.WriteJSON(resp); err != nil {
			return ErrMatchmaking.Wrap(err)
		}

		match, err := service.MatchPlayer(ctx, &player)
		if err != nil {
			return ErrMatchmaking.Wrap(err)
		}

		fmt.Println(match)
		fmt.Println("Players ------>>>>", service.players.List())
	}

	return nil
}

// List returns all players.
func (service *Service) List() map[uuid.UUID]Player {
	return service.players.List()
}

// Get returns player by user.
func (service *Service) Get(userID uuid.UUID) (Player, error) {
	player, err := service.players.Get(userID)
	return player, ErrMatchmaking.Wrap(err)
}

// Delete player by user.
func (service *Service) Delete(id uuid.UUID) error {
	return ErrMatchmaking.Wrap(service.players.Delete(id))
}

// MatchPlayer finds two players and connect they to gameplay.
func (service *Service) MatchPlayer(ctx context.Context, player *Player) (*Match, error) {
	var other *Player

	type request struct {
		Action  queue.Action `json:"action"`
		SquadID uuid.UUID    `json:"squadId"`
	}

	var reqPlayer1 request
	var reqPlayer2 request

	players := service.players.List()
	for _, p := range players {
		if p.UserID != player.UserID && p.Waiting {
			err := p.Conn.WriteJSON("ok")
			if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					err := service.players.Delete(p.UserID)
					if err != nil {
						return nil, ErrMatchmaking.Wrap(err)
					}
					continue
				} else {
					return nil, ErrMatchmaking.Wrap(err)
				}
			}

			pl := p
			other = &pl
			break
		}
	}

	if other == nil {
		// No match found, add player to waiting queue.
		player.Waiting = true
		return nil, nil
	}
	// Found a match, create a new match.
	match := &Match{
		Player1: player,
		Player2: other,
	}

	resp := queue.Response{
		Status:  http.StatusOK,
		Message: "do you confirm play?",
	}
	if err := match.Player1.Conn.WriteJSON(resp); err != nil {
		return nil, ErrMatchmaking.Wrap(err)
	}
	if err := match.Player2.Conn.WriteJSON(resp); err != nil {
		return nil, ErrMatchmaking.Wrap(err)
	}

	if err := match.Player1.Conn.ReadJSON(&reqPlayer1); err != nil {
		if strings.Contains(err.Error(), "close 1001") {
			resp.Message = "you left"
			if err = match.Player2.Conn.WriteJSON(resp); err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}
			err = service.players.Delete(match.Player1.UserID)
			if err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}
			err = service.players.Delete(match.Player2.UserID)
			if err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}
			return nil, nil
		}
		return nil, ErrMatchmaking.Wrap(err)
	}

	if err := match.Player2.Conn.ReadJSON(&reqPlayer2); err != nil {
		if strings.Contains(err.Error(), "close 1001") {
			resp.Message = "you left"
			if err = match.Player1.Conn.WriteJSON(resp); err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}
			err = service.players.Delete(match.Player1.UserID)
			if err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}
			err = service.players.Delete(match.Player2.UserID)
			if err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}
			return nil, nil
		}
		return nil, ErrMatchmaking.Wrap(err)
	}

	fmt.Println("Players ------>>>>", service.players.List())

	if reqPlayer1.Action == queue.ActionConfirm && reqPlayer2.Action == queue.ActionConfirm {
		player.Waiting = false
		other.Waiting = false

		resp := queue.Response{
			Status:  http.StatusOK,
			Message: "players found",
		}
		if err := match.Player1.Conn.WriteJSON(resp); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
		if err := match.Player2.Conn.WriteJSON(resp); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}

		startGameInformation, err := service.gameEngine.GameInformation(ctx, match.Player1.SquadID, match.Player2.SquadID)
		if err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}

		resp.Message = startGameInformation

		if err := match.Player1.Conn.WriteJSON(resp); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
		if err := match.Player2.Conn.WriteJSON(resp); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}

		type gameRequest struct {
			Action        gameengine.Action `json:"action"`
			CardID        uuid.UUID         `json:"CardId"`
			Position      int               `json:"position"`
			HasBall       bool              `json:"hasBall"`
			NewPositions  []int             `json:"newPositions"`
			FinalPosition int               `json:"finalPosition"`
		}

		startGameInformation.Rounds = 0
		for i := 1; i <= startGameInformation.Rounds; i++ {
			var req gameRequest

			if err := match.Player1.Conn.ReadJSON(&req); err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}

			cardAvailableAction, err := service.gameEngine.GameLogicByAction(ctx, startGameInformation.MatchID,
				gameengine.CardIDWithPosition{
					CardID:   req.CardID,
					Position: req.Position,
				}, req.Action, req.NewPositions, req.FinalPosition, req.HasBall)
			if err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}

			if err := match.Player1.Conn.WriteJSON(cardAvailableAction); err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}

			if err := match.Player2.Conn.ReadJSON(&req); err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}

			cardAvailableAction, err = service.gameEngine.GameLogicByAction(ctx, startGameInformation.MatchID, gameengine.CardIDWithPosition{
				CardID:   req.CardID,
				Position: req.Position,
			}, req.Action, req.NewPositions, req.FinalPosition, req.HasBall)
			if err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}

			if err := match.Player2.Conn.WriteJSON(cardAvailableAction); err != nil {
				return nil, ErrMatchmaking.Wrap(err)
			}
		}

		time.Sleep(time.Second * 10)

		var value = new(big.Int)
		value.SetString(service.queue.Config.DrawValue, 10)

		firstClient := queue.Client{
			UserID:     match.Player1.UserID,
			Connection: match.Player1.Conn,
			SquadID:    match.Player1.SquadID,
			IsPlaying:  true,
			CreatedAt:  time.Time{},
		}

		secondClient := queue.Client{
			UserID:     match.Player2.UserID,
			Connection: match.Player2.Conn,
			SquadID:    match.Player2.SquadID,
			IsPlaying:  true,
			CreatedAt:  time.Time{},
		}

		winResult := queue.WinResult{
			Client:     firstClient,
			GameResult: matches.GameResult{},
			Value:      value,
		}

		matchResultPlayer1 := matches.MatchResult{
			UserID:        match.Player1.UserID,
			QuantityGoals: 0,
			Goalscorers:   nil,
		}
		matchResultPlayer2 := matches.MatchResult{
			UserID:        match.Player2.UserID,
			QuantityGoals: 0,
			Goalscorers:   nil,
		}
		winResult.GameResult.MatchResults = append(winResult.GameResult.MatchResults, matchResultPlayer1, matchResultPlayer2)

		go service.queue.FinishWithWinResult(ctx, winResult)

		winResult.Client = secondClient

		go service.queue.FinishWithWinResult(ctx, winResult)

		return match, nil
	}

	resp.Message = "you left"
	if reqPlayer1.Action == queue.ActionReject || reqPlayer2.Action == queue.ActionReject {
		if err := match.Player1.Conn.WriteJSON(resp); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}

		if err := match.Player2.Conn.WriteJSON(resp); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}

		err := service.players.Delete(match.Player1.UserID)
		if err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
		err = service.players.Delete(match.Player2.UserID)
		if err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
	}

	return match, nil
}
