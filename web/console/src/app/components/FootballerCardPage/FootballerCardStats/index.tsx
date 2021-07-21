/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { CardStats } from '../../../types/fotballerCard';

import './index.scss';

export const FootballerCardStats: React.FC<{ props: CardStats }> = ({ props }) =>
    <div className="footballer-card-stats-item">
        <div
            className="footballer-card-stats-item__heading"
        >
            <span className="footballer-card-stats-item__heading-name"
            >
                {props.title}
            </span>
            <span
                className="footballer-card-stats-item__heading-value"
                style={{ color: props.color }}
            >
                {props.average}
            </span>
        </div>
        <div className="footballer-card-stats-item__diagram">
            <div
                className="footballer-card-stats-item__diagram-value"
                style={{ background: props.color, width: `${props.average}%` }}
            ></div>
        </div>
        {props.fields.map((item, index) =>
            <div
                key={index}
                className="footballer-card-stats-item__row"
            >
                <span className="footballer-card-stats-item__row-name"
                >
                    {item.label}
                </span>
                <span
                    className="footballer-card-stats-item__row-value"
                >
                    {item.value}
                </span>
            </div>,
        )}
    </div>;
