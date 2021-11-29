// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/fileutils"
	"ultimatedivision/pkg/pagination"
)

// ErrCards indicated that there was an error in service.
var ErrCards = errs.Class("cards service error")

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

// Create adds card in DB.
func (service *Service) Create(ctx context.Context, userID uuid.UUID, percentageQualities []int) (Card, error) {
	var (
		err  error
		card Card
	)

	if card, err = service.Generate(ctx, userID, percentageQualities); err != nil {
		return card, ErrCards.Wrap(err)
	}
	return card, service.cards.Create(ctx, card)
}

// Generate generates card.
func (service *Service) Generate(ctx context.Context, userID uuid.UUID, percentageQualities []int) (Card, error) {
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

	var isTattoo bool
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
		isTattoo = true
	}

	playerName, err := service.GeneratePlayerName()
	if err != nil {
		return Card{}, err
	}

	card := Card{
		ID:               uuid.New(),
		PlayerName:       playerName,
		Quality:          Quality(quality),
		Height:           round(rand.Float64()*(maxHeight-minHeight)+minHeight, 0.01),
		Weight:           round(rand.Float64()*(maxWeight-minWeight)+minWeight, 0.01),
		DominantFoot:     DominantFoot(searchValueByPercent(dominantFoots)),
		IsTattoo:         isTattoo,
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
		Offence:          offense,
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

	return card, nil
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

// GeneratePlayerName generates player name of card.
func (service *Service) GeneratePlayerName() (string, error) {
	var (
		fullName   string
		firstName  string
		secondName string
	)

	file, err := os.Open(service.config.PathToNamesDataset)
	if err != nil {
		return "", ErrCards.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, file.Close())
	}()

	rand.Seed(time.Now().UTC().UnixNano())

	totalCount, err := fileutils.CountLines(file)
	if err != nil {
		return "", ErrCards.Wrap(err)
	}

	randomNum := rand.Intn(totalCount) + 1

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", ErrCards.Wrap(err)
	}

	fullName, err = fileutils.ReadLine(file, randomNum)

	splitFullName := strings.Split(fullName, " ")
	if len(splitFullName) == 2 {
		firstName = splitFullName[0]
	}

	randomNum = rand.Intn(totalCount) + 1

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", ErrCards.Wrap(err)
	}

	fullName, err = fileutils.ReadLine(file, randomNum)

	splitFullName = strings.Split(fullName, " ")
	if len(splitFullName) == 2 {
		secondName = splitFullName[1]
	}

	fullName = firstName + " " + secondName
	return fullName, ErrCards.Wrap(err)
}

// Get returns card from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Card, error) {
	card, err := service.cards.Get(ctx, cardID)
	return card, ErrCards.Wrap(err)
}

// List returns all cards from DB.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	cardsListPage, err := service.cards.List(ctx, cursor)
	return cardsListPage, ErrCards.Wrap(err)
}

// ListWithFilters returns all cards from DB, taking the necessary filters.
func (service *Service) ListWithFilters(ctx context.Context, filters []Filters, cursor pagination.Cursor) (Page, error) {
	var cardsListPage Page

	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return cardsListPage, err
		}
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	cardsListPage, err := service.cards.ListWithFilters(ctx, filters, cursor)
	return cardsListPage, ErrCards.Wrap(err)
}

// ListCardIDsWithFiltersWhereActiveLot returns card ids where active lots from DB, taking the necessary filters.
func (service *Service) ListCardIDsWithFiltersWhereActiveLot(ctx context.Context, filters []Filters) ([]uuid.UUID, error) {
	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return nil, err
		}
	}
	cardsList, err := service.cards.ListCardIDsWithFiltersWhereActiveLot(ctx, filters)
	return cardsList, ErrCards.Wrap(err)
}

// ListByPlayerName returns cards from DB by player name.
func (service *Service) ListByPlayerName(ctx context.Context, filter Filters, cursor pagination.Cursor) (Page, error) {
	var cardsListPage Page
	strings.ToValidUTF8(filter.Value, "")

	// TODO: add best check
	_, err := strconv.Atoi(filter.Value)
	if err == nil {
		return cardsListPage, ErrInvalidFilter.New("%s %s", filter.Value, err)
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	cardsListPage, err = service.cards.ListByPlayerName(ctx, filter, cursor)
	return cardsListPage, ErrCards.Wrap(err)
}

// ListCardIDsByPlayerNameWhereActiveLot returns card ids where active lot from DB by player name.
func (service *Service) ListCardIDsByPlayerNameWhereActiveLot(ctx context.Context, filter Filters) ([]uuid.UUID, error) {
	strings.ToValidUTF8(filter.Value, "")

	// TODO: add best check
	_, err := strconv.Atoi(filter.Value)
	if err == nil {
		return nil, ErrInvalidFilter.New("%s %s", filter.Value, err)
	}
	cardIdsList, err := service.cards.ListCardIDsByPlayerNameWhereActiveLot(ctx, filter)
	return cardIdsList, ErrCards.Wrap(err)
}

// ListByUserID returns all user`s cards in database.
func (service *Service) ListByUserID(ctx context.Context, userID uuid.UUID) ([]Card, error) {
	userCards, err := service.cards.ListByUserID(ctx, userID)
	return userCards, ErrCards.Wrap(err)
}

// GetCardsFromSquadCards returns all card with characteristics from the squad.
func (service *Service) GetCardsFromSquadCards(ctx context.Context, id uuid.UUID) ([]Card, error) {
	cards, err := service.cards.GetSquadCards(ctx, id)

	return cards, ErrCards.Wrap(err)
}

// UpdateStatus updates status of card in database.
func (service *Service) UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error {
	return ErrCards.Wrap(service.cards.UpdateStatus(ctx, id, status))
}

// UpdateUserID updates user's id for card in database.
func (service *Service) UpdateUserID(ctx context.Context, id, userID uuid.UUID) error {
	return ErrCards.Wrap(service.cards.UpdateUserID(ctx, id, userID))
}

// Delete deletes card record in database.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return ErrCards.Wrap(service.cards.Delete(ctx, cardID))
}

