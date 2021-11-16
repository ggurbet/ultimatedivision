// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MatchScore } from './MatchScore';
import { PlayingField } from './PlayingField';
import { PlayersScore } from "./PlayersScore";

import './index.scss';

const Match: React.FC = () =>
    <div className="match">
        <div className="wrapper">
            <MatchScore />
            <PlayersScore />
            <PlayingField />
        </div>
    </div>;
export default Match;
