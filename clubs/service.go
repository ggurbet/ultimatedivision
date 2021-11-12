// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/divisions"
	"ultimatedivision/users"
)

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
	clubs     DB
	users     *users.Service
	cards     *cards.Service
	divisions divisions.DB
}

// NewService is a constructor for clubs service.
func NewService(clubs DB, users *users.Service, cards *cards.Service, divisions divisions.DB) *Service {
	return &Service{
		clubs:     clubs,
		users:     users,
		cards:     cards,
		divisions: divisions,
	}
}

// Create creates clubs.
func (service *Service) Create(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	nickname, err := service.users.GetNickNameByID(ctx, userID)
	if err != nil {
		return uuid.New(), ErrClubs.Wrap(err)
	}

	division, err := service.divisions.GetLastDivision(ctx)
	if err != nil {
		return uuid.Nil, ErrClubs.Wrap(err)
	}

	newClub := Club{
		ID:         uuid.New(),
		OwnerID:    userID,
		Name:       nickname,
		CreatedAt:  time.Now().UTC(),
		DivisionID: division.ID,
	}

	allClubs, err := service.ListByUserID(ctx, userID)
	if err != nil {
		return uuid.Nil, ErrClubs.Wrap(err)
	}

	switch {
	case len(allClubs) == 0:
		newClub.Status = StatusActive
	default:
		newClub.Status = StatusInactive
	}

	clubID, err := service.clubs.Create(ctx, newClub)
	return clubID, ErrClubs.Wrap(err)
}

// CreateSquad creates new squad for club.
func (service *Service) CreateSquad(ctx context.Context, clubID uuid.UUID) (uuid.UUID, error) {
	newSquad := Squad{
		ID:        uuid.New(),
		ClubID:    clubID,
		Formation: FourFourTwo,
		Tactic:    Balanced,
	}

	squadID, err := service.clubs.CreateSquad(ctx, newSquad)

	return squadID, ErrClubs.Wrap(err)
}

// AddSquadCard adds card to the squad.
func (service *Service) AddSquadCard(ctx context.Context, squadID uuid.UUID, newSquadCard SquadCard) error {
	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	if len(squadCards) == SquadSize {
		return ErrClubs.New("squad is full")
	}

	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newSquadCard.SquadID = squadID
	newSquadCard.Position = FormationToPosition[formation][newSquadCard.Position]

	for _, card := range squadCards {
		if card.Position != newSquadCard.Position {
			continue
		}

		err = service.clubs.DeleteSquadCard(ctx, squadID, card.CardID)
		if err != nil {
			return ErrClubs.Wrap(err)
		}
		break
	}

	return ErrClubs.Wrap(service.clubs.AddSquadCard(ctx, newSquadCard))
}

// Delete deletes card from squad.
func (service *Service) Delete(ctx context.Context, squadID, cardID uuid.UUID) error {
	return ErrClubs.Wrap(service.clubs.DeleteSquadCard(ctx, squadID, cardID))
}

// UpdateSquad updates tactic and captain of the squad.
func (service *Service) UpdateSquad(ctx context.Context, squadID uuid.UUID, tactic Tactic, captainID uuid.UUID) error {
	updatedSquad := Squad{
		ID:        squadID,
		Tactic:    tactic,
		CaptainID: captainID,
	}

	return ErrClubs.Wrap(service.clubs.UpdateTacticCaptain(ctx, updatedSquad))
}

// UpdateCardPosition updates position of card in the squad.
func (service *Service) UpdateCardPosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, newPosition Position) error {
	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	var oldPosition Position

	for _, card := range squadCards {
		if card.CardID == cardID {
			oldPosition = card.Position
			break
		}
	}

	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newPosition = FormationToPosition[formation][newPosition]

	updatedCards := make([]SquadCard, 0, 2)

	updatedSquadCard := SquadCard{
		SquadID:  squadID,
		CardID:   cardID,
		Position: newPosition,
	}

	updatedCards = append(updatedCards, updatedSquadCard)

	for _, card := range squadCards {
		if card.Position != newPosition {
			continue
		}

		card.Position = oldPosition
		updatedCards = append(updatedCards, card)
		break
	}

	return ErrClubs.Wrap(service.clubs.UpdatePositions(ctx, updatedCards))
}

