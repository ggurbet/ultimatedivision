// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Authors } from '@components/home/Authors';
import { FootballGame } from '@components/home/FootballGame';
import { Footer } from '@components/home/Footer';
import { Projects } from '@components/home/Projects';
import { Roadmap } from '@components/home/Roadmap';

import './index.scss';

const Home: React.FC = () =>
    <>
        <FootballGame />
        <Roadmap />
        <Projects />
        <Authors />
        <Footer />
    </>;
export default Home;
