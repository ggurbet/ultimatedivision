// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"net/http"

	"github.com/zeebo/errs"

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
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, service *Service) *Chore {
	return &Chore{
		log:     log,
		service: service,
		Loop:    sync.NewCycle(config.PlaceRenewalInterval),
	}
}

// Run starts the chore for re-check the expiration time of the token.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		clients := chore.service.List()
		if len(clients) >= 2 {
			for k := range clients {
				if k%2 != 0 || (clients[k] == Client{} && clients[k+1] == Client{}) {
					continue
				}

				firstUser := clients[k]
				secondUser := clients[k+1]
				firstClient := Client{
					UserID: firstUser.UserID,
					Conn:   firstUser.Conn,
				}
				secondClient := Client{
					UserID: secondUser.UserID,
					Conn:   secondUser.Conn,
				}

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

				chore.service.Finish(firstClient.UserID)
				chore.service.Finish(secondClient.UserID)

				defer func() {
					if err = firstClient.Conn.Close(); err != nil {
						chore.log.Error("could not close websocket", ErrQueue.Wrap(err))
					}
				}()
				defer func() {
					if err = secondClient.Conn.Close(); err != nil {
						chore.log.Error("could not close websocket", ErrQueue.Wrap(err))
					}
				}()

				// TODO: add to match and send result
			}
		}
		return ChoreError.Wrap(err)
	})
}
