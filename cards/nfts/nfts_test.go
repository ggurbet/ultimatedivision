// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts_test

import (
	"context"
	"testing"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/users"
)

func TestNFTs(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "3560876@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "qwerty",
		FirstName:    "Stas",
		LastName:     "Isakov",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	card1 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro yak muk",
		Quality:          "wood",
		Height:           178.8,
		Weight:           72.2,
		DominantFoot:     "left",
		IsTattoo:         false,
		Status:           cards.StatusActive,
		Type:             cards.TypeWon,
		UserID:           user1.ID,
		Tactics:          1,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offence:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
	}

	card2 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Vova",
		Quality:          "gold",
		Height:           179.9,
		Weight:           73.3,
		DominantFoot:     "right",
		IsTattoo:         true,
		Status:           cards.StatusSale,
		UserID:           user1.ID,
		Tactics:          2,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offence:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
	}

	nft1 := nfts.NFT{
		CardID:        card1.ID,
		TokenID:       1,
		Chain:         evmsignature.ChainEthereum,
		WalletAddress: evmsignature.Address("0x96216849c49358b10257cb55b28ea603c874b05e"),
	}

	nft2 := nfts.NFT{
		CardID:        card2.ID,
		TokenID:       2,
		Chain:         evmsignature.ChainPolygon,
		WalletAddress: evmsignature.Address("0x96216849c49358B10254cb55b28eA603c874b05E"),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryUsers := db.Users()
		repositoryCards := db.Cards()
		repositoryNFTs := db.NFTs()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryCards.Create(ctx, card1)
			require.NoError(t, err)

			err = repositoryCards.Create(ctx, card2)
			require.NoError(t, err)

			err = repositoryNFTs.Create(ctx, nft1)
			require.NoError(t, err)

			err = repositoryNFTs.Create(ctx, nft2)
			require.NoError(t, err)
		})

		t.Run("Get", func(t *testing.T) {
			nftGet, err := repositoryNFTs.Get(ctx, nft1.TokenID, nft1.Chain)
			require.NoError(t, err)

			compareNFTsSlice(t, []nfts.NFT{nftGet}, []nfts.NFT{nft1})
		})

		t.Run("List", func(t *testing.T) {
			nftList, err := repositoryNFTs.List(ctx)
			require.NoError(t, err)

			compareNFTsSlice(t, nftList, []nfts.NFT{nft1, nft2})
		})

		t.Run("Update", func(t *testing.T) {
			nft1.WalletAddress = "0x"
			err := repositoryNFTs.Update(ctx, nft1)
			require.NoError(t, err)

			nftGet, err := repositoryNFTs.Get(ctx, nft1.TokenID, nft1.Chain)
			require.NoError(t, err)

			compareNFTsSlice(t, []nfts.NFT{nftGet}, []nfts.NFT{nft1})
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryNFTs.Delete(ctx, card1.ID)
			require.NoError(t, err)

			nftList, err := repositoryNFTs.List(ctx)
			require.NoError(t, err)

			compareNFTsSlice(t, nftList, []nfts.NFT{nft2})
		})
	})
}

func compareNFTsSlice(t *testing.T, nft1, nft2 []nfts.NFT) {
	assert.Equal(t, len(nft1), len(nft2))

	for i := 0; i < len(nft1); i++ {
		assert.Equal(t, nft1[i].CardID, nft2[i].CardID)
		assert.Equal(t, nft1[i].TokenID, nft2[i].TokenID)
		assert.Equal(t, nft1[i].Chain, nft2[i].Chain)
		assert.Equal(t, nft1[i].WalletAddress, nft2[i].WalletAddress)
	}
}
