// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/queue"
)

const (
	// ReadBufferSize is buffer sizes for read.
	ReadBufferSize int = 1024
	// WriteBufferSize is buffer sizes for write.
	WriteBufferSize int = 1024
)

var (
	// ErrQueue is an internal error type for queue controller.
	ErrQueue = errs.Class("queue controller error")
)

// Queue is a mvc controller that handles all queue related views.
type Queue struct {
	log   logger.Logger
	queue *queue.Service
}

// NewQueue is a constructor for queue controller.
func NewQueue(log logger.Logger, queue *queue.Service) *Queue {
	queueController := &Queue{
		log:   log,
		queue: queue,
	}

	return queueController
}

// Create is an endpoint that creates queue.
func (controller *Queue) Create(w http.ResponseWriter, r *http.Request) {
	var (
		request queue.Request
		err     error
		conn    *websocket.Conn
		claims  auth.Claims
	)

	ctx := r.Context()
	upgrader := websocket.Upgrader{
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriteBufferSize,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		controller.log.Error("could not connect to websocket", ErrQueue.Wrap(err))
		return
	}

	if claims, err = auth.GetClaims(ctx); err != nil {
		controller.serveError(conn, http.StatusUnauthorized, err.Error())
		return
	}

	if err = conn.ReadJSON(&request); err != nil {
		controller.log.Error("could not read JSON from websocket", ErrQueue.Wrap(err))
		controller.serveError(conn, http.StatusBadRequest, err.Error())
		return
	}

	client := queue.Client{
		UserID:     claims.UserID,
		Connection: conn,
		SquadID:    request.SquadID,
	}

	switch request.Action {
	case queue.ActionStartSearch:
		if _, err = controller.queue.Get(client.UserID); err != nil {
			if err = controller.queue.Create(ctx, client); err != nil {
				controller.log.Error("could not create user's queue", ErrQueue.Wrap(err))
				controller.serveError(client.Connection, http.StatusInternalServerError, err.Error())
				return
			}
			controller.serveError(client.Connection, http.StatusOK, "you added!")
			return
		}
		controller.serveError(client.Connection, http.StatusBadRequest, "you have already been added!")
		return
	case queue.ActionFinishSearch:
		if _, err = controller.queue.Get(client.UserID); err == nil {
			if err = controller.queue.Finish(client.UserID); err != nil {
				controller.log.Error("could not finish search", ErrQueue.Wrap(err))
				controller.serveError(client.Connection, http.StatusInternalServerError, err.Error())
			}
			defer func() {
				controller.log.Error("could not close websocket", ErrQueue.Wrap(client.Connection.Close()))
			}()

			controller.serveError(client.Connection, http.StatusOK, "you leaved!")
			return
		}
		controller.serveError(client.Connection, http.StatusBadRequest, "you have not been added!")
		return
	default:
		controller.log.Error("wrong action", ErrQueue.Wrap(err))
		controller.serveError(client.Connection, http.StatusBadRequest, "wrong action")
		return
	}
}

// serveError replies to request with specific code and error.
func (controller *Queue) serveError(w *websocket.Conn, status int, message string) {
	response := queue.Response{Status: status, Message: message}
	if err := w.WriteJSON(response); err != nil {
		controller.log.Error("could not write to websocket", ErrQueue.Wrap(err))
		return
	}
}
