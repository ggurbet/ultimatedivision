// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useTimeCounter } from '@/app/hooks/useTimeCounter';

export const AutoCloseTimer: React.FC = () => {
    const timeCounter = useTimeCounter();

    return <div className="match-finder__timer">
        <span className="match-finder__timer__text">
            {timeCounter.minutes} : {timeCounter.seconds}
        </span>
    </div>;
};
