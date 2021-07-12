/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';

import './MarketPlaceCardsGroup.scss';

import { Card } from '../../../store/reducers/footballerCard';

export const MarketPlaceCardsGroup: React.FC<{ cards: Card[], Component: React.FC<{ card: Card, key: number }> }> = ({ cards, Component }) => {
    return (
        <div className="marketplace-cards">
            <div className="marketplace-cards__wrapper">
                {cards.map((card, index) => (
                    <Component
                        card={card}
                        key={index}
                    />
                ))}
            </div>
        </div>
    );
};
