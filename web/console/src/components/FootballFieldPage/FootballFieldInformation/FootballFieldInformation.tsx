/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import './FootballFieldInformation.scss';

import { FootballFieldInformationLine } from '../../../types/footballField';
import { FootballFieldInformationFormation }
    from './FootballFieldInformationFormation/FootballFieldInformationFormation';
import { FootballFieldInformationTactic }
    from './FootballFieldInformationTactic/FootballFieldInformationTactic';
import { FootballFieldInformationCaptain }
    from './FootballFieldInformationCaptain/FootballFieldInformationCaptain';

export const FootballFieldInformation: React.FC = () => {
    const INFORMATION_FIELDS = {
        formation: new FootballFieldInformationLine('0', 'formation', ['4-4-2', '4-2-4', '4-3-3']),
        tactics: new FootballFieldInformationLine('1', 'tactics', ['attack', 'defence', 'regular']),
        captain: new FootballFieldInformationLine('2', 'captain', ['4-4-2', '4-2-4', '4-3-3']),
    };

    return (
        <div className="football-field-information">
            <h2 className="football-field-information__title">
                information
            </h2>
            <FootballFieldInformationFormation
                props={INFORMATION_FIELDS.formation}
            />
            <FootballFieldInformationTactic
                props={INFORMATION_FIELDS.tactics}
            />
            <FootballFieldInformationCaptain
                props={INFORMATION_FIELDS.captain}
            />
        </div>
    );
};
