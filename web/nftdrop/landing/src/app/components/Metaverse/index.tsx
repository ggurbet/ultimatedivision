// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MintButton } from '@components/common/MintButton';

import webkitCards from '@static/images/metaverse/cards.png';
import webkitCardsTablet from '@static/images/metaverse/cards-tablet.png';
import webkitCardsMobile from '@static/images/metaverse/cards-mobile.png';
import cards from '@static/images/metaverse/cards.webp';
import cardsTablet from '@static/images/metaverse/cards-tablet.webp';
import cardsMobile from '@static/images/metaverse/cards-mobile.webp';

import './index.scss';

export const Metaverse: React.FC = () => {

    return (
        <section className="metaverse" id="metaverse">
            <div className="metaverse__wrapper">
                <h2 className="metaverse__title">
                    Ultimate Divison
                </h2>
                <h3 className="metaverse__subtitle">
                    Football Metaverse
                </h3>
                <picture>
                    <source media="(max-width: 600px)" srcSet={cardsMobile} type="image/webp" />
                    <source media="(max-width: 800px)" srcSet={cardsTablet} type="image/webp" />
                    <source media="(min-width: 800px)" srcSet={cards} type="image/webp" />
                    <source media="(max-width: 600px)" srcSet={webkitCardsMobile} />
                    <source media="(max-width: 800px)" srcSet={webkitCardsTablet} />
                    <img
                        className="metaverse__cards"
                        src={webkitCards}
                        alt="cards"
                        loading="lazy"
                    />
                </picture>
                <div className="metaverse__sold-scale">
                    <span className="metaverse__sold-scale__text">Cards Sold 0/10000</span>
                </div>
                <MintButton />
            </div>
        </section>
    );
};
