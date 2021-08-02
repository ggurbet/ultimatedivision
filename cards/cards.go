// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
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
type Filters map[Filter]string

// Filter defines the list of possible filters.
type Filter string

const (
	// Tactics indicates an assessment of the card's tactics.
	Tactics Filter = "tactics"
	// MinPhysique indicates an assessment of the card's minimum physique.
	MinPhysique Filter = "min_physique"
	// MaxPhysique indicates an assessment of the card's maximum physique.
	MaxPhysique Filter = "max_physique"
	// Physique indicates an assessment of the card's physique.
	Physique Filter = "physique"
	// PlayerName indicates the name of the card player name.
	PlayerName Filter = "player_name"
)

// SliceFilters entity for using group filter cards.
type SliceFilters []Filters

// Add check is empty and append value to slice.
func (s *SliceFilters) Add(name Filter, value string) {
	if value == "" {
		return
	}

	filter := Filters{
		name: value,
	}
	*s = append(*s, filter)
}

// Validate check of valid UTF-8 bytes and type.
func (f Filters) Validate() error {
	if _, found := f[Tactics]; found == true {
		strings.ToValidUTF8(f[Tactics], "")

		_, err := strconv.Atoi(f[Tactics])
		if err != nil {
			return ErrInvalidFilter.Wrap(fmt.Errorf("%s %s", f[Tactics], err))
		}
	}

	if _, found := f[MinPhysique]; found == true {
		strings.ToValidUTF8(f[MinPhysique], "")

		_, err := strconv.Atoi(f[MinPhysique])
		if err != nil {
			return ErrInvalidFilter.Wrap(fmt.Errorf("%s %s", f[MinPhysique], err))
		}
	}

	if _, found := f[MaxPhysique]; found == true {
		strings.ToValidUTF8(f[MaxPhysique], "")

		_, err := strconv.Atoi(f[MaxPhysique])
		if err != nil {
			return ErrInvalidFilter.Wrap(fmt.Errorf("%s %s", f[MaxPhysique], err))
		}
	}

	if _, found := f[PlayerName]; found == true {
		strings.ToValidUTF8(f[PlayerName], "")

		_, err := strconv.Atoi(f[PlayerName])
		if err == nil {
			return ErrInvalidFilter.Wrap(fmt.Errorf("%s %s", f[PlayerName], err))
		}
	}

	return nil
}
