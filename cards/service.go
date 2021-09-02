// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Service is handling cards related logic.
//
// architecture: Service
type Service struct {
	cards  DB
	config Config
}

// NewService is a constructor for cards service.
func NewService(cards DB, config Config) *Service {
	return &Service{
		cards:  cards,
		config: config,
	}
}

// Create add card in DB.
func (service *Service) Create(ctx context.Context, userID uuid.UUID, percentageQualities []int) (Card, error) {
	qualities := map[string]int{
		"wood":    percentageQualities[0],
		"silver":  percentageQualities[1],
		"gold":    percentageQualities[2],
		"diamond": percentageQualities[3],
	}

	minHeight := service.config.Height.Min
	maxHeight := service.config.Height.Max
	minWeight := service.config.Weight.Min
	maxWeight := service.config.Weight.Max

	var skills = map[string]map[string]int{
		"wood": {
			"elementary":  service.config.Skills.Wood.Elementary,
			"basic":       service.config.Skills.Wood.Basic,
			"medium":      service.config.Skills.Wood.Medium,
			"upperMedium": service.config.Skills.Wood.UpperMedium,
			"advanced":    service.config.Skills.Wood.Advanced,
		},
		"silver": {
			"elementary":  service.config.Skills.Silver.Elementary,
			"basic":       service.config.Skills.Silver.Basic,
			"medium":      service.config.Skills.Silver.Medium,
			"upperMedium": service.config.Skills.Silver.UpperMedium,
			"advanced":    service.config.Skills.Silver.Advanced,
		},
		"gold": {
			"elementary":    service.config.Skills.Gold.Elementary,
			"basic":         service.config.Skills.Gold.Basic,
			"medium":        service.config.Skills.Gold.Medium,
			"upperMedium":   service.config.Skills.Gold.UpperMedium,
			"advanced":      service.config.Skills.Gold.Advanced,
			"upperAdvanced": service.config.Skills.Gold.UpperMedium,
		},
		"diamond": {
			"basic":         service.config.Skills.Diamond.Basic,
			"medium":        service.config.Skills.Diamond.Medium,
			"upperMedium":   service.config.Skills.Diamond.UpperMedium,
			"advanced":      service.config.Skills.Diamond.Advanced,
			"upperAdvanced": service.config.Skills.Diamond.UpperAdvanced,
		},
	}

	RangeValueForSkills = map[string][]int{
		"elementary":    {service.config.RangeValueForSkills.MinElementary, service.config.RangeValueForSkills.MaxElementary},
		"basic":         {service.config.RangeValueForSkills.MinBasic, service.config.RangeValueForSkills.MaxBasic},
		"medium":        {service.config.RangeValueForSkills.MinMedium, service.config.RangeValueForSkills.MaxMedium},
		"upperMedium":   {service.config.RangeValueForSkills.MinUpperMedium, service.config.RangeValueForSkills.MaxUpperMedium},
		"advanced":      {service.config.RangeValueForSkills.MinAdvanced, service.config.RangeValueForSkills.MaxAdvanced},
		"upperAdvanced": {service.config.RangeValueForSkills.MinUpperAdvanced, service.config.RangeValueForSkills.MaxUpperAdvanced},
	}

	var dominantFoots = map[string]int{
		"left":  service.config.DominantFoots.Left,
		"right": service.config.DominantFoots.Right,
	}

	var isTattoos bool
	var tattoos = map[string]int{
		"gold":    service.config.Tattoos.Gold,
		"diamond": service.config.Tattoos.Diamond,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	quality := searchValueByPercent(qualities)
	tactics := generateGroupSkill(skills[quality])
	physique := generateGroupSkill(skills[quality])
	technique := generateGroupSkill(skills[quality])
	offense := generateGroupSkill(skills[quality])
	defence := generateGroupSkill(skills[quality])
	goalkeeping := generateGroupSkill(skills[quality])

	if result := searchValueByPercent(tattoos); result != "" {
		isTattoos = true
	}

	card := Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro",
		Quality:          Quality(quality),
		PictureType:      1,
		Height:           round(rand.Float64()*(maxHeight-minHeight)+minHeight, 0.01),
		Weight:           round(rand.Float64()*(maxWeight-minWeight)+minWeight, 0.01),
		SkinColor:        1,
		HairStyle:        1,
		HairColor:        1,
		Accessories:      []int{1, 2},
		DominantFoot:     DominantFoot(searchValueByPercent(dominantFoots)),
		IsTattoos:        isTattoos,
		Status:           StatusActive,
		Type:             TypeWon,
		UserID:           userID,
		Tactics:          tactics,
		Positioning:      generateSkill(tactics),
		Composure:        generateSkill(tactics),
		Aggression:       generateSkill(tactics),
		Vision:           generateSkill(tactics),
		Awareness:        generateSkill(tactics),
		Crosses:          generateSkill(tactics),
		Physique:         physique,
		Acceleration:     generateSkill(physique),
		RunningSpeed:     generateSkill(physique),
		ReactionSpeed:    generateSkill(physique),
		Agility:          generateSkill(physique),
		Stamina:          generateSkill(physique),
		Strength:         generateSkill(physique),
		Jumping:          generateSkill(physique),
		Balance:          generateSkill(physique),
		Technique:        technique,
		Dribbling:        generateSkill(technique),
		BallControl:      generateSkill(technique),
		WeakFoot:         generateSkill(technique),
		SkillMoves:       generateSkill(technique),
		Finesse:          generateSkill(technique),
		Curve:            generateSkill(technique),
		Volleys:          generateSkill(technique),
		ShortPassing:     generateSkill(technique),
		LongPassing:      generateSkill(technique),
		ForwardPass:      generateSkill(technique),
		Offense:          offense,
		FinishingAbility: generateSkill(offense),
		ShotPower:        generateSkill(offense),
		Accuracy:         generateSkill(offense),
		Distance:         generateSkill(offense),
		Penalty:          generateSkill(offense),
		FreeKicks:        generateSkill(offense),
		Corners:          generateSkill(offense),
		HeadingAccuracy:  generateSkill(offense),
		Defence:          defence,
		OffsideTrap:      generateSkill(defence),
		Sliding:          generateSkill(defence),
		Tackles:          generateSkill(defence),
		BallFocus:        generateSkill(defence),
		Interceptions:    generateSkill(defence),
		Vigilance:        generateSkill(defence),
		Goalkeeping:      goalkeeping,
		Reflexes:         generateSkill(goalkeeping),
		Diving:           generateSkill(goalkeeping),
		Handling:         generateSkill(goalkeeping),
		Sweeping:         generateSkill(goalkeeping),
		Throwing:         generateSkill(goalkeeping),
	}
	return card, ErrCards.Wrap(service.cards.Create(ctx, card))
}

