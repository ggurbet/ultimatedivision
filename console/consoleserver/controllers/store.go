// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/waitlist"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/store"
)

var (
	// ErrStore is an internal error type for store controller.
	ErrStore = errs.Class("store controller error")
)

// Store is a mvc controller that handles store related views.
type Store struct {
	log logger.Logger

	store *store.Service
}

// NewStore is a constructor for store controller.
func NewStore(log logger.Logger, store *store.Service) *Store {
	storeController := &Store{
		log:   log,
		store: store,
	}

	return storeController
}

// TransactionResponse entity describes values required to sent transaction.
type TransactionResponse struct {
	Password          evmsignature.Signature     `json:"password"`
	NFTCreateContract waitlist.NFTCreateContract `json:"nftCreateContract"`
	TokenID           int64                      `json:"tokenId"`
	Value             string                     `json:"value"`
}

// Buy is an endpoint that allows to view details of store.
func (controller *Store) Buy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrStore.Wrap(err))
		return
	}

	var request struct {
		CardID        uuid.UUID `json:"cardId"`
		WalletAddress string    `json:"walletAddress"`
		UserID        uuid.UUID `json:"userId"`
		Value         big.Int   `json:"value"`
	}

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrStore.Wrap(err))
		return
	}
	if !common.IsHexAddress(request.WalletAddress) {
		controller.serveError(w, http.StatusBadRequest, ErrStore.New("wallet address is invalid"))
		return
	}

	createNFT := waitlist.CreateNFT{
		CardID:        request.CardID,
		WalletAddress: common.HexToAddress(request.WalletAddress),
		UserID:        claims.UserID,
		Value:         request.Value,
	}

	transaction, err := controller.store.Buy(ctx, createNFT)
	if err != nil {
		switch {
		case cards.ErrNoCard.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrStore.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrStore.Wrap(err))
			controller.log.Error("could not buy card", ErrStore.Wrap(err))
		}
		return
	}

	transactionResponse := TransactionResponse{
		Password:          transaction.Password,
		NFTCreateContract: transaction.NFTCreateContract,
		TokenID:           transaction.TokenID,
		Value:             transaction.Value.String(),
	}

	if err = json.NewEncoder(w).Encode(transactionResponse); err != nil {
		controller.log.Error("failed to write json response", ErrStore.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Store) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrStore.Wrap(err))
	}
}
