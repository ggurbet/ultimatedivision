// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CasperTransactionIdentificators } from '.';
import { CasperNetworkClient } from '../api/casper';

/**
 * Exposes all casper wallet related logic.
 */
export class CasperNetworkService {
    private readonly casperWallet: CasperNetworkClient;

    /** CasperService contains http implementation of casper wallet API  */
    public constructor(casperWallet: CasperNetworkClient) {
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
    /** Gets minting signature with contract address from api */
    public async getTransaction(signature: CasperTransactionIdentificators): Promise<any> {
        await this.casperWallet.getTransaction(signature);
    }
    /** Sends deploy data to api */
    public async claim(RPCNodeAddress: string, deploy: string): Promise<void> {
        await this.casperWallet.claim(RPCNodeAddress, deploy);
    }
}
