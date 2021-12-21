// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, createContext } from 'react';

import search from '@static/img/FilterField/search.svg';
import filters from '@static/img/MarketPlacePage/filter_icon.svg';

import { Context } from '@/app/types/filters';

import './index.scss';

export const FilterContext = createContext(new Context);

export const FilterField: React.FC = ({ children }) => {
    /** Exposes default index which does not exist in array. */
    const DEFAULT_FILTER_ITEM_INDEX = -1;
    const [activeFilterIndex, setActiveFilterIndex] = useState(DEFAULT_FILTER_ITEM_INDEX);

    return (
        <FilterContext.Provider value={new Context(activeFilterIndex, setActiveFilterIndex)}>
            <section className="filter-field">
                <div className="filter-field__wrapper">
                    <div className="filter-field__use-filters">
                        <img
                            className="filter-field__use-filters__picture"
                            src={filters}
                            alt="use fitlers"
                        />
                        <span className="filter-field__use-filters__title">
                            Use filters
                        </span>
                    </div>
                    <ul className="filter-field__list">
                        <li className="filter-field__list__item">
                            <img
                                src={search}
                                alt="Filter Icon"
                                className="filter-field__list__item__picture"
                            />
                        </li>
                        {children}
                    </ul>
                </div>
            </section>
        </FilterContext.Provider>
    );
};
