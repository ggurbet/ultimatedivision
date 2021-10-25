// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cryptoutils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// EthereumSignedMessageHash defines message for signature.
const EthereumSignedMessageHash string = "19457468657265756d205369676e6564204d6573736167653a0a3332"

// CreateSignature entity describes values for create signature.
type CreateSignature struct {
	EthereumSignedMessage []byte
	AddressWallet         []byte
	AddressContract       []byte
	PrivateKey            *ecdsa.PrivateKey
}

// Address defines address type.
type Address string

// Signature defines signature type.
type Signature string

// PrivateKey defines private key type.
type PrivateKey string

// LengthPrivateKey defines length private key.
const LengthPrivateKey int = 64

// PrivateKeyV defines v of private key type.
type PrivateKeyV int

const (
	// PrivateKeyVZero indicates that the v of private key is 0.
	PrivateKeyVZero PrivateKeyV = 0
	// PrivateKeyVOne indicates that the v of private key is 1.
	PrivateKeyVOne PrivateKeyV = 1
	// PrivateKeyVTwentySeven indicates that the v of private key is 27.
	PrivateKeyVTwentySeven PrivateKeyV = 27
	// PrivateKeyVTwentyEight indicates that the v of private key is 28.
	PrivateKeyVTwentyEight PrivateKeyV = 28
)

// IsValidAddress checks if the address is valid.
func (address Address) IsValidAddress() bool {
	return common.IsHexAddress(string(address))
}

// IsValidPrivateKey validates whether each byte is valid hexadecimal private key.
func (privateKey PrivateKey) IsValidPrivateKey() bool {
	if len(string(privateKey)) != LengthPrivateKey {
		return false
	}
	for _, c := range []byte(string(privateKey)) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// GenerateSignature generates signature for user's wallet.
func GenerateSignature(addressWallet Address, addressContract Address, privateKey *ecdsa.PrivateKey) (Signature, error) {
	if !addressWallet.IsValidAddress() {
		return "", fmt.Errorf("invalid address of user's wallet")
	}
	if !addressContract.IsValidAddress() {
		return "", fmt.Errorf("invalid address of contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[2:])
	if err != nil {
		return "", err
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[2:])
	if err != nil {
		return "", err
	}

	ethereumSignedMessage, err := hex.DecodeString(EthereumSignedMessageHash)
	if err != nil {
		return "", err
	}

	createSignature := CreateSignature{
		EthereumSignedMessage: ethereumSignedMessage,
		AddressWallet:         addressWalletByte,
		AddressContract:       addressContractByte,
		PrivateKey:            privateKey,
	}

	signature, err := makeSignature(createSignature)
	if err != nil {
		return "", err
	}

	signatureWithoutEnd := string(signature)[:len(signature)-1]
	signatureString := hex.EncodeToString(signature)
	signatureLastSymbol := signatureString[len(signatureString)-1:]

	if signatureLastSymbol == fmt.Sprintf("%d", PrivateKeyVZero) {
		return Signature(hex.EncodeToString(append([]byte(signatureWithoutEnd), []byte{byte(PrivateKeyVTwentySeven)}...))), nil
	}

	if signatureLastSymbol == fmt.Sprintf("%d", PrivateKeyVOne) {
		return Signature(hex.EncodeToString(append([]byte(signatureWithoutEnd), []byte{byte(PrivateKeyVTwentyEight)}...))), nil
	}

	return "", fmt.Errorf("error private key format")
}

// makeSignature makes signature from addresses and private key.
func makeSignature(createSignature CreateSignature) ([]byte, error) {
	dataSignature := crypto.Keccak256Hash(append(createSignature.EthereumSignedMessage, crypto.Keccak256Hash(append(createSignature.AddressWallet, createSignature.AddressContract...)).Bytes()...))
	signature, err := crypto.Sign(dataSignature.Bytes(), createSignature.PrivateKey)
	return signature, err
}
