// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
	"ultimatedivision/users"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// Address defines address type.
type Address string

// ChoreConfig is the global configuration for nftsigner.
type ChoreConfig struct {
	RenewalInterval            time.Duration           `json:"renewalInterval"`
	PrivateKey                 evmsignature.PrivateKey `json:"privateKey"`
	NFTCreateContractAddress   common.Address          `json:"nftCreateContractAddress"`
	VelasSmartContractAddress  common.Address          `json:"velasSmartContractAddress"`
	CasperSmartContractAddress string                  `json:"casperSmartContractAddress"`
	CasperTokenContract        string                  `json:"casperTokenContract"`
}

// Chore requests for unsigned nft tokens and sign all of them .
//
// architecture: Chore
type Chore struct {
	loop   *thelooper.Loop
	config ChoreConfig
	nfts   *waitlist.Service
}

// NewChore instantiates Chore.
func NewChore(config ChoreConfig, db waitlist.DB) *Chore {
	return &Chore{
		loop:   thelooper.NewLoop(config.RenewalInterval),
		config: config,
		nfts:   waitlist.NewService(waitlist.Config{}, db, nil, nil, nil, nil),
	}
}

// Run starts the chore for signing unsigned nft token from ultimatedivision.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.loop.Run(ctx, func(ctx context.Context) error {
		unsignedNFTs, err := chore.nfts.ListWithoutPassword(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}
		privateKeyECDSA, err := crypto.HexToECDSA(string(chore.config.PrivateKey))
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, token := range unsignedNFTs {
			var (
				signature           evmsignature.Signature
				smartContract       common.Address
				casperTokenContract string
				casperContract      string
			)

			switch token.WalletType {
			case users.WalletTypeETH:
				smartContract = chore.config.NFTCreateContractAddress
			case users.WalletTypeVelas:
				smartContract = chore.config.VelasSmartContractAddress
			case users.WalletTypeCasper:
				casperContract = chore.config.CasperSmartContractAddress
				casperTokenContract = chore.config.CasperTokenContract
			}

			if token.Value.Cmp(big.NewInt(0)) <= 0 {
				if casperContract != "" {
					signature, err = GenerateCasperSignature(token.CasperWallet,
						casperContract, token.TokenID, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				} else {
					signature, err = evmsignature.GenerateSignatureWithValue(evmsignature.Address(token.Wallet.String()),
						evmsignature.Address(smartContract.String()), token.TokenID, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				}
			} else {
				if casperContract != "" {
					signature, err = GenerateCasperWithValueAndNonce(token.CasperWallet,
						casperTokenContract, &token.Value, token.TokenID, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				} else {
					signature, err = evmsignature.GenerateSignatureWithValueAndNonce(evmsignature.Address(token.Wallet.String()),
						evmsignature.Address(smartContract.String()), &token.Value, token.TokenID, privateKeyECDSA)
					if err != nil {
						return ChoreError.Wrap(err)
					}
				}
			}

			if err = chore.nfts.Update(ctx, token.TokenID, signature); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}

// GenerateCasperSignature generates casper signature for user's wallet with value.
func GenerateCasperSignature(addressWallet string, addressContract string, value int64, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte

	addressWalletByte, err := hex.DecodeString(addressWallet)
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(addressContract)
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ChoreError.Wrap(err)
}

// GenerateCasperWithValueAndNonce generates signature for user's wallet with value and nonce.
func GenerateCasperWithValueAndNonce(addressWallet string, addressContract string, value *big.Int, nonce int64, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte

	addressWalletByte, err := hex.DecodeString(addressWallet)
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(addressContract)
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	nonceStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nonce))
	nonceByte, err := hex.DecodeString(string(nonceStringWithZeros))
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte, nonceByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ChoreError.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ChoreError.Wrap(err)
}
