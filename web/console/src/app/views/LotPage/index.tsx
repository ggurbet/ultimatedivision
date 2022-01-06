// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';

import { RootState } from '@/app/store';
import { openMarketplaceCard } from '@/app/store/actions/marketplace';
import { FootballerCardIllustrations } from '@/app/components/common/Card/CardIllustrations';
import { FootballerCardPrice } from '@/app/components/common/Card/CardPrice';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';

import './index.scss';

const Lot: React.FC = () => {
    const dispatch = useDispatch();
    const { card } = useSelector((state: RootState) => state.marketplaceReducer);

    const { id }: { id: string } = useParams();
    /** implements opening new card */
    async function openCard() {
        try {
            await dispatch(openMarketplaceCard(id));
        } catch (error: any) {
            /** TODO: it will be reworked with notification system */
        };
    };

    useEffect(() => {
        openCard();
    }, []);

    return (
        card &&
        <div className="lot">
            <div className="lot__border">
                <div className="lot__wrapper">
                    <div className="lot__name-wrapper">
                        <h1 className="lot__name">
                            {card.playerName}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={card} />
                    <div className="lot__stats-area">
                        <FootballerCardPrice card={card} isMinted={false} />
                        <FootballerCardStatsArea card={card} />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Lot;
