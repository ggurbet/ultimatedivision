// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Doughnut } from 'react-chartjs-2';

import { CardField } from '@/card';

import './index.scss';

export const FootballerCardIllustrationsDiagram: React.FC<{
    props: CardField;
}> = ({ props }) => {
    const { value, label } = props;
    const FULL_VALUE_STATISTIC_SCALE = 2000;
    /** percent value of player stat */

    return (
        <div className="footballer-card-illustrations-diagram">
            <Doughnut
                type={Doughnut}
                data={{
                    datasets: [
                        {
                            data: [value, FULL_VALUE_STATISTIC_SCALE - +value],
                            backgroundColor: ['#3CCF5D', '#5E5EAA'],
                            borderColor: ['transparent'],
                            cutout: '80%',
                            rotation: 270,
                            circumference: 180,
                            maintainAspectRatio: true,
                        },
                    ],
                }}
                options={{
                    plugins: {
                        title: {
                            display: true,
                            text: label.toUpperCase(),
                            color: 'white',
                        },
                    },
                }}
            />
            <div className="footballer-card-illustrations-diagram__values-area">
                <span className="footballer-card-illustrations-diagram__min">
                    {0}
                </span>
                <span className="footballer-card-illustrations-diagram__value">
                    {value}
                </span>
                <span className="footballer-card-illustrations-diagram__max">
                    {FULL_VALUE_STATISTIC_SCALE}
                </span>
            </div>
        </div>
    );
};
