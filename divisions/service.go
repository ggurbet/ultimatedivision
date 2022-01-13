// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package divisions

import (
	"context"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrDivisions indicates that there was an error in the service.
var ErrDivisions = errs.Class("divisions service error")

// Service is handling divisions related logic.
//
// architecture: Service
type Service struct {
	divisions DB
	config    Config
}

// NewService is a constructor for divisions service.
func NewService(divisions DB, config Config) *Service {
	return &Service{
		divisions: divisions,
		config:    config,
	}
}

// CreateDivisions creates a divisions when app started.
func (service *Service) CreateDivisions(ctx context.Context, name []int) error {
	divisions, err := service.List(ctx)
	if err != nil {
		return ErrDivisions.Wrap(err)
	}
	var divisionNames []int
	for _, d := range divisions {
		divisionNames = append(divisionNames, d.Name)
	}
	if reflect.DeepEqual(divisionNames, name) {
		return nil
	}
	for _, divisionName := range name {
		division := Division{
			ID:             uuid.New(),
			Name:           divisionName,
			PassingPercent: service.config.PassingPercent,
			CreatedAt:      time.Now().UTC(),
		}

		err := ErrDivisions.Wrap(service.divisions.Create(ctx, division))
		if err != nil {
			return ErrDivisions.Wrap(err)
		}
	}
	return nil
}

// Create creates a division.
func (service *Service) Create(ctx context.Context, name int) error {
	division := Division{
		ID:             uuid.New(),
		Name:           name,
		PassingPercent: service.config.PassingPercent,
		CreatedAt:      time.Now().UTC(),
	}

	return ErrDivisions.Wrap(service.divisions.Create(ctx, division))
}

// List returns all divisions from DB.
func (service *Service) List(ctx context.Context) ([]Division, error) {
	divisions, err := service.divisions.List(ctx)
	return divisions, ErrDivisions.Wrap(err)
}

// Get returns division from DB.
func (service *Service) Get(ctx context.Context, divisionID uuid.UUID) (Division, error) {
	division, err := service.divisions.Get(ctx, divisionID)
	return division, ErrDivisions.Wrap(err)
}

// GetByName returns division from DB.
func (service *Service) GetByName(ctx context.Context, divisionName int) (Division, error) {
	division, err := service.divisions.GetByName(ctx, divisionName)
	return division, ErrDivisions.Wrap(err)
}

// GetLastDivision returns last division.
func (service *Service) GetLastDivision(ctx context.Context) (Division, error) {
	division, err := service.divisions.GetLastDivision(ctx)
	return division, ErrDivisions.Wrap(err)
}

// Delete deletes a division.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrDivisions.Wrap(service.divisions.Delete(ctx, id))
}
