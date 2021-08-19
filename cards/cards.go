// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/sqlsearchoperators"
)

// ErrNoCard indicated that card does not exist.
var ErrNoCard = errs.Class("card does not exist")

// ErrInvalidFilter indicated that filter does not valid.
var ErrInvalidFilter = errs.Class("invalid filter")

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
	// ListWithFilters returns all cards from the data base with filters.
	ListWithFilters(ctx context.Context, filters []Filters) ([]Card, error)
	// ListByPlayerName returns cards from DB by player name.
	ListByPlayerName(ctx context.Context, filters Filters) ([]Card, error)
	// UpdateStatus updates status card in the database.
	UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error
	// UpdateUserID updates user id card in the database.
	UpdateUserID(ctx context.Context, id, userID uuid.UUID) error
	// Delete deletes card record in the data base.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Card describes card entity.
type Card struct {
	ID               uuid.UUID    `json:"id"`
	PlayerName       string       `json:"playerName"`
	Quality          Quality      `json:"quality"`
	PictureType      int          `json:"pictureType"`
	Height           float64      `json:"height"`
	Weight           float64      `json:"weight"`
	SkinColor        int          `json:"skinColor"`
	HairStyle        int          `json:"hairStyle"`
	HairColor        int          `json:"hairColor"`
	Accessories      []int        `json:"accessories"`
	DominantFoot     DominantFoot `json:"dominantFoot"`
	IsTattoos        bool         `json:"isTattoos"`
	Status           Status       `json:"status"`
	Type             Type         `json:"type"`
	UserID           uuid.UUID    `json:"userId"`
	Tactics          int          `json:"tactics"`
	Positioning      int          `json:"positioning"`
	Composure        int          `json:"composure"`
	Aggression       int          `json:"aggression"`
	Vision           int          `json:"vision"`
	Awareness        int          `json:"awareness"`
	Crosses          int          `json:"crosses"`
	Physique         int          `json:"physique"`
	Acceleration     int          `json:"acceleration"`
	RunningSpeed     int          `json:"runningSpeed"`
	ReactionSpeed    int          `json:"reactionSpeed"`
	Agility          int          `json:"agility"`
	Stamina          int          `json:"stamina"`
	Strength         int          `json:"strength"`
	Jumping          int          `json:"jumping"`
	Balance          int          `json:"balance"`
	Technique        int          `json:"technique"`
	Dribbling        int          `json:"dribbling"`
	BallControl      int          `json:"ballControl"`
	WeakFoot         int          `json:"weakFoot"`
	SkillMoves       int          `json:"skillMoves"`
	Finesse          int          `json:"finesse"`
	Curve            int          `json:"curve"`
	Volleys          int          `json:"volleys"`
	ShortPassing     int          `json:"shortPassing"`
	LongPassing      int          `json:"longPassing"`
	ForwardPass      int          `json:"forwardPass"`
	Offense          int          `json:"offense"`
	FinishingAbility int          `json:"finishingAbility"`
	ShotPower        int          `json:"shotPower"`
	Accuracy         int          `json:"accuracy"`
	Distance         int          `json:"distance"`
	Penalty          int          `json:"penalty"`
	FreeKicks        int          `json:"freeKicks"`
	Corners          int          `json:"corners"`
	HeadingAccuracy  int          `json:"headingAccuracy"`
	Defence          int          `json:"defence"`
	OffsideTrap      int          `json:"offsideTrap"`
	Sliding          int          `json:"sliding"`
	Tackles          int          `json:"tackles"`
	BallFocus        int          `json:"ballFocus"`
	Interceptions    int          `json:"interceptions"`
	Vigilance        int          `json:"vigilance"`
	Goalkeeping      int          `json:"goalkeeping"`
	Reflexes         int          `json:"reflexes"`
	Diving           int          `json:"diving"`
	Handling         int          `json:"handling"`
	Sweeping         int          `json:"sweeping"`
	Throwing         int          `json:"throwing"`
}

// Quality defines the list of possible card qualities.
type Quality string

const (
	// QualityWood indicates that card quality is wood.
	QualityWood Quality = "wood"
	// QualitySilver indicates that card quality is silver.
	QualitySilver Quality = "silver"
	// QualityGold indicates that card quality is gold.
	QualityGold Quality = "gold"
	// QualityDiamond indicates that card quality is diamond.
	QualityDiamond Quality = "diamond"
)

