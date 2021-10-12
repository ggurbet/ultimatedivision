// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import webkitField from '@static/images/description/field.png';
import webkitFieldMobile from '@static/images/description/field-mobile.png';
import webkitFieldTablet from '@static/images/description/field-tablet.png';
import field from '@static/images/description/field.webp';
import fieldMobile from '@static/images/description/field-mobile.webp';
import fieldTablet from '@static/images/description/field-tablet.webp';

import './index.scss';

export const DescriptionAbout = () => {
    return (
        <div className="description-about" id="about">
            <picture>
                <source media="(max-width: 700px)" srcSet={fieldMobile} type="image/webp" />
                <source media="(max-width: 1100px)" srcSet={fieldTablet} type="image/webp" />
                <source media="(min-width: 1100px)" srcSet={field} type="image/webp" />
                <source media="(max-width: 700px)" srcSet={webkitFieldMobile} />
                <source media="(max-width: 1100px)" srcSet={webkitFieldTablet} />
                <img
                    className="description-about__field"
                    src={webkitField}
                    alt="field"
                    loading="lazy"
                />
            </picture>
            <div className="description-about__text-area">
                <h2 className="description-about__title">About the Game</h2>
                <p className="description-about__text">
                    Ultimate Division is a world football simulator.
                    UD players will own clubs, players and face each other
                    in weekly competitions to win cash prizes!
                    Other players can be hired as managers or coaches
                    for your Club!
                </p>
            </div>
        </div>
    );
};
