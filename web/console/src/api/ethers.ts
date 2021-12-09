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

        return new Transaction(transaction.password, transaction.tokenId, transaction.contract);
    }

    /** Gets message from API for sign */
    public async getMessage(): Promise<string> {
        const response = await this.http.get(`${this.ROOT_PATH}/auth/metamask/message-token`);

        if (!response.ok) {
            await this.handleError(response);
        };

        return await response.json();
    }

    /** Sends signed message, and logs-in */
    public async signMessage(message: SignedMessage): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/auth/metamask/login`, JSON.stringify(message));

        if (!response.ok) {
            await this.handleError(response);
        };
    }
}
