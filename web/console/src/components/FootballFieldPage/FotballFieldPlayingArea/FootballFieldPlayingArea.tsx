/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';
import './FootballFieldPlayingArea.scss';

import { RootState } from '../../../store';

import { PlayingFormation } from './PlayingAreaFormation/PlayingFormation';

export const FootballFieldPlayingArea: React.FC = () => {
    const formation = useSelector((state: RootState) => state.fieldReducer.options.formation);
    const cardData = useSelector((state: RootState) => state.fieldReducer);

    return (
        <div className="football-field-playing-area" id="playingArea">
            <PlayingFormation props={cardData} formation={formation} />
        </div>
    )
}
