// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DescriptionAbout } from '@components/WelcomePage/Description/DescriptionAbout';
import { DescriptionCards } from '@components/WelcomePage/Description/DescriptionCards';
import { DescriptionPay } from '@components/WelcomePage/Description/DescriptionPay';

import './index.scss';

export const Description = () => {
    return (
        <section className="description">
            <DescriptionAbout />
            <DescriptionCards />
            <DescriptionPay />
        </section>
    );
};
