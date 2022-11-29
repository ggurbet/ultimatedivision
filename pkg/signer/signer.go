// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package signer

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrCreateSignature indicates that an error occurred while creating a signature.
var ErrCreateSignature = errs.Class("signature package error")

// Address defines address type.
type Address string

// Signature defines signature type.
type Signature string

// GenerateSignatureWithValue generates signature for user's wallet with value.
func GenerateSignatureWithValue(addressWallet Address, addressContract Address, value uuid.UUID, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte
	addressWalletByte, err := hex.DecodeString(string(addressWallet)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// GenerateCasperSignatureWithValue generates signature for user's wallet with value.
func GenerateCasperSignatureWithValue(addressWallet Address, addressContract Address, value uuid.UUID, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte
	addressWalletByte, err := hex.DecodeString(string(addressWallet))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// GenerateSignatureWithValueAndNonce generates signature for casper user's wallet with value and nonce.
func GenerateSignatureWithValueAndNonce(addressWallet Address, addressContract Address, value *big.Int, nonce int64, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte
	addressWalletByte, err := hex.DecodeString(string(addressWallet)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	nonceStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nonce))
	nonceByte, err := hex.DecodeString(string(nonceStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte, nonceByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// GenerateCasperSignatureWithValueAndNonce generates signature for casper user's wallet with value and nonce.
func GenerateCasperSignatureWithValueAndNonce(addressWallet Address, addressContract Address, value *big.Int, nonce int64, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte
	addressWalletByte, err := hex.DecodeString(string(addressWallet))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	nonceStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nonce))
	nonceByte, err := hex.DecodeString(string(nonceStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte, nonceByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// IsValidAddress checks if the address is valid.
func (address Address) IsValidAddress() error {
	if !common.IsHexAddress(string(address)) {
		return ErrCreateSignature.New("")
	}
	return nil
}
