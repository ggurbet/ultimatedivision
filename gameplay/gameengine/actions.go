// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"github.com/google/uuid"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/clubs"
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

// Config contains config values related to game.
type Config struct {
	LeftSide  LeftSide  `json:"leftSide"`
	RightSide RightSide `json:"rightSide"`
	Rounds    int       `json:"rounds"`
}

// LeftSide contains config values of the left side team positions.
type LeftSide struct {
	Positions `json:"positions"`
}

// RightSide contains config values of the right side team positions.
type RightSide struct {
	Positions `json:"positions"`
}

// Positions contains config values of all players positions.
type Positions struct {
	Goalkeeper      int `json:"goalkeeper"`
	LeftBack        int `json:"leftBack"`
	CenterBackLeft  int `json:"centerBackLeft"`
	CenterBackRight int `json:"centerBackRight"`
	RightBack       int `json:"rightBack"`
	LeftMid         int `json:"leftMid"`
	CenterMidLeft   int `json:"centerMidLeft"`
	CenterMidRight  int `json:"centerMidRight"`
	RightMid        int `json:"rightMid"`
	ForwardLeft     int `json:"forwardLeft"`
	ForwardRight    int `json:"forwardRight"`
}

// CardAvailableAction defines in which position card could be placed and which action it could do there.
type CardAvailableAction struct {
	Action        Action    `json:"action"`
	CardID        uuid.UUID `json:"cardId"`
	FieldPosition []int     `json:"fieldPosition"`
}

// CardWithPosition defines card with position in the field.
type CardWithPosition struct {
	cards.Card     `json:"card"`
	avatars.Avatar `json:"avatar"`
	FieldPosition  int `json:"fieldPosition"`
}

// MatchRepresentation defines user1 and user2 cards with positions,
// ball position at the moment and available actions for user cards.
type MatchRepresentation struct {
	MatchID                uuid.UUID             `json:"matchId"`
	User1CardsWithPosition []CardWithPosition    `json:"user1CardsWithPosition"`
	User2CardsWithPosition []CardWithPosition    `json:"user2CardsWithPosition"`
	BallPosition           int                   `json:"ballPosition"`
	CardAvailableAction    []CardAvailableAction `json:"cardAvailableAction"`
	User1ClubInformation   clubs.Club            `json:"user1ClubInformation"`
	User2ClubInformation   clubs.Club            `json:"user2ClubInformation"`
	User1SquadInformation  clubs.Squad           `json:"user1SquadInformation"`
	User2SquadInformation  clubs.Squad           `json:"user2SquadInformation"`
	Rounds                 int                   `json:"rounds"`
	UserSide               int                   `json:"userSide"`
}

// ActionResult defines result of action.
type ActionResult struct {
	Message interface{} `json:"message"`
	CardIDWithPosition
	BallPosition int
	CardAvailableAction
}
