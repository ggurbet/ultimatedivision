/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { useSelector } from 'react-redux';
import './FootballFieldPlayingArea.scss';

import { RootState } from '../../../store';

import { PlayingFormation_442 }
    from './PlayingAreaFormations/PlayingFormation_442/PlayingFormation_442';
import { PlayingFormation_424 }
    from './PlayingAreaFormations/PlayingFormation_424/PlayingFormation_424';
import { PlayingFormation_433 }
    from './PlayingAreaFormations/PlayingFormation_433/PlayingFormation_433';

export const FootballFieldPlayingArea: React.FC = () => {
    const formation = useSelector((state: RootState) => state.fieldReducer.options.formation);
    const cardData = useSelector((state: RootState) => state.fieldReducer);

    const formationSelection = () => {
        switch (formation) {
            case '4-4-2':
                return <PlayingFormation_442 props={cardData} />
            case '4-2-4':
                return <PlayingFormation_424 props={cardData} />;
            case '4-3-3':
                return <PlayingFormation_433 props={cardData} />
        };
    }

    return (
        <div className="football-field-playing-area">
            {formationSelection()}
        </div>
    )
}
