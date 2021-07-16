// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../store';

import './index.scss';

import { Radar } from 'react-chartjs-2';

export const FootballerCardIllustrationsRadar: React.FC = () => {
    const FIRST_CARD_INDEX = 0;
    const stats = useSelector((state: RootState) =>
        state.cardReducer[FIRST_CARD_INDEX].stats);

    return (
        <div className="footballer-card-illustrations-radar">
            <Radar
                type={Radar}
                data={{
                    labels: ['TAC', 'PHY', 'TEC', 'OFF', 'DEF', 'GK'],
                    datasets: [{
                        backgroundColor: '#66FF8866',
                        data: stats.map(item => item.average),
                    }],
                }}
                options={{
                    animations: {
                        tension: {
                            duration: 1000,
                            easing: 'linear',
                            from: 0,
                            to: 0,
                        },
                    },
                    plugins: {
                        legend: {
                            display: false,
                        },
                        interaction: {
                            display: false,
                        },
                    },
                    scale: {
                        ticks: {
                            maxTicksLimit: 2,
                        },
                    },
                    scales: {
                        r: {
                            ticks: {
                                display: false,
                            },
                            pointLabels: {
                                color: '#afafaf',
                            },
                            angleLines: {
                                type: 'dashed',
                                color: '#515180',
                            },
                            grid: {
                                color: '#515180',
                            },
                        },
                    },
                }
                }
            />
        </div>
    );
};
