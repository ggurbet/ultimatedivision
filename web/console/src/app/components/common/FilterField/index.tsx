// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { AnyAction, Dispatch } from 'redux';
import rectangle
    from '@static/img/FilterField/rectangle.svg';
import search
    from '@static/img/FilterField/search.svg';
import star
    from '@static/img/FilterField/star.svg';
import fut
    from '@static/img/FilterField/fut.svg';
import eye
    from '@static/img/FilterField/eye.svg';
import parametres
    from '@static/img/FilterField/parametres.svg';
import filters
    from '@static/img/MarketPlacePage/filter_icon.svg';
import { FilterFieldDropdown } from './FilterFieldDropdown';

import './index.scss';

export const FilterField: React.FC<{
    title: string;
    thunk: (lowRange: string, topRange: string) => (dispatch: Dispatch<AnyAction>) => Promise<void>;
}> = ({ title, thunk }) => {
    const [searchData, setSearchData] = useState('');

    /** Class for each filter field item */
    class FilterItem {
        /** labe is text filling of field, image - src for image  */
        constructor(
            public label: string,
            public image: string,
        ) { }
    }

    const handleSerchChange = (event: any) => {
        setSearchData(event.target.value);
    };

    const filterFieldTitles: FilterItem[] = [
        new FilterItem('Version', rectangle),
        new FilterItem('Positions', rectangle),
        new FilterItem('Nations', rectangle),
        new FilterItem('Leagues', rectangle),
        new FilterItem('WRF', rectangle),
        new FilterItem('Stats', rectangle),
        new FilterItem('', star),
        new FilterItem('PS', fut),
        new FilterItem('T&S', rectangle),
        new FilterItem('', eye),
        new FilterItem('', star),
        new FilterItem('RPP', rectangle),
        new FilterItem('', parametres),
        new FilterItem('Misc', rectangle),
    ];

    return (
        <section className="filter-field">
            <h1 className="filter-field__title">
                {title}
            </h1>
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
                        <input
                            value={searchData}
                            placeholder="Search"
                            className="filter-field__list__item__search"
                            onChange={handleSerchChange}
                        />
                    </li>
                    {filterFieldTitles.map((item, index) =>
                        <li
                            className="filter-field__list__item"
                            key={index}
                        >
                            <FilterFieldDropdown props={item} />
                        </li>
                    )}
                </ul>
            </div>
        </section >
    );
};
