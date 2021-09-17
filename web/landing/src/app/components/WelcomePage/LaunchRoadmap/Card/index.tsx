// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Doughnut } from 'react-chartjs-2';

import box from '@static/images/launchRoadmap/box1.svg';

import './index.scss';

export const Card: React.FC<{
    card: {
        title: string,
        subTitle: string,
        description: string,
        value: number,
    }
}> = ({ card }) => {
    const LOWER_BREAKPOINT = card.value;
    const UPPER_BREAKPOINT = 100 - LOWER_BREAKPOINT;
    const LOWER_BREAKPOINT_COLOR = '#37FB63';
    const UPPER_BREAKPOINT_COLOR = '#323c92';
    return (
        <div className="card">
            <h1 className="card__title">
                {card.title}
            </h1>
            <div className="card__diagram">
                <Doughnut data={{
                    labels: [],
                    datasets: [{
                        data: [LOWER_BREAKPOINT, UPPER_BREAKPOINT],
                        backgroundColor: [
                            LOWER_BREAKPOINT_COLOR,
                            UPPER_BREAKPOINT_COLOR,
                        ],
                        borderColor: 'transparent',
                    }],
                    hoverOffset: 16,
                }}
                    options={{
                        plugins: {
                            tooltip: {
                                backgroundColor: 'transparent',
                                displayColors: false,
                                padding: {
                                    left: 135,
                                    right: 355,
                                    top: 270,
                                    bottom: 280,
                                },
                            },
                        },
                    }}
                />
                <p className="card__diagram__value">
                    {card.value}%
                </p>
            </div>
            <p className="card__description">
                {card.description}
            </p>
            <div className="card__box">
                <img
                    className="card__box__present"
                    src={box}
                    alt="utlimate box"
                />
                <p className="card__box__subtitle">
                    {card.subTitle}
                </p>
            </div>
        </div>
    )
};
