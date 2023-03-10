// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { createContext, useState } from 'react';

import { Context } from '@/app/types/filters';

import { setScrollAble } from '@/app/internal/setScrollAble';

import { SearchIcon } from '@/app/static/img/FilterField/search';
import { ArrowBackIcon } from '@/app/static/img/FilterField/arrowBack';
import filterIcon from '@static/img/FieldPage/filter-icon.svg';

import './index.scss';

export const FilterContext = createContext(new Context());

export const FilterField: React.FC = ({ children }) => {
    /** Exposes default index which does not exist in array. */
    const DEFAULT_FILTER_ITEM_INDEX = -1;
    const [activeFilterIndex, setActiveFilterIndex] = useState(DEFAULT_FILTER_ITEM_INDEX);
    const [isActiveMobileCardsFilter, setIsActiveMobileCardsFilter] = useState(false);

    const openCardsFilter = () => {
        setIsActiveMobileCardsFilter(true);
        setScrollAble(false);
    };

    const closeCardsFilter = () => {
        setIsActiveMobileCardsFilter(false);
        setScrollAble(true);
    };

    return (
        <FilterContext.Provider value={new Context(activeFilterIndex, setActiveFilterIndex)}>
            <section className="filter-field">
                <div className="filter-field__filters">
                    <div className="filter-field__list__item filter-field__list__item__search filter-field__list__item__mobile-search">
                        <SearchIcon/>
                        Search
                    </div>

                    <div className="filter-field__use-filters" onClick={() => openCardsFilter()}>
                        <img className="filter-field__use-filters__picture" src={filterIcon} alt="use fitlers" />
                        <span className="filter-field__use-filters__title">Card filter</span>
                    </div>
                </div>
                {isActiveMobileCardsFilter ?
                    <div className={`filter-field__mobile ${isActiveMobileCardsFilter ? 'filter-field__mobile--active' : ''} `}>
                        <div className="filter-field__mobile__content">
                            <div className="filter-field__mobile__top-side">
                                <span onClick={() => closeCardsFilter()}
                                    className="filter-field__mobile__top-side__arrow-left">
                                    <ArrowBackIcon/>
                                </span>
                                <h2 className="filter-field__mobile__top-side__title">
                                    Filter
                                </h2>
                            </div>
                            <ul className="filter-field__mobile__list">
                                {children}
                            </ul>
                        </div>
                    </div> :
                    <ul className="filter-field__list">
                        <li className="filter-field__list__item filter-field__list__item__search">
                            <SearchIcon/>
                            Search
                        </li>
                        {children}
                    </ul>}
            </section>
        </FilterContext.Provider>
    );
};
