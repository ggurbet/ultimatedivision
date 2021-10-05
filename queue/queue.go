// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"
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
	Get(UserID uuid.UUID) (Client, error)
	// List returns clients from database.
	List() []Client
	// Delete deletes client record in database.
	Delete(UserID uuid.UUID)
}

// Client entity describes the value of connect with the client.
type Client struct {
	UserID uuid.UUID
	Conn   *websocket.Conn
}

// Request entity describes values sent by client.
type Request struct {
	Action Action `json:"action"`
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
)

// Response entity describes values sent to user.
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Config defines configuration for queue.
type Config struct {
	PlaceRenewalInterval time.Duration `json:"placeRenewalInterval"`
}

// ReadJSON reads request sent by client.
func (client *Client) ReadJSON() (Request, error) {
	var request Request
	if err := client.Conn.ReadJSON(&request); err != nil {
		if err = client.WriteJSON(http.StatusBadRequest, err.Error()); err != nil {
			return request, ErrWrite.Wrap(ErrQueue.Wrap(err))
		}
		return request, ErrRead.Wrap(ErrQueue.Wrap(err))
	}
	return request, nil
}

// WriteJSON writes response to client.
func (client *Client) WriteJSON(status int, message string) error {
	if err := client.Conn.WriteJSON(Response{Status: status, Message: message}); err != nil {
		return ErrWrite.Wrap(ErrQueue.Wrap(err))
	}
	return nil
}
