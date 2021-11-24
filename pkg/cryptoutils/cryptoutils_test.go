// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cryptoutils_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/pkg/cryptoutils"
)

func TestCryptoUtils(t *testing.T) {
	value1 := map[string]string{
		"addressWallet":   "0xe2B32824733d350845c056CedD73c491FC4C1585",
		"addressContract": "0x0c80417acb4b309725de29b1d950bca974120996",
		"privateKey":      "5aefce0a2d473f59578fa7dee6a122d6509af1e0f79fcbee700dfcfeddabe4cc",
		"signature":       "707fb93c61be8d54c6d1fdf4b83c8642831c480194f7cc93ebdd6fe1ac7474ae63efd077cf6398bf00dc0f7ea96be9f3f9a05dfac1382c4d2f1bb11ec46148491b",
	}

	value2 := map[string]string{
		"addressWallet":   "0xe2B32824733d350845c056CedD73c491FC4C1585",
		"addressContract": "0x02a061be81ee0d7dbbd972bc7edee30b7b102a40",
		"privateKey":      "5aefce0a2d473f59578fa7dee6a122d6509af1e0f79fcbee700dfcfeddabe4cc",
		"signature":       "c2ab46d5981fe2fe951a63240af99a639ca13ea9afa1ee640d418690eae178215571ef9d318cc5b79d2c9e184e036f513aea7edcd6517cca75c4790d9ead45fc1b",
	}
	var tokenID int64 = 2

	t.Run("GenerateSignature", func(t *testing.T) {
		privateKeyECDSA, err := crypto.HexToECDSA(value1["privateKey"])
		require.NoError(t, err)

		signature, err := cryptoutils.GenerateSignature(
			cryptoutils.Address(value1["addressWallet"]),
			cryptoutils.Address(value1["addressContract"]),
			privateKeyECDSA,
		)
		require.NoError(t, err)
		assert.Equal(t, signature, cryptoutils.Signature(value1["signature"]))
	})

	t.Run("GenerateSignatureWithToken", func(t *testing.T) {
		privateKeyECDSA, err := crypto.HexToECDSA(value2["privateKey"])
		require.NoError(t, err)

		signature, err := cryptoutils.GenerateSignatureWithToken(
			cryptoutils.Address(value2["addressWallet"]),
			cryptoutils.Address(value2["addressContract"]),
			tokenID,
			privateKeyECDSA,
		)
		require.NoError(t, err)
		assert.Equal(t, signature, cryptoutils.Signature(value2["signature"]))
	})
}
