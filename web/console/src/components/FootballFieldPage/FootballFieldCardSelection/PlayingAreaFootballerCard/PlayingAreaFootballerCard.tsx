/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './PlayingAreaFootballerCard.scss';

import { Card }
    from '../../../../store/reducers/footballerCard';
import { RootState } from '../../../../store';
import { useDispatch, useSelector } from 'react-redux';
import { addCard, removeCard }
    from '../../../../store/reducers/footballField';
import { useState } from 'react';
import { FootballCardStyle }
    from '../../../../utils/footballField';

export const PlayingAreaFootballerCard: React.FC<{ card: Card, index?: number, place?: string }> = ({ card, index, place }) => {
    const dispatch = useDispatch();
    const chosenCard = useSelector((state: RootState) => state.fieldReducer.options.chosedCard);
    const [visibility, changeVisibility] = useState(false);
    const style = new FootballCardStyle(visibility).style;

    return (
        <div
            onClick={place ? () => changeVisibility(prev => !prev) : () => dispatch(addCard(card, chosenCard))}
            className="football-field-card"
        >
            <img
                className="football-field-card__background"
                src={card.mainInfo.backgroundType}
                alt='background img'
                draggable={false}
            />
            <img className="football-field-card__face-picture"
                src={card.mainInfo.playerFace}
                draggable={false}
                alt="Player face" />
            <span className="football-field-card__name">
                {card.mainInfo.lastName}
            </span>
            <ul className="football-field-card__list">
                {card.stats.map(
                    (property, index) => {
                        return (
                            <li
                                className="football-field-card__list__item"
                                key={index}>
                                {
                                    /**
                                    * get only average value of player's game property
                                    */
                                    `${property.abbreviated} ${property.average} `
                                }
                            </li>
                        );
                    }
                )}
            </ul>
            <div
                style={{ display: style }}
                onClick={() => dispatch(removeCard(index))}
                className="football-field-card__control">
                &#10006; delete a player
            </div>
        </div>
    );
};

