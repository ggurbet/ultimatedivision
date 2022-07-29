// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import arrowIcon from '@static/img/FieldPage/arrow.svg';
import arrowActiveIcon from '@static/img/FieldPage/arrow-active.svg';

import './index.scss';
import { DropdownStyle } from '@/app/internal/dropdownStyle';

/** FilterByParameterWrapper is common wrapper component for each filter component.*/
export const FilterByParameterWrapper: React.FC<{
    showComponent: () => void;
    isVisible: boolean;
    title: string;
}> = ({ showComponent, children, isVisible, title }) =>
    <li className="filter-field__list__item">
        <div className="filter-item">
            <span className={'filter-item__title'} onClick={showComponent}>
                {title}
            </span>
            <img
                className="filter-item__picture"
                src={isVisible ? arrowActiveIcon : arrowIcon}
                alt="filter icon"
                style={isVisible ? { transform: new DropdownStyle(true).triangleRotate } : {}}
            />
            <div className={`filter-item__dropdown${isVisible ? '-active' : '-inactive'}`}>{children}</div>
        </div>
    </li>;

