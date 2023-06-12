// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Buffer } from 'buffer';
import { JsonTypes } from 'typedjson';
import { CLPublicKey, CLValueBuilder, DeployUtil, RuntimeArgs, decodeBase16 } from 'casper-js-sdk';

import { CasperNetworkClient } from '@/api/casper';
import { CasperMatchTransaction } from '@/matches';
import { ToastNotifications } from '@/notifications/service';
import {
    ACCEPT_OFFER_PAYMENT_AMOUNT,
    ACCOUNT_HASH_PREFIX,
    APPROVE_NFT_PAYMENT_AMOUNT,
    APPROVE_TOKEN_PAYMENT_AMOUNT,
    BUY_OFFER_PAYMENT_AMOUNT,
    BidsMakeOfferTransaction,
    CHAIN_NAME,
    CREATE_LOT_PAYMENT_AMOUNT,
    CasperSeasonRewardTransaction,
    CasperTransactionApprove,
    CasperTransactionIdentificators,
    GAS_PRICE,
    MAKE_OFFER_PAYMENT_AMOUNT,
    MINT_ONE_PAYMENT_AMOUNT,
    MarketCreateLotTransaction,
    OfferTransaction,
    PAYMENT_AMOUNT,
    TOKEN_PAYMENT_AMOUNT,
    TTL,
} from '@/casper/types';

enum CasperRuntimeArgs {
    RECIPIENT = 'recipient',
    TOKEN_ID = 'token_id',
    NFT_CONTRACT_HASH = 'nft_contract_hash',
    MIN_BID_PRICE = 'min_bid_price',
    REDEMPRION_PRICE = 'redemption_price',
    AUCTION_DURATION = 'auction_duration',
    VALUE = 'value',
    NONCE = 'nonce',
    SIGNATURE = 'signature',
    SPENDER = 'spender',
    OFFER_PRICE = 'offer_price',
    ERC20_CONTRACT = 'erc20_contract',
    AMOUNT = 'amount'
}

/** CasperTransactionService describes casper transaction entity. */
class CasperTransactionService {
    private readonly paymentAmount: number = PAYMENT_AMOUNT;
    private readonly gasPrice: number = GAS_PRICE;
    private readonly ttl: number = TTL;
    private readonly client: CasperNetworkClient = new CasperNetworkClient();
    public walletAddress: string = '';

    /** default CasperTransactionService implementation */
    constructor(walletAddress: string) {
        this.walletAddress = walletAddress;
    }

    /** Gets minting signature with contract address from api */
    async getTransaction(signature: CasperTransactionIdentificators): Promise<any> {
        return await this.client.getTransaction(signature);
    }

    /** Converts contract hash to bytes  */
    public static async convertContractHashToBytes(contractHash: string): Promise<Uint8Array> {
        return await Uint8Array.from(Buffer.from(contractHash, 'hex'));
    }

    /** Signs a contract */
    public async contractSign(
        entryPoint: string,
        runtimeArgs: RuntimeArgs,
        paymentAmount: number,
        contractAddress: string
    ): Promise<JsonTypes> {
        const contractHashToBytes = await CasperTransactionService.convertContractHashToBytes(contractAddress);

        const walletAddressConverted = CLPublicKey.fromHex(this.walletAddress);

        const deployParams = new DeployUtil.DeployParams(walletAddressConverted, CHAIN_NAME, this.gasPrice, this.ttl);

        const deploy = DeployUtil.makeDeploy(
            deployParams,
            DeployUtil.ExecutableDeployItem.newStoredContractByHash(
                contractHashToBytes,
                entryPoint,
                runtimeArgs),
            DeployUtil.standardPayment(paymentAmount)
        );

        const deployJson = DeployUtil.deployToJson(deploy);

        const signature = await window.casperlabsHelper.sign(deployJson, this.walletAddress, contractAddress);

        return signature;
    }

