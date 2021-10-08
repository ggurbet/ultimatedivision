// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import { dataAosLogoAnimation } from '@utils/dataAosLogoAnimation';
import { aosDelayLogoAnimation } from '@utils/aosDelayLogoAnimation';

import consensys from '@static/images/projects/consensys.svg';
import nem from '@static/images/projects/nem.svg';
import storj from '@static/images/projects/storj.svg';
import near from '@static/images/projects/near.svg';
import elixir from '@static/images/projects/elixir.svg';
import affinidi from '@static/images/projects/affinidi.svg';
import trustana from '@static/images/projects/trustana.svg';
import bloom from '@static/images/projects/bloom.svg';

import './index.scss';

export const Projects: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 1500
        });
    }, []);

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

                <h2 className="projects__title" data-aos="fade-left">
                    The game was created by a team involved in the development
                    of well-know crypto projects
                </h2>
                <div className="projects__area">
                    {logos.map((logo, index) => (
                        <div
                            key={index}
                            data-aos={
                                dataAosLogoAnimation(index)
                            }
                            data-aos-delay={
                                aosDelayLogoAnimation(index)
                            }
                            data-aos-duration={500}
                            data-aos-easing="ease-in-out-cubic"
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
            </div>
        </section>
    );
};
