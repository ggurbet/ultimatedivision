// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"

	"ultimatedivision/console/connections"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
)

const (
	// ReadBufferSize is buffer sizes for read.
	ReadBufferSize int = 1024
	// WriteBufferSize is buffer sizes for write.
	WriteBufferSize int = 1024
)

var (
	// ErrConnections is an internal error type for connections controller.
	ErrConnections = errs.Class("connections controller error")
)

// Connections is a mvc controller that handles all connections related views.
type Connections struct {
	log logger.Logger

	connection *connections.Service
}

// NewConnections is a constructor for connections controller.
func NewConnections(log logger.Logger, connection *connections.Service) *Connections {
	connectionsController := &Connections{
		log:        log,
		connection: connection,
	}

	return connectionsController
}

// Connect is an endpoint that creates websocket connection.
func (controller *Connections) Connect(w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var err error
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrConnections.Wrap(err))
		return
	}

	err = controller.connection.Close(claims.UserID)
	if err != nil {
		if !connections.ErrNoConnection.Has(err) {
			controller.log.Error(fmt.Sprintf("could not close old connection to websocket for user %x", claims.UserID), ErrConnections.Wrap(err))
			controller.serveError(w, http.StatusInternalServerError, ErrConnections.Wrap(err))
			return
		}
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriteBufferSize,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		controller.log.Error("could not connect to websocket", ErrConnections.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrConnections.Wrap(err))
		return
	}

	if err = controller.connection.Create(claims.UserID, conn); err != nil {
		controller.log.Error(fmt.Sprintf("could not create connection for user %x", claims.UserID), ErrConnections.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrConnections.Wrap(err))
		return
	}
}

// serveError replies to request with specific code and error.
func (controller *Connections) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrConnections.Wrap(err))
	}
}
