// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import cards from '@static/images/Description/cardsGroup.png';

import './index.scss';

export const DescriptionCards = () => {
    return (
        <div className="description-cards" id="cards">
            <div className="description-cards__text-area">
                <h2 className="description-cards__title">
                    The Player Cards - Become UD Founder
                </h2>
                <p className="description-cards__text">
                    Each player in your club is an NFT - build a squad of 11 NFTs,
                    for your team to compete week in, week out.
                    These NFT’s have the player’s stats, which determine how strong
                    each player is. If you want to score a limited Founder Collection
                    NFT, then make sure you are ready for date.
                </p>
            </div>
            <div className="description-cards__wrapper">
                <img
                    className="description-cards__card"
                    src={cards}
                    alt="cards"
                />
            </div>
        </div>
    );
};
