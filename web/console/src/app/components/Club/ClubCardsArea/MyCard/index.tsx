// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { Link } from 'react-router-dom';

import { Card } from '@/card';

import { createLot } from '@/app/store/actions/marketplace';
import { CreatedLot } from '@/app/types/marketplace';

import { PlayerCard } from '@components/common/PlayerCard';

/** TODO: replace it by class fields */
import confirmIcon from '@static/img/MarketPlacePage/MyCard/ok.svg';
import priceGoldIcon from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';


import './index.scss';

export const MyCard: React.FC<{ card: Card }> = ({ card }) => {
    const dispatch = useDispatch();

    const [controlVisibility, changeControlVisibility] = useState<boolean>(false);

    const handleControls = (e: any) => {
        e.preventDefault();
        changeControlVisibility(prev => !prev);
    };
    const handleSelling = (e: any) => {
        e.stopPropagation();
        e.nativeEvent.stopImmediatePropagation();
        /** TODO: create interface for adding selling parameters */
        /* eslint-disable */
        dispatch(createLot(new CreatedLot(card.id, 200, 200, 1)));
        /* eslint-enable */
        changeControlVisibility(false);
    };

    return (
        <div
            className="club-card"
            onContextMenu={handleControls}
        >
            <Link
                className="club-card__link"
                to={`/card/${card.id}`}
            >
                <img
                    className="club-card__confirm-icon"
                    src={confirmIcon}
                    alt="Confirm icon"
                />
                <img
                    className="club-card__price-gold"
                    src={priceGoldIcon}
                    alt="Price icon"
                />
                <PlayerCard
                    card={card}
                    parentClassName={'club-card'}
                />
            </Link>
            {controlVisibility &&
                <div className="club-card__control"
                    onClick={(e) => handleSelling(e)}>
                    Sell card
                </div>
            }
        </div>
    );
};
