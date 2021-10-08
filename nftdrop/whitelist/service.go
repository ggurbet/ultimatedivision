// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/zeebo/errs"
)

// ErrWhitelist indicated that there was an error in service.
var ErrWhitelist = errs.Class("whitelist service error")

// Service is handling whitelist related logic.
//
// architecture: Service
type Service struct {
	whitelist DB
}

// NewService is a constructor for whitelist service.
func NewService(whitelist DB) *Service {
	return &Service{
		whitelist: whitelist,
	}
}

// Create adds whitelist in the database.
func (service *Service) Create(ctx context.Context, address Address) error {
	whitelist := Whitelist{
		Address: address,
		// TODO: generate password
		Password: []byte{},
	}

	return ErrWhitelist.Wrap(service.whitelist.Create(ctx, whitelist))
}

// GetByAddress returns whitelist by address from the database.
func (service *Service) GetByAddress(ctx context.Context, address Address) (Whitelist, error) {
	whitelist, err := service.whitelist.GetByAddress(ctx, address)
	return whitelist, ErrWhitelist.Wrap(err)
}

// List returns all whitelist from the database.
func (service *Service) List(ctx context.Context) ([]Whitelist, error) {
	whitelistRecords, err := service.whitelist.List(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}

// Delete deletes whitelist.
func (service *Service) Delete(ctx context.Context, address Address) error {
	return ErrWhitelist.Wrap(service.whitelist.Delete(ctx, address))
}
