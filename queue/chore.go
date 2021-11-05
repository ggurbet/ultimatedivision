// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/sync"
)

var (
	// ChoreError represents place chore error type.
	ChoreError = errs.Class("expiration place chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log     logger.Logger
	service *Service
	Loop    *sync.Cycle
	matches *matches.Service
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, service *Service, matches *matches.Service) *Chore {
	return &Chore{
		log:     log,
		service: service,
		Loop:    sync.NewCycle(config.PlaceRenewalInterval),
		matches: matches,
	}
}

// Run starts the chore for re-check the expiration time of the token.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		clients := chore.service.List()

		if len(clients) >= 2 {
			for k := range clients {
				isEvenNumber := (k%2 != 0)
				if isEvenNumber {
					continue
				}

				isEmptyClients := (clients[k] == Client{} && clients[k+1] == Client{})
				isEqualDivisions := (clients[k].DivisionID != clients[k+1].DivisionID)
				if isEmptyClients || isEqualDivisions {
					continue
				}

				firstClient := clients[k]
				secondClient := clients[k+1]

				if err := firstClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
					return ChoreError.Wrap(err)
				}

				firstRequest, err := firstClient.ReadJSON()
				if err != nil {
					return ChoreError.Wrap(err)
				}
				secondRequest, err := secondClient.ReadJSON()
				if err != nil {
					return ChoreError.Wrap(err)
				}

				if firstRequest.Action != ActionConfirm && firstRequest.Action != ActionReject {
					if err := firstClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
						return ChoreError.Wrap(err)
					}
				}
				if secondRequest.Action != ActionConfirm && secondRequest.Action != ActionReject {
					if err := firstClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
						return ChoreError.Wrap(err)
					}
				}

				if firstRequest.Action == ActionReject || secondRequest.Action == ActionReject {
					if err := firstClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
						return ChoreError.Wrap(err)
					}
					if err := secondClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
						return ChoreError.Wrap(err)
					}
					continue
				}

				squadCardsFirstClient, err := chore.service.clubs.ListSquadCards(ctx, firstClient.SquadID)
				if err != nil {
					return ChoreError.Wrap(err)
				}
				if len(squadCardsFirstClient) != clubs.SquadSize {
					if err := firstClient.WriteJSON(http.StatusInternalServerError, "squad is not full"); err != nil {
						return ChoreError.Wrap(err)
					}
				}

				squadCardsSecondClient, err := chore.service.clubs.ListSquadCards(ctx, secondClient.SquadID)
				if err != nil {
					return ChoreError.Wrap(err)
				}
				if len(squadCardsSecondClient) != clubs.SquadSize {
					if err := secondClient.WriteJSON(http.StatusInternalServerError, "squad is not full"); err != nil {
						return ChoreError.Wrap(err)
					}
				}

				matchesID, err := chore.matches.Create(ctx, firstClient.SquadID, secondClient.SquadID, firstClient.UserID, secondClient.UserID)
				if err != nil {
					if err := firstClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
						return ChoreError.Wrap(err)
					}
					if err := secondClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
						return ChoreError.Wrap(err)
					}
				}

				resultMatch, err := chore.matches.GetMatchResult(ctx, matchesID)
				if err != nil {
					if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not get result of match"); err != nil {
						return ChoreError.Wrap(err)
					}
				}
				if err := firstClient.WriteJSON(http.StatusOK, resultMatch); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusOK, resultMatch); err != nil {
					return ChoreError.Wrap(err)
				}

				if err = chore.service.Finish(firstClient.UserID); err != nil {
					return ChoreError.Wrap(err)
				}
				if err = chore.service.Finish(secondClient.UserID); err != nil {
					return ChoreError.Wrap(err)
				}

				defer func() {
					if err = firstClient.Connection.Close(); err != nil {
						chore.log.Error("could not close websocket", ErrQueue.Wrap(err))
					}
				}()
				defer func() {
					if err = secondClient.Connection.Close(); err != nil {
						chore.log.Error("could not close websocket", ErrQueue.Wrap(err))
					}
				}()
			}
		}
		return ChoreError.Wrap(err)
	})
}
