// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { RootState } from '@/app/store';

import './index.scss';

export const MatchScore: React.FC = () => {
    const { firstTeamGoalsCrored, secondTeamGoalsScored } = useSelector((state: RootState) => state.matchesReducer);

    return <div className="match__score">
        <div className="match__score__board">
            <div className="match__score__board__gradient"></div>
            <div className="match__score__board__timer">90:00</div>
            <div className="match__score__board__result">
                <div className="match__score__board-team-1">{firstTeamGoalsCrored}</div>
                <div className="match__score__board-dash">-</div>
                <div className="match__score__board-team-2">{secondTeamGoalsScored}</div>
            </div>
        </div>
    </div>;
};
