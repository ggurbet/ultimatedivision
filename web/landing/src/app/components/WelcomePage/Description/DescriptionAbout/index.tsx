// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';
import field from '@static/images/Description/field.png';
import fieldMirrored from '@static/images/Description/fieldMirrored.png';
import cardStats from '@static/images/Description/cardStats.svg';

export const DescriptionAbout = () => {
    return (
        <div className="description-about">
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
                    Ultimate Division is football world simulator.
                    Players can own clubs, compete with each other in weekly
                    competitions and earn money by winning.
                    Other players can be hired as managers and coaches for your own club.
                    Each assets in the game is NFT and brings profit when put to use.
                </p>
            </div>
        </div>
    );
};
