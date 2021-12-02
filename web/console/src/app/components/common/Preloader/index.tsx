// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import ball from '@static/img/Preloader/ball.svg';
import ultimateLogo from '@static/img/Preloader/ultimatedivision-logo.svg';

import './index.scss';

const Preloader: React.FC = () =>
    <div className="preloader">
        <img
            className="preloader__logo"
            src={ultimateLogo}
            alt="UltimateDivision logo"
        />
        <div className="preloader__load">
            <div className="preloader__load-ball-background">
                <img
                    className="preloader__load-ball"
                    src={ball}
                    alt="Ball"
                />
            </div>
            <div className="preloader__load-text">
                <span>L</span>
                <span>O</span>
                <span>A</span>
                <span>D</span>
                <span>I</span>
                <span>N</span>
                <span>G</span>
            </div>
        </div>
    </div>;
export default Preloader;
