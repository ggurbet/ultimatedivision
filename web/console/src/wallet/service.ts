// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import MetaMaskOnboarding from '@metamask/onboarding';

import { ServicePlugin } from '@/app/plugins/service';

import CasperTransactionService from '@/casper';
import { User } from '@/users';
import { walletTypes } from '.';
import { ethers } from 'ethers';
import { ToastNotifications } from '@/notifications/service';

/**
 * Exposes all wallet service related logic.
 */
class WalletService {
    // @ts-ignore
    public metamaskProvider = window.ethereum && new ethers.providers.Web3Provider(window.ethereum);
    public metamaskService = ServicePlugin.create();
    public onboarding = new MetaMaskOnboarding();
    public user: User = new User();

    /** default MintingService implementation */
    constructor(user: User) {
        this.user = user;
    }

    /** Mints chosed card with metamask */
    private async metamaskMint(id: string) {
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore .
                await this.metamaskProvider.request({
                    method: 'eth_requestAccounts',
                });
                await this.metamaskService.sendTransaction(id);
            } catch (error: any) {
                ToastNotifications.metamaskError(error);
            }
        } else {
            this.onboarding.startOnboarding();
        }
    };

    /** Mints chosed card with casper */
    private async casperMint(id: string) {
        const casperTransactionService = new CasperTransactionService(this.user.casperWalletId);

        await casperTransactionService.mint(id);
    };

    /** Mints chosed card with velas */
    private static velasMint() { };

    /** Mints chosed card. */
    public async mintNft(id: string) {
        switch (this.user.walletType) {
        case walletTypes.VELAS_WALLET_TYPE:
            await WalletService.velasMint();
            break;
        case walletTypes.CASPER_WALLET_TYPE:
            await this.casperMint(id);
            break;
        case walletTypes.METAMASK_WALLET_TYPE:
            await this.metamaskMint(id);
            break;
        default:
            break;
        }
    }

    /** Mints token with metamask wallet. */
    private metamaskMintToken(messageEvent: any) {
        this.metamaskService.mintUDT(messageEvent.message.transaction);
    };

    /** Mints token with casper wallet. */
    private casperMintToken(messageEvent: any) {
        const casperTransactionService = new CasperTransactionService(this.user.casperWalletId);

        casperTransactionService.mintUDT(messageEvent.message.casperTransaction, messageEvent.message.rpcNodeAddress);
    };

    /** Mints token with velas wallet. */
    private static velasMintToken() { };

    /** Mints token. */
    public mintToken(messageEvent: any) {
        switch (this.user.walletType) {
        case walletTypes.VELAS_WALLET_TYPE:
            WalletService.velasMintToken();
            break;
        case walletTypes.CASPER_WALLET_TYPE:
            this.casperMintToken(messageEvent);
            break;
        case walletTypes.METAMASK_WALLET_TYPE:
            this.metamaskMintToken(messageEvent);
            break;
        default:
            break;
        }
    };
}

export default WalletService;
