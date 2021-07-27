/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */
import { Card } from '@/app/store/reducers/footballerCard';

import './index.scss';

export const MarketPlaceCardsGroup: React.FC<{ cards: Card[]; Component: React.FC<{ card: Card; key: number }> }> = ({ cards, Component }) =>
    <div className="marketplace-cards">
        <div className="marketplace-cards__wrapper">
            {cards.map((card, index) =>
                <Component
                    card={card}
                    key={index}
                />,
            )}
        </div>
    </div>;