// PictureType defines the list of possible card picture types.
var PictureType = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// SkinColor defines the list of possible card skin colors.
var SkinColor = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// HairStyle defines the list of possible card hairstyles.
var HairStyle = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// HairColor defines the list of possible card hair colors.
var HairColor = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// Accessory defines the list of possible card accessories.
var Accessory = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// DominantFoot defines the list of possible card dominant foots.
type DominantFoot string

const (
	// DominantFootLeft indicates that dominant foot of the footballer is left.
	DominantFootLeft DominantFoot = "left"
	// DominantFootRight indicates that dominant foot of the footballer is right.
	DominantFootRight DominantFoot = "right"
)

// Status defines the list of possible card statuses.
type Status int

const (
	// StatusActive indicates that the card can be used in a team and sold.
	StatusActive Status = 0
	// StatusSale indicates that the card is sold and can't used by the team.
	StatusSale Status = 1
)

// Type defines the list of possible card Typees.
type Type string

const (
	// TypeWon indicates that the card won in a lootbox.
	TypeWon Type = "won"
	// TypeBought indicates that the card bought on the marketplaced.
	TypeBought Type = "bought"
)

// RangeValueForSkills defines the list of possible group skills.
var RangeValueForSkills = map[string][]int{}

// Config defines values needed by generate cards.
type Config struct {
	Height struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"height"`

	Weight struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"weight"`

	DominantFoots struct {
		Left  int `json:"left"`
		Right int `json:"right"`
	} `json:"dominantFoots"`

	Skills struct {
		Wood struct {
			Elementary  int `json:"elementary"`
			Basic       int `json:"basic"`
			Medium      int `json:"medium"`
			UpperMedium int `json:"upperMedium"`
			Advanced    int `json:"advanced"`
		} `json:"wood"`
		Silver struct {
			Elementary  int `json:"elementary"`
			Basic       int `json:"basic"`
			Medium      int `json:"medium"`
			UpperMedium int `json:"upperMedium"`
			Advanced    int `json:"advanced"`
		} `json:"silver"`
		Gold struct {
			Elementary    int `json:"elementary"`
			Basic         int `json:"basic"`
			Medium        int `json:"medium"`
			UpperMedium   int `json:"upperMedium"`
			Advanced      int `json:"advanced"`
			UpperAdvanced int `json:"upperAdvanced"`
		} `json:"gold"`
		Diamond struct {
			Basic         int `json:"basic"`
			Medium        int `json:"medium"`
			UpperMedium   int `json:"upperMedium"`
			Advanced      int `json:"advanced"`
			UpperAdvanced int `json:"upperAdvanced"`
		} `json:"diamond"`
	} `json:"skills"`

	RangeValueForSkills struct {
		MinElementary    int `json:"minElementary"`
		MaxElementary    int `json:"maxElementary"`
		MinBasic         int `json:"minBasic"`
		MaxBasic         int `json:"maxBasic"`
		MinMedium        int `json:"minMedium"`
		MaxMedium        int `json:"maxMedium"`
		MinUpperMedium   int `json:"minUpperMedium"`
		MaxUpperMedium   int `json:"maxUpperMedium"`
		MinAdvanced      int `json:"minAdvanced"`
		MaxAdvanced      int `json:"maxAdvanced"`
		MinUpperAdvanced int `json:"minUpperAdvanced"`
		MaxUpperAdvanced int `json:"maxUpperAdvanced"`
	} `json:"rangeValueForSkills"`

	Tattoos struct {
		Gold    int `json:"gold"`
		Diamond int `json:"diamond"`
	} `json:"tattoos"`
}

// PercentageQualities entity for probabilities generate cards.
type PercentageQualities struct {
	Wood    int `json:"wood"`
	Silver  int `json:"silver"`
	Gold    int `json:"gold"`
	Diamond int `json:"diamond"`
}

// Filters entity for using filter cards.
type Filters struct {
	Name           Filter
	Value          string
	SearchOperator sqlsearchoperators.SearchOperator
}

// Filter defines the list of possible filters.
type Filter string

