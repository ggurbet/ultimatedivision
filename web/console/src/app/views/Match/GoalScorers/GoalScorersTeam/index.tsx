// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';

import { Goal } from '@/matches';
import { SquadCard } from '@/club';
import { RootState } from '@/app/store';

export const GoalScorersTeam: React.FC<{ goals: Goal[] }> = ({ goals }) => {
    const [cardId, setCardId] = useState<number | null>(null);
    /** COUNTER is counter that describes index number of each scored goal. */
    const COUNTER: number = 1;

    return (
        <>
            {goals.map((goal: Goal, index: number) =>
                <div
                    className="player"
                    key={index}
                    onMouseLeave={() => setCardId(null)}
                    onMouseEnter={() => setCardId(index)}
                >
                    {cardId === index &&
                        <PlayerCard
                            className="player__card"
                            id={goal.card.id}
                        />
                    }
                    <div className="player__wrapper">
                        {
                            <PlayerCard
                                className="player__logo"
                                id={goal.card.id}
                            />
                        }
                    </div>
                    <span className="player__name">{goal.card.playerName}</span>
                    <span className="player__goal-time">{goal.minute}</span>
                    <div className="player__goals">{index + COUNTER}</div>
                </div>
            )}
        </>
    );
};
