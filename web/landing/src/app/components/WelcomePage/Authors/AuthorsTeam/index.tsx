// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import Aos from 'aos';

import chikenfish from '@static/images/authorsPage/authors/chikenfish.png';
import boosty from '@static/images/authorsPage/authors/boosty.png';
import './index.scss';

export const AuthorsTeam: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 1000,
        });
    });

    return (
        <div className="authors-team">
            <a
                className="authors-team__link"
                href="https://boostylabs.com/"
                target="_blank"
                rel="noreferrer"
            >
                <img
                    className="authors-team__image"
                    src={boosty}
                    alt="boostylabs logo"
                />
            </a>
            <a
                className="authors-team__link"
                href="https://chickenfish.games/"
                target="_blank"
                rel="noreferrer"
            >
                <img
                    className="authors-team__image"
                    src={chikenfish}
                    alt="chikenfish logo"
                />
            </a>
        </div>
    );
};