const (
	// FilterTactics indicates an tactics of the card.
	FilterTactics Filter = "tactics"
	// FilterPositioning indicates an positioning of the card.
	FilterPositioning Filter = "positioning"
	// FilterComposure indicates an composure of the card.
	FilterComposure Filter = "composure"
	// FilterAggression indicates an aggression of the card.
	FilterAggression Filter = "aggression"
	// FilterVision indicates an vision of the card.
	FilterVision Filter = "vision"
	// FilterAwareness indicates an awareness of the card.
	FilterAwareness Filter = "awareness"
	// FilterCrosses indicates an crosses of the card.
	FilterCrosses Filter = "crosses"
	// FilterPhysique indicates an physique of the card.
	FilterPhysique Filter = "physique"
	// FilterAcceleration indicates an acceleration of the card.
	FilterAcceleration Filter = "acceleration"
	// FilterRunningSpeed indicates an runningSpeed of the card.
	FilterRunningSpeed Filter = "runningSpeed"
	// FilterReactionSpeed indicates an reactionSpeed of the card.
	FilterReactionSpeed Filter = "reactionSpeed"
	// FilterAgility indicates an agility of the card.
	FilterAgility Filter = "agility"
	// FilterStamina indicates an stamina of the card.
	FilterStamina Filter = "stamina"
	// FilterStrength indicates an strength of the card.
	FilterStrength Filter = "strength"
	// FilterJumping indicates an jumping of the card.
	FilterJumping Filter = "jumping"
	// FilterBalance indicates an balance of the card.
	FilterBalance Filter = "balance"
	// FilterTechnique indicates an technique of the card.
	FilterTechnique Filter = "technique"
	// FilterDribbling indicates an dribbling of the card.
	FilterDribbling Filter = "dribbling"
	// FilterBallControl indicates an ballControl of the card.
	FilterBallControl Filter = "ballControl"
	// FilterWeakFoot indicates an weakFoot of the card.
	FilterWeakFoot Filter = "weakFoot"
	// FilterSkillMoves indicates an skillMoves of the card.
	FilterSkillMoves Filter = "skillMoves"
	// FilterFinesse indicates an finesse of the card.
	FilterFinesse Filter = "finesse"
	// FilterCurve indicates an curve of the card.
	FilterCurve Filter = "curve"
	// FilterVolleys indicates an volleys of the card.
	FilterVolleys Filter = "volleys"
	// FilterShortPassing indicates an shortPassing of the card.
	FilterShortPassing Filter = "shortPassing"
	// FilterLongPassing indicates an longPassing of the card.
	FilterLongPassing Filter = "longPassing"
	// FilterForwardPass indicates an forwardPass of the card.
	FilterForwardPass Filter = "forwardPass"
	// FilterOffense indicates an offense of the card.
	FilterOffense Filter = "offense"
	// FilterFinishingAbility indicates an finishingAbility of the card.
	FilterFinishingAbility Filter = "finishingAbility"
	// FilterShotPower indicates an shotPower of the card.
	FilterShotPower Filter = "shotPower"
	// FilterAccuracy indicates an accuracy of the card.
	FilterAccuracy Filter = "accuracy"
	// FilterDistance indicates an distance of the card.
	FilterDistance Filter = "distance"
	// FilterPenalty indicates an penalty of the card.
	FilterPenalty Filter = "penalty"
	// FilterFreeKicks indicates an freeKicks of the card.
	FilterFreeKicks Filter = "freeKicks"
	// FilterCorners indicates an corners of the card.
	FilterCorners Filter = "corners"
	// FilterHeadingAccuracy indicates an headingAccuracy of the card.
	FilterHeadingAccuracy Filter = "headingAccuracy"
	// FilterDefence indicates an defence of the card.
	FilterDefence Filter = "defence"
	// FilterOffsideTrap indicates an offsideTrap of the card.
	FilterOffsideTrap Filter = "offsideTrap"
	// FilterSliding indicates an sliding of the card.
	FilterSliding Filter = "sliding"
	// FilterTackles indicates an tackles of the card.
	FilterTackles Filter = "tackles"
	// FilterBallFocus indicates an ballFocus of the card.
	FilterBallFocus Filter = "ballFocus"
	// FilterInterceptions indicates an interceptions of the card.
	FilterInterceptions Filter = "interceptions"
	// FilterVigilance indicates an vigilance of the card.
	FilterVigilance Filter = "vigilance"
	// FilterGoalkeeping indicates an goalkeeping of the card.
	FilterGoalkeeping Filter = "goalkeeping"
	// FilterReflexes indicates an reflexes of the card.
	FilterReflexes Filter = "reflexes"
	// FilterDiving indicates an diving of the card.
	FilterDiving Filter = "diving"
	// FilterHandling indicates an handling of the card.
	FilterHandling Filter = "handling"
	// FilterSweeping indicates an sweeping of the card.
	FilterSweeping Filter = "sweeping"
	// FilterThrowing indicates an throwing of the card.
	FilterThrowing Filter = "throwing"
	// FilterQuality indicates an quality of the card.
	FilterQuality Filter = "quality"
	// FilterHeight indicates an height of the card.
	FilterHeight Filter = "height"
	// FilterWeight indicates an weight of the card.
	FilterWeight Filter = "weight"
	// FilterDominantFoot indicates an dominant foot of the card.
	FilterDominantFoot Filter = "dominantFoot"
	// FilterType indicates an type of the card.
	FilterType Filter = "type"
	// FilterPlayerName indicates the name of the card player name.
	FilterPlayerName Filter = "player_name"
)

