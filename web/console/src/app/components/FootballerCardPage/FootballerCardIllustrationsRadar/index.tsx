// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/app/types/fotballerCard';

import { Radar } from 'react-chartjs-2';

import './index.scss';

export const FootballerCardIllustrationsRadar: React.FC<{ card: Card }> = ({ card }) => {
    const stats = card.stats;
    const labels = ['TAC', 'PHY', 'TEC', 'OFF', 'DEF', 'GK'];

    return (
        <div className="footballer-card-illustrations-radar">
            <Radar
                type={Radar}
                data={{
                    labels: stats.map((item, index) =>
                        `${labels[index]} ${item.average}`),
                    datasets: [{
                        backgroundColor: '#66FF8866',
                        data: stats.map(item => item.average),
                    }],
                }}
                options={{
                    elements: {
                        line: {
                            borderWidth: 1,
                            borderColor: '#66ff8791',
                        },
                    },
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
                                borderDash: [7, 8],
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
