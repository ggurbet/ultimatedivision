// Copyright (C) 2021-2023 Creditor Corp. Group.
// See LICENSE for copying information.

package casper

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/casper-ecosystem/casper-golang-sdk/keypair"
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
	"github.com/casper-ecosystem/casper-golang-sdk/serialization"
	"github.com/casper-ecosystem/casper-golang-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	contract "ultimatedivision/pkg/contractcasper"
)

// Transfer describes sign func to sign transaction and casper client to send transaction.
type Transfer struct {
	casper contract.Casper

	sign func([]byte) ([]byte, error)
}

// NewTransfer is constructor for Transfer.
func NewTransfer(casper contract.Casper, sign func([]byte) ([]byte, error)) *Transfer {
	return &Transfer{
		casper: casper,
		sign:   sign,
	}
}

// MintOneRequest describes values to calls MintOne method.
type MintOneRequest struct {
	PublicKey              keypair.PublicKey
	ChainName              string
	StandardPayment        int64
	NFTContractPackageHash string
	TokenID                string
	Recipient              string
}

// MintOne mints NFT for user.
func (t *Transfer) MintOne(ctx context.Context, req MintOneRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	recipient := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.Recipient,
	}
	recipientBytes, err := serialization.Marshal(recipient)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
		"recipient": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(recipientBytes),
		},
	}

	keyOrder := []string{"token_id", "recipient"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.NFTContractPackageHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "mint_one", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// ApproveNFTRequest describes values to calls ApproveNFT method.
type ApproveNFTRequest struct {
	PublicKey              keypair.PublicKey
	ChainName              string
	StandardPayment        int64
	NFTContractPackageHash string
	Spender                string
	TokenID                string
}

// ApproveNFT approves NFT to send from one user to another.
func (t *Transfer) ApproveNFT(ctx context.Context, req ApproveNFTRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	spender := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.Spender,
	}
	spenderBytes, err := serialization.Marshal(spender)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"spender": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(spenderBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"spender", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.NFTContractPackageHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "approve", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// CreateListingRequest describes values to calls CreateListing method.
type CreateListingRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
	MinBidPrice        *big.Int
	RedemptionPrice    *big.Int
	AuctionDuration    *big.Int
}

// CreateListing creates listing for NFT.
func (t *Transfer) CreateListing(ctx context.Context, req CreateListingRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	minBidPrice := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.MinBidPrice,
	}
	minBidPriceBytes, err := serialization.Marshal(minBidPrice)
	if err != nil {
		return "", err
	}

	redemptionPrice := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.RedemptionPrice,
	}
	redemptionPriceBytes, err := serialization.Marshal(redemptionPrice)
	if err != nil {
		return "", err
	}

	auctionDuration := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.AuctionDuration,
	}
	auctionDurationBytes, err := serialization.Marshal(auctionDuration)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
		"min_bid_price": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(minBidPriceBytes),
		},
		"redemption_price": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(redemptionPriceBytes),
		},
		"auction_duration": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(auctionDurationBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id", "min_bid_price", "redemption_price", "auction_duration"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "create_listing", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// ApproveTokensRequest describes values to calls ApproveTokens method.
type ApproveTokensRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	TokensContractHash string
	Spender            common.Hash
	Amount             *big.Int
}

// ApproveTokens approves tokens to send from one user to another.
func (t *Transfer) ApproveTokens(ctx context.Context, req ApproveTokensRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	var spenderHashBytes [32]byte
	copy(spenderHashBytes[:], req.Spender.Bytes())

	spender := types.CLValue{
		Type: types.CLTypeKey,
		Key: &types.Key{
			Type: types.KeyTypeHash,
			Hash: spenderHashBytes,
		},
	}
	spenderBytes, err := serialization.Marshal(spender)
	if err != nil {
		return "", err
	}

	amount := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.Amount,
	}
	amountBytes, err := serialization.Marshal(amount)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"spender": {
			Tag:         types.CLTypeKey,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(spenderBytes),
		},
		"amount": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(amountBytes),
		},
	}

	keyOrder := []string{"spender", "amount"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.TokensContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "approve", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// MakeOfferRequest describes values to calls MakeOffer method.
type MakeOfferRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
	Amount             *big.Int
}

// MakeOffer makes offer for lot.
func (t *Transfer) MakeOffer(ctx context.Context, req MakeOfferRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	amount := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.Amount,
	}
	amountBytes, err := serialization.Marshal(amount)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
		"offer_price": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(amountBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id", "offer_price"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "make_offer", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// AcceptOfferRequest describes values to calls AcceptOffer method.
type AcceptOfferRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
}

// AcceptOffer accepts offer for lot.
func (t *Transfer) AcceptOffer(ctx context.Context, req AcceptOfferRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "accept_offer", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// BuyListingRequest describes values to calls BuyListing method.
type BuyListingRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
}

// BuyListing buys offer for lot at full price.
func (t *Transfer) BuyListing(ctx context.Context, req BuyListingRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "buy_listing", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// FinalListingRequest describes values to calls FinalListing method.
type FinalListingRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
}

// FinalListing finals listing for lot.
func (t *Transfer) FinalListing(ctx context.Context, req FinalListingRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "final_listing", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}
