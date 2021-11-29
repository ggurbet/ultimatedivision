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
	"ultimatedivision/seasons"
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
	seasons *seasons.Service
	clubs   *clubs.Service
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, service *Service, matches *matches.Service, seasons *seasons.Service, clubs *clubs.Service) *Chore {
	return &Chore{
		log:     log,
		service: service,
		Loop:    sync.NewCycle(config.PlaceRenewalInterval),
		matches: matches,
		seasons: seasons,
		clubs:   clubs,
	}
}

// Run starts the chore for re-check the expiration time of the token.
func (chore *Chore) Run(ctx context.Context) (err error) {
	firstRequestChan := make(chan Request)
	secondRequestChan := make(chan Request)

	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		clients := chore.service.ListNotPlayingUsers()

		if len(clients) >= 2 {
			for k := range clients {
				isEvenNumber := (k%2 != 0)
				if isEvenNumber {
					continue
				}

				go func(clients []Client, k int) {
					firstClient := clients[k]
					secondClient := clients[k+1]

					if err = chore.service.UpdateIsPlaying(firstClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
					}
					if err = chore.service.UpdateIsPlaying(secondClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
					}

					if err := firstClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
					}
					if err := secondClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
					}

					go func() {
						request, err := firstClient.ReadJSON()
						if err != nil {
							chore.log.Error("could not read json", ChoreError.Wrap(err))
						}
						firstRequestChan <- request
					}()

					go func() {
						request, err := secondClient.ReadJSON()
						if err != nil {
							chore.log.Error("could not read json", ChoreError.Wrap(err))
						}
						secondRequestChan <- request
					}()

					var firstRequest, secondRequest Request
					for {
						select {
						case firstRequest = <-firstRequestChan:
							if (firstRequest != Request{}) {
								if firstRequest.Action != ActionConfirm && firstRequest.Action != ActionReject {
									if err := firstClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
										chore.log.Error("could not write json", ChoreError.Wrap(err))
									}

									if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									return
								}
							}
						case secondRequest = <-secondRequestChan:
							if (secondRequest != Request{}) {
								if secondRequest.Action != ActionConfirm && secondRequest.Action != ActionReject {
									if err := secondClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
										chore.log.Error("could not write json", ChoreError.Wrap(err))
									}

									if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									return
								}
							}
						}

						if (firstRequest == Request{} && secondRequest == Request{}) {
							continue
						}

						if firstRequest.Action == ActionReject || secondRequest.Action == ActionReject {
							if err := firstClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
							}
							if err := secondClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
							}

							if err = chore.service.Finish(firstClient.UserID); err != nil {
								chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
							}
							if err = chore.service.Finish(secondClient.UserID); err != nil {
								chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
							}
							return
						}

						if (firstRequest == Request{} || secondRequest == Request{}) {
							continue
						}

						if err = chore.Play(ctx, firstClient, secondClient); err != nil {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							chore.log.Error("could not play game", ChoreError.Wrap(err))
						}
						return
					}
				}(clients, k)
			}
		}
		return ChoreError.Wrap(err)
	})
}

// Play method contains all the logic for playing matches.
func (chore *Chore) Play(ctx context.Context, firstClient, secondClient Client) error {
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

	firstClientSquad, err := chore.clubs.GetSquad(ctx, firstClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	firstClientClub, err := chore.clubs.Get(ctx, firstClientSquad.ClubID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	season, err := chore.seasons.GetSeasonByDivisionID(ctx, firstClientClub.DivisionID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	matchesID, err := chore.matches.Create(ctx, firstClient.SquadID, secondClient.SquadID, firstClient.UserID, secondClient.UserID, season.ID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	matchResult, err := chore.matches.GetMatchResult(ctx, matchesID)
	if err != nil {
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not get result of match"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	var (
		firstClientResult  []matches.MatchResult
		secondClientResult []matches.MatchResult
	)

	switch {
	case firstClient.UserID == matchResult[0].UserID:
		firstClientResult = matchResult
		secondClientResult = matches.Swap(matchResult)
	case secondClient.UserID == matchResult[0].UserID:
		firstClientResult = matches.Swap(matchResult)
		secondClientResult = matchResult
	}

	if err := firstClient.WriteJSON(http.StatusOK, firstClientResult); err != nil {
		return ChoreError.Wrap(err)
	}
	if err := secondClient.WriteJSON(http.StatusOK, secondClientResult); err != nil {
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

	return nil
}

// Close closes the chore chore for re-check the expiration time of the token.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
