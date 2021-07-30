//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.
import { useState } from 'react';
import { useDispatch } from 'react-redux';

import { PlayerCard } from '@components/PlayerCard';

import { Card } from '@/app/store/reducers/footballerCard';
import { removeCard } from '@/app/store/reducers/footballField';

import './index.scss';

export const PlayingAreaFootballerCard: React.FC<{ card: Card; index?: number; place?: string }> = ({ card, index, place }) => {
    const dispatch = useDispatch();
    const [visibility, changeVisibility] = useState(false);
    const style = visibility ? 'block' : 'none';
    /** remove player card implementation */
    function handleDeletion(e: any) {
        e.preventDefault();
        dispatch(removeCard(index));
    }

    return (
        <div
            onClick={() => changeVisibility(prev => !prev)}
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
