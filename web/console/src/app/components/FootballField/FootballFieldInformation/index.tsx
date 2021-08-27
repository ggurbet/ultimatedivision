// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { FootballFieldInformationLine } from '@/app/types/footballField';
import { FootballFieldInformationCaptain } from
    '@components/FootballField/FootballFieldInformation/FootballFieldInformationCaptain';
import { FootballFieldInformationFormation } from
    '@components/FootballField/FootballFieldInformation/FootballFieldInformationFormation';
import { FootballFieldInformationTactic } from
    '@components/FootballField/FootballFieldInformation/FootballFieldInformationTactic';

import './index.scss';

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
