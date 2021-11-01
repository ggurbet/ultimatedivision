// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import cards from '@static/images/description/cardsGroup.webp';
import mobileCards from '@static/images/description/mobile-cards.png';
import webkitCards from '@static/images/description/cardsGroup.png';

import './index.scss';

export const DescriptionCards = () => {
    return (
        <div className="description-cards" id="cards">
            <div className="description-cards__text-area">
                <h2 className="description-cards__title">
                    The Player Cards - Become UD Founder
                </h2>
                <p className="description-cards__text">
                    Each player in your club is an NFT - build a squad of 11
                    NFTs, for your team to compete week in, week out. These
                    NFT’s have the player’s stats, which determine how strong
                    each player is. If you want to score a limited Founder
                    Collection NFT, then make sure you are ready for date.
                </p>
            </div>
            <div className="description-cards__wrapper">
                <picture>
                    <source
                        media="(min-width: 601px)"
                        srcSet={cards}
                        type="image/webp"
                    />
                    <source media="(max-width: 600px)" srcSet={mobileCards} />
                    <img
                        className="description-cards__card"
                        src={webkitCards}
                        alt="cards"
                        loading="lazy"
                    />
                </picture>
            </div>
        </div>
    );
};
