// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MatchScore } from './MatchScore';
import { PlayingField } from './PlayingField';
import { GoalScorers } from './GoalScorers';

import './index.scss';

const Match: React.FC = () =>
    <div className="match">
        <div className="wrapper">
            <MatchScore />
            <GoalScorers />
            <PlayingField />
        </div>
    </div>;
export default Match;
