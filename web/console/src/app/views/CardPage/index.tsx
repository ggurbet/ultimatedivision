// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router';

import { RootState } from '@/app/store';
import { openUserCard } from '@/app/store/actions/cards';
import { FootballerCardIllustrations } from '@/app/components/common/Card/CardIllustrations';
import { FootballerCardPrice } from '@/app/components/common/Card/CardPrice';
import { FootballerCardStatsArea } from '@/app/components/common/Card/CardStatsArea';
import { FootballerCardInformation } from '@/app/components/common/Card/CardInformation';

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
            </div>
        </div>
    );
};

export default Card;
