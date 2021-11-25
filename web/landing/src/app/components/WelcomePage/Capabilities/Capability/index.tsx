// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import './index.scss';

export const Capability: React.FC<{
    title: string,
    description: string,
    icon: string,
    id: number
}> = ({
    title,
    description,
    icon,
    id
}) => {
    useEffect(() => {
        Aos.init({
            duration: 500,
        });
    });

    return (
        <div
            className="capability"
            data-aos="fade-left-capability"
            data-aos-delay={200 * (id - 1)}
        >
            <div className="capability__icon">
                <img
                    className="capability__image"
                    alt="capability"
                    src={icon}
                />
            </div>
            <div className="capability__information">
                <h3
                    className="capability__title">
                    {title.toUpperCase()}
                </h3>
                <p className="capability__description">
                    {description}
                </p>
            </div>
        </div>
    );
};
