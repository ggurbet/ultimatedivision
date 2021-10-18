// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import './index.scss';

export const RoadmapPoint: React.FC<{
    item: {
        date: string,
        points: string[],
        id: number,
        done: boolean
    }
}> = ({ item }) => {

    return (
        <div className="roadmap-point">
            <p className="roadmap-point__date">
                {item.date}
            </p>
            <ul className="roadmap-point__list">
                {item.points.map((point, index) => (
                    <li
                        className="roadmap-point__item"
                        key={index}
                    >
                        <span className="roadmap-point__item__description">
                            {point}
                        </span>
                    </li>
                ))}
            </ul>
        </div>
    );
};

