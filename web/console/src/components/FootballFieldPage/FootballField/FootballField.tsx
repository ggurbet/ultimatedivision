/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './FootballField.scss';

import { FootballFieldPlayingArea }
    from '../FotballFieldPlayingArea/FootballFieldPlayingArea';
import { FootballFieldInformation }
    from '../FootballFieldInformation/FootballFieldInformation';

export const FootballField = () => {
    return (
        <div className="football-field">
            <h1 className="football-field__title">Football Field</h1>
            <div className="football-field__wrapper">
                <FootballFieldPlayingArea />
                <FootballFieldInformation />
            </div>
        </div>
    )
}
