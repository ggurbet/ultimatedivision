// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

// ErrSeasons indicates that there was an error in the service.
var ErrSeasons = errs.Class("seasons service error")

// Service is handling seasons related logic.
//
// architecture: Service
type Service struct {
	seasons          DB
	divisions        *divisions.Service
	matches          *matches.Service
	config           Config
	clubs            *clubs.Service
	users            *users.Service
	currencywaitlist *currencywaitlist.Service
}

// NewService is a constructor for seasons service.
func NewService(seasons DB, config Config, divisions *divisions.Service, matches *matches.Service, clubs *clubs.Service, users *users.Service, currencywaitlist *currencywaitlist.Service) *Service {
	return &Service{
		seasons:          seasons,
		divisions:        divisions,
		config:           config,
		matches:          matches,
		clubs:            clubs,
		users:            users,
		currencywaitlist: currencywaitlist,
	}
}

// Create creates a season.
func (service *Service) Create(ctx context.Context) error {
	divisions, err := service.divisions.List(ctx)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	for _, division := range divisions {
		season := Season{
			DivisionID: division.ID,
			StartedAt:  time.Now().UTC(),
			EndedAt:    time.Time{},
		}

		if err = service.seasons.Create(ctx, season); err != nil {
			return ErrSeasons.Wrap(err)
		}
	}

	return nil
}

// CreateReward creates a rewards in the end of a season.
func (service *Service) CreateReward(ctx context.Context, reward Reward) error {
	_, err := service.currencywaitlist.Create(ctx, reward.UserID, reward.Value, reward.Nonce)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}
	return ErrSeasons.Wrap(service.seasons.CreateReward(ctx, reward))
}

// EndSeason changes status when season end.
func (service *Service) EndSeason(ctx context.Context, id int) error {
	return ErrSeasons.Wrap(service.seasons.EndSeason(ctx, id))
}

// List returns all seasons from DB.
func (service *Service) List(ctx context.Context) ([]Season, error) {
	seasons, err := service.seasons.List(ctx)
	return seasons, ErrSeasons.Wrap(err)
}

// GetCurrentSeasons returns all current seasons from DB.
func (service *Service) GetCurrentSeasons(ctx context.Context) ([]Season, error) {
	seasons, err := service.seasons.GetCurrentSeasons(ctx)
	return seasons, ErrSeasons.Wrap(err)
}

// Get returns season from DB.
func (service *Service) Get(ctx context.Context, seasonID int) (Season, error) {
	season, err := service.seasons.Get(ctx, seasonID)
	return season, ErrSeasons.Wrap(err)
}

// GetRewardByUserID returns user reward by id from DB.
func (service *Service) GetRewardByUserID(ctx context.Context, userID uuid.UUID) (Reward, error) {
	season, err := service.seasons.GetRewardByUserID(ctx, userID)
	return season, ErrSeasons.Wrap(err)
}

// Delete deletes a season.
func (service *Service) Delete(ctx context.Context, id int) error {
	return ErrSeasons.Wrap(service.seasons.Delete(ctx, id))
}

// GetSeasonByDivisionID returns season by division id.
func (service *Service) GetSeasonByDivisionID(ctx context.Context, divisionID uuid.UUID) (Season, error) {
	season, err := service.seasons.GetSeasonByDivisionID(ctx, divisionID)
	return season, ErrSeasons.Wrap(err)
}

// GetDivision returns division by name.
func (service *Service) GetDivision(ctx context.Context, divisionName int) (divisions.Division, error) {
	return service.divisions.GetByName(ctx, divisionName)
}

// GetAllClubsStatistics returns all clubs statistics by division.
func (service *Service) GetAllClubsStatistics(ctx context.Context, division divisions.Division) ([]matches.Statistic, error) {
	currentSeasons, err := service.GetCurrentSeasons(ctx)
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}
	var season Season
	for _, currentSeason := range currentSeasons {
		if currentSeason.DivisionID == division.ID {
			season = currentSeason
			break
		}
	}

	clubs, err := service.clubs.ListByDivision(ctx, division)
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}

	var statistics []matches.Statistic
	for _, club := range clubs {
		statistic, err := service.matches.GetStatistic(ctx, club, season.ID)
		if err != nil {
			return nil, ErrSeasons.Wrap(err)
		}
		if statistic.MatchPlayed >= matches.MinNumberOfMatches {
			statistics = append(statistics, statistic)
		}
	}

	return statistics, nil
}

// UpdateClubsToNewDivision updates clubs to new division.
func (service *Service) UpdateClubsToNewDivision(ctx context.Context) error {
	divisions, err := service.divisions.List(ctx)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	var totalPassingClubs float64
	for _, division := range divisions {
		clubsStatisticsByDivision, err := service.GetAllClubsStatistics(ctx, division)
		if err != nil {
			return ErrSeasons.Wrap(err)
		}

		for _, statistic := range clubsStatisticsByDivision {
			userProfile, err := service.users.GetProfile(ctx, statistic.Club.OwnerID)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			var reward Reward

			switch userProfile.WalletType {
			case users.WalletTypeCasper:
				var nonce int64
				nonce, err = service.currencywaitlist.GetNonceByWallet(ctx, userProfile.CasperWalletID)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				reward = Reward{
					UserID:              userProfile.ID,
					WalletAddress:       common.Address{},
					CasperWalletAddress: userProfile.CasperWalletID,
					WalletType:          userProfile.WalletType,
					Value:               *big.NewInt(10),
					Nonce:               nonce,
					Signature:           "",
				}
			default:
				reward = Reward{
					UserID:              userProfile.ID,
					WalletAddress:       userProfile.Wallet,
					CasperWalletAddress: "",
					WalletType:          userProfile.WalletType,
					Value:               *big.NewInt(10),
					Nonce:               0,
					Signature:           "",
				}
			}

			err = service.CreateReward(ctx, reward)
			if err != nil {
				return ChoreError.Wrap(err)
			}
		}

		var percent float64
		percent = 100 / float64(len(clubsStatisticsByDivision))
		if percent < float64(division.PassingPercent) {
			totalPassingClubs = float64(division.PassingPercent) / percent
		} else {
			totalPassingClubs = 1
		}
		sortStatistics := clubsStatisticsByDivision
		sort.Slice(sortStatistics, func(i, j int) bool {
			return sortStatistics[i].Points < sortStatistics[j].Points
		})

		if len(sortStatistics) > 0 {
			topStatisticsClubs := sortStatistics[len(sortStatistics)-int(totalPassingClubs):]
			lowStatisticsClubs := sortStatistics[:int(totalPassingClubs)]

			divisionHigher, err := service.divisions.GetByName(ctx, division.Name-1)
			if err != nil {
				return ErrSeasons.Wrap(err)
			}
			for _, statistic := range topStatisticsClubs {
				err := service.clubs.UpdateClubToNewDivision(ctx, statistic.Club.ID, divisionHigher.ID)
				if err != nil {
					return ErrSeasons.Wrap(err)
				}
			}

			lastDivision, err := service.divisions.GetLastDivision(ctx)
			if err != nil {
				return ErrSeasons.Wrap(err)
			}

			if division.Name < lastDivision.Name {
				divisionLower, err := service.divisions.GetByName(ctx, division.Name+1)
				if err != nil {
					return ErrSeasons.Wrap(err)

				}
				for _, statistic := range lowStatisticsClubs {
					err := service.clubs.UpdateClubToNewDivision(ctx, statistic.Club.ID, divisionLower.ID)
					if err != nil {
						return ErrSeasons.Wrap(err)
					}
				}
			}
		}
	}

	return nil
}
