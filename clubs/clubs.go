// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ErrNoClub indicated that club does not exist.
var ErrNoClub = errs.Class("club does not exist")

// ErrNoPlayer indicated that player does not exist.
var ErrNoPlayer = errs.Class("players does not exist")

// DB is exposing access to clubs db.
//
// architecture: DB
type DB interface {
	// Create creates club in the database.
	Create(ctx context.Context, club Club) error
	// Update updates club in the database.
	Update(ctx context.Context, club Club) error
	// GetClub gets team from database by user id.
	GetClub(ctx context.Context, userID uuid.UUID) (Club, error)
	// ListCards returns all cards from club.
	ListCards(ctx context.Context, userID uuid.UUID) ([]Player, error)
	// Add add new card to the team.
	Add(ctx context.Context, userID uuid.UUID, card cards.Card, capitan uuid.UUID, position Position) error
	// UpdateCapitan updates capitan in the team.
	UpdateCapitan(ctx context.Context, capitan uuid.UUID, userID uuid.UUID) error
	// GetCapitan returns id of clubs capitan.
	GetCapitan(ctx context.Context, userID uuid.UUID) (uuid.UUID, error)
}

// Club defines club entity.
type Club struct {
	UserID    uuid.UUID `json:"userID"`
	Formation Formation `json:"formation"`
	Tactic    Tactic    `json:"tactic"`
}

// Player defines card entity which in top11.
type Player struct {
	UserID   uuid.UUID `json:"userID"`
	CardID   uuid.UUID `json:"cardID"`
	Position Position  `json:"position"`
	Capitan  uuid.UUID `json:"capitan"`
}

// Formation defines a list of possible formations.
type Formation int

// TODO: add others formations.
const (
	// FourFourTwo defines 4-4-2 scheme.
	FourFourTwo Formation = 1
	// FourTwoFour defines 4-2-4 scheme.
	FourTwoFour Formation = 2
	// FourTwoTwoTwo defines 4-2-2-2 scheme.
	FourTwoTwoTwo Formation = 3
)

// Tactic defines a list of possible tactics.
type Tactic int

const (
	// Attack defines attacking style.
	Attack Tactic = 1
	// Defence defines defensive style.
	Defence Tactic = 2
	// Regular balance between attack and defense.
	Regular Tactic = 3
)

// Position defines a list of possible positions.
type Position int

const (
	// GK defines goalkeeper.
	GK Position = 1
	// CB defines central defenders.
	CB Position = 2
	// LB defines left defenders.
	LB Position = 3
	// RB defines right defenders.
	RB Position = 4
	// CM defines central midfielder.
	CM Position = 5
	// LM defines left midfielder.
	LM Position = 6
	// RM defines right midfielder.
	RM Position = 7
	// CAM defines central attacking midfielder.
	CAM Position = 8
	// LWB defines left attacking defenders.
	LWB Position = 9
	// RWB defines right attacking defenders.
	RWB Position = 10
	// RW defines right forward.
	RW Position = 11
	// LW defines left forward.
	LW Position = 12
	// ST defines central forward.
	ST Position = 13
)
