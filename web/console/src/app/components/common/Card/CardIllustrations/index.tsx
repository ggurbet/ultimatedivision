// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardIllustrationsDiagramsArea } from '@/app/components/common/Card/CardIllustrationsDiagramsArea';
import { FootballerCardIllustrationsRadar } from '@/app/components/common/Card/CardIllustrationsRadar';

import { Card } from '@/card';

import './index.scss';

export const FootballerCardIllustrations: React.FC<{ card: Card }> = ({
    card,
}) =>
    <div className="footballer-card-illustrations">
        <div className="footballer-card-illustrations__card">
            <img
                className="footballer-card-illustrations__card__background-type"
                /** TODO: check for undefined will removed after correct Card type */
                src={card.style && card.style.background}
                alt="background img"
                draggable={false}
            />
            <div className="footballer-card-illustrations__card__wrapper">
                <img
                    className="footballer-card-illustrations__card__wrapper-face-picture"
                    src={card.face}
                    alt="Player face"
                    draggable={false}
                />
            </div>
            <span className="footballer-card-illustrations__card__name">
                {card.playerName}
            </span>
            <ul className="footballer-card-illustrations__card__list">
                {card.statsArea.map((property, index) =>
                    <li
                        className="footballer-card-illustrations__card__list__item"
                        key={index}
                    >
                        {
                            /**
                             * get only average value of player's game property
                             */
                            `${property.abbreviated} ${property.average} `
                        }
                    </li>
                )}
            </ul>
        </div>
        <div className="footballer-card-illustrations__divider"></div>
        <FootballerCardIllustrationsRadar card={card} />
        <div className="footballer-card-illustrations__divider"></div>
        <FootballerCardIllustrationsDiagramsArea card={card} />
    </div>;

