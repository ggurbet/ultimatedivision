// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';

/**
 * VelasClient is a http implementation of velas-wallet API.
 * Exposes all velas-related functionality.
 */
export class VelasClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** sends data to register user with velas wallet */
    public async register(walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        const path = `${this.ROOT_PATH}/auth/velas/register`;
        const response = await this.http.post(path, JSON.stringify({ walletAddress, accessToken, expiresAt }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** sends address to get nonce to login user */
    public async nonce(address: string): Promise<string> {
        const walletType = 'velas_wallet_address';

        const path = `${this.ROOT_PATH}/auth/velas/nonce?address=${address}&walletType=${walletType}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const status = response.json();

        return status;
    }
    /** sends data to login user with velas wallet */
    public async login(nonce: string, walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        const path = `${this.ROOT_PATH}/auth/velas/login`;
        const response = await this.http.post(path, JSON.stringify({ walletAddress, accessToken, expiresAt, nonce }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** gets token to login user with velas wallet */
    public async csrfToken(): Promise<string> {
        const path = 'https://velas.ultimatedivision.com/csrf';
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result.token;
    }

    /** gets creds to fill velas vaclient */
    public async vaclientCreds(): Promise<any> {
        const path = `${this.ROOT_PATH}/auth/velas/vaclient`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result;
    }
}
