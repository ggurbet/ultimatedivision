// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { BoostyLogo, ChickenfishLogo } from '@static/img/gameLanding/authorsPage/authors';

import './index.scss';

export const Authors: React.FC = () =>
    <section className="authors">
        <div className="authors__wrapper">
            <span className="authors__title">
                    Created by
            </span>
            <div className="authors__created-by">
                <a
                    className="authors__chikenfish"
                    href="https://chickenfish.games/"
                    target="_blank" rel="noreferrer"
                >
                    <ChickenfishLogo />
                    <span className="authors__chikenfish__text">CHICKENFISH GAMES</span>
                </a>
                <a
                    className="authors__boostylabs"
                    href="https://boostylabs.com/"
                    target="_blank" rel="noreferrer"
                >
                    <BoostyLogo />
                    <span className="authors__boostylabs__text">BOOSTY LABS</span>
                </a>
            </div>
        </div>
    </section>;