// GetSquadByClubID returns squad of club.
func (service *Service) GetSquadByClubID(ctx context.Context, clubID uuid.UUID) (Squad, error) {
	squad, err := service.clubs.GetSquadByClubID(ctx, clubID)
	return squad, ErrClubs.Wrap(err)
}

// GetSquad returns squad.
func (service *Service) GetSquad(ctx context.Context, squadID uuid.UUID) (Squad, error) {
	squad, err := service.clubs.GetSquad(ctx, squadID)
	return squad, ErrClubs.Wrap(err)
}

// ListSquadCards returns all cards from the squad.
func (service *Service) ListSquadCards(ctx context.Context, squadID uuid.UUID) ([]SquadCard, error) {
	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return squadCards, ErrClubs.Wrap(err)
	}

	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	squadCards = convertPositions(squadCards, formation)

	if len(squadCards) < SquadSize {
		for i := 0; i < SquadSize; i++ {
			var isPositionInTheSquad bool
			for _, card := range squadCards {
				if card.Position == Position(i) {
					isPositionInTheSquad = true
					break
				}
			}

			if isPositionInTheSquad == true {
				continue
			}

			var squadCard = SquadCard{
				SquadID:  squadID,
				Position: Position(i),
			}

			squadCards = append(squadCards, squadCard)
		}
	}

	sortSquadCards(squadCards)

	return squadCards, ErrClubs.Wrap(err)
}

// ListByUserID returns user's clubs.
func (service *Service) ListByUserID(ctx context.Context, userID uuid.UUID) ([]Club, error) {
	club, err := service.clubs.ListByUserID(ctx, userID)
	return club, ErrClubs.Wrap(err)
}

// Get returns club.
func (service *Service) Get(ctx context.Context, clubID uuid.UUID) (Club, error) {
	club, err := service.clubs.Get(ctx, clubID)
	return club, ErrClubs.Wrap(err)
}

