// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { VelasClient } from '../api/velas';

/**
 * Exposes all velas wallet related logic.
 */
export class VelasService {
    private readonly velasWallet: VelasClient;

    /** VelasService contains http implementation of velas wallet API  */
    public constructor(velasWallet: VelasClient) {
        this.velasWallet = velasWallet;
    }

    /** sends data to register user with velas wallet */
    public async register(walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        await this.velasWallet.register(walletAddress, accessToken, expiresAt);
    }

    /** sends address to get nonce to login user */
    public async nonce(address: string): Promise<string> {
        return await this.velasWallet.nonce(address);
    }

    /** sends data to login user with velas wallet */
    public async login(nonce: string, walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        await this.velasWallet.login(nonce, walletAddress, accessToken, expiresAt);
    }

    /** gets token to login user with velas wallet */
    public async csrfToken(): Promise<string> {
        return await this.velasWallet.csrfToken();
    }

    /** gets creds to fill velas vaclient */
    public async vaclientCreds(): Promise<any> {
        return await this.velasWallet.vaclientCreds();
    }
}
