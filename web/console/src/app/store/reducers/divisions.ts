// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { DivisionsState, CurrentDivisionSeasons } from '@/divisions';
import {
    GET_CURRENT_DIVISION_SEASONS,
    GET_DIVISION_SEASONS_STATISTICS,
    SET_ACTIVE_DIVISION,
} from '@/app/store/actions/divisions';

/** First divisions index from list. */
const FIRST_DIVISIONS_INDEX: number = 0;

export const divisionsReducer = (
    divisionsState: DivisionsState = new DivisionsState(),
    action: any = {}
) => {
    switch (action.type) {
    case GET_CURRENT_DIVISION_SEASONS:
        return {
            ...divisionsState,
            currentDivisionSeasons: action.currentDivisionSeasons,
            activeDivision: action.currentDivisionSeasons.length
                ? action.currentDivisionSeasons[FIRST_DIVISIONS_INDEX]
                : new CurrentDivisionSeasons(),
        };
    case GET_DIVISION_SEASONS_STATISTICS:
        return {
            ...divisionsState,
            seasonsStatistics: action.seasonsStatistics,
        };
    case SET_ACTIVE_DIVISION:
        return {
            ...divisionsState,
            activeDivision: action.activeDivision,
        };
    default:
        return divisionsState;
    }
};
