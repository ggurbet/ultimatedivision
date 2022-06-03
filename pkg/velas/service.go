// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package velas

// Service is handling velas related logic.
//
// architecture: Service
type Service struct {
	config Config
}

// NewService is a constructor for velas service.
func NewService(config Config) *Service {
	return &Service{
		config: config,
	}
}

// Get returns va client fields.
func (service *Service) Get() VAClientFields {
	return VAClientFields{
		ClientID:                   service.config.ClientID,
		RedirectURI:                service.config.RedirectURI,
		AccountProviderHost:        service.config.AccountProviderHost,
		NetworkAPIHost:             service.config.NetworkAPIHost,
		TransactionsSponsorAPIHost: service.config.TransactionsSponsorAPIHost,
		TransactionsSponsorPubKey:  service.config.TransactionsSponsorPubKey,
	}
}
