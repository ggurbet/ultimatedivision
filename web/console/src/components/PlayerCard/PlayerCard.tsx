/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { NavLink } from 'react-router-dom';

import { Card } from '../../store/reducers/footballerCard';
import { RouteConfig } from '../../routes';

export const PlayerCard: React.FC<{ card: Card; parentClassName: string }> = ({
    card, parentClassName,
}) =>
    <>
        <img
            className={parentClassName + '__background-type'}
            src={card.mainInfo.backgroundType}
            alt='background img'
            draggable={false}
        />
        <img
            className={parentClassName + '__face-picture'}
            src={card.mainInfo.playerFace}
            alt="Player face"
            draggable={false}
        />
        <NavLink to={RouteConfig.FootballerCard.path} >
            <span className={parentClassName + '__name'}>
                {card.mainInfo.lastName}
            </span>
        </NavLink>
        <ul className={parentClassName + '__list'}>
            {card.stats.map(
                (property, index) =>
                    <li
                        className={parentClassName + '__list__item'}
                        key={index}>
                        {
                            /**
                                * get only average value of player's game property
                                */
                            `${property.abbreviated} ${property.average} `
                        }
                    </li>
            )}
        </ul>
    </>
    ;
