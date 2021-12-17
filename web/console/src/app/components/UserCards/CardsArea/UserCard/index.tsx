// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch } from 'react-redux';
import { Link } from 'react-router-dom';

import { PlayerCard } from '@components/common/PlayerCard';

import confirmIcon from '@static/img/MarketPlacePage/MyCard/ok.svg';
import priceGoldIcon from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';

import { createLot } from '@/app/store/actions/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { Card } from '@/card';

import './index.scss';
import { useEffect, useState } from 'react';

export const UserCard: React.FC<{
    card: Card;
    position: number;
    sellingCardIndex: number;
    setIndex: React.Dispatch<React.SetStateAction<number>>;
}> = ({
    card,
    position,
    sellingCardIndex,
    setIndex,
}) => {
    /** Default index which does not exist in array */
    const DEFAULT_INDEX = -1;
    const dispatch = useDispatch();
    const [sellButtonVisibility, setVisibility] = useState(false);
    const isVisible = sellButtonVisibility && position === sellingCardIndex;
    useEffect(() => {
        position !== sellingCardIndex &&
                setVisibility(false);
    }, [sellingCardIndex]);

    const handleControls = (e: React.MouseEvent<HTMLInputElement>, position: number) => {
        e.preventDefault();
        setIndex(position);
        setVisibility(prev => !prev);
    };

    const handleSelling = (e: React.MouseEvent<HTMLInputElement>) => {
        e.stopPropagation();
        e.nativeEvent.stopImmediatePropagation();
        /** TODO: create interface for adding selling parameters */
        /* eslint-disable */
            dispatch(createLot(new CreatedLot(card.id, 200, 200, 1)));
            /* eslint-enable */
        setIndex(DEFAULT_INDEX);
        setVisibility(false);
    };

    return (
        <div
            className="user-card"
            onContextMenu={(e: React.MouseEvent<HTMLInputElement>) => handleControls(e, position)}
        >
            <Link
                className="user-card__link"
                to={`/card/${card.id}`}
            >
                <img
                    className="user-card__confirm-icon"
                    src={confirmIcon}
                    alt="Confirm icon"
                />
                <img
                    className="user-card__price-gold"
                    src={priceGoldIcon}
                    alt="Price icon"
                />
                <PlayerCard
                    id={card.id}
                    className={'user-card__image'}
                />
            </Link>
            {isVisible &&
                    <div className="user-card__control"
                        onClick={(e: React.MouseEvent<HTMLInputElement>) => handleSelling(e)}>
                        Sell card
                    </div>
            }
        </div>
    );
};
