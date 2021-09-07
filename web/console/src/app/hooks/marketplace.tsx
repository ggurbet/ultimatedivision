// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { marketplaceCards } from '../store/actions/cards';

export const useMarketplace = () => {
    const dispatch = useDispatch();

    /** Calls method get from  ClubClient */
    async function getCards() {
        await dispatch(marketplaceCards());
    };

    useEffect(() => {
        getCards();
    }, []);

    return useSelector((state: RootState) => state.cardsReducer.marketplace);
};
