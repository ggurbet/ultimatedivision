// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';

import Aos from 'aos';

import { RoadmapCircle } from '@components/WelcomePage/Roadmap/RoadmapCircle';

import footballPlayer from '@static/images/roadmap/footballPlayer.png';
import ball from '@static/images/roadmap/footballPlayer.png';

import './index.scss';

export const Roadmap: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    }, []);

    const dataList = [
        {
            id: 1,
            date: '01.07.2021',
            title: 'Whitepaper and project preparation',
            points: [],
            done: true,
        },
        {
            id: 2,
            date: '01.09.2021',
            title: 'MVP launch',
            points: [
                'Smart contracts',
                'NFT assets',
                'Marketplace',
                'Player Boxes',
                'Player Cards',
            ],
            done: false,
        },
        {
            id: 3,
            date: '01.11.2021',
            title: 'Football Clubs',
            points: [
                'FC Management',
                'Squad building',
                'Strategies',
                'Player-to-player contracts',
            ],
            done: false,
        },
        {
            id: 4,
            date: '01.12.2021',
            title: 'Gameplay and Leagues',
            points: [
                'P2P gameplay',
                'Division placement and progression',
                'Weekly Rewards'
            ],
            done: false,
        },
        {
            id: 5,
            date: '01.02.2022',
            title: 'Club Roles and more P2E',
            points: [
                'Managers',
                'In-game coaches',
                'Smart-contract club management',
            ],
            done: false,
        },
    ];

    return (
        <section className="roadmap">
            <img
                className="roadmap__footbal-player-img"
                src={footballPlayer}
                alt="football player"
            />
            <img
                className="roadmap__ball-img"
                src={ball}
                alt="ball"
            />
            <div
                className="roadmap__road"
                data-aos="zoom-out-down">
                {dataList.map((item) => (
                    <RoadmapCircle
                        key={item.id} item={item} />
                ))}
            </div>
        </section>
    );
};
