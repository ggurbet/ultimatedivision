// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"

	"github.com/google/uuid"
)

// Service is handling cards related logic.
//
// architecture: Service
type Service struct {
	cards DB
}

// NewService is a constructor for cards service.
func NewService(cards DB) *Service {
	return &Service{
		cards: cards,
	}
}

// Create add card in DB.
func (service *Service) Create(ctx context.Context, userID uuid.UUID) error {

	// TODO:logic generate card
	card := Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro",
		Quality:          "bronze",
		PictureType:      1,
		Height:           178.8,
		Weight:           72.2,
		SkinColor:        1,
		HairStyle:        1,
		HairColor:        1,
		Accessories:      []int{1, 2},
		DominantFoot:     "left",
		UserID:           userID,
		Tactics:          1,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offense:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
	}
	return service.cards.Create(ctx, card)
}

// Get returns card from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Card, error) {
	return service.cards.Get(ctx, cardID)
}

// List returns all cards from DB.
func (service *Service) List(ctx context.Context) ([]Card, error) {
	return service.cards.List(ctx)
}

// ListWithFilters returns all cards from DB, taking the necessary filters.
func (service *Service) ListWithFilters(ctx context.Context, filters []Filters) ([]Card, error) {
	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return nil, err
		}
	}
	return service.cards.ListWithFilters(ctx, filters)
}

// Delete destroy card in DB.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return service.cards.Delete(ctx, cardID)
}
