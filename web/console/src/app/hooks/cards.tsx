// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';
import { SetStateAction, useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../store';

export const useCards = () => {
    const cardService = useSelector((state: RootState) => state.cardsReducer.cardService);

    type Data = {
        data: null | Card[];
        isLoading: boolean;
    };

    const [data, handleData] = useState<SetStateAction<Data>>({ data: null, isLoading: true });

    /** Calls method get from  ClubClient */
    async function getDataFromApi() {
        const cards = await cardService.get();

        handleData({
            // @ts-ignore
            data: cards,
            isLoading: false,
        });
    };

    useEffect(() => {
        getDataFromApi();
    }, []);

    return data;
};
