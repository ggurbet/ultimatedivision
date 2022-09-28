// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useLocation } from 'react-router-dom';

import { Authors } from '@components/home/Authors';
import { FootballGame } from '@components/home/FootballGame';
import { Footer } from '@components/home/Footer';
import { Projects } from '@components/home/Projects';
import { Roadmap } from '@components/home/Roadmap';
import { VideoGame } from '@components/home/VideoGame';
import { GameInfo } from '@components/home/GameInfo';
import Navbar from '@components/home/HomeNavbar';

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
            <Projects />
            <Authors />
            <Footer />
        </>
    );
};

export default Home;
