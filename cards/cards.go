// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoCard indicated that card does not exist.
var ErrNoCard = errs.Class("card does not exist")

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create add card in the data base.
	Create(ctx context.Context, card Card) error
	// Get returns card by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Card, error)
	// List returns all cards from the data base.
	List(ctx context.Context) ([]Card, error)
}

// Card describes card entity.
type Card struct {
	ID               uuid.UUID     `json:"id"`
	PlayerName       string        `json:"playerName"`
	Quality          Quality       `json:"quality"`
	PictureType      PictureType   `json:"pictureType"`
	Height           float32       `json:"height"`
	Weight           float32       `json:"weight"`
	SkinColor        SkinColor     `json:"skinColor"`
	HairStyle        HairStyle     `json:"hairStyle"`
	HairColor        HairColor     `json:"hairColor"`
	Accessories      []Accessories `json:"accessories"`
	DominantFoot     DominantFoot  `json:"dominantFoot"`
	UserID           uuid.UUID     `json:"userId"`
	Positioning      int           `json:"positioning"`
	Composure        int           `json:"composure"`
	Aggression       int           `json:"aggression"`
	Vision           int           `json:"vision"`
	Awareness        int           `json:"awareness"`
	Crosses          int           `json:"crosses"`
	Acceleration     int           `json:"acceleration"`
	RunningSpeed     int           `json:"runningSpeed"`
	ReactionSpeed    int           `json:"reactionSpeed"`
	Agility          int           `json:"agility"`
	Stamina          int           `json:"stamina"`
	Strength         int           `json:"strength"`
	Jumping          int           `json:"jumping"`
	Balance          int           `json:"balance"`
	Dribbling        int           `json:"dribbling"`
	BallControl      int           `json:"ballControl"`
	WeakFoot         int           `json:"weakFoot"`
	SkillMoves       int           `json:"skillMoves"`
	Finesse          int           `json:"finesse"`
	Curve            int           `json:"curve"`
	Volleys          int           `json:"volleys"`
	ShortPassing     int           `json:"shortPassing"`
	LongPassing      int           `json:"longPassing"`
	ForwardPass      int           `json:"forwardPass"`
	FinishingAbility int           `json:"finishingAbility"`
	ShotPower        int           `json:"shotPower"`
	Accuracy         int           `json:"accuracy"`
	Distance         int           `json:"distance"`
	Penalty          int           `json:"penalty"`
	FreeKicks        int           `json:"freeKicks"`
	Corners          int           `json:"corners"`
	HeadingAccuracy  int           `json:"headingAccuracy"`
	OffsideTrap      int           `json:"offsideTrap"`
	Sliding          int           `json:"sliding"`
	Tackles          int           `json:"tackles"`
	BallFocus        int           `json:"ballFocus"`
	Interceptions    int           `json:"interceptions"`
	Vigilance        int           `json:"vigilance"`
	Reflexes         int           `json:"reflexes"`
	Diving           int           `json:"diving"`
	Handling         int           `json:"handling"`
	Sweeping         int           `json:"sweeping"`
	Throwing         int           `json:"throwing"`
}

// Quality defines the list of possible card qualities.
type Quality string

const (
	// QualityWood indicates that card quality is wood.
	QualityWood Quality = "wood"
	// QualityBronze indicates that card quality is bronze.
	QualityBronze Quality = "bronze"
	// QualitySilver indicates that card quality is silver.
	QualitySilver Quality = "silver"
	// QualityGold indicates that card quality is gold.
	QualityGold Quality = "gold"
	// QualityDiamond indicates that card quality is diamond.
	QualityDiamond Quality = "diamond"
)

// PictureType defines the list of possible card picture types.
type PictureType string

// SkinColor defines the list of possible card skin colors.
type SkinColor int

// HairStyle defines the list of possible card hairstyles.
type HairStyle int

// HairColor defines the list of possible card hair colors.
type HairColor int

// Accessories defines the list of possible card accessories.
type Accessories int

// DominantFoot defines the list of possible card dominant foots.
type DominantFoot string

const (
	// DominantFootLeft indicates that dominant foot of the footballer is left.
	DominantFootLeft DominantFoot = "left"
	// DominantFootRight indicates that dominant foot of the footballer is right.
	DominantFootRight DominantFoot = "right"
)
