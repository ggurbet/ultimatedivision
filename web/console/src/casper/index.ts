// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Buffer } from 'buffer';
import { JsonTypes } from 'typedjson';
import { CLPublicKey, CLValueBuilder, DeployUtil, RuntimeArgs } from 'casper-js-sdk';

import { CasperNetworkClient } from '@/api/casper';
import { CasperMatchTransaction } from '@/matches';
import { ToastNotifications } from '@/notifications/service';
import { MarketplaceClient } from '@/api/marketplace';
import { MarketCreateLotTransaction } from '@/marketplace';

enum CasperRuntimeArgs {
    SIGNATURE = 'signature',
    TOKEN_ID = 'token_id'
}

/** Describes parameters for transaction */
export class CasperTransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public casperWallet: string,
        public cardId: string
    ) { }
}

/** Describes parameters for casper token transaction */
export class CasperTokenContract {
    /** default CasperTokenContract implementation */
    constructor(
        public address: string = '0',
        public addressMethod: string = ''
    ) { }
}

/** Transaction describes transaction entity of match response. */
export class CasperSeasonRewardTransaction {
    /** Transaction contains of nonce, signature hash udtContract and value. */
    constructor(
        public ID: string,
        public userId: string,
        public seasonID: string,
        public walletAddress: string,
        public casperWalletAddress: string,
        public walleType: string,
        public status: number,
        public nonce: number,
        public signature: string,
        public value: string,
        public casperTokenContract: {
            address: string;
            addressMethod: string;
        },
        public rpcNodeAddress: string,
    ) { }
};

export const ACCOUNT_HASH_PREFIX = 'account-hash-';
const CHAIN_NAME = 'casper-test';

const TTL = 1800000;
const PAYMENT_AMOUNT = 50000000000;
const GAS_PRICE = 1;

const TOKEN_PAYMENT_AMOUNT = 6000000000;
const LOT_PAYMENT_AMOUNT = 40000000000;

/** CasperTransactionService describes casper transaction entity. */
class CasperTransactionService {
    private readonly paymentAmount: number = PAYMENT_AMOUNT;
    private readonly gasPrice: number = GAS_PRICE;
    private readonly ttl: number = TTL;
    private readonly client: CasperNetworkClient = new CasperNetworkClient();
    private readonly marketPlaceClient: MarketplaceClient = new MarketplaceClient();
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
                [CasperRuntimeArgs.SIGNATURE]: CLValueBuilder.string(nftWaitlist.password),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(nftWaitlist.tokenId),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('claim', runtimeArgs, this.paymentAmount, nftWaitlist.nftCreateCasperContract.address);

            await this.client.claim(nftWaitlist.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Mints a token */
    async mintUDT(transaction: CasperMatchTransaction | CasperSeasonRewardTransaction, rpcNodeAddress: string): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                'value': CLValueBuilder.u256(transaction.value),
                'nonce': CLValueBuilder.u64(transaction.nonce),
                'signature': CLValueBuilder.string(transaction.signature),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('claim', runtimeArgs, TOKEN_PAYMENT_AMOUNT, transaction.casperTokenContract.address);

            await this.client.claim(rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Creates a lot */
    async createLot(transaction: MarketCreateLotTransaction): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                'nft_contract_hash': CLValueBuilder.string(transaction.contractHash),
                'token_id': CLValueBuilder.string(transaction.tokenId),
                'min_bid_price': CLValueBuilder.u256(transaction.minBidPrice),
                'redemption_price': CLValueBuilder.u256(transaction.redemptionPrice),
                'auction_duration': CLValueBuilder.u256(transaction.auctionDuration),

            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('create_listing', runtimeArgs, LOT_PAYMENT_AMOUNT, transaction.address);

            await this.client.claim(transaction.rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Accepts offer */
    async acceptOffer(transaction: any): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                'nft_contract_hash': CLValueBuilder.string('hash-4ff1e5e37b8720e8049bfff88676d8e27c1037c02e1172a1006c6d2a535607da'),
                'token_id': CLValueBuilder.string('746c85ba-583c-4c45-9af7-dce858c4e121'),
                'erc20_contract': CLValueBuilder.byteArray(transaction.ercContract),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('accept_offer', runtimeArgs, LOT_PAYMENT_AMOUNT, transaction.address);

            await this.client.claim(transaction.rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }

    /** Accepts offer */
    async makeOffer(transaction: any): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                'nft_contract_hash': CLValueBuilder.string('hash-4ff1e5e37b8720e8049bfff88676d8e27c1037c02e1172a1006c6d2a535607da'),
                'token_id': CLValueBuilder.string('746c85ba-583c-4c45-9af7-dce858c4e121'),
                'erc20_contract': CLValueBuilder.byteArray(transaction.erc20Contract),
                'offer_price': CLValueBuilder.byteArray(transaction.ercContract),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('accept_offer', runtimeArgs, LOT_PAYMENT_AMOUNT, transaction.address);

            await this.client.claim(transaction.rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (error: any) {
            ToastNotifications.casperError(`${error.error}`);
        }
    }
}

export default CasperTransactionService;
