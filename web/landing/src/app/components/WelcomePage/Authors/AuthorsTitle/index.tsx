// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';

import Aos from 'aos';

import './index.scss';

export const AuthorsTitle: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    });

    return (
        <div className="authors-heading" data-aos="fade-left">
            <p className="authors-heading__information">
                {'PROJECT CREATED BY '}
                <a
                    href="https://chickenfish.games/"
                    target="_blank"
                    rel="noreferrer"
                    className="authors-heading__information-modificator"
                >
                    {'CHIKENFISH GAMES'}
                </a>
                {' AND '}
                <a
                    href="https://boostylabs.com/"
                    target="_blank"
                    rel="noreferrer"
                    className="authors-heading__information-modificator"
                >
                    {'BOOSTY LABS'}
                </a>
            </p>
        </div>
    );
};
