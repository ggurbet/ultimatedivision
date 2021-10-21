// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { EthersClient } from '@/api/ethers';
import { buildHash } from '@utils/ethers';
import { ethers } from 'ethers';

export class Service {
    private readonly provider;
    private readonly client = new EthersClient();

    public constructor(ethereumProvider: any) {
        this.provider = ethereumProvider;
    }

    /** Get nft adress and password. */
    public async getAddress(wallet: string) {
        return await this.client.getAddress(wallet);
    }
    /** Gets current wallet address. */
    public async getWallet() {
        const signer = await this.provider.getSigner();

        return await signer.getAddress();
    }

    /** Get last lot id. */
    public async getLastTokenId(wallet: string, abi: any[]) {
        const address = await this.getAddress(wallet);
        const contract = await new ethers.Contract(address.smartContractAddress.nft, abi);
        const signer = await this.provider.getSigner();
        const connect = await contract.connect(signer);
        const totalSupply = await connect.functions.totalSupply();

        return parseInt(totalSupply[0]._hex, 16);
    };

    /** Send smart contract transaction. */
    public async sendTransaction(wallet: string, totalSupply: number, abi: any[]) {
        const signer = await this.provider.getSigner();
        const address = await this.getAddress(wallet);
        const contract = await new ethers.Contract(address.smartContractAddress.nftSale, abi);
        const connect = await contract.connect(signer);
        const currentPrice = await connect.functions.getCurrentPrice();
        const data = `0xd399cc1a${buildHash((totalSupply + 1).toString(16))}${buildHash(40)}${buildHash(60)}${buildHash(address.password.slice(-2))}${address.password.slice(0, address.password.length - 2)}`;
        const gasLimit = await signer.estimateGas({
            to: address.smartContractAddress.nftSale,
            data,
            value: currentPrice[0],
        })
        const transaction = await signer.sendTransaction({
            to: address.smartContractAddress.nftSale,
            data,
            gasLimit,
            chainId: 3,
            value: currentPrice[0]
        });
    };

    public async getBalance(id: string) {
        try {
            const balance = await this.provider.getBalance(id);

            return ethers.utils.formatEther(balance);
        } catch (error: any) {
            /* eslint-disable-next-line */
            console.log(error.message)
        }
    }
};
