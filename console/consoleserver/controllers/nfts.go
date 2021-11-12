// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/users"
)

// ErrNFTs is an internal error type for NFTs controller.
var ErrNFTs = errs.Class("NFTs controller error")

// NFTs is a mvc controller that handles all NFTs related views.
type NFTs struct {
	log logger.Logger

	nfts *nfts.Service
}

// NewNFTs is a constructor for NFTs controller.
func NewNFTs(log logger.Logger, nfts *nfts.Service) *NFTs {
	nftsController := &NFTs{
		log:  log,
		nfts: nfts,
	}

	return nftsController
}

// Create is an endpoint that creates nft token.
func (controller *NFTs) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrNFTs.Wrap(err))
		return
	}

	var createNFT nfts.CreateNFT

	if err = json.NewDecoder(r.Body).Decode(&createNFT); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrNFTs.Wrap(err))
		return
	}
	createNFT.UserID = claims.UserID

	err = controller.nfts.Create(ctx, createNFT)
	if err != nil {
		controller.log.Error("could not create nft token", ErrNFTs.Wrap(err))

		if users.ErrNoUser.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrNFTs.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrNFTs.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *NFTs) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrCards.Wrap(err))
	}
}
