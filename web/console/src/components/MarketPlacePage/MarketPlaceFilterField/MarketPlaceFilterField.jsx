/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';

import './marketPlaceFilterField.scss';

export const MarketPlaceFilterField = () => {
    const filterFieldTitles = useSelector(
        (state) => state.filterFieldTitles
    );

    return (
        <section className="marketplace-filter">
            <h1 className="marketplace-filter__title">
                MARKETPLACE
            </h1>
            <div className="marketplace-filter__wrapper">
                <ul className="marketplace-filter__list">
                    {filterFieldTitles.map(item => {
                        return (
                            <li key={filterFieldTitles.indexOf(item)}
                                className="marketplace-filter__list__item">
                                {item.title}
                                <img
                                    src={item.src}
                                    alt="Filter icon"
                                    className="marketplace-filter__list__item__picture"
                                />
                            </li>
                        );
                    })}
                </ul>
            </div>
        </section>
    );
};

