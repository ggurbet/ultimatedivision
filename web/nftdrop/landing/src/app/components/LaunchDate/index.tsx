// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { Modal } from './Modal';

import ball from '@static/images/launchDate/ball.webp';
import ballMobile from '@static/images/launchDate/ballMobile.webp';
import webkitBall from '@static/images/launchDate/ball.png';
import webkitBallMobile from '@static/images/launchDate/ballMobile.png';

import './index.scss';

export const LaunchDate: React.FC = () => {

    const [isShowModal, setIsShowModal] = useState(false);

    const handleModal = () => setIsShowModal(prev => !prev);

    return (
        <>
            <section className="launch-date">
                <div className="launch-date__wrapper">
                    <picture>
                        <source media="(max-width: 800px)" srcSet={ballMobile} type="image/webp" />
                        <source media="(min-width: 800px)" srcSet={ball} type="image/webp" />
                        <source media="(max-width: 800px)" srcSet={webkitBallMobile} />
                        <img
                            className="launch-date__ball"
                            src={webkitBall}
                            alt="ultimate division ball"
                            loading="lazy"
                        />
                    </picture>
                    <div className="launch-date__information">
                        <p
                            className="launch-date__information__subtitle"
                        >
                            Launch Date
                        </p>
                        <h1
                            className="launch-date__information__title"
                        >
                            1 November 20:00
                        </h1>
                        <div
                            data-aos="fade-left"
                            data-aos-delay={600}
                            className="launch-date__information__remind"
                            onClick={handleModal}
                        >
                            Remind Me
                        </div>
                    </div>
                </div>
            </section>
            {isShowModal && <Modal handleModal={handleModal} />}
        </>
    );
};