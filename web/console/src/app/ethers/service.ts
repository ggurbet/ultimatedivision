// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

import { EthersClient } from '@/api/ethers';
import { buildHash } from '../internal/ethers';
import { Transaction } from '.';
import { TransactionIdentificators } from '@/app/ethers';
import { web3Provider } from '@/app/plugins/service';
import { MatchTransaction } from '../../matches';

const CHAIN_ID = 4;

/** Service for ethers methods */
export class Service {
    private readonly provider;
    private readonly client = new EthersClient();

    /** Applies ethereum provider for internal methons */
    public constructor(ethereumProvider: typeof web3Provider) {
        this.provider = ethereumProvider;
    }

    /** Returns metamask walllet address */
    public async getWallet() {
        const signer = await this.provider.getSigner();

        return await signer.getAddress();
    }

    /** Signs message and creates message raw signature */
    public async signMessage(message: string) {
        const signer = await this.provider.getSigner();

        return await signer.signMessage(message);
    }

    /** Gets transaction from api */
    public async getTransaction(signature: TransactionIdentificators): Promise<Transaction> {
        return await this.client.getTransaction(signature);
    }


    /** Returns required fields for metamask login */
    public async login(message: string) {
        const signedMessage = await this.signMessage(message);
        return signedMessage;
    }

    public async getNonce(UDTContractAddress: string, abi: any) {
        const signer = await this.provider.getSigner();
        const address = await this.getWallet();

        const contract = await new ethers.Contract(UDTContractAddress, abi);
        const connect = await contract.connect(signer);

        const nonce = await connect.functions.claimNonce(address);

        const FIRTS_NONCE_ELEMENT: number = 0;
        const HEX_TYPE: number = 16;

        /* eslint-disable */
        return parseInt(nonce[FIRTS_NONCE_ELEMENT]._hex, HEX_TYPE);
    };

    /** Mints UDT. */
    public async mintUDT(transaction: MatchTransaction) {
        const signer = await this.provider.getSigner();
        const address = await this.getWallet();

        const value = await ethers.utils.parseEther('0');
        /* eslint-disable */
        const data = transaction.value && `${transaction.udtContract.addressMethod}${buildHash(Number(transaction.value).toString(16))}${buildHash(40)}${buildHash(60)}${buildHash(transaction.signature.slice(-2))}${transaction.signature.slice(0, transaction.signature.length - 2)}`;

        const gasLimit = await signer.estimateGas({
            to: transaction.udtContract.address,
            data,
        });

        const CHAN_ID: number = 4;

        const transactionUDT = await signer.sendTransaction({
            to: transaction.udtContract.address,
            data,
            gasLimit,
            chainId: CHAN_ID,
        });
    };

    /** Sends smart contract transaction. */
    public async sendTransaction(
        cardId: string,
    ) {
        const walletAddress = await this.getWallet();
        const signer = await this.provider.getSigner();
        const address = await this.getTransaction(new TransactionIdentificators(walletAddress, cardId));
        /* eslint-disable */
        const data = `${address.nftCreateContract.mintWithSignatureSelector}${buildHash(40)}${buildHash(address.tokenId.toString(16))}${buildHash(60)}${buildHash(
            address.password.slice(-2)
        )}${address.password.slice(0, address.password.length - 2)}`;

        const gasLimit = await signer.estimateGas({
            to: address.nftCreateContract.address,
            data,
        });

        await signer.sendTransaction({
            to: address.nftCreateContract.address,
            data,
            gasLimit,
            chainId: address.nftCreateContract.chainId,
        });
    }

    public async getBalance(id: string) {
        const balance = await this.provider.getBalance(id);

        return ethers.utils.formatEther(balance);
    }
}
