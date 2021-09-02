// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Link } from 'react-router-dom';

import { RouteConfig } from '@/app/router';

import { Card } from '@/card';

export const PlayerCard: React.FC<{ card: Card; parentClassName: string }> = ({
    card, parentClassName,
}) =>
    <>
        <img
            className={`${parentClassName}__background-type`}
            src={card.style?.background}
            alt="background img"
            draggable={false}
        />
        <img
            className={`${parentClassName}__face-picture`}
            src={card.face}
            alt="Player face"
            draggable={false}
        />
        <Link
            to={{
                pathname: RouteConfig.FootballerCard.path,
                state: {
                    card,
                },
            }}
        >
            <span className={`${parentClassName}__name`}>
                {card.playerName}
            </span>
        </Link>
        <ul className={`${parentClassName}__list`}>
            {card.statsArea.map(
                (property, index) =>
                    <li
                        className={`${parentClassName}__list__item`}
                        key={index}>
                        {
                            /**
                                * get only average value of player's game property
                                */
                            `${property.abbreviated} ${property.average} `
                        }
                    </li>,
            )}
        </ul>
    </>;
