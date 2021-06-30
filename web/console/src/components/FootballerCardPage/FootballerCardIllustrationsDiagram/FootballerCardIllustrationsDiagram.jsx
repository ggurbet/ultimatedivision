// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import './FootballerCardIllustrationsDiagram.scss';
import { Doughnut } from 'react-chartjs-2';
/*eslint-disable*/
export const FootballerCardIllustrationsDiagram = ({
    name,
    min,
    max,
    value,
}) => {

    const percent = (Math.round((value - min) / max * 100))
    return (
        <div className="footballer-card-illustrations-diagram">
            <Doughnut
                data={{
                    datasets: [{
                        data: [percent, (100 - percent)],
                        backgroundColor: ['#3CCF5D', '#5E5EAA'],
                        borderColor: [
                            'transparent'
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
                        }
                    }
                }}
            />
            <div className="footballer-card-illustrations-diagram__values-area">
                <span className="footballer-card-illustrations-diagram__min">{min}</span>
                <span className="footballer-card-illustrations-diagram__value">{value}</span>
                <span className="footballer-card-illustrations-diagram__max">{max}</span>
            </div>
        </div>
    )
}
