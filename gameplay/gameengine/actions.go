// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"time"

	"github.com/google/uuid"

	"ultimatedivision/cards"
)

// Action defines list of possible player action in the field.
type Action string

const (
	// ActionMove defines move action by player.
	ActionMove Action = "move"
	// ActionMoveWithBall defines move action by player with ball.
	ActionMoveWithBall Action = "moveWithBall"
	// ActionPass defines pass by player to another player.
	ActionPass Action = "pass"
	// ActionCrossPass defines passing the ball by throwing it into the air in the direction of a player on his team.
	ActionCrossPass Action = "crossPass"
	// ActionPassThrough defines pass in free zone on the move often between players of the other team.
	ActionPassThrough Action = "passTrough"
	// ActionDirectShot defines direct shot.
	ActionDirectShot Action = "directShot"
	// ActionCurlShot defines curl shot.
	ActionCurlShot Action = "curlShot"
	// ActionTakeawayShot defines powerful shot from the box.
	ActionTakeawayShot Action = "takeawayShot"
	// ActionTackle defines tackling the ball from an opponent.
	ActionTackle Action = "tackle"
	// ActionSlidingTackle defines tackle by sliding on the field.
	ActionSlidingTackle Action = "slidingTackle"
	// ActionDribbling defines action when player move with some feints ot tricks.
	ActionDribbling Action = "dribbling"
	// ActionFeints defines action when player show feints.
	ActionFeints Action = "feints"
)

// GameConfig contains config values related to game.
type GameConfig struct {
}

// CardAvailableAction defines in which position card could be placed and which action it could do there.
type CardAvailableAction struct {
	CardID        uuid.UUID `json:"cardId"`
	Action        Action    `json:"action"`
	FieldPosition int       `json:"fieldPosition"`
}

// MatchRepresentation defines user1 and user2 cards with positions,
// ball position at the moment and available actions for user cards.
type MatchRepresentation struct {
	User1CardsWithPosition []int                 `json:"user1CardsWithPosition"`
	User2CardsWithPosition []int                 `json:"user2CardsWithPosition"`
	BallPosition           int                   `json:"ballPosition"`
	Actions                []MakeAction          `json:"actions"`
	CardAvailableAction    []CardAvailableAction `json:"cardAvailableAction"`
}

// CardWithPosition defines card with position in the field.
type CardWithPosition struct {
	cards.Card    `json:"card"`
	FieldPosition int `json:"fieldPosition"`
}

// MakeAction defines fields that describes football action.
type MakeAction struct {
	CardsLayout       []CardWithPosition `json:"cardsLayout"`
	BallPosition      int                `json:"ballPosition"`
	PlayerID          uuid.UUID          `json:"playerId"`
	Action            Action             `json:"action"`
	ReceiverPlayerID  uuid.UUID          `json:"receiverPlayerId"`
	OpponentPlayerIDs []uuid.UUID        `json:"opponentPlayerIds"`
	ActionTime        time.Time          `json:"actionTime"`
}
