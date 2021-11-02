// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DescriptionCards } from '@components/Description/DescriptionCards';
import { DescriptionPay } from '@components/Description/DescriptionPay';

import './index.scss';

export const Description = () => {

    return (
        <section className="description">
            <div className="description__wrapper">
                <div className="description__container">
                    <DescriptionCards />
                </div>
                <div className="description__container">
                    <DescriptionPay />
                </div>
            </div>
        </section>
    );
};