// EffectivenessGK determines the effectiveness of the card in the GK position.
func (service *Service) EffectivenessGK(card Card) float64 {
	return service.config.CardEfficiencyParameters.GK.Goalkeeping*float64(card.Goalkeeping) +
		service.config.CardEfficiencyParameters.GK.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.GK.Tactics*float64(card.Tactics)
}

// EffectivenessCD determines the effectiveness of the card in the CD position.
func (service *Service) EffectivenessCD(card Card) float64 {
	return service.config.CardEfficiencyParameters.CD.Defence*float64(card.Defence) +
		service.config.CardEfficiencyParameters.CD.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.CD.Tactics*float64(card.Tactics)
}

// EffectivenessLBorRB determines the effectiveness of the card in the LB/RB position.
func (service *Service) EffectivenessLBorRB(card Card) float64 {
	return service.config.CardEfficiencyParameters.LBorRB.Defence*float64(card.Defence) +
		service.config.CardEfficiencyParameters.LBorRB.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.LBorRB.Tactics*float64(card.Tactics) +
		service.config.CardEfficiencyParameters.LBorRB.Technique*float64(card.Technique)
}

// EffectivenessCDM determines the effectiveness of the card in the CDM position.
func (service *Service) EffectivenessCDM(card Card) float64 {
	return service.config.CardEfficiencyParameters.CDM.Defence*float64(card.Defence) +
		service.config.CardEfficiencyParameters.CDM.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.CDM.Tactics*float64(card.Tactics) +
		service.config.CardEfficiencyParameters.CDM.Technique*float64(card.Technique) +
		service.config.CardEfficiencyParameters.CDM.Offence*float64(card.Offence)
}

// EffectivenessCM determines the effectiveness of the card in the CM position.
func (service *Service) EffectivenessCM(card Card) float64 {
	return service.config.CardEfficiencyParameters.CM.Defence*float64(card.Defence) +
		service.config.CardEfficiencyParameters.CM.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.CM.Tactics*float64(card.Tactics) +
		service.config.CardEfficiencyParameters.CM.Technique*float64(card.Technique) +
		service.config.CardEfficiencyParameters.CM.Offence*float64(card.Offence)
}

// EffectivenessCAM determines the effectiveness of the card in the CAM position.
func (service *Service) EffectivenessCAM(card Card) float64 {
	return service.config.CardEfficiencyParameters.CAM.Defence*float64(card.Defence) +
		service.config.CardEfficiencyParameters.CAM.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.CAM.Tactics*float64(card.Tactics) +
		service.config.CardEfficiencyParameters.CAM.Technique*float64(card.Technique) +
		service.config.CardEfficiencyParameters.CAM.Offence*float64(card.Offence)
}

// EffectivenessRMorLM determines the effectiveness of the card in the LM/RM position.
func (service *Service) EffectivenessRMorLM(card Card) float64 {
	return service.config.CardEfficiencyParameters.RMorLM.Defence*float64(card.Defence) +
		service.config.CardEfficiencyParameters.RMorLM.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.RMorLM.Tactics*float64(card.Tactics) +
		service.config.CardEfficiencyParameters.RMorLM.Technique*float64(card.Technique) +
		service.config.CardEfficiencyParameters.RMorLM.Offence*float64(card.Offence)
}

// EffectivenessRWorLW determines the effectiveness of the card in the LW/RW position.
func (service *Service) EffectivenessRWorLW(card Card) float64 {
	return service.config.CardEfficiencyParameters.RWorLW.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.RWorLW.Tactics*float64(card.Tactics) +
		service.config.CardEfficiencyParameters.RWorLW.Technique*float64(card.Technique) +
		service.config.CardEfficiencyParameters.RWorLW.Offence*float64(card.Offence)
}

// EffectivenessST determines the effectiveness of the card in the ST position.
func (service *Service) EffectivenessST(card Card) float64 {
	return service.config.CardEfficiencyParameters.ST.Physique*float64(card.Physique) +
		service.config.CardEfficiencyParameters.ST.Tactics*float64(card.Tactics) +
		service.config.CardEfficiencyParameters.ST.Technique*float64(card.Technique) +
		service.config.CardEfficiencyParameters.ST.Offence*float64(card.Offence)
}
