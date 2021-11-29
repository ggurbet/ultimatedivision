// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Goal } from '@/matches';

import player from '@/app/static/img/match/player.svg';

export const GoalScorersTeam: React.FC<{ goals: Goal[] }> = ({ goals }) => {
    /** COUNTER is counter that describes index number of each scored goal. */
    const COUNTER: number = 1;

    return <>
        {
            goals.map((goal: Goal, index: number) =>
                <div className="player" key={index}>
                    <img
                        src={player}
                        alt="player avatar"
                    ></img>
                    <span className="player__name">{goal.card}</span>
                    <span className="player__goal-time">
                        {goal.minute}
                    </span>
                    <div className="player__goals">{index + COUNTER}</div>
                </div>
            )}
    </>;
};
