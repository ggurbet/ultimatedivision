// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CasperClient } from '../api/casper';

/**
 * Exposes all casper wallet related logic.
 */
export class CasperService {
    private readonly casperWallet: CasperClient;

    /** CasperService contains http implementation of casper wallet API  */
    public constructor(casperWallet: CasperClient) {
        this.casperWallet = casperWallet;
    }

    /** sends data to register user with casper wallet */
    public async register(walletAddress: string): Promise<void> {
        await this.casperWallet.register(walletAddress);
    }
    /** sends address to get casper nonce to login user */
    public async nonce(address: string): Promise<string> {
        return await this.casperWallet.nonce(address);
    }
    /** sends data to login user with casper wallet */
    public async login(nonce: string, walletAddress: string): Promise<void> {
        await this.casperWallet.login(nonce, walletAddress);
    }
}
