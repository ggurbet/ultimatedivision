// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';

/**
 * CasperClient is a http implementation of casper-wallet API.
 * Exposes all casper wallet related functionality.
 */
export class CasperClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** Sends signed message and registers user */
    public async register(walletAddress: string): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/auth/casper/register`, JSON.stringify(walletAddress));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Gets message from API for sign with casper */
    public async nonce(walletAddress: string): Promise<string> {
        const path = `${this.ROOT_PATH}/auth/casper/nonce?address=${walletAddress}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }

    /** Sends signed message, and logs-in */
    public async login(nonce: string, signature: string): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/auth/casper/login`, JSON.stringify({ nonce, signature }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
