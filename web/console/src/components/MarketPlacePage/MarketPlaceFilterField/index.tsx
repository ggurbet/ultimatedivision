/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';

import './index.scss';

import rectangle
    from '../../../img/MarketPlacePage/marketPlaceFilterField/rectangle.png';
import search
    from '../../../img/MarketPlacePage/marketPlaceFilterField/search.png';
import star
    from '../../../img/MarketPlacePage/marketPlaceFilterField/star.png';
import fut
    from '../../../img/MarketPlacePage/marketPlaceFilterField/fut.png';
import eye
    from '../../../img/MarketPlacePage/marketPlaceFilterField/eye.png';
import stars
    from '../../../img/MarketPlacePage/marketPlaceFilterField/stars.png';
import parametres
    from '../../../img/MarketPlacePage/marketPlaceFilterField/parametres.png';

export const MarketPlaceFilterField: React.FC<{ title: string }> = ({ title }) => {
    const filterFieldTitles: Array<{ title: string; src: string }> = [
        {
            title: 'Search',
            src: search,
        },
        {
            title: 'Version',
            src: rectangle,
        },
        {
            title: 'Positions',
            src: rectangle,
        },
        {
            title: 'Nations',
            src: rectangle,
        },
        {
            title: 'Leagues',
            src: rectangle,
        },
        {
            title: 'WRF',
            src: rectangle,
        },
        {
            title: 'Stats',
            src: rectangle,
        },
        {
            title: '',
            src: star,
        },
        {
            title: 'PS',
            src: fut,
        },
        {
            title: 'T&S',
            src: rectangle,
        },
        {
            title: '',
            src: eye,
        },
        {
            title: '',
            src: stars,
        },
        {
            title: 'RPP',
            src: rectangle,
        },
        {
            title: '',
            src: parametres,
        }
        ,
        {
            title: 'Misc',
            src: rectangle,
        },
    ];

    return (
        <section className="marketplace-filter">
            <h1 className="marketplace-filter__title">
                {title}
            </h1>
            <div className="marketplace-filter__wrapper">
                <ul className="marketplace-filter__list">
                    {filterFieldTitles.map((item, index) =>
                        <li key={index}
                            className="marketplace-filter__list__item">
                            {item.title}
                            <img
                                src={item.src}
                                alt="Filter icon"
                                className="marketplace-filter__list__item__picture"
                            />
                        </li>
                    )}
                </ul>
            </div>
        </section>
    );
};
