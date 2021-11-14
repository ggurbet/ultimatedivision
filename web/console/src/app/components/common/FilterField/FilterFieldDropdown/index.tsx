// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import './index.scss';

export const FilterFieldDropdown: React.FC<{
    props: { label: string; image: string };
}> = ({ props }) => {
    const { label, image } = props;
    const [shouldDropdownShow, handleShowing] = useState(false);

    return (
        <div
            className="filter-item"
            onClick={() => handleShowing((prev) => !prev)}
        >
            <span className="filter-item__title">{label}</span>
            <img
                className="filter-item__picture"
                src={image}
                alt={image && 'filter icon'}
            />
            <div
                className={`filter-item__dropdown${
                    shouldDropdownShow ? '-active' : '-inactive'
                }`}
            ></div>
        </div>
    );
};
