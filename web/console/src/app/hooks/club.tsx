// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { userCards } from '../store/actions/cards';

export const useClub = () => {
    const dispatch = useDispatch();

    /** Calls method get from  ClubClient */
    async function getCards() {
        await dispatch(userCards());
    };

    useEffect(() => {
        getCards();
    }, []);

    return useSelector((state: RootState) => state.cardsReducer.club);
};