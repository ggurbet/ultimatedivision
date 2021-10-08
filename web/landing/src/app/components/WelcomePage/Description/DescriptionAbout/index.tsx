// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';
import field from '@static/images/Description/field.png';
import fieldMirrored from '@static/images/Description/fieldMirrored.png';
import cardStats from '@static/images/Description/cardStats.svg';

export const DescriptionAbout = () => {
    return (
        <div className="description-about" id="about">
            <div className="description-about__image-area">
                <img
                    src={field}
                    alt="field image"
                />
                <img
                    className="description-about__mirrored"
                    src={fieldMirrored}
                    alt="field mirrored"
                />
                <img
                    className="description-about__stats-image"
                    src={cardStats}
                    alt="card stats image"
                />
            </div>
            <div className="description-about__text-area">
                <h2 className="description-about__title">About the Game</h2>
                <p className="description-about__text">
                    Ultimate Division is a world football simulator.
                    UD players will own clubs, players and face each other
                    in weekly competitions to win cash prizes!
                    Other players can be hired as managers or coaches
                    for your Club!
                </p>
            </div>
        </div>
    );
};
