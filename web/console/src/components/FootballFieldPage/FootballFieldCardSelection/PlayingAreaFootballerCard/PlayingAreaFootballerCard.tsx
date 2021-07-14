/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { PlayerCard } from '../../../PlayerCard/PlayerCard';

import { Card }
    from '../../../../store/reducers/footballerCard';
import { RootState } from '../../../../store';
import { addCard, removeCard }
    from '../../../../store/reducers/footballField';
import { FootballCardStyle }
    from '../../../../utils/footballField';

import './PlayingAreaFootballerCard.scss';

export const PlayingAreaFootballerCard: React.FC<{ card: Card; index?: number; place?: string }> = ({ card, index, place }) => {
    const dispatch = useDispatch();
    const chosenCard = useSelector((state: RootState) => state.fieldReducer.options.chosedCard);
    const [visibility, changeVisibility] = useState(false);
    const style = new FootballCardStyle(visibility).style;
    /** remove player card implementation */
    function handleDeletion(e: any) {
        e.preventDefault();
        dispatch(removeCard(index));
    }

    return (
        <div
            onClick={place ? () => changeVisibility(prev => !prev) : () => dispatch(addCard(card, chosenCard))}
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
