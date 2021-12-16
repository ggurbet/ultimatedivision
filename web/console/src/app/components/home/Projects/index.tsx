// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import affinidi from '@static/img/gameLanding/projects/affinidi.svg';
import elixir from '@static/img/gameLanding/projects/elixir.svg';
import bloom from '@static/img/gameLanding/projects/bloom.svg';
import consensys from '@static/img/gameLanding/projects/consensys.svg';
import near from '@static/img/gameLanding/projects/near.svg';
import nem from '@static/img/gameLanding/projects/nem.svg';
import storj from '@static/img/gameLanding/projects/storj.svg';
import trustana from '@static/img/gameLanding/projects/trustana.svg';

import './index.scss';

export const Projects: React.FC = () => {
    const logos = [
        consensys,
        nem,
        storj,
        near,
        elixir,
        affinidi,
        trustana,
        bloom,
    ];

    return (
        <section className="projects">
            <div className="projects__wrapper">
                <h2 className="projects__title">
                    The game was created by a team involved in the development
                    of well-known crypto projects
                </h2>
                <div className="projects__area">
                    {logos.map((logo, index) =>
                        <div
                            key={index}
                            className="projects__area__item"
                        >
                            <img
                                className="projects__area__item__logo"
                                key={index}
                                src={logo}
                                alt="logo"
                            />
                        </div>

                    )}
                </div>
            </div >
        </section >
    );
};
