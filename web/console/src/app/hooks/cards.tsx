// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect} from 'react';
import { useDispatch } from 'react-redux';

export const useCards = (thunk: any) => {

    const dispatch = useDispatch();

    /** Calls method get from  ClubClient */
    async function getDataFromApi() {
        await dispatch(thunk());
    };

    useEffect(() => {
        getDataFromApi();
    }, []);

};
