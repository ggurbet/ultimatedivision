// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import './index.scss';

export const RoadmapPoint: React.FC<{
    date: string,
    points: string[],
    id: number,
    done: boolean
}> = ({
    date,
    points,
    id
}) => {
    useEffect(() => {
        Aos.init({
            duration: 1000,
        });
    }, []);

    return (
        <div
            className="roadmap-point"
            data-aos={id % 2 === 0 ? 'zoom-in-left-custom' : 'zoom-in-right-custom'}
            data-aos-delay={200 * id}
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
        >
            <p className="roadmap-point__date">
                {date}
            </p>
            <ul className="roadmap-point__list">
                {points.map((point) => (
                    <li
                        className="roadmap-point__item"
                        key={points.indexOf(point)}
                    >
                        {point}
                    </li>
                ))}
            </ul>
        </div>
    );
};

