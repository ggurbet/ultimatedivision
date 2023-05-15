// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { Lot } from '@/marketplace';
import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';

const DEFAULT_TIME_VALUE = 0;
const HOURS = 24;
const MINUTES_AND_SECONDS = 60;
const INTERVAL = 1000;
const CONVERT_MILLISECONDS = 1000;
const TIME_TO_CHECK_IS_TIME_ENDED = 7;

export const MarketplaceTimer: React.FC<{
    lot: Lot;
    setIsEndTime: React.Dispatch<React.SetStateAction<boolean>>;
    isEndTime: boolean;
    className?: string;
}>
    = ({ lot, setIsEndTime, isEndTime, className }) => {
        const [hours, setHours] = useState(DEFAULT_TIME_VALUE);
        const [minutes, setMinutes] = useState(DEFAULT_TIME_VALUE);
        const [seconds, setSeconds] = useState(DEFAULT_TIME_VALUE);

        const marketplaceClient = new MarketplaceClient();
        const marketplaceService = new Marketplaces(marketplaceClient);

        const getTime = async(deadline: string) => {
            const time = Date.parse(deadline) - Date.now();

            const hours = Math.floor(time / (CONVERT_MILLISECONDS * MINUTES_AND_SECONDS * MINUTES_AND_SECONDS) % HOURS);
            const minutes = Math.floor(time / CONVERT_MILLISECONDS / MINUTES_AND_SECONDS % MINUTES_AND_SECONDS);
            const seconds = Math.floor(time / CONVERT_MILLISECONDS % MINUTES_AND_SECONDS);

            if (hours <= DEFAULT_TIME_VALUE
                && minutes <= DEFAULT_TIME_VALUE
                && seconds < TIME_TO_CHECK_IS_TIME_ENDED
                && !isEndTime) {
                const endTime = await marketplaceService.endTime(lot.cardId);

                setIsEndTime(endTime);
            }

            setHours(hours);
            setMinutes(minutes);
            setSeconds(seconds);
        };

        useEffect(() => {
            const interval = setInterval(() => getTime(lot.endTime), INTERVAL);

            return () => clearInterval(interval);
        }, [lot]);

        return (
            <div className={`${className ? className : ''}`}>
                {hours} : {minutes} : {seconds}
            </div>
        );
    };
