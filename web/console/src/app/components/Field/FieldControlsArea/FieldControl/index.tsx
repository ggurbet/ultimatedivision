// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import triangle from '@static/img/FieldPage/triangle.svg';

import { RootState } from '@/app/store';
import { Control } from '@/app/types/club';
import { DropdownStyle } from '@/app/internal/dropdownStyle';

import './index.scss';

export const FieldControl: React.FC<{ props: Control }> = ({ props }) => {
    const squad = useSelector((state: RootState) => state.clubsReducer.activeClub.squad);
    const [optionVisibility, changeVisibility] = useState(false);
    const optionStyle = new DropdownStyle(optionVisibility);

    const dispatch = useDispatch();

    return (
        <div className="field-control">
            <div
                className="field-control__heading"
                onClick={() => changeVisibility((prev) => !prev)}
            >
                <h4 className="field-control__title">{props.title}</h4>
                <img
                    className="field-control__image"
                    src={triangle}
                    style={{ transform: optionStyle.triangleRotate }}
                    alt="triangle img"
                    id={`triangle-${props.id}`}
                />
            </div>
            <ul
                style={{ height: optionStyle.listHeight }}
                className="field-control__list"
                id="0"
            >
                {props.options.map((item, index) =>
                    <li
                        key={index}
                        className="field-control__item"
                        onClick={() => dispatch(props.action(squad, item))}
                    >
                        {item}
                    </li>
                )}
            </ul>
        </div>
    );
};
