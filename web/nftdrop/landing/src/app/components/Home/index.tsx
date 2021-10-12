// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import { ScrollTop } from '../ScrollTop';

import webkitFootballer from '@static/images/home/footballer.png';
import webkitFootballerTablet from '@static/images/home/footballerTablet.png';
import webkitFootballerMobile from '@static/images/home/footballerMobile.png';
import footballer from '@static/images/home/footballer.webp';
import footballerTablet from '@static/images/home/footballerTablet.webp';
import footballerMobile from '@static/images/home/footballerMobile.webp';

import discord from '@static/images/home/discord.svg';
import twitter from '@static/images/home/twitter.svg';

import './index.scss';

export const Home: React.FC = () => {

    return (
        <section className="home" id="home">
            <div className="home__wrapper">
                <div className="home__text-area"

                >
                    <h1 className="home__value">10 000</h1>
                    <h3 className="home__title">Unique Collectible Player Cards.</h3>
                    <p className="home__description">
                        Get one to become UD founder and join the Play-to-Earn game.
                        Build your club in the metaverse.
                    </p>
                    <div className="home__buttons-wrapper">
                        <a
                            className="home__discord"
                            href="https://discord.com/invite/ultimatedivision"
                            target="_blank"
                            rel="noreferrer"
                        >
                            <img
                                className="home__discord__logo"
                                src={discord}
                                alt="discord logo"
                            />
                            <span className="home__discord__text">
                                Join Discord
                            </span>
                        </a>
                        <a
                            className="home__twitter"
                            href="https://twitter.com/UltimateDivnft"
                            target="_blank"
                            rel="noreferrer"
                        >
                            <img
                                className="home__twitter__logo"
                                src={twitter}
                                alt="twitter logo"
                            />
                        </a>
                    </div>
                </div>
                <picture>
                    <source
                        media="(max-width: 500px)"
                        srcSet={footballerMobile} type="image/webp"
                    />
                    <source
                        media="(max-width: 800px)"
                        srcSet={footballerTablet} type="image/webp"
                    />
                    <source
                        media="(min-width: 800px)"
                        srcSet={footballer}
                        type="image/webp"
                    />
                    <source
                        media="(max-width: 500px)"
                        srcSet={webkitFootballerMobile}
                    />
                    <source
                        media="(max-width: 800px)"
                        srcSet={webkitFootballerTablet}
                    />
                    <img
                        className="home__player-image"
                        src={webkitFootballer}
                        alt="Player Illustration"
                    />
                </picture>
                <div className="home__mobile-wrapper">
                    <div className="home__description-mobile">
                        Get one to become UD founder and join the Play-to-Earn game.
                        Build your club in the metaverse.
                        <div className="home__buttons-wrapper__mobile">
                            <a
                                className="home__discord"
                                href="https://discord.com/invite/ultimatedivision"
                                target="_blank"
                                rel="noreferrer"
                            >
                                <img
                                    className="home__discord__logo"
                                    src={discord}
                                    alt="discord logo"
                                />
                                <span className="home__discord__text">
                                    Join Discord
                                </span>
                            </a>
                            <a
                                className="home__twitter"
                                href="https://twitter.com/UltimateDivnft"
                                target="_blank"
                                rel="noreferrer"
                            >
                                <img
                                    className="home__twitter__logo"
                                    src={twitter}
                                    alt="twitter logo"
                                />
                            </a>
                        </div>
                    </div>
                </div>
                <ScrollTop />
            </div>
        </section >
    );
};
