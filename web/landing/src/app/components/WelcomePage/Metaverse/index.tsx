// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import cards from '@static/images/metaverse/cards.svg';

import './index.scss';
import { MintButton } from '@components/common/MintButton';

export const Metaverse: React.FC = () => {
    return (
        <section className="metaverse" id="metaverse">
            <div className="metaverse__wrapper">
                <h2 className="metaverse__title"
                    data-aos="fade-right"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    Ultimate Divison
                </h2>
                <h3 className="metaverse__subtitle"
                    data-aos="fade-right"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    Football Metaverse
                </h3>
                <img
                    className="metaverse__cards"
                    data-aos="fade-right"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                    src={cards}
                    alt="cards"
                    loading="lazy"
                />
                <div className="metaverse__sold-scale"
                    data-aos="fade-left"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <span className="metaverse__sold-scale__text">Cards Sold 0/10000</span>
                </div>
                <MintButton />
            </div>
        </section>
    );
};