    /** Mints a nft */
    async mint(cardId: string): Promise<void> {
        try {
            const accountHash = CLPublicKey.fromHex(this.walletAddress).toAccountHashStr();
            const accountHashConverted = accountHash.replace(ACCOUNT_HASH_PREFIX, '');

            const nftWaitlist = await this.getTransaction(new CasperTransactionIdentificators(accountHashConverted, cardId));

            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(nftWaitlist.tokenId),
                [CasperRuntimeArgs.RECIPIENT]: CLValueBuilder.string(`account-hash-${accountHashConverted}`),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('mint_one', runtimeArgs, MINT_ONE_PAYMENT_AMOUNT, nftWaitlist.nftCreateCasperContract.address);

            await this.client.sendTx(nftWaitlist.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Mints tokens */
    async mintUDT(transaction: CasperMatchTransaction | CasperSeasonRewardTransaction, rpcNodeAddress: string): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.VALUE]: CLValueBuilder.u256(transaction.value),
                [CasperRuntimeArgs.NONCE]: CLValueBuilder.u64(transaction.nonce),
                [CasperRuntimeArgs.SIGNATURE]: CLValueBuilder.string(transaction.signature),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('claim', runtimeArgs, TOKEN_PAYMENT_AMOUNT, transaction.casperTokenContract.address);

            await this.client.sendTx(rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Approves nft minting */
    async approveNftMinting(transaction: CasperTransactionApprove): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(transaction.tokenId),
                [CasperRuntimeArgs.SPENDER]: CLValueBuilder.string(transaction.approveNftSpender),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('approve', runtimeArgs, APPROVE_NFT_PAYMENT_AMOUNT, transaction.NFTContractAddress);

            await this.client.sendTx(transaction.addressNodeServer, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Approves getting tokens  */
    async approveTokenRevard(transaction: CasperTransactionApprove): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.SPENDER]: CLValueBuilder.key(
                    CLValueBuilder.byteArray(decodeBase16(transaction.approveTokensSpender))
                ),
                [CasperRuntimeArgs.AMOUNT]: CLValueBuilder.u256(transaction.amount),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('approve', runtimeArgs, APPROVE_TOKEN_PAYMENT_AMOUNT, transaction.tokenRewardContractAddress);

            await this.client.sendTx(transaction.addressNodeServer, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }


    /** Creates a lot */
    async createLot(transaction: MarketCreateLotTransaction): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string(transaction.contractHash),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(transaction.tokenId),
                [CasperRuntimeArgs.MIN_BID_PRICE]: CLValueBuilder.u256(transaction.minBidPrice),
                [CasperRuntimeArgs.REDEMPRION_PRICE]: CLValueBuilder.u256(transaction.redemptionPrice),
                [CasperRuntimeArgs.AUCTION_DURATION]: CLValueBuilder.u256(transaction.auctionDuration),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('create_listing', runtimeArgs, CREATE_LOT_PAYMENT_AMOUNT, transaction.address);

            await this.client.sendTx(transaction.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Accepts offer */
    async acceptOffer(transaction: OfferTransaction): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string(transaction.contractHash),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(transaction.tokenId),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('accept_offer', runtimeArgs, ACCEPT_OFFER_PAYMENT_AMOUNT, transaction.address);

            await this.client.sendTx(transaction.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Makes offer */
    async makeOffer(transaction: BidsMakeOfferTransaction): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string(transaction.contractHash),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(transaction.tokenId),
                [CasperRuntimeArgs.OFFER_PRICE]: CLValueBuilder.u256(transaction.offerPrice),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('make_offer', runtimeArgs, MAKE_OFFER_PAYMENT_AMOUNT, transaction.address);

            await this.client.sendTx(transaction.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Buys listing */
    async buyListing(transaction: OfferTransaction): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.NFT_CONTRACT_HASH]: CLValueBuilder.string(transaction.contractHash),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(transaction.tokenId),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('buy_listing', runtimeArgs, BUY_OFFER_PAYMENT_AMOUNT, transaction.address);

            await this.client.sendTx(transaction.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }
}

export default CasperTransactionService;
