/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React, { useState } from 'react';
import { useDispatch } from 'react-redux';

import { FotballFieldInformationLine } from '../../../../types/fotballerCard';

import triangle from '../../../../img/FootballFieldPage/triangle.png';
import { handleFormations } from '../../../../store/reducers/footballField';
import { ListStyle } from '../../../../utils/footballField';
import { TriangleStyle } from '../../../../utils/footballField';

import './index.scss';

export const FootballFieldInformationFormation: React.FC<{ props: FotballFieldInformationLine }> = ({ props }) => {
    const [optionVisibility, changeVisibility] = useState(true);

    const LIST_HEIGHT = new ListStyle(optionVisibility);
    const TRIANGLE_ROTATE = new TriangleStyle(optionVisibility);

    const dispatch = useDispatch();

    return (
        <div className="football-field-information-option">
            <div
                className="football-field-information-option__heading"
                onClick={() => changeVisibility(prev => !prev)}
            >
                <h4 className="football-field-information-option__title">
                    {props.title}
                </h4>
                <img
                    className="football-field-information-option__image"
                    src={triangle}
                    style={{ transform: TRIANGLE_ROTATE.style }}
                    alt="triangle img"
                    id={`triangle-${props.id}`}
                />
            </div>
            <ul
                style={{ height: LIST_HEIGHT.style }}
                className="football-field-information-option__list"
                id={props.id}
            >
                {props.options.map((item, index) =>
                    <li
                        key={index}
                        className="football-field-information-option__item"
                        onClick={() => dispatch(handleFormations(item))}
                    >
                        {item}
                    </li>
                )}
            </ul>
        </div>
    );
};
