// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import { AnimationImage } from '@components/common/AnimationImage';
import { RoadmapPoint } from './RoadmapPoint';

import animationFlagData from '@static/images/roadmap/flag-animation/data.json';

import footer from '@static/images/roadmap/bottom.svg';

import './index.scss';

export const Roadmap: React.FC = () => {
    const dataList = [
        {
            id: 1,
            date: '2021 Q4',
            points: [
                'Ultimate Division Metaverse MVP launch',
                'Play to Earn mechanics available to players',
                'UDT (Ultimate Division Token) introduced',
            ],
            done: true,
        },
        {
            id: 2,
            date: '2022 Q1',
            points: [
                'Game in full swing',
                'Management roles available',
                'Club owners can tokenize their clubs and sell shares to other players',
                'UD partnership with Top-5 Leagues',
            ],
            done: false,
        },
        {
            id: 3,
            date: '2022 Q2',
            points: [
                'Advanced gameplay introduced',
                'Local competitions launched',
                'DAO governance adopted',
            ],
            done: false,
        },
    ];

    return (
        <section className="roadmap">
            <div className="roadmap__wrapper">
                <h2 className="roadmap__title">Development Roadmap</h2>
                <div className="roadmap__road">
                    {dataList.map((item) => (
                        <RoadmapPoint key={item.id} item={item} />
                    ))}
                </div>
                <AnimationImage
                    className={'roadmap__flag'}
                    heightFrom={1000}
                    heightTo={-300}
                    loop={true}
                    animationData={animationFlagData}
                    animationImages={[]}
                    isNeedScrollListener={false}
                />
            </div>
            <img
                className="roadmap__bottom"
                src={footer}
                alt="bottom texture"
            />
        </section>
    );
};
