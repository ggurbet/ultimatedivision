// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';

import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '@/app/store';
import { openMarketplaceCard } from '@/app/store/actions/cards';

import { useParams } from 'react-router-dom';

import { FootballerCardIllustrations } from '@/app/components/common/Card/CardIllustrations';
import { FootballerCardPrice } from '@/app/components/common/Card/CardPrice';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';
import { FootballerCardInformation } from '@/app/components/common/Card/CardInformation';

import './index.scss';

const Lot: React.FC = () => {
    const dispatch = useDispatch();
    const cardData = useSelector((state: RootState) => state.cardsReducer.openedCard);
    // @ts-ignore
    const { id } = useParams();

    useEffect(() => {
        dispatch(openMarketplaceCard(id));
    }, []);

    return (
        cardData &&
        <div className="lot">
            <div className="lot__border">
                <div className="lot__wrapper">
                    <div className="lot__name-wrapper">
                        <h1 className="lot__name">
                            {cardData.playerName}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={cardData} />
                    <div className="lot__stats-area">
                        <FootballerCardPrice card={cardData} />
                        <FootballerCardStatsArea card={cardData} />
                        <FootballerCardInformation card={cardData} />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Lot;
