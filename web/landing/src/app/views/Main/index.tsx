// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Capabilities } from '@components/WelcomePage/Capabilities';
import { Roadmap } from '@components/WelcomePage/Roadmap';
import { Projects } from '@components/WelcomePage/Projects';
import { Authors } from '@components/WelcomePage/Authors';
import { Footer } from '@components/WelcomePage/Footer';
import { LaunchRoadmap } from '@components/WelcomePage/LaunchRoadmap';

const Main: React.FC = () => {
    return (
        <>
            <Capabilities />
            <LaunchRoadmap />
            <Roadmap />
            <Projects />
            <Authors />
            <Footer />
        </>
    );
};

export default Main;
