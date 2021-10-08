// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import './index.scss';

export const RoadmapPoint: React.FC<{
    item: {
        date: string,
        points: string[],
        id: number,
        done: boolean
    }
}> = ({
    item
}) => {
    useEffect(() => {
        Aos.init({
            duration: 1000,
        });
    }, []);

    return (
        <div
            className="roadmap-point"
            data-aos={item.id % 2 === 0 ? 'zoom-in-left-custom' : 'zoom-in-right-custom'}
            data-aos-delay={200 * item.id}
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
        >
            <p className="roadmap-point__date">
                {item.date}
            </p>
            <ul className="roadmap-point__list">
                {item.points.map((point, index) => (
                    <li
                        className="roadmap-point__item"
                        key={index}
                    >
                        {point}
                    </li>
                ))}
            </ul>
        </div>
    );
};

