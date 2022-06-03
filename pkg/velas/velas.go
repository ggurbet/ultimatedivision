// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package velas

// APIRequest for velas response.
type APIRequest struct {
	WalletAddress string `json:"walletAddress"`
	AccessToken   string `json:"accessToken"`
	ExpiresAt     int64  `json:"expiresAt"`
}

// VAClientFields for velas va client fields from config.
type VAClientFields struct {
	ClientID                   string `json:"clientId"`
	RedirectURI                string `json:"redirectUri"`
	AccountProviderHost        string `json:"accountProviderHost"`
	NetworkAPIHost             string `json:"networkApiHost"`
	TransactionsSponsorAPIHost string `json:"transactionsSponsorApiHost"`
	TransactionsSponsorPubKey  string `json:"transactionsSponsorPubKey"`
}

// Config defines configuration for velas va client.
type Config struct {
	ClientID                   string `json:"clientId"`
	RedirectURI                string `json:"redirectUri"`
	AccountProviderHost        string `json:"accountProviderHost"`
	NetworkAPIHost             string `json:"networkApiHost"`
	TransactionsSponsorAPIHost string `json:"transactionsSponsorApiHost"`
	TransactionsSponsorPubKey  string `json:"transactionsSponsorPubKey"`
}
