// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/remotefilestorage/storj"
	contract "ultimatedivision/pkg/contractcasper"
	"ultimatedivision/pkg/imageprocessing"
	"ultimatedivision/users"
)

// ErrWaitlist indicated that there was an error in service.
var ErrWaitlist = errs.Class("waitlist service error")

// Service is handling waitList related logic.
//
// architecture: Service
type Service struct {
	config   Config
	waitList DB
	cards    *cards.Service
	avatars  *avatars.Service
	users    *users.Service
	nfts     *nfts.Service
	events   *http.Client
}

// NewService is a constructor for waitlist service.
func NewService(config Config, waitList DB, cards *cards.Service, avatars *avatars.Service, users *users.Service, nfts *nfts.Service) *Service {
	eventsClient := &http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}

	return &Service{
		config:   config,
		waitList: waitList,
		cards:    cards,
		avatars:  avatars,
		users:    users,
		nfts:     nfts,
		events:   eventsClient,
	}
}

// Create creates nft for wait list.
func (service *Service) Create(ctx context.Context, createNFT CreateNFT) (Transaction, error) {
	var transaction Transaction

	user, err := service.users.Get(ctx, createNFT.UserID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if len(createNFT.WalletAddress.String()) == 0 {
		createNFT.WalletAddress = user.Wallet
	}

	if len(createNFT.CasperWallet) == 0 {
		createNFT.CasperWallet = user.CasperWallet
	}

	card, err := service.cards.Get(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if createNFT.Value.Cmp(big.NewInt(0)) <= 0 {
		if card.UserID != createNFT.UserID {
			return transaction, ErrWaitlist.New("this card does not belongs to user")
		}
	}

	if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
		switch item.WalletType {
		case users.WalletTypeVelas:
			transaction = Transaction{
				Password:          item.Password,
				NFTCreateContract: NFTCreateContract(service.config.NFTCreateVelasContract),
				TokenID:           item.TokenID,
				Value:             item.Value,
				WalletType:        item.WalletType,
			}
		case users.WalletTypeCasper:
			transaction = Transaction{
				Password:                item.Password,
				NFTCreateCasperContract: service.config.NFTCreateCasperContract,
				TokenID:                 item.TokenID,
				Value:                   item.Value,
				WalletType:              item.WalletType,
				RPCNodeAddress:          service.config.RPCNodeAddress,
			}
		default:
			transaction = Transaction{
				Password:          item.Password,
				NFTCreateContract: service.config.NFTCreateContract,
				TokenID:           item.TokenID,
				Value:             item.Value,
				WalletType:        item.WalletType,
			}
		}

		return transaction, nil
	}

	image, err := service.avatars.GetImage(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	client, err := storj.NewClient(service.config.FileStorage)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	// TODO: add transaction and mb lock db.
	lastTokenID, err := service.GetLastTokenID(ctx)
	if err != nil {
		if !ErrNoItem.Has(err) {
			return transaction, ErrWaitlist.Wrap(err)
		}
	}

	nextTokenID := lastTokenID + 1

	if err = client.Upload(ctx, service.config.Bucket, fmt.Sprintf("%d.%s", nextTokenID, imageprocessing.TypeFilePNG), image); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	nft := service.nfts.Generate(ctx, card, fmt.Sprintf(service.config.URLToAvatar, nextTokenID))
	fileMetadata, err := json.MarshalIndent(nft, "", " ")
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if err = client.Upload(ctx, service.config.Bucket, fmt.Sprintf("%d.%s", nextTokenID, imageprocessing.TypeFileJSON), fileMetadata); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if user.WalletType != users.WalletTypeCasper {
		if err = service.users.UpdateWalletAddress(ctx, createNFT.WalletAddress, createNFT.UserID, user.WalletType); err != nil {
			if !users.ErrWalletAddressAlreadyInUse.Has(err) {
				return transaction, ErrWaitlist.Wrap(err)
			}
		}
	}

	item := Item{
		TokenID:          uuid.New(),
		CardID:           createNFT.CardID,
		Wallet:           createNFT.WalletAddress,
		WalletType:       user.WalletType,
		CasperWallet:     user.CasperWallet,
		CasperWalletHash: user.CasperWalletHash,
		Value:            createNFT.Value,
	}

	if err = service.waitList.Create(ctx, item); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	for range time.NewTicker(time.Millisecond * service.config.WaitListCheckSignature).C {
		if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
			switch item.WalletType {
			case users.WalletTypeVelas:
				transaction = Transaction{
					Password:          item.Password,
					NFTCreateContract: NFTCreateContract(service.config.NFTCreateVelasContract),
					TokenID:           item.TokenID,
					Value:             item.Value,
					WalletType:        item.WalletType,
				}
			case users.WalletTypeCasper:
				transaction = Transaction{
					Password:                item.Password,
					NFTCreateCasperContract: service.config.NFTCreateCasperContract,
					TokenID:                 item.TokenID,
					Value:                   item.Value,
					WalletType:              item.WalletType,
					RPCNodeAddress:          service.config.RPCNodeAddress,
				}
			default:
				transaction = Transaction{
					Password:          item.Password,
					NFTCreateContract: service.config.NFTCreateContract,
					TokenID:           item.TokenID,
					Value:             item.Value,
					WalletType:        item.WalletType,
				}
			}
			break
		}
	}

	return transaction, err
}

// GetByTokenID returns nft for wait list by token id.
func (service *Service) GetByTokenID(ctx context.Context, tokenNumber int64) (Item, error) {
	nft, err := service.waitList.GetByTokenID(ctx, tokenNumber)
	return nft, ErrWaitlist.Wrap(err)
}

// GetByCardID returns nft for wait list by card id.
func (service *Service) GetByCardID(ctx context.Context, cardID uuid.UUID) (Item, error) {
	nft, err := service.waitList.GetByCardID(ctx, cardID)
	return nft, ErrWaitlist.Wrap(err)
}

// GetLastTokenID returns id of latest nft for wait list.
func (service *Service) GetLastTokenID(ctx context.Context) (int64, error) {
	lastID, err := service.waitList.GetLastTokenID(ctx)
	return lastID, ErrWaitlist.Wrap(err)
}

// List returns all nft for wait list.
func (service *Service) List(ctx context.Context) ([]Item, error) {
	allNFT, err := service.waitList.List(ctx)
	return allNFT, ErrWaitlist.Wrap(err)
}

// ListWithoutPassword returns nft for wait list without password.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Item, error) {
	nftWithoutPassword, err := service.waitList.ListWithoutPassword(ctx)
	return nftWithoutPassword, ErrWaitlist.Wrap(err)
}

