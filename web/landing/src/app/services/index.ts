// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import MetaMaskOnboarding from "@metamask/onboarding";
import { ethers } from 'ethers';

export class Service {
    private readonly provider;

    public constructor(ethereumProvider: any) {
        this.provider = ethereumProvider;
    }

    public async connectMetamask() {
        const onboarding = new MetaMaskOnboarding();

        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            // @ts-ignore
            return await window.ethereum.request({ method: 'eth_requestAccounts' });

        } else {
            onboarding.startOnboarding();
        }
    }

    public async sendTransaction(adress: string, amount: string) {
        try {
            const signer = this.provider.getSigner();

            //throws error when adress is not valid
            ethers.utils.getAddress(adress);

            const transaction = await signer.sendTransaction({
                to: adress,
                value: ethers.utils.parseEther(amount)
            });
        } catch (error: any) {
            /* eslint-disable-next-line */
            console.log(error.message)
        }
    }

    public async getBalance(id: string) {
        try {
            const balance = await this.provider.getBalance(id);

            return ethers.utils.formatEther(balance);
        } catch (error: any) {
            /* eslint-disable-next-line */
            console.log(error.message)
        }
    }
}
