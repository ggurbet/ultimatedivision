/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */
import { useSelector } from 'react-redux';

import { PlayingFormation } from './PlayingAreaFormation';

import { RootState } from '../../../store';

import './index.scss';

export const FootballFieldPlayingArea: React.FC = () => {
    const formation = useSelector((state: RootState) => state.fieldReducer.options.formation);
    const cardData = useSelector((state: RootState) => state.fieldReducer);

    return (
        <div className="football-field-playing-area" id="playingArea">
            <PlayingFormation props={cardData} formation={formation} />
        </div>
    );
};
