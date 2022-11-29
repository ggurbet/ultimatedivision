// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package signer_test

import (
	"math/big"
	"testing"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/pkg/signer"
)

func TestSignature(t *testing.T) {
	privateKey := "5aefce0a2d473f59578fa7dee6a122d6509af1e0f79fcbee700dfcfeddabe4cc"

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	require.NoError(t, err)

	tokenID := uuid.MustParse("94b94d50-d001-4f88-b7cf-1763b39044b1")

	contractAddress := signer.Address("c3d2bdedf7f309e2908ecf90d0dfb44acf2a8077cf053a05779fb15bbbfdbfb9")
	casperContractAddress := signer.Address("7238b9f78d4c88232c56a8e7609a437d3c865fff5dbb9d1f8aeab0508f1a5bc1")

	wallet := "0x56f088767D91badc379155290c4205c7b917a36E"
	casperWallet := "9060c0820b5156b1620c8e3344d17f9fad5108f5dc2672f2308439e84363c88e"

	expectedSignature := evmsignature.Signature("e0a3a3bcb0de3941ed598c49d227a61e5d4e2d1eccaa554b22655882a43503ff333b3590d8032fb212bddd4ce4dea24a42006389288872dc79df30427e64610d1c")
	expectedCasperSignature := evmsignature.Signature("41e3fb37d58cf7e5dbcb3daa11ca02ed540c7c0c01c934ba1bad76f206d8e7b57cc56738e6add643343cf61a4d1dfffada9a1580fc14d56e76aee0f29805a2d21c")
	expectedNonceSignature := evmsignature.Signature("d1aee17e45885b5cb29a155a6303784ac0c8a840ed3135f611252540af1d8a6a5a12fac77cba974eb0e4383a74551bd516c1f3e15bdc7c583a3ea12118e38ec71b")
	expectedCasperNonceSignature := evmsignature.Signature("c9cac1d7d80fbc431b54fce157d7d863d52147db4f27270deb3c3bc41e1fc2482178571df8be156d656d46f9a81132173561eecd6b3313ae511312fa7e8c4fbb1b")

	value := *big.NewInt(100)
	nonce := int64(5)

	t.Run("GenerateSignatureWithValue", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValue(signer.Address(wallet), contractAddress, tokenID, privateKeyECDSA)
		assert.Equal(t, expectedSignature, signature)
		require.NoError(t, err)
	})

	t.Run("GenerateCasperSignatureWithValue", func(t *testing.T) {
		signature, err := signer.GenerateCasperSignatureWithValue(signer.Address(casperWallet), casperContractAddress, tokenID, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedCasperSignature, signature)
	})

	t.Run("GenerateSignatureWithValueAndNonce", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(wallet), contractAddress, &value, nonce, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedNonceSignature, signature)
	})

	t.Run("GenerateCasperSignatureWithValueAndNonce", func(t *testing.T) {
		signature, err := signer.GenerateCasperSignatureWithValueAndNonce(signer.Address(casperWallet), contractAddress, &value, nonce, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedCasperNonceSignature, signature)
	})

	t.Run("Negative GenerateSignatureWithValueAndNonce", func(t *testing.T) {
		nonce = -5
		_, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(casperWallet), contractAddress, &value, nonce, privateKeyECDSA)
		require.Error(t, err)

		assert.True(t, signer.ErrCreateSignature.Has(err))
	})

}
