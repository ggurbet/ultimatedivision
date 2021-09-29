// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { Card } from '@/card';

import { RootState } from '@/app/store';

import { MyCard } from './MyCard';

import './index.scss';

export const ClubCardsArea: React.FC = () => {
    const { cards } =
        useSelector((state: RootState) => state.cardsReducer.cardsPage);

    return <div className="club-cards">
        <div className="club-cards__wrapper">
            {cards.map((card: Card, index: number) =>
                <MyCard
                    card={card}
                    key={index}
                />,
            )}
        </div>
    </div>;
};
