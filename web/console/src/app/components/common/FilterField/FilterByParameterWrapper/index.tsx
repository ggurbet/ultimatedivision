// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ArrowActiveIcon, ArrowIcon } from '@/app/static/img/FilterField/arrows';

import { DropdownStyle } from '@/app/internal/dropdownStyle';

import './index.scss';

/** FilterByParameterWrapper is common wrapper component for each filter component.*/
export const FilterByParameterWrapper: React.FC<{
    showComponent: () => void;
    isVisible: boolean;
    title: string;
}> = ({ showComponent, children, isVisible, title }) =>
    <li className="filter-field__list__item">
        <div className="filter-item">
            <div className="filter-item__content" onClick={showComponent}>
                <span className="filter-item__title">
                    {title}
                </span>
                <span
                    className={`filter-item__picture ${isVisible ?'filter-item__picture__visible':''}` }
                    style={isVisible ? { transform: new DropdownStyle(true).triangleRotate } : {}}
                >
                    {isVisible ? <ArrowIcon /> : <ArrowActiveIcon/>}
                </span>
            </div>
            <div className={`filter-item__dropdown${isVisible ? '-active' : '-inactive'}`}>{children}</div>
        </div>
    </li>;

