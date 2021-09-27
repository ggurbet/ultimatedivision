// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { BoostyLogo, ChickenfishLogo } from '@static/images/authorsPage/authors';
import './index.scss';

export const Authors: React.FC = () => (
    <section className="authors">
        <div className="authors__wrapper">
            <span className="authors__wrapper-title"  
                data-aos="fade-right"
                data-aos-duration="600"
                data-aos-easing="ease-in-out-cubic"
            >
                Created by
            </span>
            <div className="authors__wrapper-created-by"  
                data-aos="fade-right"
                data-aos-duration="600"
                data-aos-easing="ease-in-out-cubic"
            >
                <ChickenfishLogo />
                <BoostyLogo />
            </div>
            <div className="authors__wrapper-created-by__name"  
                data-aos="fade-right"
                data-aos-duration="600"
                data-aos-easing="ease-in-out-cubic"
            >
                <span className="chickenfish">CHICKENFISH GAMES</span>
                <span className="boosty-labs">BOOSTY LABS</span>
            </div>
        </div>
    </section>
);
