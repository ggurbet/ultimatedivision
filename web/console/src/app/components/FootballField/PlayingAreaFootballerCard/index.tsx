// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { PlayerCard } from '@components/common/PlayerCard';
import { Card } from '@/card';
import { removeCard } from '@/app/store/actions/club';

import './index.scss';

export const PlayingAreaFootballerCard: React.FC<{ card: Card; index?: number; place?: string }> = ({ card, index, place }) => {
    const dispatch = useDispatch();
    const [visibility, changeVisibility] = useState(false);
    const style = visibility ? 'block' : 'none';

    /** show/hide delete block, preventing scroll to cardSelection */
    function handleVisibility(e: any) {
        e.stopPropagation();
        changeVisibility(prev => !prev);
    }
    /** remove player card implementation */
    function handleDeletion(e: any) {
        e.stopPropagation();
        e.preventDefault();
        dispatch(removeCard(index));
    }

    return (
        <div
            onClick={handleVisibility}
            className="football-field-card"
        >
            <div
                className="football-field-card__wrapper"
                style={{ display: style }}
            ></div>
            <PlayerCard
                card={card}
                parentClassName="football-field-card"
            />
            <div
                style={{ display: style }}
                onClick={handleDeletion}
                className="football-field-card__control">
                &#10006; delete a player
            </div>
        </div >
    );
};
