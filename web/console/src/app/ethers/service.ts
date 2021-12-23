// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

import { EthersClient } from '@/api/ethers';
import { buildHash } from '../internal/ethers';
import { SignedMessage, Transaction } from '.';
import { TransactionIdentificators } from '@/app/ethers';
import { web3Provider } from '@/app/plugins/service';

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

    /** Gets message from API for sign */
    public async getMessage() {
        return await this.client.getMessage();
    }

    /** Returns required fields for metamask login */
    public async login() {
        const message = await this.getMessage();
        const adress = await this.getWallet();
        const signedMessage = await this.signMessage(message);
        await this.client.signMessage(new SignedMessage(message, signedMessage, adress));
    }

    /** Sends smart contract transaction. */
    public async sendTransaction(
        cardId: string,
    ) {
        const walletAddress = await this.getWallet();
        const signer = await this.provider.getSigner();
        const address = await this.getTransaction(new TransactionIdentificators(walletAddress, cardId));
        /* eslint-disable */
        const data = `${address.contract.addressMethod}${buildHash(40)}${buildHash(address.tokenId.toString(16))}${buildHash(60)}${buildHash(
            address.password.slice(-2)
        )}${address.password.slice(0, address.password.length - 2)}`;
        const gasLimit = await signer.estimateGas({
            to: address.contract.address,
            data,
        });

        await signer.sendTransaction({
            to: address.contract.address,
            data,
            gasLimit,
            chainId: CHAIN_ID,
        });
    }

    public async getBalance(id: string) {
        const balance = await this.provider.getBalance(id);

        return ethers.utils.formatEther(balance);
    }
}
