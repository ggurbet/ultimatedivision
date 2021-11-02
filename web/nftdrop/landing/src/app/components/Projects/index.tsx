// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import affinidi from '@static/images/projects/affinidi.svg';
import elixir from '@static/images/projects/elixir.svg';
import bloom from '@static/images/projects/bloom.svg';
import consensys from '@static/images/projects/consensys.svg';
import near from '@static/images/projects/near.svg';
import nem from '@static/images/projects/nem.svg';
import storj from '@static/images/projects/storj.svg';
import trustana from '@static/images/projects/trustana.svg';

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
                    {logos.map((logo, index) => (
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

                    ))}
                </div>
            </div >
        </section >
    );
};
