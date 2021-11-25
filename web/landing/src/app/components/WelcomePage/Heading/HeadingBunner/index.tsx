// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';

import ball from '@static/images/Heading/ball.png';
import logo from '@static/images/Heading/ntflGame.png';
import logoText from '@static/images/Heading/logoText.png';

import './index.scss';

export const HeadingBunner: React.FC<{ offsetTop: number }> = ({
    offsetTop
}) => (
    <div className="heading-container" onScroll={() => scroll}>
        <img
            className="heading-container__logo"
            src={logo}
            alt="logo" />
        <img
            className="heading-container__logo-text"
            src={logoText}
            alt="logo"
        />
        <img
            style={{ top: `${offsetTop * -0.3}px` }}
            src={ball}
            alt="ball"
            className="heading-container__ball"
        />
    </div>
);
