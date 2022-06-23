// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/pagination"
)

// ErrNoCard indicated that card does not exist.
var ErrNoCard = errs.Class("card does not exist")

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create adds card in the data base.
	Create(ctx context.Context, card Card) error
	// Get returns card by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Card, error)
	// GetByPlayerName returns card by player name from DB.
	GetByPlayerName(ctx context.Context, playerName string) (Card, error)
	// List returns all cards from the data base.
	List(ctx context.Context, cursor pagination.Cursor) (Page, error)
	// ListByUserID returns cards by user id from the database.
	ListByUserID(ctx context.Context, id uuid.UUID, cursor pagination.Cursor) (Page, error)
	// ListByTypeUnordered returns cards where type is unordered from the database.
	ListByTypeUnordered(ctx context.Context) ([]Card, error)
	// ListWithFilters returns cards with filters from the database.
	ListWithFilters(ctx context.Context, filters []Filters, cursor pagination.Cursor) (Page, error)
	// ListCardIDsWithFiltersWhereActiveLot returns card ids where active lots from DB, taking the necessary filters.
	ListCardIDsWithFiltersWhereActiveLot(ctx context.Context, filters []Filters) ([]uuid.UUID, error)
	// ListByUserIDAndPlayerName returns cards from DB by user id and player name.
	ListByUserIDAndPlayerName(ctx context.Context, userID uuid.UUID, filters Filters, cursor pagination.Cursor) (Page, error)
	// ListCardIDsByPlayerNameWhereActiveLot returns card ids where active lot from DB by player name.
	ListCardIDsByPlayerNameWhereActiveLot(ctx context.Context, filter Filters) ([]uuid.UUID, error)
	// GetSquadCards returns all card with characteristics from the squad from the database.
	GetSquadCards(ctx context.Context, id uuid.UUID) ([]Card, error)
	// UpdateStatus updates status card in the database.
	UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error
	// UpdateStatus updates type of card in the database.
	UpdateType(ctx context.Context, id uuid.UUID, typeCard Type) error
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
	Height           float64      `json:"height"`
	Weight           float64      `json:"weight"`
	DominantFoot     DominantFoot `json:"dominantFoot"`
	IsTattoo         bool         `json:"isTattoo"`
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
	Offence          int          `json:"offence"`
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

// QualityToValue describes quality-to-value ratio.
var QualityToValue = map[Quality]int{
	QualityWood:    0,
	QualitySilver:  1,
	QualityGold:    2,
	QualityDiamond: 3,
}

// GetValueOfQuality returns value of card by key.
func (quality Quality) GetValueOfQuality() int {
	return QualityToValue[quality]
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

// Type defines the list of possible card types.
type Type string

const (
	// TypeWon indicates that the card won in a lootbox.
	TypeWon Type = "won"
	// TypeBought indicates that the card bought in the marketplace.
	TypeBought Type = "bought"
	// TypeUnordered indicates that the card unordered in the store.
	TypeUnordered Type = "unordered"
	// TypeOrdered indicates that the card ordered in the store.
	TypeOrdered Type = "ordered"
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

	pagination.Cursor `json:"cursor"`

	// CardEfficiencyParameters coefficients for calculating the efficiency of the card.
	CardEfficiencyParameters struct {
		GK struct {
			Goalkeeping float64 `json:"goalkeeping"`
			Physique    float64 `json:"physique"`
			Tactics     float64 `json:"tactics"`
		} `json:"gk"`
		CD struct {
			Defence  float64 `json:"defence"`
			Physique float64 `json:"physique"`
			Tactics  float64 `json:"tactics"`
		} `json:"cd"`
		LBorRB struct {
			Defence   float64 `json:"defence"`
			Physique  float64 `json:"physique"`
			Tactics   float64 `json:"tactics"`
			Technique float64 `json:"technique"`
		} `json:"lbOrRb"`
		CDM struct {
			Defence   float64 `json:"defence"`
			Physique  float64 `json:"physique"`
			Tactics   float64 `json:"tactics"`
			Technique float64 `json:"technique"`
			Offence   float64 `json:"offence"`
		} `json:"cdm"`
		CM struct {
			Defence   float64 `json:"defence"`
			Physique  float64 `json:"physique"`
			Tactics   float64 `json:"tactics"`
			Technique float64 `json:"technique"`
			Offence   float64 `json:"offence"`
		} `json:"cm"`
		CAM struct {
			Defence   float64 `json:"defence"`
			Physique  float64 `json:"physique"`
			Tactics   float64 `json:"tactics"`
			Technique float64 `json:"technique"`
			Offence   float64 `json:"offence"`
		} `json:"cam"`
		RMorLM struct {
			Defence   float64 `json:"defence"`
			Physique  float64 `json:"physique"`
			Tactics   float64 `json:"tactics"`
			Technique float64 `json:"technique"`
			Offence   float64 `json:"offence"`
		} `json:"rmOrLm"`
		RWorLW struct {
			Physique  float64 `json:"physique"`
			Tactics   float64 `json:"tactics"`
			Technique float64 `json:"technique"`
			Offence   float64 `json:"offence"`
		} `json:"rwOrLw"`
		ST struct {
			Physique  float64 `json:"physique"`
			Tactics   float64 `json:"tactics"`
			Technique float64 `json:"technique"`
			Offence   float64 `json:"offence"`
		} `json:"st"`
	} `json:"cardEfficiencyParameters"`
	PathToNamesDataset string `json:"pathToNamesDataset"`
}

// PercentageQualities entity for probabilities generate cards.
type PercentageQualities struct {
	Wood    int `json:"wood"`
	Silver  int `json:"silver"`
	Gold    int `json:"gold"`
	Diamond int `json:"diamond"`
}

// Page holds card page entity which is used to show listed page of cards.
type Page struct {
	Cards []Card          `json:"cards"`
	Page  pagination.Page `json:"page"`
}
