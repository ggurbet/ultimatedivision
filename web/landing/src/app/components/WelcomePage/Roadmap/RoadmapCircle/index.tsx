// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import { RoadmapPoint } from '@components/WelcomePage/Roadmap/RoadmapPoint';

import doneImg from '@static/images/Roadmap/roadmapDone.svg';

import './index.scss';

export const RoadmapCircle: React.FC<{
    item: {
        date: string,
        title: string,
        points: string[],
        done: boolean,
        id: number,
    }
}> = ({ item }) => {
    useEffect(() => {
        Aos.init({
            duration: 1000,
        });
    });

    return (
        <div
            className="roadmap-circle"
            style={
                item.done
                    ? { backgroundImage: `url(${doneImg})` }
                    : { background: '#1E4175' }}
        >
            <RoadmapPoint
                key={item.id}
                {...item}
            />
            <div
                className="roadmap-circle__pseudo-element"
                data-aos={item.id % 2 === 0 ? 'showRight' : 'showLeft'}
                data-aos-easing="ease-in-out-cubic"
                data-aos-duration="700"
                data-aos-delay={200 * item.id}
            ></div>
        </div>
    );
};
