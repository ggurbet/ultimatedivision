// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"math/big"
	"net/http"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/matches"
)

// ErrNoClient indicated that client does not exist.
var ErrNoClient = errs.Class("client does not exist")

// ErrRead indicates a read error.
var ErrRead = errs.Class("error read from websocket")

// ErrWrite indicates a write error.
var ErrWrite = errs.Class("error write to websocket")

// DB is exposing access to clients database.
//
// architecture: DB
type DB interface {
	// Create adds client in database.
	Create(client Client)
	// Get returns client from database.
	Get(userID uuid.UUID) (Client, error)
	// List returns clients from database.
	List() []Client
	// ListNotPlayingUsers returns clients who don't play game from database.
	ListNotPlayingUsers() []Client
	// UpdateIsPlaying updates is playing status of client in database.
	UpdateIsPlaying(userID uuid.UUID, IsPlaying bool) error
	// Delete deletes client record in database.
	Delete(userID uuid.UUID) error
}

// Client entity describes the value of connect with the client.
type Client struct {
	UserID     uuid.UUID
	Connection *websocket.Conn
	SquadID    uuid.UUID
	IsPlaying  bool
}

// Request entity describes values sent by client.
type Request struct {
	Action        Action               `json:"action"`
	SquadID       uuid.UUID            `json:"squadId"`
	WalletAddress evmsignature.Address `json:"walletAddress"`
	Nonce         int64                `json:"nonce"`
}

// Action defines list of possible clients action.
type Action string

const (
	// ActionStartSearch indicates that the client starts the search.
	ActionStartSearch Action = "startSearch"
	// ActionFinishSearch indicates that the client finishes the search.
	ActionFinishSearch Action = "finishSearch"
	// ActionConfirm indicates that the client confirms the game.
	ActionConfirm Action = "confirm"
	// ActionReject indicates that the client rejects the game.
	ActionReject Action = "reject"
	// ActionAllowAddress indicates that the client allows to add address of wallet.
	ActionAllowAddress Action = "allowAddress"
	// ActionForbidAddress indicates that the client is forbidden to add wallet address.
	ActionForbidAddress Action = "forbidAddress"
)

// Response entity describes values sent to user.
type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

// Config defines configuration for queue.
type Config struct {
	PlaceRenewalInterval time.Duration         `json:"placeRenewalInterval"`
	WinValue             string                `json:"winValue"`
	DrawValue            string                `json:"drawValue"`
	UDTContract          evmsignature.Contract `json:"udtContract"`
}

// ReadJSON reads request sent by client.
func (client *Client) ReadJSON() (Request, error) {
	var request Request
	if err := client.Connection.ReadJSON(&request); err != nil {
		if err = client.WriteJSON(http.StatusBadRequest, err.Error()); err != nil {
			return request, ErrWrite.Wrap(ErrQueue.Wrap(err))
		}
		return request, ErrRead.Wrap(ErrQueue.Wrap(err))
	}
	return request, nil
}

// WriteJSON writes response to client.
func (client *Client) WriteJSON(status int, message interface{}) error {
	if err := client.Connection.WriteJSON(Response{Status: status, Message: message}); err != nil {
		return ErrWrite.Wrap(ErrQueue.Wrap(err))
	}
	return nil
}

// WinResult entity describes values which send to user after win game.
type WinResult struct {
	Client     Client             `json:"client"`
	GameResult matches.GameResult `json:"gameResult"`
	Value      *big.Int           `json:"value"`
}
