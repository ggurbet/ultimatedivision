// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package jsonrpc

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// Transaction entity describes the values required to send the transaction.
type Transaction struct {
	JSONRPC Version             `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  []interface{}       `json:"params"`
	ID      cryptoutils.ChainID `json:"id"`
}

// Parameter entity describes parameters of transaction.
type Parameter struct {
	To   cryptoutils.Address `json:"to"`
	Data cryptoutils.Hex     `json:"data"`
}

// Response entity describes values of the response from server.
type Response struct {
	ID      cryptoutils.ChainID `json:"id"`
	JSONRPC Version             `json:"jsonrpc"`
	Result  cryptoutils.Hex     `json:"result"`
}

// NewTransaction is a constructor for transaction entity.
func NewTransaction(method string, params Parameter, id cryptoutils.ChainID) Transaction {
	return Transaction{
		JSONRPC: VersionTwo,
		Method:  method,
		Params:  []interface{}{&params, cryptoutils.BlockTagLatest},
		ID:      id,
	}
}

// Version defines the list of possible json rpc version of server.
type Version string

// VersionTwo indicates that json rpc version of server is 2.0.
const VersionTwo Version = "2.0"

// Method defines the list of possible http methods.
type Method string

// MethodPost indicates that http method is POST.
// HTTP-method POST is designed to send data to the server in the body of the request.
const MethodPost Method = "POST"

// EthCall indicates that this name of contact method is eth_call.
const EthCall string = "eth_call"

// Send sends request to server and returns result of response.
func Send(url string, payload io.Reader) (cryptoutils.Address, error) {
	client := &http.Client{}
	req, err := http.NewRequest(string(MethodPost), url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)

	return cryptoutils.Address(cryptoutils.HexPrefix + response.Result[cryptoutils.LengthOneBlockInputValue-cryptoutils.LengthAddress+cryptoutils.LengthHexPrefix:]), err
}