// Update updates signature to nft token.
func (service *Service) Update(ctx context.Context, tokenID uuid.UUID, password evmsignature.Signature) error {
	return ErrWaitlist.Wrap(service.waitList.Update(ctx, tokenID, password))
}

// Delete deletes nft for wait list.
func (service *Service) Delete(ctx context.Context, tokenIDs []int64) error {
	return ErrWaitlist.Wrap(service.waitList.Delete(ctx, tokenIDs))
}

// GetNodeEvents is real time events streaming from blockchain.
func (service *Service) GetNodeEvents(ctx context.Context) (MintData, error) {
	var body io.Reader
	req, err := http.NewRequest(http.MethodGet, service.config.EventNodeAddress, body)
	if err != nil {
		return MintData{}, ErrWaitlist.Wrap(err)
	}

	resp, err := service.events.Do(req)
	if err != nil {
		defer func() {
			err = errs.Combine(err, resp.Body.Close())
		}()
		return MintData{}, ErrWaitlist.Wrap(err)
	}

	for {
		reader := bufio.NewReader(resp.Body)
		rawBody, err := reader.ReadBytes('\n')
		if err != nil {
			return MintData{}, ErrWaitlist.Wrap(err)
		}

		rawBody = []byte(strings.Replace(string(rawBody), "data:", "", 1))
		var event contract.Event
		_ = json.Unmarshal(rawBody, &event)

		transforms := event.DeployProcessed.ExecutionResult.Success.Effect.Transforms
		if len(transforms) == 0 {
			continue
		}

		var tokenID int64
		var walletAddress string
		for _, transform := range transforms {
			for _, i2 := range transform.Transform[WriteCLValueKey][Parsed] {
				switch i2.Key {
				case "token_id":
					tokenID, err = strconv.ParseInt(i2.Value, 10, 0)
					if err != nil {
						return MintData{}, ErrWaitlist.New("could not convert token_id from string to int64")
					}
				case "recipient":
					walletAddress = strings.ReplaceAll(i2.Value, "Key::Account(", "")
					walletAddress = strings.ReplaceAll(walletAddress, ")", "")
				default:
					continue
				}
			}
		}

		if tokenID != 0 && walletAddress != "" {
			return MintData{
				TokenID:       tokenID,
				WalletAddress: walletAddress,
			}, ErrWaitlist.Wrap(err)
		}
	}
}

// RunCasperCheckMintEvent runs a task to check and create the casper nft assignment.
func (service *Service) RunCasperCheckMintEvent(ctx context.Context) (err error) {
	event, err := service.GetNodeEvents(ctx)
	if err != nil {
		return err
	}

	if event.WalletAddress == "" {
		nftWaitList, err := service.GetByTokenID(ctx, event.TokenID)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		toAddress := common.HexToAddress(nftWaitList.CasperWalletHash)
		nft := nfts.NFT{
			CardID:        nftWaitList.CardID,
			Chain:         evmsignature.ChainEthereum,
			TokenID:       event.TokenID,
			WalletAddress: toAddress,
		}

		if err = service.nfts.Create(ctx, nft); err != nil {
			return ChoreError.Wrap(err)
		}

		user, err := service.users.GetByCasperHash(ctx, nftWaitList.CasperWalletHash)
		if err != nil {
			if err = service.nfts.Delete(ctx, nft.CardID); err != nil {
				return ChoreError.Wrap(err)
			}

			if err = service.cards.UpdateUserID(ctx, nft.CardID, uuid.Nil); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		if err = service.nfts.Update(ctx, nft); err != nil {
			return ChoreError.Wrap(err)
		}

		if err = service.cards.UpdateUserID(ctx, nft.CardID, user.ID); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	return ChoreError.Wrap(err)
}
