// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardWithStats } from '@/card';

export const PlayerCard: React.FC<{ card: CardWithStats; parentClassName: string }> = ({
    card,
    parentClassName,
}) =>
    <>
        <img
            className={`${parentClassName}__background-type`}
            /** TODO: check for undefined will removed after correct Card type */
            src={card.style && card.style.background}
            alt="background img"
            draggable={false}
        />
        <img
            className={`${parentClassName}__face-picture`}
            src={card.face}
            alt="Player face"
            draggable={false}
        />
        <span className={`${parentClassName}__name`}>{card.playerName}</span>
        <ul className={`${parentClassName}__list`}>
            {card.statsArea.map((property, index) =>
                <li className={`${parentClassName}__list__item`} key={index}>
                    {
                        /**
             * get only average value of player's game property
             */
                        `${property.abbreviated} ${property.average} `
                    }
                </li>
            )}
        </ul>
    </>;

