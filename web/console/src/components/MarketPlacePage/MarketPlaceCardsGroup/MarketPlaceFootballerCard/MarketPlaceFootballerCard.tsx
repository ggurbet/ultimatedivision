/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';

import { PlayerCard }
    from '../../../PlayerCard/PlayerCard';

import { Card } from '../../../../store/reducers/footballerCard';

import './MarketPlaceFootballerCard.scss';

export const MarketPlaceFootballerCard: React.FC<{ card: Card, place?: string }> = ({ card, place }) => {
    return (
        <div
            className="marketplace-playerCard"
        >
            <PlayerCard
                card={card}
                parentClassName={"marketplace-playerCard"}
            />
            <div className="marketplace-playerCard__price">
                <img className="marketplace-playerCard__price__picture"
                    src={card.mainInfo.priceIcon}
                    alt="Player price" />
                <span className="marketplace-playerCard__price__current">
                    {card.mainInfo.price}
                </span>
                <img className="marketplace-playerCard__price__status"
                    src={card.mainInfo.priceStatus}
                    alt="Price status" />
            </div>
        </div>
    );
};
