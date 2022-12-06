// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { toast } from 'react-toastify';
import MetaMaskOnboarding from '@metamask/onboarding';

import { QueueClient } from '@/api/queue';
import { UDT_ABI } from '@/app/ethers';
import { RootState } from '@/app/store';
import { ServicePlugin } from '@/app/plugins/service';
import { getCurrentQueueClient, queueActionAllowAddress, queueCasperActionAllowAddress } from '@/queue/service';
import { setCurrentUser } from '@/app/store/actions/users';
import CasperTransactionService from '@/casper';

import coin from '@static/img/match/money.svg';

import './index.scss';
import { CLPublicKey } from 'casper-js-sdk';

const VELAS_WALLET_TYPE = 'velas_wallet_address';
const CASPER_WALLET_TYPE = 'casper_wallet_address';
const METAMASK_WALLET_TYPE = 'wallet_address';

export const MatchScore: React.FC = () => {
    const dispatch = useDispatch();

    const [queueClient, setQueueClient] = useState<QueueClient | null>(null);

    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);

    const service = ServicePlugin.create();

    const { squad } = useSelector((state: RootState) => state.clubsReducer.activeClub);

    const { matchResults, transaction } = useSelector((state: RootState) => state.matchesReducer);

    const { question } = useSelector((state: RootState) => state.matchesReducer);

    const user = useSelector((state: RootState) => state.usersReducer.user);

    /** FIRST_TEAM_INDEX is variable that describes first team index in teams array. */
    const FIRST_TEAM_INDEX: number = 0;
    /** SECOND_TEAM_INDEX is variable that describes second team index in teams array. */
    const SECOND_TEAM_INDEX: number = 1;

    /** Variable describes that it needs alllow to add address or forbid add adress. */
    const CONFIRM_ADD_WALLET: string = 'do you allow us to take your address?';

    /** sets user info */
    async function setUser() {
        try {
            await dispatch(setCurrentUser());
        } catch (error: any) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    }

    /** Adds metamask wallet address for earning reward. */
    const addMetamaskWallet = async() => {
        /** Code which indicates that 'eth_requestAccounts' already processing */
        const METAMASK_RPC_ERROR_CODE = -32002;
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore .
                await window.ethereum.request({
                    method: 'eth_requestAccounts',
                });

                const wallet = await service.getWallet();

                const currentQueueClient = getCurrentQueueClient();

                const nonce = await service.getNonce(transaction.udtContract.address, UDT_ABI);

                setQueueClient(currentQueueClient);

                queueActionAllowAddress(wallet, nonce);
            } catch (error: any) {
                error.code === METAMASK_RPC_ERROR_CODE
                    ? toast.error('Please open metamask manually!', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    })
                    : toast.error('Something went wrong', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    });
            }
        } else {
            onboarding.startOnboarding();
        }
    };

    /** Adds casper wallet address for earning reward. */
    const addCasperWallet = () => {
        try {
            const ACCOUNT_HASH_PREFIX = 'account-hash-';

            const currentQueueClient = getCurrentQueueClient();

            setQueueClient(currentQueueClient);

            const accountHash = CLPublicKey.fromHex(user.casperWallet).toAccountHashStr();
            const accountHashConverted = accountHash.replace(ACCOUNT_HASH_PREFIX, '');

            queueCasperActionAllowAddress(accountHashConverted, user.walletType, squad.id);
        }
        catch (error: any) {
            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    /** Adds velas wallet address for earning reward. */
    const addVelasWallet = async() => {};

    /** Adds wallets addresses for earning reward. */
    const addWallet = async() => {
        try {
            const addingWallets = new Map();

            const addingWalletsTypes = [
                {
                    walletType: VELAS_WALLET_TYPE,
                    mint: addVelasWallet,
                },
                {
                    walletType: CASPER_WALLET_TYPE,
                    mint: addCasperWallet,
                },
                {
                    walletType: METAMASK_WALLET_TYPE,
                    mint: addMetamaskWallet,
                },
            ];

            addingWalletsTypes.forEach(addingWalletsType =>
                addingWallets.set(addingWalletsType.walletType, addingWalletsType.mint));

            await addingWallets.get(user.walletType)();
        } catch (e) {
            toast.error('Invalid transaction', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    /** Mints token with casper wallet. */
    const casperMint = async(messageEvent: any) => {
        const casperTransactionService = new CasperTransactionService(user.casperWallet);

        await casperTransactionService.mintUDT(messageEvent.message.casperTransaction, messageEvent.message.rpcNodeAddress);
    };

    /** Mints token with metamask wallet. */
    const metamaskMint = async(messageEvent: any) => {
        await service.mintUDT(messageEvent.message.transaction);
    };

    /** Mints token with velas wallet. */
    const velasMint = () => {};

    if (queueClient) {
        queueClient.ws.onmessage = async({ data }: MessageEvent) => {
            const messageEvent = JSON.parse(data);

            const mintingTokens = new Map();

            const mintingTokensTypes = [
                {
                    walletType: VELAS_WALLET_TYPE,
                    mint: velasMint,
                },
                {
                    walletType: CASPER_WALLET_TYPE,
                    mint: casperMint,
                },
                {
                    walletType: METAMASK_WALLET_TYPE,
                    mint: metamaskMint,
                },
            ];

            mintingTokensTypes.forEach(mintingTokensType =>
                mintingTokens.set(mintingTokensType.walletType, mintingTokensType.mint));

            await mintingTokens.get(user.walletType)(messageEvent);
        };
    }

    useEffect(() => {
        setUser();
    }, []);

    return (
        <div className="match__score">
            <div className="match__score__board">
                <div className="match__score__board__gradient"></div>
                <div className="match__score__board__timer">90:00</div>
                <div className="match__score__board__result">
                    <div className="match__score__board-team-1">
                        {matchResults[FIRST_TEAM_INDEX].quantityGoals}
                    </div>
                    <div className="match__score__board-dash">-</div>
                    <div className="match__score__board-team-2">
                        {matchResults[SECOND_TEAM_INDEX].quantityGoals}
                    </div>
                </div>
                {question === CONFIRM_ADD_WALLET && <div className="match__score__board__coins">
                    <img
                        className="match__score__board__coins-image"
                        src={coin}
                        alt="Coin"
                    />
                    <span className="match__score__board__coins-value">
                        {transaction.value}
                    </span>
                    <button
                        className="match__score__board__coins__btn"
                        onClick={addWallet}
                    >
                        <span className="match__score__board__coins__btn-text">
                            GET
                        </span>
                    </button>
                </div>
                }
            </div>
        </div>
    );
};
