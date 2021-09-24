// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import { dataAosLogoAnimation } from '@utils/dataAosLogoAnimation';
import { aosDelayLogoAnimation } from '@utils/aosDelayLogoAnimation';

import consensys from '@static/images/Projects/consensys.png';
import nem from '@static/images/Projects/nem.png';
import storj from '@static/images/Projects/storj.png';
import near from '@static/images/Projects/near.png';
import elixir from '@static/images/Projects/elixir.png';
import affinidi from '@static/images/Projects/affinidi.png';
import trustana from '@static/images/Projects/trustana.png';
import bloom from '@static/images/Projects/bloom.png';

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
            <h2 className="projects__title" data-aos="fade-left">
                The game was created by a team involved in the development
                of well-know crypto projects
            </h2>
            <div className="projects__area">
                {logos.map((logo, index) => (
                    <div
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
        </section>
    );
};
