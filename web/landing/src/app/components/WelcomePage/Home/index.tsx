// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import React from 'react';
import PlayerIllustration from '@static/images/home/Player-Illustration.png';
import { ScrollTop } from '../ScrollTop';

import './index.scss';

export const Home: React.FC = () => {

    return (
        <section className="home" id="home">
            <div className="home__wrapper">
                <div className="home__text-area"
                    data-aos="fade-right"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <h2 className="home__value">10 000</h2>
                    <h3 className="home__title">Unique Collectible Player Cards.</h3>
                    <p className="home__description">
                        Get one to become UD founder and join the Play-to-Earn game.
                        Build your club in the metaverse.
                    </p>
                </div>
                <img
                    src={PlayerIllustration}
                    alt="Player Illustration"
                    className="home__player-image"
                    data-aos="fade-left"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                />
                <p className="home__description-mobile">
                        Get one to become UD founder and join the Play-to-Earn game.
                        Build your club in the metaverse.
                </p>
                <ScrollTop />
            </div>
        </section >
    );
};
