// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import './FootballerCardIllustrationsDiagram.scss';
import { Doughnut } from 'react-chartjs-2';
import { Diagram } from '../../../types/fotballerCard';

export const FootballerCardIllustrationsDiagram: React.FC<{ props: Diagram }> = ({ props }) => {
    const { name, min, max, value } = props;
    const FULL_VALUE_STATISTIC_SCALE = 100;
    /** percent value of player stat */
    const statsValue = Math.round((value - min) / max * FULL_VALUE_STATISTIC_SCALE);

    return (
        <div className="footballer-card-illustrations-diagram">
            <Doughnut
                type={Doughnut}
                data={{
                    datasets: [{
                        data: [statsValue, FULL_VALUE_STATISTIC_SCALE - statsValue],
                        backgroundColor: ['#3CCF5D', '#5E5EAA'],
                        borderColor: [
                            'transparent',
                        ],
                        cutout: '80%',
                        rotation: 270,
                        circumference: 180,
                        maintainAspectRatio: true,
                    }],
                }}
                options={{
                    plugins: {
                        title: {
                            display: true,
                            text: name.toUpperCase(),
                            color: 'white',
                        },
                    },
                }}
            />
            <div className="footballer-card-illustrations-diagram__values-area">
                <span className="footballer-card-illustrations-diagram__min">{min}</span>
                <span className="footballer-card-illustrations-diagram__value">{value}</span>
                <span className="footballer-card-illustrations-diagram__max">{max}</span>
            </div>
        </div>
    );
};