// Validate check of valid UTF-8 bytes and type.
func (f Filters) Validate() error {
	if f.Name == FilterTactics || f.Name == FilterPositioning || f.Name == FilterComposure || f.Name == FilterAggression ||
		f.Name == FilterVision || f.Name == FilterAwareness || f.Name == FilterCrosses || f.Name == FilterPhysique ||
		f.Name == FilterAcceleration || f.Name == FilterRunningSpeed || f.Name == FilterReactionSpeed || f.Name == FilterAgility ||
		f.Name == FilterStamina || f.Name == FilterStrength || f.Name == FilterJumping || f.Name == FilterBalance ||
		f.Name == FilterTechnique || f.Name == FilterDribbling || f.Name == FilterBallControl || f.Name == FilterWeakFoot ||
		f.Name == FilterSkillMoves || f.Name == FilterFinesse || f.Name == FilterCurve || f.Name == FilterVolleys ||
		f.Name == FilterShortPassing || f.Name == FilterLongPassing || f.Name == FilterForwardPass || f.Name == FilterOffense ||
		f.Name == FilterFinishingAbility || f.Name == FilterShotPower || f.Name == FilterAccuracy || f.Name == FilterDistance ||
		f.Name == FilterPenalty || f.Name == FilterFreeKicks || f.Name == FilterCorners || f.Name == FilterHeadingAccuracy ||
		f.Name == FilterDefence || f.Name == FilterOffsideTrap || f.Name == FilterSliding || f.Name == FilterTackles ||
		f.Name == FilterBallFocus || f.Name == FilterInterceptions || f.Name == FilterVigilance || f.Name == FilterGoalkeeping ||
		f.Name == FilterReflexes || f.Name == FilterDiving || f.Name == FilterHandling || f.Name == FilterSweeping || f.Name == FilterThrowing {
		strings.ToValidUTF8(f.Value, "")

		_, err := strconv.Atoi(f.Value)
		if err != nil {
			return ErrInvalidFilter.New("%s %s", f.Value, err)
		}
		return nil
	}

	if f.Name == FilterHeight || f.Name == FilterWeight {
		strings.ToValidUTF8(f.Value, "")

		_, err := strconv.ParseFloat(f.Value, 64)
		if err != nil {
			return ErrInvalidFilter.New("%s %s", f.Value, err)
		}
		return nil
	}

	if f.Name == FilterQuality {
		strings.ToValidUTF8(f.Value, "")

		quality := Quality(f.Value)
		if quality == QualityWood || quality == QualitySilver || quality == QualityGold || quality == QualityDiamond {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, " is not an indicator of quality card")
	}

	if f.Name == FilterDominantFoot {
		strings.ToValidUTF8(f.Value, "")

		dominantFoot := DominantFoot(f.Value)
		if dominantFoot == DominantFootLeft || dominantFoot == DominantFootRight {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, " is not an indicator of dominant foot card")
	}

	if f.Name == FilterType {
		strings.ToValidUTF8(f.Value, "")

		filterType := Type(f.Value)
		if filterType == TypeWon || filterType == TypeBought {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, " is not an indicator of type card")
	}

	return ErrInvalidFilter.New("invalid name parameter - " + string(f.Name))
}
