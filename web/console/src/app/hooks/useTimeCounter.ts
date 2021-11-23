// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

/** hook useTimeCounter uses for timing counter. */
export const useTimeCounter = () => {
    const COUNTER: number = 1;
    /** Variables describes default values for seconds and minutes hooks. */
    const DEFAULT_SECONDS_VALUE: number = 0;
    const DEFAULT_MINUTES_VALUE: number = 1;

    /** Describes upper seconds breakpoint for counting minutes. */
    const UPPER_SECONDS_BREAKPOINT: number = 59;
    /** Describes lower seconds breakpoint to change single-digit number to two-digit. */
    const LOWER_SECONDS_BREAKPOINT: number = 9;
    /** Describes lower minutes breakpoint to change single-digit number to two-digit. */
    const LOWER_MINUTES_BREAKPOINT: number = 9;

    const [seconds, setSeconds] = useState<number>(DEFAULT_SECONDS_VALUE);
    const [minutes, setMinutes] = useState<number>(DEFAULT_MINUTES_VALUE);

    const [timeCounter, setTimeCounter] = useState({
        minutes: '00',
        seconds: '00',
    });

    useEffect(() => {
        /** DELAY is time delay in milliseconds for resetting timer */
        const DELAY: number = 1000;

        const timer = setInterval(() => {
            setSeconds(seconds + COUNTER);

            if (seconds > UPPER_SECONDS_BREAKPOINT) {
                minutes > LOWER_MINUTES_BREAKPOINT ? setTimeCounter({ minutes: `${minutes}`, seconds: '00' }) :
                    setTimeCounter({ minutes: `0${minutes}`, seconds: '00' });

                setMinutes(minutes + COUNTER);
                setSeconds(COUNTER);

                return;
            };

            if (seconds <= LOWER_SECONDS_BREAKPOINT) {
                setTimeCounter({ ...timeCounter, seconds: `0${seconds}` });

                return;
            };

            setTimeCounter({ ...timeCounter, seconds: `${seconds}` });
        }, DELAY);

        return () => clearInterval(timer);
    }, [seconds, minutes, timeCounter]);

    return timeCounter;
};
