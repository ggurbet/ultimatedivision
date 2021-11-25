// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import './index.scss';

export const HeadingModal: React.FC<{ handleCloseModal: any }> = ({
    handleCloseModal
}) => (
    <div className="heading-modal">
        <div className="heading-modal__window">
            <div
                className="heading-modal__close"
                onClick={handleCloseModal}
            >
                <p
                    className="heading-modal__close-text"
                >
                    &#215;
                </p>
            </div>
            <div className="heading-modal__description">
                <p>
                    LEAVE YOUR CONTACT DETAILS TO JOIN OUR BETA
                </p>
            </div>
            <div>
                <input
                    placeholder="Email"
                    className="heading-modal__send"
                />
            </div>
            <div
                className="heading-modal__wrapper"
                onClick={handleCloseModal}
            >
                <p className="heading-modal__confirm">
                    PLAY
                </p>
            </div>
        </div>
    </div>
);
