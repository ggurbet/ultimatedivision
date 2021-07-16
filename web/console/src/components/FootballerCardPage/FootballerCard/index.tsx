/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import { FootballerCardIllustrations } from '../FootballerCardIllustrations';
import { FootballerCardPrice } from '../FootballerCardPrice';
import { FootballerCardStatsArea } from '../FootballerCardStatsArea';
import { FootballerCardInformation } from '../FootballerCardInformation';

import './index.scss';

export const FootballerCard: React.FC = () => {
    /** TODO: Route config with cards ID */
    const FIRST_CARD_INDEX = 0;
    const cardName = useSelector((state: RootState) =>
        state.cardReducer[FIRST_CARD_INDEX].overalInfo[FIRST_CARD_INDEX].value);

    return (
        <div className="footballer-card">
            <div className="footballer-card__wrapper">
                <div className="footballer-card__name-wrapper">
                    <h1 className="footballer-card__name">
                        {cardName}
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
