// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import MetaMaskOnboarding from '@metamask/onboarding';
import { toast } from 'react-toastify';

import coin from '@static/img/match/money.svg';

import { QueueClient } from '@/api/queue';
import { UDT_ABI } from '@/app/ethers';
import { RootState } from '@/app/store';
import { ServicePlugin } from '@/app/plugins/service';
import { getCurrentQueueClient, queueActionAllowAddress } from '@/queue/service';

import './index.scss';

export const MatchScore: React.FC = () => {
    const [queueClient, setQueueClient] = useState<QueueClient | null>(null);

    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);
    const service = ServicePlugin.create();

    const { matchResults, transaction } = useSelector((state: RootState) => state.matchesReducer);

    const { question } = useSelector((state: RootState) => state.matchesReducer);

    /** FIRST_TEAM_INDEX is variable that describes first team index in teams array. */
    const FIRST_TEAM_INDEX: number = 0;
    /** SECOND_TEAM_INDEX is variable that describes second team index in teams array. */
    const SECOND_TEAM_INDEX: number = 1;

    /** Variable describes that it needs alllow to add address or forbid add adress. */
    const CONFIRM_ADD_WALLET: string = 'do you allow us to take your address?';

    /** Adds metamask wallet address for earning reward. */
    const addWallet = async() => {
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

    if (queueClient) {
        queueClient.ws.onmessage = ({ data }: MessageEvent) => {
            const messageEvent = JSON.parse(data);
            service.mintUDT(messageEvent.message.transaction);
        };
    }

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
