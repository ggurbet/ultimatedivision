// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import './index.scss';

export const HeadingButton: React.FC<{ handleShowModal: any }> = ({
    handleShowModal
}) => {
    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    });

    return (
        <div className="heading-join-wrapper" data-aos="fade-left">
            <p
                className="heading-join-wrapper__confirm"
                onClick={() => handleShowModal()}
            >
                PLAY
            </p>
        </div>
    );
};
