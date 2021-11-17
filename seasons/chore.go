// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/sync"
)

var (
	// ChoreError represents season chore error type.
	ChoreError = errs.Class("expiration season chore error")
)

// Chore requests access for season service, re-requests it after expiration time.
//
// architecture: Chore
type Chore struct {
	seasons *Service
	Loop    *sync.Cycle
}

// NewChore instantiates Chore.
func NewChore(config Config, service *Service) *Chore {
	return &Chore{
		seasons: service,
		Loop:    sync.NewCycle(config.SeasonTime),
	}
}

// Run starts the chore for re-check the expiration time of the season.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		seasons, err := chore.seasons.List(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, season := range seasons {
			if season.EndedAt.IsZero() {
				err := chore.seasons.EndSeason(ctx, season.ID)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			}
		}

		err = chore.seasons.Create(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		return ChoreError.Wrap(err)
	})
}

// Close closes the chore for re-check the expiration time of the season.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
