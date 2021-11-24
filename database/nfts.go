// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts"
	"ultimatedivision/pkg/cryptoutils"
)

// ensures that nftsDB implements nfts.DB.
var _ nfts.DB = (*nftsDB)(nil)

// ErrNFTs indicates that there was an error in the database.
var ErrNFTs = errs.Class("ErrNFTs repository error")

// nftsDB provide access to nfts DB.
//
// architecture: Database
type nftsDB struct {
	conn *sql.DB
}

// Create creates nft in the database.
func (nftsDB *nftsDB) Create(ctx context.Context, nft nfts.NFT) error {
	query := `INSERT INTO nfts(card_id, token_id, chain, wallet_address)
	          VALUES($1,$2,$3,$4)`

	_, err := nftsDB.conn.ExecContext(ctx, query, nft.CardID, nft.TokenID, nft.Chain, nft.WalletAddress)
	return ErrNFTs.Wrap(err)
}

// Get returns nft by token id and chain from database.
func (nftsDB *nftsDB) Get(ctx context.Context, tokenID int64, chain cryptoutils.Chain) (nfts.NFT, error) {
	query := `
		SELECT 
			card_id, token_id, chain, wallet_address 
		FROM 
			nfts
		WHERE 
			token_id = $1 AND chain = $2`

	var nft nfts.NFT
	row := nftsDB.conn.QueryRowContext(ctx, query, tokenID, chain)

	err := row.Scan(&nft.CardID, &nft.TokenID, &nft.Chain, &nft.WalletAddress)
	if errors.Is(err, sql.ErrNoRows) {
		return nft, nfts.ErrNoNFTs.Wrap(err)
	}

	return nft, ErrNFTs.Wrap(err)
}

// List returns nfts from database.
func (nftsDB *nftsDB) List(ctx context.Context) ([]nfts.NFT, error) {
	var nftList []nfts.NFT
	query := `SELECT * FROM nfts`

	rows, err := nftsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nftList, ErrNFTs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var nft nfts.NFT

		if err = rows.Scan(&nft.CardID, &nft.TokenID, &nft.Chain, &nft.WalletAddress); err != nil {
			return nftList, ErrNFTs.Wrap(err)
		}
		nftList = append(nftList, nft)
	}

	return nftList, ErrNFTs.Wrap(rows.Err())
}

// Update updates users wallet address for nft token in the database.
func (nftsDB *nftsDB) Update(ctx context.Context, nft nfts.NFT) error {
	query := `UPDATE nfts
	          SET wallet_address = $1
	          WHERE chain = $2 AND token_id = $3`

	result, err := nftsDB.conn.ExecContext(ctx, query, nft.WalletAddress, nft.Chain, nft.TokenID)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err != nil {
		return ErrNFTs.Wrap(err)
	}
	if rowNum == 0 {
		return nfts.ErrNoNFTs.New("nft does not exist")
	}

	return ErrNFTs.Wrap(err)
}

// Delete deletes nft token in the database.
func (nftsDB *nftsDB) Delete(ctx context.Context, cardID uuid.UUID) error {
	result, err := nftsDB.conn.ExecContext(ctx, "DELETE FROM nfts WHERE card_id = $1", cardID)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err == nil && rowNum == 0 {
		return nfts.ErrNoNFTs.New("nft does not exist")
	}

	return ErrNFTs.Wrap(err)
}
