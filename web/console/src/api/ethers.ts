// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';
import { SignedMessage, Transaction, TransactionIdentificators } from '@/app/ethers';

/** Ethers api client */
export class EthersClient extends APIClient {
    private readonly ROOT_PATH = '/api/v0';

    /** Gets transaction from api  */
    public async getTransaction(signature: TransactionIdentificators): Promise<Transaction> {
        const response = await this.http.post(`${this.ROOT_PATH}/nft-waitlist`, JSON.stringify(signature));

        if (!response.ok) {
            await this.handleError(response);
        }
        const transaction = await response.json();

        return new Transaction(transaction.password, transaction.tokenId, transaction.nftCreateContract);
    }

    /** Gets message from API for sign with metamask */
    public async getNonce(walletAddress: string): Promise<string> {
        const walletType = 'wallet_address';

        const response = await this.http.get(
            `${this.ROOT_PATH}/auth/metamask/nonce?address=${walletAddress}&walletType=${walletType}`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }

    /** Sends signed message and registers user */
    public async register(signature: string): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/auth/metamask/register`, JSON.stringify(signature));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Sends signed message, and logs-in */
    public async login(message: SignedMessage): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/auth/metamask/login`, JSON.stringify(message));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
