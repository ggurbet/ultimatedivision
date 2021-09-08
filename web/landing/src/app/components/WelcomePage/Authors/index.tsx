// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { AuthorsTitle } from '@components/WelcomePage/Authors/AuthorsTitle';
import { AuthorsTeam } from '@components/WelcomePage/Authors/AuthorsTeam';

import './index.scss';

export const Authors: React.FC = () => (
    <section className="authors">
        <div className="authors__wrapper">
            <AuthorsTitle />
            <AuthorsTeam />
        </div>
    </section>
);
