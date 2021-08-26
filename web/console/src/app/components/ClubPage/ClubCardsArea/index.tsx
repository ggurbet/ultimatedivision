// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/cards';

import { MyCard } from './MyCard';

import './index.scss';

export const ClubCardsArea: React.FC<{cards: Card[]}> = ({ cards }) =>
    <div className="club-cards">
        <div className="club-cards__wrapper">
            {cards.map((card, index) =>
                <MyCard
                    card={card}
                    key={index}
                />,
            )}
        </div>
    </div>;

