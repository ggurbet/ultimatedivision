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
	// List returns all seasons from the database.
	List(ctx context.Context) ([]Season, error)
	// Get returns season by id from the database.
	Get(ctx context.Context, id int) (Season, error)
	// GetCurrentSeasons returns all current seasons from the database.
	GetCurrentSeasons(ctx context.Context) ([]Season, error)
	// GetSeasonByDivisionID returns season by division id from the database.
	GetSeasonByDivisionID(ctx context.Context, divisionID uuid.UUID) (Season, error)
	// GetRewardByUserID returns user reward by id from the database.
	GetRewardByUserID(ctx context.Context, userID uuid.UUID) (Reward, error)
	// ListOfUnpaidRewardsByUserID returns all unpaid season rewards from the database by user id.
	ListOfUnpaidRewardsByUserID(ctx context.Context, userID uuid.UUID) ([]Reward, error)
	// Delete deletes a season in the database.
	Delete(ctx context.Context, id int) error
}

// StatusReward defines the list of possible reward statuses.
type StatusReward int

const (
	// StatusUnPaid indicates that reward is unpaid.
	StatusUnPaid StatusReward = 0
	// StatusPaid indicates that reward is paid.
	StatusPaid StatusReward = 1
)

// Season describes seasons entity.
type Season struct {
	ID         int       `json:"id"`
	DivisionID uuid.UUID `json:"divisionId"`
	StartedAt  time.Time `json:"startedAt"`
	EndedAt    time.Time `json:"endedAt"`
}

// Config defines configuration for seasons.
type Config struct {
	SeasonTime          time.Duration         `json:"seasonTime"`
	CasperTokenContract evmsignature.Contract `json:"casperTokenContract"`
	RPCNodeAddress      string                `json:"rpcNodeAddress"`
}

// SeasonStatistics returns statistics of clubs in season.
type SeasonStatistics struct {
	Division   divisions.Division  `json:"division"`
	Statistics []matches.Statistic `json:"statistics"`
}

// Reward entity describes values which send to user after season ends.
type Reward struct {
	ID                  uuid.UUID        `json:"ID"`
	UserID              uuid.UUID        `json:"userId"`
	SeasonID            int              `json:"seasonID"`
	WalletAddress       common.Address   `json:"walletAddress"`
	CasperWalletAddress string           `json:"casperWalletAddress"`
	CasperWalletHash    string           `json:"CasperWalletHash"`
	WalletType          users.WalletType `json:"walleType"`
	Status              StatusReward     `json:"status"`
	Value               big.Int          `json:"value"`
}

// RewardWithTransaction entity describes values reward with transaction.
type RewardWithTransaction struct {
	Reward
	Nonce               int64                  `json:"nonce"`
	Signature           evmsignature.Signature `json:"signature"`
	Value               string                 `json:"value"`
	CasperTokenContract evmsignature.Contract  `json:"casperTokenContract"`
	RPCNodeAddress      string                 `json:"rpcNodeAddress"`
}
