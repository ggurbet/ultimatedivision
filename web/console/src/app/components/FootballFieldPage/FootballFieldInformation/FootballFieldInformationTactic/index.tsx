/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { useState } from 'react';
import { useDispatch } from 'react-redux';

import { FotballFieldInformationLine } from '@/app/types/fotballerCard';
import { handleTactics } from '@/app/store/reducers/footballField';

import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';

export const FootballFieldInformationTactic: React.FC<{ props: FotballFieldInformationLine }> = ({ props }) => {
    const [optionVisibility, changeVisibility] = useState(true);

    const listHeight = optionVisibility ? '0' : '90px';
    const triangleRotate = optionVisibility ? 'rotate(-90deg)' : 'rotate(0deg)';

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
                    style={{ transform: triangleRotate }}
                    alt="triangle img"
                    id={`triangle-${props.id}`}
                />
            </div>
            <ul
                style={{ height: listHeight }}
                className="football-field-information-option__list"
                id={props.id}
            >
                {props.options.map((item, index) =>
                    <li
                        key={index}
                        className="football-field-information-option__item"
                        onClick={() => dispatch(handleTactics)}
                    >
                        {item}
                    </li>,
                )}
            </ul>
        </div>
    );
};
