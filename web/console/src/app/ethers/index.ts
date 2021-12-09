// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

const DEFAULT_VALUE = 0;
/** Part of transaction props */
export class Contract {
    /** Includes address and addressMethod fields */
    constructor(
        public address: string = '',
        public addressMethod: string = ''
    ) { }
}
/** Transaction props from api */
export class Transaction {
    /** Includes password, tokenId and contract fields */
    constructor(
        public password: string = '',
        public tokenId: number = DEFAULT_VALUE,
        public contract: Contract = new Contract()
    ) { }
}

/** Abi method interface */
interface MethodAbi {
    internalType: string;
    name: string;
    type: string;
}

/** Abi block interface */
interface Abi {
    inputs: MethodAbi[];
    outputs: MethodAbi[];
    name: string;
    stateMutability: string;
    type: string;
}

/** Desctibes parameters for transaction */
export class TransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public walletAddress: string,
        public cardId: string
    ) { }
}

/** Describes parameters for sign message */
export class SignedMessage {
    /** Includes message from API, signed message, and wallet address */
    constructor(
        public message: string,
        public hash: string,
        public address: string
    ) { }
}

/** Smart conract interface */
export const GAME_ABI: Array<Partial<Abi>> = [
    {
        'inputs': [
            {
                'internalType': 'address',
                'name': '_nftAddress',
                'type': 'address',
            },
        ],
        'stateMutability': 'nonpayable',
        'type': 'constructor',
    },
    {
        'inputs': [
            {
                'internalType': 'bytes',
                'name': '_signature',
                'type': 'bytes',
            },
            {
                'internalType': 'uint256',
                'name': 'tokenID',
                'type': 'uint256',
            },
        ],
        'name': 'buyWithSignature',
        'outputs': [],
        'stateMutability': 'nonpayable',
        'type': 'function',
    },
    {
        'inputs': [],
        'name': 'nft',
        'outputs': [
            {
                'internalType': 'contract INFT',
                'name': '',
                'type': 'address',
            },
        ],
        'stateMutability': 'view',
        'type': 'function',
    },
    {
        'inputs': [],
        'name': 'verifyAddress',
        'outputs': [
            {
                'internalType': 'address',
                'name': '',
                'type': 'address',
            },
        ],
        'stateMutability': 'view',
        'type': 'function',
    },
];
