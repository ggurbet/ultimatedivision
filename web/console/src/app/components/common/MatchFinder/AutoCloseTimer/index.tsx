// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useDispatch } from 'react-redux';

import { useTimeCounter } from '@/app/hooks/useTimeCounter';
import { startSearchingMatch } from '@/app/store/actions/clubs';

export const AutoCloseTimer: React.FC = () => {
    const dispatch = useDispatch();
    const timeCounter = useTimeCounter();

    /** DELAY is time delay in milliseconds for closing MatchFinder component */
    const DELAY: number = 30000;
    useEffect(() => {
        setTimeout(() => {
            dispatch(startSearchingMatch(false));
        }, DELAY);
    }, []);

    return <div className="match-finder__timer">
        <span className="match-finder__timer__text">
            {timeCounter.minutes} : {timeCounter.seconds}
        </span>
    </div>;
};
