// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Roadmap } from '@components/WelcomePage/Roadmap';
import { Projects } from '@components/WelcomePage/Projects';
import { Footer } from '@components/WelcomePage/Footer';
import { LaunchRoadmap } from '@components/WelcomePage/LaunchRoadmap';
import { Navbar } from '@components/WelcomePage/NavBar';
import { Home } from '@components/WelcomePage/Home';
import { LaunchDate } from '@components/WelcomePage/LaunchDate';
import { Description } from '@components/WelcomePage/Description';
import { Metaverse } from '@components/WelcomePage/Metaverse';
import { Authors } from '@components/WelcomePage/Authors';

export const App = () => (
    <main className="main">
        <Navbar />
        <Home />
        <LaunchDate />
        <Metaverse />
        <Description />
        <LaunchRoadmap />
        <Roadmap />
        <Projects />
        <Authors />
        <Footer />
    </main>
);
