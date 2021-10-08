// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { FieldControl } from '@/app/types/club';
import { FootballFieldControl } from
    '@/app/components/FootballField/FootballFieldControlsArea/FootballFieldControl';
import { setCaptain, setFormation, setTactic } from '@/app/store/actions/club';

import './index.scss';

export const FootballFieldControlsArea: React.FC = () => {
    const CONTROLS_FIELDS = [
        new FieldControl('0', 'formation', setFormation,
            [
                '4-4-2',
                '4-2-4',
                '4-2-2-2',
                '4-3-1-2',
                '4-3-3',
                '4-2-3-1',
                '4-3-2-1',
                '4-1-3-2',
                '5-3-2',
                '4-5-2',
            ]
        ),
        new FieldControl('1', 'tactics', setTactic, ['attack', 'defence', 'balanced']),
        new FieldControl('2', 'captain', setCaptain, ['Captain 1', 'Captain 2', 'Captain 3']),
    ];

    return (
        <div className="football-field-controls-area">
            <h2 className="football-field-controls-area__title">
                information
            </h2>
            {CONTROLS_FIELDS.map((item, index) =>
                <FootballFieldControl
                    key={index}
                    props={item}
                />
            )}
        </div>
    );
};
