// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cryptoutils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// EthereumSignedMessageHash defines message for signature.
const EthereumSignedMessageHash string = "19457468657265756d205369676e6564204d6573736167653a0a3332"

// CreateSignature entity describes values for create signature.
type CreateSignature struct {
	EthereumSignedMessage []byte
	Values                [][]byte
	PrivateKey            *ecdsa.PrivateKey
}

// Address defines address type.
type Address string

// CreateValidAddress creates valid address.
func CreateValidAddress(address Hex) Address {
	return Address(HexPrefix + address[LengthOneBlockInputValue-LengthAddress+LengthHexPrefix:])
}

// Hex defines hex type.
type Hex string

// Signature defines signature type.
type Signature string

// PrivateKey defines private key type.
type PrivateKey string

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

// Contract entity describes addresses of contract and method.
type Contract struct {
	Address       Address `json:"address"`
	AddressMethod Hex     `json:"addressMethod"`
}

// Chain defines the list of possible chains in blockchain.
type Chain string

const (
	// ChainEthereum indicates that chain is ethereum.
	ChainEthereum Chain = "ethereum"
	// ChainPolygon indicates that chain is polygon.
	ChainPolygon Chain = "polygon"
)

// ChainID defines the list of possible number chains in blockchain.
type ChainID int

const (
	// ChainIDRinkeby indicates that chain id is 4.
	ChainIDRinkeby ChainID = 4
)

// HexPrefix defines the prefix of hex type.
const HexPrefix Hex = "0x"

// Length defines the list of possible lengths of blockchain elements.
type Length int

const (
	// LengthPrivateKey defines length private key.
	LengthPrivateKey Length = 64
	// LengthOneBlockInputValue defines the length of one block of input data.
	LengthOneBlockInputValue Length = 64
	// LengthAddress defines the length of address.
	LengthAddress Length = 40
	// LengthHexPrefix defines the length of hex prefix.
	LengthHexPrefix Length = 2
)

// BlockTag defines the list of possible block tags in blockchain.
type BlockTag string

// BlockTagLatest indicates that the last block will be used.
const BlockTagLatest BlockTag = "latest"

// IsValidAddress checks if the address is valid.
func (address Address) IsValidAddress() bool {
	return common.IsHexAddress(string(address))
}

// IsValidPrivateKey validates whether each byte is valid hexadecimal private key.
func (privateKey PrivateKey) IsValidPrivateKey() bool {
	if Length(len(string(privateKey))) != LengthPrivateKey {
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

// Data entity describes values for data field in transacton.
type Data struct {
	AddressContractMethod Hex
	TokenID               int64
}

// NewDataHex is a constructor for data entity, but returns hex string.
func NewDataHex(data Data) Hex {
	tokenID := createHexStringFixedLength(new(big.Int).SetInt64(data.TokenID))
	return data.AddressContractMethod + tokenID
}

// GenerateSignature generates signature for user's wallet.
func GenerateSignature(addressWallet Address, addressContract Address, privateKey *ecdsa.PrivateKey) (Signature, error) {
	var values [][]byte
	if !addressWallet.IsValidAddress() {
		return "", fmt.Errorf("invalid address of user's wallet")
	}
	if !addressContract.IsValidAddress() {
		return "", fmt.Errorf("invalid address of contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[LengthHexPrefix:])
	if err != nil {
		return "", err
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[LengthHexPrefix:])
	if err != nil {
		return "", err
	}

	ethereumSignedMessage, err := hex.DecodeString(EthereumSignedMessageHash)
	if err != nil {
		return "", err
	}

	values = append(values, addressWalletByte, addressContractByte)
	createSignature := CreateSignature{
		EthereumSignedMessage: ethereumSignedMessage,
		Values:                values,
		PrivateKey:            privateKey,
	}

	signatureByte, err := makeSignature(createSignature)
	if err != nil {
		return "", err
	}

	return reformSignature(signatureByte)
}

// GenerateSignatureWithValue generates signature for user's wallet with value.
func GenerateSignatureWithValue(addressWallet Address, addressContract Address, value int64, privateKey *ecdsa.PrivateKey) (Signature, error) {
	var values [][]byte
	if !addressWallet.IsValidAddress() {
		return "", fmt.Errorf("invalid address of user's wallet")
	}
	if !addressContract.IsValidAddress() {
		return "", fmt.Errorf("invalid address of contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[LengthHexPrefix:])
	if err != nil {
		return "", err
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[LengthHexPrefix:])
	if err != nil {
		return "", err
	}

	ethereumSignedMessage, err := hex.DecodeString(EthereumSignedMessageHash)
	if err != nil {
		return "", err
	}

	valueStringWithZeros := createHexStringFixedLength(new(big.Int).SetInt64(value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", err
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte)
	createSignature := CreateSignature{
		EthereumSignedMessage: ethereumSignedMessage,
		Values:                values,
		PrivateKey:            privateKey,
	}

	signatureByte, err := makeSignature(createSignature)
	if err != nil {
		return "", err
	}

	return reformSignature(signatureByte)
}

// GenerateSignatureWithValueAndNonce generates signature for user's wallet with value and nonce.
func GenerateSignatureWithValueAndNonce(addressWallet Address, addressContract Address, value *big.Int, nonce int64, privateKey *ecdsa.PrivateKey) (Signature, error) {
	var values [][]byte
	if !addressWallet.IsValidAddress() {
		return "", fmt.Errorf("invalid address of user's wallet")
	}
	if !addressContract.IsValidAddress() {
		return "", fmt.Errorf("invalid address of contract")
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[LengthHexPrefix:])
	if err != nil {
		return "", err
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[LengthHexPrefix:])
	if err != nil {
		return "", err
	}

	ethereumSignedMessage, err := hex.DecodeString(EthereumSignedMessageHash)
	if err != nil {
		return "", err
	}

	valueStringWithZeros := createHexStringFixedLength(value)
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", err
	}

	nonceStringWithZeros := createHexStringFixedLength(new(big.Int).SetInt64(nonce))
	nonceByte, err := hex.DecodeString(string(nonceStringWithZeros))
	if err != nil {
		return "", err
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte, nonceByte)
	createSignature := CreateSignature{
		EthereumSignedMessage: ethereumSignedMessage,
		Values:                values,
		PrivateKey:            privateKey,
	}

	signatureByte, err := makeSignature(createSignature)
	if err != nil {
		return "", err
	}

	return reformSignature(signatureByte)
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

	return "", fmt.Errorf("error private key format")
}

// createHexStringFixedLength creates srings with fixed length and number in hex formate in the end.
func createHexStringFixedLength(value *big.Int) Hex {
	valueString := fmt.Sprintf("%x", value)
	var zeroString string
	for i := 0; i < (int(LengthOneBlockInputValue) - len(valueString)); i++ {
		zeroString += "0"
	}

	return Hex(zeroString + valueString)
}

// makeSignatureWithToken makes signature from addresses, private key and token id.
func makeSignature(createSignature CreateSignature) ([]byte, error) {
	var allValues []byte
	for _, value := range createSignature.Values {
		allValues = append(allValues, value...)
	}
	dataSignature := crypto.Keccak256Hash(append(createSignature.EthereumSignedMessage, crypto.Keccak256Hash(allValues).Bytes()...))
	signature, err := crypto.Sign(dataSignature.Bytes(), createSignature.PrivateKey)
	return signature, err
}
