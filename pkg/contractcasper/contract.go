// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package contract

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/casper-ecosystem/casper-golang-sdk/keypair"
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
	"github.com/zeebo/errs"
)

// ErrContract indicates that there was an error in the contract package.
var ErrContract = errs.Class("contract package")

type (
	// Event describes event structure in casper network.
	Event struct {
		DeployProcessed DeployProcessed `json:"DeployProcessed"`
	}
	// DeployProcessed describes all about deploy.
	DeployProcessed struct {
		DeployHash      string          `json:"deploy_hash"`
		Account         string          `json:"account"`
		BlockHash       string          `json:"block_hash"`
		ExecutionResult ExecutionResult `json:"execution_result"`
	}
	// ExecutionResult describes result.
	ExecutionResult struct {
		Success Success `json:"Success"`
	}
	// Success describes success result.
	Success struct {
		Effect Effect `json:"effect"`
	}
	// Effect describes.
	Effect struct {
		Transforms []Transform `json:"transforms"`
	}
	// Transform describes transform data.
	Transform struct {
		Key       string                      `json:"key"`
		Transform map[string]map[string][]Map `json:"transform"`
	}

	// Map describes struct with keys and values.
	Map struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)

type (
	// EventWithBytes describes event with bytes structure in casper network.
	EventWithBytes struct {
		DeployProcessed DeployProcessed2 `json:"DeployProcessed"`
	}
	// DeployProcessed2 describes all about deploy.
	DeployProcessed2 struct {
		DeployHash       string           `json:"deploy_hash"`
		Account          string           `json:"account"`
		BlockHash        string           `json:"block_hash"`
		ExecutionResult2 ExecutionResult2 `json:"execution_result"`
	}
	// ExecutionResult2 describes result.
	ExecutionResult2 struct {
		Success2 Success2 `json:"Success"`
	}
	// Success2 describes success result.
	Success2 struct {
		Effect2 Effect2 `json:"effect"`
	}
	// Effect2 describes.
	Effect2 struct {
		Transforms2 []Transform2 `json:"transforms"`
	}
	// Transform2 describes transform data.
	Transform2 struct {
		Key       string      `json:"key"`
		Transform interface{} `json:"transform"`
	}
)

// DeployProcessedNew describes transform data.
type DeployProcessedNew struct {
	DeployHash      string    `json:"deploy_hash"`
	Account         string    `json:"account"`
	Timestamp       time.Time `json:"timestamp"`
	TTL             string    `json:"ttl"`
	Dependencies    []string  `json:"dependencies"`
	BlockHash       string    `json:"block_hash"`
	ExecutionResult struct {
		Success struct {
			Effect struct {
				Operations []string `json:"operations"`
				Transforms []struct {
					Key       string      `json:"key"`
					Transform interface{} `json:"transform"`
				} `json:"transforms"`
			} `json:"effect"`
		} `json:"Success"`
	} `json:"execution_result"`
}

// Casper exposes access to the casper sdk methods.
type Casper interface {
	// PutDeploy deploys a contract or sends a transaction and returns deployment hash.
	PutDeploy(deploy sdk.Deploy) (string, error)
	// GetBlockNumberByHash returns block number by deploy hash.
	GetBlockNumberByHash(hash string) (int, error)
}

// StringNetworkAddress describes an address for some network.
type StringNetworkAddress struct {
	NetworkName string
	Address     string
}

// SendTxRequest describes values to initiate transaction.
type SendTxRequest struct {
	Deploy              string
	RPCNodeAddress      string
	CasperWalletAddress string
	CasperWalletHash    string
}

// SendTxResponse describes tx hash.
type SendTxResponse struct {
	Txhash string
}

