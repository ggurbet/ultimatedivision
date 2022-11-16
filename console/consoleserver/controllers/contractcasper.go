package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	contractcasper "ultimatedivision/pkg/contractcasper"
)

var (
	// ErrContractCasper is an internal error type for contract casper controller.
	ErrContractCasper = errs.Class("contract casper controller error")
)

// ContractCasper is a mvc controller that handles all contract casper related views.
type ContractCasper struct {
	log logger.Logger
}

// NewContractCasper constructor for contract.
func NewContractCasper(log logger.Logger) *ContractCasper {
	return &ContractCasper{
		log: log,
	}
}

// Claim sends transaction to claim method.
func (contract *ContractCasper) Claim(w http.ResponseWriter, r *http.Request) {
	var req contractcasper.ClaimRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		contract.serveError(w, http.StatusBadRequest, ErrContractCasper.Wrap(err))
		return
	}

	resp, err := contractcasper.Claim(r.Context(), req)
	if err != nil {
		contract.serveError(w, http.StatusInternalServerError, ErrContractCasper.Wrap(err))
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		contract.log.Error("failed to write json response", err)
		return
	}
}

// serveError replies to the request with specific code and error message.
func (contract *ContractCasper) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()
	if err = json.NewEncoder(w).Encode(response); err != nil {
		contract.log.Error("failed to write json error response", ErrContractCasper.Wrap(err))
	}
}
