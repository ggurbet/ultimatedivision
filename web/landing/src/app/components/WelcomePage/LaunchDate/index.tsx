// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect } from 'react';
import Aos from 'aos';

import { Modal } from './Modal';

import ball from '@static/images/headingPage/ball.png';

import './index.scss';

export const LaunchDate: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 500,
        });
    });

    const [isShowModal, setIsShowModal] = useState(false);

    const handleModal = () => setIsShowModal(prev => !prev);

    return (
        <>
            <section className="launch-date">
                <div className="launch-date__wrapper">
                    <img
                        className="launch-date__ball"
                        src={ball}
                        alt="ultimate division ball"
                    />
                    <div className="launch-date__information">
                        <p
                            data-aos="fade-left"
                            data-aos-delay={200}
                            className="launch-date__information__subtitle"
                        >
                            Launch Date
                        </p>
                        <h1
                            data-aos-delay={400}
                            data-aos="fade-left"
                            className="launch-date__information__title"
                        >
                            20 September 20:00
                        </h1>
                        <a
                            data-aos="fade-left"
                            data-aos-delay={600}
                            className="launch-date__information__remind"
                            onClick={handleModal}
                        >
                            Remind Me
                        </a>
                    </div>
                </div>
            </section>
            {isShowModal && <Modal handleModal={handleModal} />}
        </>
    );
};