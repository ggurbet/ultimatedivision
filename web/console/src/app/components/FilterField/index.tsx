// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

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
import stars
    from '@static/img/FilterField/stars.svg';
import parametres
    from '@static/img/FilterField/parametres.svg';

import './index.scss';
import { FilterFieldDropdown } from './FilterFieldDropdown';

export const FilterField: React.FC<{ title: string }> = ({ title }) => {
    const [searchData, setSearchData] = useState('');

    const handleSerchChange = (event: any) => {
        setSearchData(event.target.value);
    };

    const filterFieldTitles: Array<{ label: string; src: string }> = [
        {
            label: 'Version',
            src: rectangle,
        },
        {
            label: 'Positions',
            src: rectangle,
        },
        {
            label: 'Nations',
            src: rectangle,
        },
        {
            label: 'Leagues',
            src: rectangle,
        },
        {
            label: 'WRF',
            src: rectangle,
        },
        {
            label: 'Stats',
            src: rectangle,
        },
        {
            label: '',
            src: star,
        },
        {
            label: 'PS',
            src: fut,
        },
        {
            label: 'T&S',
            src: rectangle,
        },
        {
            label: '',
            src: eye,
        },
        {
            label: '',
            src: stars,
        },
        {
            label: 'RPP',
            src: rectangle,
        },
        {
            label: '',
            src: parametres,
        },
        {
            label: 'Misc',
            src: rectangle,
        },
    ];

    return (
        <section className="filter-field">
            <h1 className="filter-field__title">
                {title}
            </h1>
            <div className="filter-field__wrapper">
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
