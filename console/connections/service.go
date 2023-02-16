// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package connections

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"
)

// ErrConnections indicates that there was an error in the service.
var ErrConnections = errs.Class("connections service error")

// Service is handling connections related logic.
//
// architecture: Service
type Service struct {
	connections DB
}

// NewService is a constructor for connections service.
func NewService(connections DB) *Service {
	return &Service{
		connections: connections,
	}
}

// Create creates a connection by user.
func (service *Service) Create(userID uuid.UUID, connection *websocket.Conn) error {
	return ErrConnections.Wrap(service.connections.Create(userID, connection))
}

// List returns all connections.
func (service *Service) List() map[uuid.UUID]*websocket.Conn {
	return service.connections.List()
}

// Get returns connection by user.
func (service *Service) Get(userID uuid.UUID) (*websocket.Conn, error) {
	connection, err := service.connections.Get(userID)
	return connection, ErrConnections.Wrap(err)
}

// Close closes a connection by user.
func (service *Service) Close(id uuid.UUID) error {
	connection, err := service.connections.Get(id)
	if err != nil {
		return ErrConnections.Wrap(err)
	}

	err = connection.Close()
	if err != nil {
		return ErrConnections.Wrap(err)
	}

	return ErrConnections.Wrap(service.connections.Delete(id))
}
