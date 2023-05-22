// Copyright (C) 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package eventparsing_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"ultimatedivision/pkg/eventparsing"
)

func TestEventParsing(t *testing.T) {
	eventData := eventparsing.EventData{
		Bytes: "7c000000013c0c1847d1c410338ab9b4ee0919c181cf26085997ff9c797e8a1ae5b02ddf2306000000474f45524c49280000003330393566393535646137303062393632313563666663396263363461623265363965623764616203a0860100daa2b596e0a496b04933e241e0567f2bcbecc829aa57d88cab096c28fd07dee2",
	}

	expectedEventTyte := 1
	expectedTokenContractAddress := "3c0c1847d1c410338ab9b4ee0919c181cf26085997ff9c797e8a1ae5b02ddf23"
	expectedChainName := "GOERLI"
	expectedChainAddress := "3095f955da700b96215cffc9bc64ab2e69eb7dab"
	expectedAmount := 10520065
	expectedUserWalletAddress := "daa2b596e0a496b04933e241e0567f2bcbecc829aa57d88cab096c28fd07dee2"

	tokenEventData := eventparsing.EventData{
		Bytes: "0300000001014a0d0720000000ab2a7124e519694282d8f98a637153b9295dc93671d4630f518641286b6a3a8a2400000033653130393335302d626662322d343262642d613931352d386331336161363061623966",
	}

	tokenID, _ := uuid.Parse("3e109350-bfb2-42bd-a915-8c13aa60ab9f")

	t.Run("GetEventType", func(t *testing.T) {
		actualEventType, err := eventData.GetEventType()
		require.NoError(t, err)
		require.Equal(t, expectedEventTyte, actualEventType)
	})

	t.Run("GetTokenID", func(t *testing.T) {
		GetTokenID, err := eventData.GetTokenID(tokenEventData)
		require.NoError(t, err)
		require.Equal(t, tokenID, GetTokenID)
	})

	t.Run("GetTokenContractAddress", func(t *testing.T) {
		actualTokenContractAddress := eventData.GetTokenContractAddress()
		require.Equal(t, expectedTokenContractAddress, actualTokenContractAddress)
	})

	t.Run("GetChainName", func(t *testing.T) {
		actualChainName, err := eventData.GetChainName()
		require.NoError(t, err)
		require.Equal(t, expectedChainName, actualChainName)
	})

	t.Run("GetChainAddress", func(t *testing.T) {
		actualChainAddress, err := eventData.GetChainAddress()
		require.NoError(t, err)
		require.Equal(t, expectedChainAddress, actualChainAddress)
	})

	t.Run("GetAmount", func(t *testing.T) {
		actualAmount, err := eventData.GetAmount()
		require.NoError(t, err)
		require.Equal(t, expectedAmount, actualAmount)
	})

	t.Run("GetUserWalletAddress", func(t *testing.T) {
		actualUserWalletAddress := eventData.GetUserWalletAddress()
		require.Equal(t, expectedUserWalletAddress, actualUserWalletAddress)
	})
}
