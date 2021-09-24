// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import React from 'react';
import PlayerIllustration from '@static/images/home/Player-Illustration.png';

import './index.scss';

export const Home: React.FC = () => {

    return (
        <section className="ultimatedivision-home" id="home">
            <div className="wrapper">
                <div className="ultimatedivision-home__text-left">
                    <span className="value">10 000</span>
                    <span className="title">Unique Collectible Player Cards.</span>
                    <span className="description">
                        Get one to become UD founder and join the Play-to-Earn game. 
                        Build your club in the metaverse.
                    </span>
                </div>
                <picture className="ultimatedivision-home__player-image">
                    <img src={PlayerIllustration} alt="Player Illustration"></img>
                </picture>
            </div>
        </section>
    );
};
