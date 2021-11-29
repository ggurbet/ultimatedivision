// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';
import { Transaction } from '@/app/ethers';
import { TransactionIdentificators } from '@/app/types/ethers';

/** Whitelist api client */
export class WhitelistClient extends APIClient {
    private readonly ROOT_PATH = '/api/v0/nft-waitlist';

    /** Gets transaction from api  */
    public async getTransaction(signature: TransactionIdentificators): Promise<Transaction> {
        const response = await this.http.post(`${this.ROOT_PATH}`, JSON.stringify(signature));

        if (!response.ok) {
            await this.handleError(response);
        }
        const transaction = await response.json();

        return new Transaction(transaction.password, transaction.tokenId, transaction.contract);
    }
}
