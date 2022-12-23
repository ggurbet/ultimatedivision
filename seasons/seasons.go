// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons

import (
	"context"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/users"
)

// ErrNoSeason indicated that season does not exist.
var ErrNoSeason = errs.Class("season does not exist")

// DB exposes access to seasons db.
//
// architecture: DB
type DB interface {
	// Create creates a season and writes to the database.
	Create(ctx context.Context, season Season) error
	// CreateReward creates a season reward and writes to the database.
	CreateReward(ctx context.Context, reward Reward) error
	// EndSeason updates a status in the database when season ended.
	EndSeason(ctx context.Context, id int) error
	// List returns all seasons from the data base.
	List(ctx context.Context) ([]Season, error)
	// Get returns season by id from the data base.
	Get(ctx context.Context, id int) (Season, error)
	// GetCurrentSeasons returns all current seasons from the data base.
	GetCurrentSeasons(ctx context.Context) ([]Season, error)
	// GetSeasonByDivisionID returns season by division id from the data base.
	GetSeasonByDivisionID(ctx context.Context, divisionID uuid.UUID) (Season, error)
	// GetRewardByUserID returns user reward by id from the data base.
	GetRewardByUserID(ctx context.Context, userID uuid.UUID) (Reward, error)
	// ListRewards returns all seasons rewards from the data base.
	ListRewards(ctx context.Context) ([]Reward, error)
	// Delete deletes a season in the database.
	Delete(ctx context.Context, id int) error
}

// Status defines the list of possible season statuses.
type Status int

// Season describes seasons entity.
type Season struct {
	ID         int       `json:"id"`
	DivisionID uuid.UUID `json:"divisionId"`
	StartedAt  time.Time `json:"startedAt"`
	EndedAt    time.Time `json:"endedAt"`
}

// Config defines configuration for seasons.
type Config struct {
	SeasonTime time.Duration `json:"seasonTime"`
}

// SeasonStatistics returns statistics of clubs in season.
type SeasonStatistics struct {
	Division   divisions.Division  `json:"division"`
	Statistics []matches.Statistic `json:"statistics"`
}

// Reward entity describes values which send to user after season ends.
type Reward struct {
	ID                  uuid.UUID              `json:"ID"`
	UserID              uuid.UUID              `json:"userId"`
	SeasonID            int                    `json:"seasonID"`
	WalletAddress       common.Address         `json:"walletAddress"`
	CasperWalletAddress string                 `json:"casperWalletAddress"`
	Status              int                    `json:"status"`
	Value               big.Int                `json:"value"`
	Nonce               int64                  `json:"nonce"`
	WalletType          users.WalletType       `json:"walleType"`
	Signature           evmsignature.Signature `json:"signature"`
}
