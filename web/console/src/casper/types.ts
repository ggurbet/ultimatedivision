// Copyright (C) 2023 Creditor Corp. Group.
// See LICENSE for copying information.

export const ACCOUNT_HASH_PREFIX = 'account-hash-';
export const CHAIN_NAME = 'casper-test';

const DEFAULT_MIN_PRICE = 3000;
const DEFAULT_DURATION = 1;
const DEFAULT_REDEMPTION_PRRICE = 30000;
const DEFAULT_AMOUNT = 0;

export const TTL = 1800000;
export const GAS_PRICE = 1;

export const PAYMENT_AMOUNT = 50000000000;
export const TOKEN_PAYMENT_AMOUNT = 6000000000;
export const MAKE_OFFER_PAYMENT_AMOUNT = 4400000000;
export const CREATE_LOT_PAYMENT_AMOUNT = 5700000000;
export const APPROVE_TOKEN_PAYMENT_AMOUNT = 100000000;
export const APPROVE_NFT_PAYMENT_AMOUNT = 2500000000;
export const BUY_OFFER_PAYMENT_AMOUNT = 12000000000;
export const ACCEPT_OFFER_PAYMENT_AMOUNT = 10000000000;
export const MINT_ONE_PAYMENT_AMOUNT = 4100000000;

/** Describes parameters for transaction */
export class CasperTransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public casperWallet: string,
        public cardId: string
    ) { }
}

/** Describes parameters for transaction */
export class CasperTransactionApprove {
    /** Includes wallet address, and card id */
    constructor(
        public addressNodeServer: string = '',
        public amount: number = DEFAULT_AMOUNT,
        public NFTContractAddress: string = '',
        public tokenRewardContractAddress: string = '',
        public tokenId: string = '',
        public approveTokensSpender: string = '',
        public approveNftSpender: string = '',
    ) { }
}

/** Describes parameters for casper token transaction */
export class CasperTokenContract {
    /** default CasperTokenContract implementation */
    constructor(
        public address: string = '',
        public addressMethod: string = ''
    ) { }
}

/** Transaction describes transaction entity of match response. */
export class CasperSeasonRewardTransaction {
    /** Transaction describes transaction entity */
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

/** Class defines creating lot */
export class MarketCreateLotTransaction {
    /** Includes creating lot entity */
    constructor(
        public address: string = '',
        public rpcNodeAddress: string = '',
        public tokenId: string = '',
        public contractHash: string = '',
        public minBidPrice: number = DEFAULT_MIN_PRICE,
        public auctionDuration: number = DEFAULT_DURATION,
        public redemptionPrice: number = DEFAULT_REDEMPTION_PRRICE,
    ) { }
};

/** Class defines bids make offer transaction entity */
export class BidsMakeOfferTransaction {
    /** Includes address,rpcNodeAddress,tokenId, contractHash and offer price */
    constructor(
        public address: string = '',
        public rpcNodeAddress: string = '',
        public tokenId: string = '',
        public contractHash: string = '',
        public tokenContractHash: string = '',
        public offerPrice: number = DEFAULT_DURATION,
    ) { }
};

/** Class defines offer transaction entity */
export class OfferTransaction {
    /** Includes address,rpcNodeAddress,tokenId and contractHash */
    constructor(
        public address: string = '',
        public rpcNodeAddress: string = '',
        public tokenId: string = '',
        public contractHash: string = '',
        public tokenContractHash: string = '',
    ) { }
};
