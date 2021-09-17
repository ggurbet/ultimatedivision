
// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Roadmap } from '@components/WelcomePage/Roadmap';
import { Projects } from '@components/WelcomePage/Projects';
import { Authors } from '@components/WelcomePage/Authors';
import { Footer } from '@components/WelcomePage/Footer';
import { Description } from '@components/WelcomePage/Description';
import { LaunchRoadmap } from '@components/WelcomePage/LaunchRoadmap';

const Main: React.FC = () => {
    return (
        <>
            <Description />
            <LaunchRoadmap />
            <Roadmap />
            <Projects />
            <Authors />
            <Footer />
        </>
    );
};

export default Main;