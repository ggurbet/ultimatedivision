// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoClub indicated that club does not exist.
var ErrNoClub = errs.Class("club does not exist")

// ErrNoSquad indicated that squad does not exist.
var ErrNoSquad = errs.Class("squad does not exist")

// DB is exposing access to clubs db.
//
// architecture: DB
type DB interface {
	// Create creates club in the database.
	Create(ctx context.Context, club Club) error
	// CreateSquad creates squad for clubs in the database.
	CreateSquad(ctx context.Context, squad Squads) error
	// List returns all the clubs owned by the user.
	List(ctx context.Context, userID uuid.UUID) ([]Club, error)
	// GetSquad returns squad.
	GetSquad(ctx context.Context, squadID uuid.UUID) (Squads, error)
	// GetCapitan returns id of capitan.
	GetCapitan(ctx context.Context, squadID uuid.UUID) (uuid.UUID, error)
	// ListSquadCards returns all cards from squad.
	ListSquadCards(ctx context.Context, squadID uuid.UUID) ([]SquadCards, error)
	// Add add new card to the squad.
	Add(ctx context.Context, squadCards SquadCards) error
	// UpdateTacticFormation updates tactic and formation in the squad.
	UpdateTacticFormation(ctx context.Context, squad Squads) error
	// UpdateCapitan updates capitan in the squad.
	UpdateCapitan(ctx context.Context, capitan uuid.UUID, squadID uuid.UUID) error
	// UpdatePosition updates position of card in the squad.
	UpdatePosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, newPosition Position) error
}

// Club defines club entity.
type Club struct {
	ID        uuid.UUID `json:"id"`
	OwnerID   uuid.UUID `json:"ownerId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Squads describes squads of clubs.
type Squads struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ClubID    uuid.UUID `json:"clubId"`
	Formation Formation `json:"formation"`
	Tactic    Tactic    `json:"tactic"`
}

// SquadCards defines all cards from squad.
type SquadCards struct {
	ID       uuid.UUID `json:"id"`
	CardID   uuid.UUID `json:"cardId"`
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
	// Balanced balance between attack and defense.
	Balanced Tactic = 3
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
