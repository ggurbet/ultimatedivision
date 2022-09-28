// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useLocation } from 'react-router-dom';

import { FootballGame } from '@components/home/FootballGame';
import { Roadmap } from '@components/home/Roadmap';
import { Partnerships } from '@/app/components/home/Partnerships';
import { VideoGame } from '@components/home/VideoGame';
import { GameInfo } from '@components/home/GameInfo';
import Navbar from '@components/home/HomeNavbar';

import banner from '@static/img/gameLanding/banner.png';

import './index.scss';

const Home: React.FC = () => {
    /** Current path from hook */
    const location = useLocation();
    const currentPath = location.pathname;

    return (
        <>
            {currentPath === '/' && <Navbar />}
            <FootballGame />
            <GameInfo />
            <VideoGame/>
            <Roadmap />
            <Partnerships />
            <div className="home__banner">
                <img src={banner} className="home__banner__image" alt="banner" />
            </div>
        </>
    );
};

export default Home;
