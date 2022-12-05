// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Buffer } from 'buffer';
import { toast } from 'react-toastify';
import { JsonTypes } from 'typedjson';
import { CLValueBuilder, RuntimeArgs, CLPublicKey, DeployUtil, Signer } from 'casper-js-sdk';

import { CasperNetworkClient } from '@/api/casper';
import { CasperMatchTransaction } from '@/matches';

enum CasperRuntimeArgs {
    SIGNATURE = 'signature',
    TOKEN_ID = 'token_id'
}

/** Desctibes parameters for transaction */
export class CasperTransactionIdentificators {
    /** Includes wallet address, and card id */
    constructor(
        public casperWallet: string,
        public cardId: string
    ) { }
}

const ACCOUNT_HASH_PREFIX = 'account-hash-';

const TTL = 1800000;
const PAYMENT_AMOUNT = 50000000000;
const GAS_PRICE = 1;

/** CasperTransactionService describes casper transaction entity. */
class CasperTransactionService {
    private readonly paymentAmount: number = PAYMENT_AMOUNT;
    private readonly gasPrice: number = GAS_PRICE;
    private readonly ttl: number = TTL;
    private readonly client: any = new CasperNetworkClient();
    public walletAddress: string = '';

    /** default VelasTransactionService implementation */
    constructor(walletAddress: string) {
        this.walletAddress = walletAddress;
    }

    /** Gets minting signature with contract address from api */
    async getTransaction(signature: CasperTransactionIdentificators): Promise<any> {
        return await this.client.getTransaction(signature);
    }

    /** Converts contract hash to bytes  */
    public static async convertContractHashToBytes(contractHash: string): Promise<Uint8Array> {
        return await Uint8Array.from(Buffer.from(contractHash, 'hex'));
    }

    /** Signs a contract */
    public async contractSign(
        entryPoint: string,
        runtimeArgs: RuntimeArgs,
        paymentAmount: number,
        contractAddress: string
    ): Promise<JsonTypes> {
        const contractHashToBytes = await CasperTransactionService.convertContractHashToBytes(contractAddress);

        try {
            const walletAddressConverted = CLPublicKey.fromHex(this.walletAddress);

            const deployParams = new DeployUtil.DeployParams(walletAddressConverted, 'casper-test', this.gasPrice, this.ttl);

            const deploy = DeployUtil.makeDeploy(
                deployParams,
                DeployUtil.ExecutableDeployItem.newStoredContractByHash(
                    contractHashToBytes,
                    entryPoint,
                    runtimeArgs),
                DeployUtil.standardPayment(paymentAmount)
            );

            const deployJson = DeployUtil.deployToJson(deploy);

            const signature = await window.casperlabsHelper.sign(deployJson, this.walletAddress, contractAddress);

            return signature;
        }
        catch (e) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }

        return false;
    }

    /** Mints a nft */
    async mint(cardId: string): Promise<void> {
        try {
            const accountHash = CLPublicKey.fromHex(this.walletAddress).toAccountHashStr();
            const accountHashConverted = accountHash.replace(ACCOUNT_HASH_PREFIX, '');

            const nftWaitlist = await this.getTransaction(new CasperTransactionIdentificators(accountHashConverted, cardId));

            const runtimeArgs = RuntimeArgs.fromMap({
                [CasperRuntimeArgs.SIGNATURE]: CLValueBuilder.string(nftWaitlist.password),
                [CasperRuntimeArgs.TOKEN_ID]: CLValueBuilder.string(nftWaitlist.tokenId),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('claim', runtimeArgs, this.paymentAmount, nftWaitlist.nftCreateCasperContract.address);

            await this.client.claim(nftWaitlist.rpcNodeAddress, JSON.stringify(signature));
        }
        catch (e) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }

    /** Mints a token */
    async mintUDT(transaction: CasperMatchTransaction, rpcNodeAddress: string): Promise<void> {
        try {
            const runtimeArgs = RuntimeArgs.fromMap({
                'signature': CLValueBuilder.string(transaction.signature),
                'value': CLValueBuilder.u256(transaction.value),
                'nonce': CLValueBuilder.u64(transaction.nonce),
            });

            const isConnected = window.casperlabsHelper.isConnected();

            if (!isConnected) {
                await window.casperlabsHelper.requestConnection();
            }

            const signature = await this.contractSign('claim', runtimeArgs, this.paymentAmount, transaction.casperTokenContract.address);

            await this.client.claim(rpcNodeAddress, JSON.stringify(signature), this.walletAddress);
        }
        catch (e) {
            toast.error('Invalid transaction', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }
}

export default CasperTransactionService;
