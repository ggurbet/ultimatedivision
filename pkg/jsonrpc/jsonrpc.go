// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package jsonrpc

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// Transaction entity describes the values required to send the transaction.
type Transaction struct {
	JSONRPC Version             `json:"jsonrpc"`
	Method  MethodEth           `json:"method"`
	Params  interface{}         `json:"params"`
	ID      cryptoutils.ChainID `json:"id"`
}

// NewTransaction is a constructor for transaction entity.
func NewTransaction(method MethodEth, params interface{}, id cryptoutils.ChainID) Transaction {
	return Transaction{
		JSONRPC: VersionTwo,
		Method:  method,
		Params:  params,
		ID:      id,
	}
}

// Parameter entity describes parameters of transaction.
type Parameter struct {
	To   cryptoutils.Address `json:"to"`
	Data cryptoutils.Hex     `json:"data"`
}

// CreateFilter entity describes the values required to create filter.
type CreateFilter struct {
	ToBlock cryptoutils.BlockTag `json:"toBlock"`
	Address cryptoutils.Address  `json:"address"`
	Topics  []cryptoutils.Hex    `json:"topics"`
}

// ResponseAddress entity describes values of the response from server, where the result is equal to the address.
type ResponseAddress struct {
	ID      cryptoutils.ChainID `json:"id"`
	JSONRPC Version             `json:"jsonrpc"`
	Result  cryptoutils.Hex     `json:"result"`
}

// ResponseEvents entity describes values of the response from server, where the result is equal to the list events.
type ResponseEvents struct {
	ID      cryptoutils.ChainID `json:"id"`
	JSONRPC Version             `json:"jsonrpc"`
	Result  []Event             `json:"result"`
}

// Event entity describes event which happens in blockchain when nft is minted or transferred.
type Event struct {
	LogIndex         cryptoutils.Hex     `json:"logIndex"`
	BlockNumber      cryptoutils.Hex     `json:"blockNumber"`
	BlockHash        cryptoutils.Hex     `json:"blockHash"`
	TransactionHash  cryptoutils.Hex     `json:"transactionHash"`
	TransactionIndex cryptoutils.Hex     `json:"transactionIndex"`
	Address          cryptoutils.Address `json:"address"`
	Data             cryptoutils.Hex     `json:"data"`
	Topics           []cryptoutils.Hex   `json:"topics"`
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
func GetOwnersWalletAddress(body io.ReadCloser) (cryptoutils.Address, error) {
	var (
		response ResponseAddress
		err      error
	)
	defer func() {
		err = errs.Combine(err, body.Close())
	}()

	err = json.NewDecoder(body).Decode(&response)
	return cryptoutils.CreateValidAddress(response.Result), err
}

// GetAddressOfFilter returns address of new filter.
func GetAddressOfFilter(body io.ReadCloser) (cryptoutils.Address, error) {
	var (
		response ResponseAddress
		err      error
	)
	defer func() {
		err = errs.Combine(err, body.Close())
	}()

	err = json.NewDecoder(body).Decode(&response)
	return cryptoutils.Address(response.Result), err
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