// searchValueByPercent search value string by percent.
func searchValueByPercent(generateMap map[string]int) string {
	rand := rand.Intn(99) + 1
	var sum int

	for k, v := range generateMap {
		sum += v
		if rand <= sum {
			return k
		}
	}
	return ""
}

// generateGroupSkill search value string by percent and generate assessment in the appropriate range.
func generateGroupSkill(generateMap map[string]int) int {
	skillValue := RangeValueForSkills[searchValueByPercent(generateMap)]
	difference := skillValue[1] - skillValue[0]
	rand := rand.Intn(difference) + 1
	return skillValue[0] + rand
}

// generateSkill generate assessment in the range +-10.
func generateSkill(value int) int {
	rand := rand.Intn(20) - 10
	result := value + rand
	if result < 1 {
		result = 1
	} else if result > 100 {
		result = 100
	}
	return result
}

// round rounds float64 the specified range.
func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Get returns card from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Card, error) {
	card, err := service.cards.Get(ctx, cardID)
	return card, ErrCards.Wrap(err)
}

// List returns all cards from DB.
func (service *Service) List(ctx context.Context) ([]Card, error) {
	cards, err := service.cards.List(ctx)
	return cards, ErrCards.Wrap(err)
}

// ListWithFilters returns all cards from DB, taking the necessary filters.
func (service *Service) ListWithFilters(ctx context.Context, filters []Filters) ([]Card, error) {
	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return nil, err
		}
	}
	cards, err := service.cards.ListWithFilters(ctx, filters)
	return cards, ErrCards.Wrap(err)
}

// ListByPlayerName returns cards from DB by player name.
func (service *Service) ListByPlayerName(ctx context.Context, filter Filters) ([]Card, error) {
	strings.ToValidUTF8(filter.Value, "")

	// TODO: add best check
	_, err := strconv.Atoi(filter.Value)
	if err == nil {
		return nil, ErrInvalidFilter.New("%s %s", filter.Value, err)
	}

	cards, err := service.cards.ListByPlayerName(ctx, filter)
	return cards, ErrCards.Wrap(err)
}

// UpdateStatus updates card status.
func (service *Service) UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error {
	return ErrCards.Wrap(service.cards.UpdateStatus(ctx, id, status))
}

// UpdateUserID updates card status.
func (service *Service) UpdateUserID(ctx context.Context, id, userID uuid.UUID) error {
	return ErrCards.Wrap(service.cards.UpdateUserID(ctx, id, userID))
}

// Delete destroy card in DB.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return ErrCards.Wrap(service.cards.Delete(ctx, cardID))
}
