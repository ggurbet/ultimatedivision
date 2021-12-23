// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router';
import MetaMaskOnboarding from '@metamask/onboarding';
import { toast } from 'react-toastify';

import { RootState } from '@/app/store';
import { openUserCard } from '@/app/store/actions/cards';
import { FootballerCardIllustrations } from '@/app/components/common/Card/CardIllustrations';
import { FootballerCardPrice } from '@/app/components/common/Card/CardPrice';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';
import { FootballerCardInformation } from '@/app/components/common/Card/CardInformation';
import { ServicePlugin } from '@/app/plugins/service';

import './index.scss';

const Card: React.FC = () => {
    const dispatch = useDispatch();
    const { card } = useSelector((state: RootState) => state.cardsReducer);

    const { id }: { id: string } = useParams();
    /** implements opening new card */
    async function openCard() {
        try {
            await dispatch(openUserCard(id));
        } catch (error: any) {
            /** TODO: it will be reworked with notification system */
        };
    };
    useEffect(() => {
        openCard();
    }, []);

    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);
    const service = ServicePlugin.create();

    /** Mints chosed card */
    const mint = async() => {
        /** Code which indicates that 'eth_requestAccounts' already processing */
        const METAMASK_RPC_ERROR_CODE = -32002;
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore
                await window.ethereum.request({
                    method: 'eth_requestAccounts',
                });
                await service.sendTransaction(id);
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

    return (
        card &&
        <div className="card">
            <div className="card__border">
                <div className="card__wrapper">
                    <div className="card__name-wrapper">
                        <h1 className="card__name">
                            {card.playerName}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={card} />
                    <div className="card__stats-area">
                        <FootballerCardPrice card={card} />
                        <FootballerCardStatsArea card={card} />
                        <FootballerCardInformation card={card} />
                    </div>
                </div>
                <button className="card__mint" onClick={mint}>Mint</button>
            </div>
        </div>
    );
};

export default Card;
