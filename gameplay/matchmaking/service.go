// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matchmaking

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/console/connections"
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
}

// NewService is a constructor for matchmaking service.
func NewService(players DB, connections *connections.Service) *Service {
	return &Service{
		players:     players,
		connections: connections,
	}
}

// Create creates a player by user.
func (service *Service) Create(userID uuid.UUID) error {
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

	if req.Action == queue.ActionStartSearch {
		if err = service.players.Create(player); err != nil {
			return ErrMatchmaking.Wrap(err)
		}

		if err = conn.WriteJSON("you added"); err != nil {
			return ErrMatchmaking.Wrap(err)
		}

		match, err := service.MatchPlayer(&player)
		if err != nil {
			return ErrMatchmaking.Wrap(err)
		}

		fmt.Println(match)
	}

	fmt.Println(service.players.List())

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
func (service *Service) MatchPlayer(player *Player) (*Match, error) {
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

	if err := match.Player1.Conn.WriteJSON("do you confirm play?"); err != nil {
		return nil, ErrMatchmaking.Wrap(err)
	}
	if err := match.Player2.Conn.WriteJSON("do you confirm play?"); err != nil {
		return nil, ErrMatchmaking.Wrap(err)
	}

	if err := match.Player1.Conn.ReadJSON(&reqPlayer1); err != nil {
		return nil, ErrMatchmaking.Wrap(err)
	}

	if err := match.Player2.Conn.ReadJSON(&reqPlayer2); err != nil {
		return nil, ErrMatchmaking.Wrap(err)
	}

	if reqPlayer1.Action == queue.ActionConfirm && reqPlayer2.Action == queue.ActionConfirm {
		player.Waiting = false
		other.Waiting = false

		if err := match.Player1.Conn.WriteJSON(match); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
		if err := match.Player2.Conn.WriteJSON(match); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
	}

	if reqPlayer1.Action == queue.ActionReject {
		if err := match.Player1.Conn.WriteJSON("you are still in search"); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
	}

	if reqPlayer2.Action == queue.ActionReject {
		if err := match.Player2.Conn.WriteJSON("you are still in search"); err != nil {
			return nil, ErrMatchmaking.Wrap(err)
		}
	}

	return match, nil
}
