// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cryptoutils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"
)

// ErrCreateSignature indicates that there was an error in the cryptoutils package.
var ErrCreateSignature = errs.Class("create signature error")

// CreateSignature entity describes values for create signature.
type CreateSignature struct {
	Values     [][]byte
	PrivateKey *ecdsa.PrivateKey
}

// Signature defines signature type.
type Signature string

// EthereumSignedMessage defines message for signature.
const EthereumSignedMessage string = "\x19Ethereum Signed Message:\n"

// SignHash is a function that calculates a hash for the given message.
func SignHash(data []byte) []byte {
	msg := fmt.Sprintf(EthereumSignedMessage+"%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

// GenerateSignature generates signature for user's wallet.
func GenerateSignature(addressWallet Address, addressContract Address, privateKey *ecdsa.PrivateKey) (Signature, error) {
	var values [][]byte
	if !addressWallet.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of user's wallet")
	}
	if !addressContract.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte)
	createSignature := CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := makeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := reformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// GenerateSignatureWithValue generates signature for user's wallet with value.
func GenerateSignatureWithValue(addressWallet Address, addressContract Address, value int64, privateKey *ecdsa.PrivateKey) (Signature, error) {
	var values [][]byte
	if !addressWallet.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of user's wallet")
	}
	if !addressContract.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := createHexStringFixedLength(new(big.Int).SetInt64(value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte)
	createSignature := CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := makeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := reformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// GenerateSignatureWithValueAndNonce generates signature for user's wallet with value and nonce.
func GenerateSignatureWithValueAndNonce(addressWallet Address, addressContract Address, value *big.Int, nonce int64, privateKey *ecdsa.PrivateKey) (Signature, error) {
	var values [][]byte
	if !addressWallet.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of user's wallet")
	}
	if !addressContract.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := createHexStringFixedLength(value)
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	nonceStringWithZeros := createHexStringFixedLength(new(big.Int).SetInt64(nonce))
	nonceByte, err := hex.DecodeString(string(nonceStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte, nonceByte)
	createSignature := CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := makeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := reformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// GenerateSignatureWithTokenIDAndValue generates signature for user's wallet with tokenID and value.
func GenerateSignatureWithTokenIDAndValue(addressWallet Address, addressSaleContract Address, addressNFTContract Address, tokenID int64, value *big.Int, privateKey *ecdsa.PrivateKey) (Signature, error) {
	var values [][]byte
	if !addressWallet.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of user's wallet")
	}
	if !addressSaleContract.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of sale contract")
	}
	if !addressNFTContract.IsValidAddress() {
		return "", ErrCreateSignature.New("invalid address of nft contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressSaleContractByte, err := hex.DecodeString(string(addressSaleContract)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressNFTContractByte, err := hex.DecodeString(string(addressNFTContract)[LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	tokenIDStringWithZeros := createHexStringFixedLength(new(big.Int).SetInt64(tokenID))
	tokenIDByte, err := hex.DecodeString(string(tokenIDStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := createHexStringFixedLength(value)
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressSaleContractByte, addressNFTContractByte, tokenIDByte, valueByte)
	createSignature := CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := makeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := reformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// makeSignatureWithToken makes signature from addresses, private key and token id.
func makeSignature(createSignature CreateSignature) ([]byte, error) {
	var allValues []byte
	for _, value := range createSignature.Values {
		allValues = append(allValues, value...)
	}
	dataSignature := SignHash(crypto.Keccak256Hash(allValues).Bytes())
	signature, err := crypto.Sign(dataSignature, createSignature.PrivateKey)
	return signature, ErrCreateSignature.Wrap(err)
}

// reformSignature reforms last two byte of signature from 00, 01 to 1b, 1c.
func reformSignature(signatureByte []byte) (Signature, error) {
	signatureWithoutEnd := string(signatureByte)[:len(signatureByte)-1]
	signatureString := hex.EncodeToString(signatureByte)
	signatureLastSymbol := signatureString[len(signatureString)-1:]

	if signatureLastSymbol == fmt.Sprintf("%d", PrivateKeyVZero) {
		return Signature(hex.EncodeToString(append([]byte(signatureWithoutEnd), []byte{byte(PrivateKeyVTwentySeven)}...))), nil
	}

	if signatureLastSymbol == fmt.Sprintf("%d", PrivateKeyVOne) {
		return Signature(hex.EncodeToString(append([]byte(signatureWithoutEnd), []byte{byte(PrivateKeyVTwentyEight)}...))), nil
	}

	return "", ErrCreateSignature.New("error private key format")
}
