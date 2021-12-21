// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import rectangle from '@static/img/FilterField/rectangle.svg';
import reverseRectangle from '@static/img/FilterField/reverseRectangle.svg';

import './index.scss';

/** FilterByParameterWrapper is common wrapper component for each filter component.*/
export const FilterByParameterWrapper: React.FC<{
    showComponent: () => void;
    isVisible: boolean;
    title: string;
}> = ({ showComponent, children, isVisible, title }) =>
    <li className="filter-field__list__item">
        <div
            className="filter-item"
        >
            <span
                className={`filter-item__title${isVisible ? '-active' : '-inactive'}`}
                onClick={showComponent}
            >
                {title}
            </span>
            <img
                className="filter-item__picture"
                src={isVisible ? reverseRectangle : rectangle}
                alt="filter icon"
            />
            <div className={`filter-item__dropdown${isVisible ? '-active' : '-inactive'}`} >
                {children}
            </div>
        </div>
    </li>;


