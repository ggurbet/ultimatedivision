// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import './index.scss';

export const RoadmapPoint: React.FC<{
    date: string,
    title: string,
    points: string[],
    id: number,
    done: boolean
}> = ({
    date,
    title,
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
            data-aos={id % 2 === 0 ? 'zoom-in-left' : 'zoom-in-right-custom'}
            data-aos-delay={200 * id}
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
            className="roadmap-point"
        >
            <p className="roadmap-point__date">
                {date}
            </p>
            <h2 className="roadmap-point__title">
                {title}
            </h2>
            <ul className="roadmap-point__list">
                {points.map((point) => (
                    <li data-aos={
                        id % 2 === 0 ? 'fade-left' : 'fade-right-custom'
                    }
                    data-aos-easing="ease-in-out-cubic"
                    data-aos-duration="700"
                    data-aos-delay={
                        points.length > 2
                            ? 250 * (points.indexOf(point) + 1)
                            : 500 * (points.indexOf(point) + 1)
                    }
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

