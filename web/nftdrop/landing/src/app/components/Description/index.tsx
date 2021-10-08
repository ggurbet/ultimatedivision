// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DescriptionAbout } from '@components/Description/DescriptionAbout';
import { DescriptionCards } from '@components/Description/DescriptionCards';
import { DescriptionPay } from '@components/Description/DescriptionPay';

import './index.scss';

export const Description = () => {
    return (
        <section className="description">
            <div className="description__wrapper">
                <div className="description__container"
                    data-aos="fade-right"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <DescriptionAbout />
                </div>
                <div className="description__container"
                    data-aos="fade-left"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <DescriptionCards />
                </div>
                <div className="description__container"
                    data-aos="fade-right"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <DescriptionPay />
                </div>
            </div>
        </section>
    );
};
