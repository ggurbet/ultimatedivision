// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import rectangle from '@static/img/FootballFieldPage/triangle.svg';
import search from '@static/img/FootballFieldPage/search.svg';

import './index.scss';

export const FilterField: React.FC = () => {
    const filterFieldTitles: Array<{ title: string; src: string }> = [
        {
            title: 'Card quality',
            src: rectangle,
        },
        {
            title: 'Overal rating',
            src: rectangle,
        },
        {
            title: 'Player`s position',
            src: rectangle,
        },

    ];

    return (
        <section className="football-field-filter">
            <div className="football-field-filter__wrapper">
                <ul className="football-field-filter__list">
                    <li className="football-field-filter__list__item">
                        <form action="" className="football-field-filter__form">
                            <button type="submit"
                                className="football-field-filter__submit"
                            >
                                <img
                                    src={search}
                                    alt="Filter icon"
                                    className="football-field-filter__item__search-picture"
                                />
                            </button>
                            <input
                                type="text"
                                placeholder="Player`s name"
                                className="football-field-filter__input"
                            />
                        </form>
                    </li>
                    {filterFieldTitles.map((item, index) =>
                        <li key={index}
                            className="football-field-filter__list__item">
                            {item.title}
                            <img
                                src={item.src}
                                alt="Filter icon"
                                className="football-field-filter__list__item__picture"
                            />
                        </li>,
                    )}
                </ul>
            </div>
        </section>
    );
};
