const roadmapData = [
    {
        id: '1',
        date: '01.05.2021',
        title: 'Whitepaper and project preparation',
        points: [],
        done: true,
    },
    {
        id: '2',
        date: '01.07.2021',
        title: 'MVP launch',
        points: [
            'Smart contracts',
            'NFT assets',
            'Market place',
            'Player Chests - in progress',
            'Player Cards',
        ],
        done: false,
    },
    {
        id: '3',
        date: '01.09.2021',
        title: 'Football Clubs',
        points: [
            'FC Managemen',
            'Team building',
            'Strategies',
            'Player-to-player contracts',
        ],
        done: false,
    },
    {
        id: '4',
        date: '01.10.2021',
        title: 'Gameplay and Leagues',
        points: [
            'P2P gameplay',
            'Division placement and progression',
            'Rewards'
        ],
        done: false,
    },
    {
        id: '5',
        date: '01.12.2021',
        title: 'Club Roles and more P2E',
        points: [
            'Managers',
            `In-game coach and other jobs to manage
            clubs based on smart contracts between players`,
        ],
        done: false,
    },
];

/* eslint-disable */

export const roadmapReducer = (roadmapState = roadmapData, action) => {
    return roadmapState;
};
