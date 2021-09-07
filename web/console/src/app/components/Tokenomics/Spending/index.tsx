// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Doughnut } from 'react-chartjs-2';

import './index.scss';

const Spending: React.FC = () => {
    const dataDoughnutStats = [
        {
            value: 'Play to earn',
            points: 4182,
            style: {
                'backgroundColor': '#A1B9CA',
            },
        },
        {
            value: 'Private Sale',
            points: 2010,
            style: {
                'backgroundColor': '#A1CAC5',
            },
        },
        {
            value: 'Staking Rewards',
            points: 500,
            style: {
                'backgroundColor': '#B3CAA1',
            },
        },
        {
            value: 'Public Sale',
            points: 300,
            style: {
                'backgroundColor': '#F9E1BF',
            },
        },
        {
            value: 'Core Team',
            points: 129,
            style: {
                'backgroundColor': '#F39B9B',
            },
        },
        {
            value: 'UD Fund',
            points: 3201,
            style: {
                'backgroundColor': '#CAA1C8',
            },
        },
        {
            value: 'Advisors',
            points: 542,
            style: {
                'backgroundColor': '#BDC0FF',
            },
        },
    ];

    return (
        <div className="spending">
            <h1 className="spending__title">
                Ultimate Division Tokenomics
            </h1>
            <p className="spending__description">
                UDT tokens will be unlocked within a 5-year schedule following the initial sale.
                Originally, 20% of tokens will be in circulation
            </p>
            <h2 className="spending__subtitle">
                UDT Spending
            </h2>
            <p className="spending__description">
                UDT can be spent in-game to obtain packs that contain new player cards.
                New player cards will have a high demand to form better squads.
                New cards will also have intrinsic collector value.
                <br /><br />
                UDT will also be the main circulating currency on the marketplace and will be used in smart
                contracts between players.
                <br /><br />
                Some cosmetic and consumable in-games items will also require a steady UDT flow for every player.
                <br /><br />
                Overall, UDT will be necessary in-game currency that facilitates player interaction and rewards
                winners and high-performing teams.
            </p>
            <div className="spending__diagrams">
                <div className="spending__diagrams__doughnut">
                    <Doughnut
                        type={Doughnut}
                        data={{
                            datasets: [{
                                data: dataDoughnutStats.map((dataDoughnutStat) => dataDoughnutStat.points),
                                label: dataDoughnutStats.map((dataDoughnutStat) => dataDoughnutStat.value),
                                backgroundColor: dataDoughnutStats.map((dataDoughnutStat) => dataDoughnutStat.style.backgroundColor),
                                borderColor: [
                                    'transparent',
                                ],
                                cutout: '80%',
                                rotation: 0,
                                esponsive: true,
                                maintainAspectRatio: true,
                                hoverOffset: 16,
                            }],
                            labels: dataDoughnutStats.map((dataDoughnutStat) => dataDoughnutStat.value),
                        }
                        }
                        options={{
                            layout: {
                                padding: '10',
                            },
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
                                legend: {
                                    position: 'right',
                                    labels: {
                                        color: 'white',
                                        font: {
                                            size: 16,
                                        },
                                        usePointStyle: true,
                                        padding: 30,
                                    },
                                },
                            },
                        }}
                    />
                </div>
            </div>
        </div>
    );
};

export default Spending;