// ChangeFormation is a method that change formation and card position.
func (service *Service) ChangeFormation(ctx context.Context, newFormation Formation, squadID uuid.UUID) error {
	var cardsWithNewPositions map[Position]uuid.UUID

	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	if len(squadCards) == 0 {
		return ErrClubs.Wrap(service.clubs.UpdateFormation(ctx, newFormation, squadID))
	}

	err = service.clubs.UpdateFormation(ctx, newFormation, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	cardsWithNewPositions, err = service.CardsWithNewPositions(ctx, squadCards, FormationToPosition[newFormation])
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	var squadCardsWithNewPositions []SquadCard
	for position, card := range cardsWithNewPositions {
		squadCard := SquadCard{
			Position: position,
			SquadID:  squadID,
			CardID:   card,
		}

		squadCardsWithNewPositions = append(squadCardsWithNewPositions, squadCard)
	}

	return ErrClubs.Wrap(service.clubs.UpdatePositions(ctx, squadCardsWithNewPositions))
}

// CalculateEffectivenessOfSquad calculates effectiveness of user's squad.
func (service *Service) CalculateEffectivenessOfSquad(ctx context.Context, squadCards []SquadCard) (float64, error) {
	var effectiveness float64

	if len(squadCards) == 0 {
		return float64(0), nil
	}

	cardsFromSquad, err := service.cards.GetCardsFromSquadCards(ctx, squadCards[0].SquadID)
	if err != nil {
		return float64(0), ErrClubs.Wrap(err)
	}

	for index, squadCard := range squadCards {
		switch squadCard.Position {
		case GK:
			effectiveness += service.cards.EffectivenessGK(cardsFromSquad[index])
		case LB, RB, LWB, RWB:
			effectiveness += service.cards.EffectivenessLBorRB(cardsFromSquad[index])
		case CCD, LCD, RCD:
			effectiveness += service.cards.EffectivenessCD(cardsFromSquad[index])
		case CCDM, LCDM, RCDM:
			effectiveness += service.cards.EffectivenessCDM(cardsFromSquad[index])
		case CCM, LCM, RCM:
			effectiveness += service.cards.EffectivenessCM(cardsFromSquad[index])
		case CCAM, LCAM, RCAM:
			effectiveness += service.cards.EffectivenessCAM(cardsFromSquad[index])
		case LM, RM:
			effectiveness += service.cards.EffectivenessRMorLM(cardsFromSquad[index])
		case LW, RW:
			effectiveness += service.cards.EffectivenessRWorLW(cardsFromSquad[index])
		case CST, RST, LST:
			effectiveness += service.cards.EffectivenessST(cardsFromSquad[index])
		}
	}

	return effectiveness, nil
}

// UpdateStatus updates status of club.
func (service *Service) UpdateStatus(ctx context.Context, userID, clubID uuid.UUID, newStatus Status) error {
	allUserClubs, err := service.clubs.ListByUserID(ctx, userID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	for i := 0; i < len(allUserClubs); i++ {
		if allUserClubs[i].ID != clubID {
			allUserClubs[i].Status = StatusInactive
			continue
		}

		allUserClubs[i].Status = StatusActive
	}

	return ErrClubs.Wrap(service.clubs.UpdateStatuses(ctx, allUserClubs))
}

// RemoveIndex removes element from the slice.
func RemoveIndex(s []SquadCard, index int) []SquadCard {
	return append(s[:index], s[index+1:]...)
}

// EffectiveCardForPosition determines the effective card in the position.
func (service *Service) EffectiveCardForPosition(ctx context.Context, position Position, squadCards []SquadCard) (cards.Card, int, error) {
	cardCoefficients := make(map[float64]cards.Card)
	var index int

	for _, squadCard := range squadCards {
		card, err := service.cards.Get(ctx, squadCard.CardID)
		if err != nil {
			return card, index, ErrClubs.Wrap(err)
		}
		switch position {
		case GK:
			coefficient := service.cards.EffectivenessGK(card)
			cardCoefficients[coefficient] = card
		case CST, LST, RST:
			coefficient := service.cards.EffectivenessST(card)
			cardCoefficients[coefficient] = card
		case LW, RW:
			coefficient := service.cards.EffectivenessRWorLW(card)
			cardCoefficients[coefficient] = card
		case RM, LM:
			coefficient := service.cards.EffectivenessRMorLM(card)
			cardCoefficients[coefficient] = card
		case CCAM, RCAM, LCAM:
			coefficient := service.cards.EffectivenessCAM(card)
			cardCoefficients[coefficient] = card
		case CCM, LCM, RCM:
			coefficient := service.cards.EffectivenessCM(card)
			cardCoefficients[coefficient] = card
		case CCDM, LCDM, RCDM:
			coefficient := service.cards.EffectivenessCDM(card)
			cardCoefficients[coefficient] = card
		case LB, RB, RWB, LWB:
			coefficient := service.cards.EffectivenessLBorRB(card)
			cardCoefficients[coefficient] = card
		case CCD, LCD, RCD:
			coefficient := service.cards.EffectivenessCD(card)
			cardCoefficients[coefficient] = card
		}
	}

	var max float64

	for coeff := range cardCoefficients {
		max = coeff
		if coeff > max {
			max = coeff
		}
	}

	for key, v := range squadCards {
		if v.CardID == cardCoefficients[max].ID {
			index = key
		}
	}

	return cardCoefficients[max], index, nil
}

// CardsWithNewPositions returns cards with new position by new formation.
func (service *Service) CardsWithNewPositions(ctx context.Context, cards []SquadCard, positions []Position) (map[Position]uuid.UUID, error) {
	positionMap := make(map[Position]uuid.UUID)
	maxCards := SquadSize
	for _, position := range positions {
		card, index, err := service.EffectiveCardForPosition(ctx, position, cards)
		if err != nil {
			return positionMap, ErrClubs.Wrap(err)
		}
		positionMap[position] = card.ID
		if len(cards) == maxCards {
			cards = RemoveIndex(cards, index)
		}
		maxCards--
	}

	return positionMap, nil
}