// SendTx initiates transaction.
func SendTx(ctx context.Context, req SendTxRequest) (SendTxResponse, error) {
	request := struct {
		Deploy struct {
			Hash      sdk.Hash                  `json:"hash"`
			Header    *sdk.DeployHeader         `json:"header"`
			Payment   *sdk.ExecutableDeployItem `json:"payment"`
			Session   *sdk.ExecutableDeployItem `json:"session"`
			Approvals []struct {
				Signer    string `json:"signer"`
				Signature string `json:"signature"`
			} `json:"approvals"`
		}
	}{}

	err := json.Unmarshal([]byte(req.Deploy), &request)
	if err != nil {
		return SendTxResponse{}, ErrContract.Wrap(err)
	}

	pubKeyData, err := hex.DecodeString(request.Deploy.Approvals[0].Signer[2:])
	if err != nil {
		return SendTxResponse{}, ErrContract.Wrap(err)
	}

	signer := keypair.PublicKey{
		Tag:        request.Deploy.Header.Account.Tag,
		PubKeyData: pubKeyData,
	}

	signatureData, err := hex.DecodeString(request.Deploy.Approvals[0].Signature[2:])
	if err != nil {
		return SendTxResponse{}, ErrContract.Wrap(err)
	}

	signature := keypair.Signature{
		Tag:           request.Deploy.Header.Account.Tag,
		SignatureData: signatureData,
	}

	approval := sdk.Approval{
		Signer:    signer,
		Signature: signature,
	}

	deploy := sdk.Deploy{
		Hash:      request.Deploy.Hash,
		Header:    request.Deploy.Header,
		Payment:   request.Deploy.Payment,
		Session:   request.Deploy.Session,
		Approvals: []sdk.Approval{approval},
	}

	casperClient := sdk.NewRpcClient(req.RPCNodeAddress)
	deployResp, err := casperClient.PutDeploy(deploy)
	if err != nil {
		return SendTxResponse{}, ErrContract.Wrap(err)
	}

	resp := SendTxResponse{
		Txhash: deployResp.Hash,
	}

	return resp, nil
}

// rpcClient is a implementation of connector_service.Casper.
type rpcClient struct {
	client *sdk.RpcClient

	rpcNodeAddress string
}

// GetBlockNumberByHash returns block number by deploy hash.
func (r *rpcClient) GetBlockNumberByHash(hash string) (int, error) {
	blockResp, err := r.client.GetBlockByHash(hash)
	return blockResp.Header.Height, err
}

// JSONPutDeployRes describes result of put_deploy tx.
type JSONPutDeployRes struct {
	Hash string `json:"deploy_hash"`
}

// PutDeploy deploys a contract or sends a transaction and returns deployment hash.
func (r *rpcClient) PutDeploy(deploy sdk.Deploy) (string, error) {
	resp, err := r.rpcCall("account_put_deploy", map[string]interface{}{
		"deploy": deploy,
	})
	if err != nil {
		return "", err
	}

	var result JSONPutDeployRes
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return "", fmt.Errorf("failed to put deploy: %w", err)
	}

	return result.Hash, err
}

func (r *rpcClient) rpcCall(method string, params interface{}) (_ sdk.RpcResponse, err error) {
	var rpcResponse sdk.RpcResponse

	body, err := json.Marshal(sdk.RpcRequest{
		Version: "2.0",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return sdk.RpcResponse{}, ErrContract.Wrap(err)
	}

	resp, err := http.Post(r.rpcNodeAddress, "application/json", bytes.NewReader(body))
	if err != nil {
		return sdk.RpcResponse{}, fmt.Errorf("failed to make request: %v", err)
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return sdk.RpcResponse{}, fmt.Errorf("failed to get response body: %v", err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return sdk.RpcResponse{}, fmt.Errorf("request failed, status code - %d, response - %s", resp.StatusCode, string(b))
	}

	err = json.Unmarshal(b, &rpcResponse)
	if err != nil {
		return sdk.RpcResponse{}, fmt.Errorf("failed to parse response body: %v", err)
	}

	if rpcResponse.Error != nil {
		return rpcResponse, fmt.Errorf("rpc call failed, code - %d, message - %s", rpcResponse.Error.Code, rpcResponse.Error.Message)
	}

	return rpcResponse, nil
}

// New is constructor for rpcClient.
func New(rpcNodeAddress string) Casper {
	client := sdk.NewRpcClient(rpcNodeAddress)
	return &rpcClient{
		client:         client,
		rpcNodeAddress: rpcNodeAddress,
	}
}
