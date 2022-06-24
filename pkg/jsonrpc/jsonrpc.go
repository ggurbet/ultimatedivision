// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package jsonrpc

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"
)

// Transaction entity describes the values required to send the transaction.
type Transaction struct {
	JSONRPC Version              `json:"jsonrpc"`
	Method  MethodEth            `json:"method"`
	Params  interface{}          `json:"params"`
	ID      evmsignature.ChainID `json:"id"`
}

// NewTransaction is a constructor for transaction entity.
func NewTransaction(method MethodEth, params interface{}, id evmsignature.ChainID) Transaction {
	return Transaction{
		JSONRPC: VersionTwo,
		Method:  method,
		Params:  params,
		ID:      id,
	}
}

// Parameter entity describes parameters of transaction.
type Parameter struct {
	To   common.Address   `json:"to"`
	Data evmsignature.Hex `json:"data"`
}

// CreateFilter entity describes the values required to create filter.
type CreateFilter struct {
	ToBlock evmsignature.BlockTag `json:"toBlock"`
	Address common.Address        `json:"address"`
	Topics  []evmsignature.Hex    `json:"topics"`
}

// ResponseAddress entity describes values of the response from server, where the result is equal to the address.
type ResponseAddress struct {
	ID      evmsignature.ChainID `json:"id"`
	JSONRPC Version              `json:"jsonrpc"`
	Result  evmsignature.Hex     `json:"result"`
}

// ResponseEvents entity describes values of the response from server, where the result is equal to the list events.
type ResponseEvents struct {
	ID      evmsignature.ChainID `json:"id"`
	JSONRPC Version              `json:"jsonrpc"`
	Result  []Event              `json:"result"`
}

// Event entity describes event which happens in blockchain when nft is minted or transferred.
type Event struct {
	LogIndex         evmsignature.Hex   `json:"logIndex"`
	BlockNumber      evmsignature.Hex   `json:"blockNumber"`
	BlockHash        evmsignature.Hex   `json:"blockHash"`
	TransactionHash  evmsignature.Hex   `json:"transactionHash"`
	TransactionIndex evmsignature.Hex   `json:"transactionIndex"`
	Address          common.Address     `json:"address"`
	Data             evmsignature.Hex   `json:"data"`
	Topics           []evmsignature.Hex `json:"topics"`
}

// Version defines the list of possible json rpc version of server.
type Version string

// VersionTwo indicates that json rpc version of server is 2.0.
const VersionTwo Version = "2.0"

// MethodHTTP defines the list of possible http methods.
type MethodHTTP string

// MethodHTTPPost indicates that http method is POST.
// HTTP-method POST is designed to send data to the server in the body of the request.
const MethodHTTPPost MethodHTTP = "POST"

// MethodEth defines the list of possible contract methods.
type MethodEth string

const (
	// MethodEthCall indicates that this name of contact method is eth_call.
	MethodEthCall MethodEth = "eth_call"
	// MethodEthNewFilter indicates that this name of contact method is eth_newFilter.
	MethodEthNewFilter MethodEth = "eth_newFilter"
	// MethodEthGetFilterChanges indicates that this name of contact method is eth_getFilterChanges.
	MethodEthGetFilterChanges MethodEth = "eth_getFilterChanges"
)

// Send sends request to server and returns body of response.
func Send(url string, transaction Transaction) (io.ReadCloser, error) {
	payloadJSON, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(payloadJSON))

	client := &http.Client{}
	req, err := http.NewRequest(string(MethodHTTPPost), url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		defer func() {
			err = errs.Combine(err, resp.Body.Close())
		}()
		return nil, err
	}

	return resp.Body, nil
}

// GetOwnersWalletAddress returns owner's wallet address.
func GetOwnersWalletAddress(body io.ReadCloser) (common.Address, error) {
	var (
		response ResponseAddress
		err      error
	)
	defer func() {
		err = errs.Combine(err, body.Close())
	}()

	err = json.NewDecoder(body).Decode(&response)
	validWalletAddress := evmsignature.CreateValidAddress(response.Result)
	return common.HexToAddress(string(validWalletAddress)), err
}

// GetAddressOfFilter returns address of new filter.
func GetAddressOfFilter(body io.ReadCloser) (common.Address, error) {
	var (
		response ResponseAddress
		err      error
	)
	defer func() {
		err = errs.Combine(err, body.Close())
	}()

	err = json.NewDecoder(body).Decode(&response)
	return common.HexToAddress(string(response.Result)), err
}

// ListEvents returns list events of filter.
func ListEvents(body io.ReadCloser) ([]Event, error) {
	var (
		response ResponseEvents
		err      error
	)
	defer func() {
		err = errs.Combine(err, body.Close())
	}()

	err = json.NewDecoder(body).Decode(&response)
	return response.Result, err
}

// GetEvents sends request and returns events of filter.
func GetEvents(addressNodeServer string, transaction Transaction) ([]Event, error) {
	body, err := Send(addressNodeServer, transaction)
	if err != nil {
		return nil, err
	}

	events, err := ListEvents(body)
	return events, err
}
