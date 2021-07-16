/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { useSelector } from 'react-redux';

import { FootballerCardStats }
    from '../FootballerCardStats';

import { RootState } from '../../../store';

import './index.scss';

export const FootballerCardStatsArea: React.FC = () => {
    const FIRST_CARD_INDEX = 0;
    const stats = useSelector((state: RootState) =>
        state.cardReducer[FIRST_CARD_INDEX].stats);

    return (
        <div className="footballer-card-stats">
            {stats.map((item, index) =>
                <FootballerCardStats
                    key={index}
                    props={item}
                />,
            )}
        </div>
    );
};
