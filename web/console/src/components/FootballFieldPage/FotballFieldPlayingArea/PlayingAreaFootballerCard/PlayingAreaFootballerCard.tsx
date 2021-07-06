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
import { handleCard }
    from '../../../../store/reducers/footballField';

export const PlayingAreaFootballerCard: React.FC<{ card: Card, place?:string }> = ({ card, place }) => {

    const dispatch = useDispatch();
    const chosenCard = useSelector((state: RootState) => state.fieldReducer.options.chosedCard);

    return (
        <div
            onClick={place? () => {} : () => dispatch(handleCard(card, chosenCard))}
            className="football-field-card"
            data-background={card.mainInfo.backgroundType}
        >
            <img
                className="football-field-card__background"
                src={card.mainInfo.backgroundType}
                alt='background img'
            />
            <img className="football-field-card__face-picture"
                src={card.mainInfo.facePicture}
                alt="Player face" />
            <span className="football-field-card__name">
                {card.overalInfo[0].value}
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
                                    `${property.average} ${property.abbr}`
                                }
                            </li>
                        );
                    }
                )}
            </ul>
        </div>
    );
};

