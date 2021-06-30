/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import { useSelector } from 'react-redux';

import { FootballerCardIllustrations } from
    '../FootballerCardIllustrations/FootballerCardIllustrations';
import { FootballerCardPrice } from
    '../FootballerCardPrice/FootballerCardPrice';
import { FootballerCardStatsArea } from
    '../FootballerCardStatsArea/FootballerCardStatsArea';
import { FootballerCardInformation } from
    '../FootballerCardInformation/FootballerCardInformation';

import './FootballerCard.scss';

export const FootballerCard = () => {
    const cardData = useSelector(
        state => state.footballerCard[0].overalInfo[0]
    );

    return (
        <div className="footballer-card">
            <div className="footballer-card__wrapper">
                <div className="footballer-card__name-wrapper">
                    <h1 className="footballer-card__name">
                        {cardData.value}
                    </h1>
                </div>
                <FootballerCardIllustrations />
                <div className="footballer-card__stats-area">
                    <FootballerCardPrice />
                    <FootballerCardStatsArea />
                    <FootballerCardInformation />
                </div>
            </div>
        </div>
    );
};
