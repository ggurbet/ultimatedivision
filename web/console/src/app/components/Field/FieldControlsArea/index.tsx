// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FieldControl } from '@/app/components/Field/FieldControlsArea/FieldControl';

import { setCaptain, setFormation, setTactic } from '@/app/store/actions/clubs';
import { Control } from '@/app/types/club';

import './index.scss';

export const FieldControlsArea: React.FC = () => {
    const CONTROLS_FIELDS = [
        new Control('0', 'formation', setFormation, [
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
        ]),
        new Control('1', 'tactics', setTactic, [
            'attack',
            'defence',
            'balanced',
        ]),
        new Control('2', 'captain', setCaptain, [
            'Captain 1',
            'Captain 2',
            'Captain 3',
        ]),
    ];

    return (
        <div className="field-controls-area">
            <h2 className="field-controls-area__title">information</h2>
            {CONTROLS_FIELDS.map((item, index) =>
                <FieldControl key={index} props={item} />
            )}
        </div>
    );
};
