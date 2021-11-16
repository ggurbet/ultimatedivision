// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import FootballField from '@/app/static/img/match/playing-field.svg';

import './index.scss';

export const PlayingField: React.FC = () =>
    <div className="match__playing-field">
        <img src={FootballField} alt="Playing football field"></img>
        <div className="match__playing-field__gradient"></div>
    </div>;

