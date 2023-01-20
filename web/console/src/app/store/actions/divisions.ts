// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';
import { DivisionsClient } from '@/api/divisions';
import { CurrentDivisionSeasons, DivisionSeasonsStatistics, SeasonRewardTransaction } from '@/divisions';
import { DivisionsService } from '@/divisions/service';

export const GET_CURRENT_DIVISION_SEASONS = 'GET_CURRENT_DIVISION_SEASONS';
export const GET_DIVISION_SEASONS_STATISTICS =
    'GET_DIVISION_SEASONS_STATISTICS';

export const SET_ACTIVE_DIVISION = 'SET_ACTIVE_DIVISION';

/** handles gets current seasons divisions */
export const getCurrentDivisionSeasons = (
    currentDivisionSeasons: CurrentDivisionSeasons[]
) => ({
    type: GET_CURRENT_DIVISION_SEASONS,
    currentDivisionSeasons,
});

/** handles gets divisions matches statistics */
export const getDivisionSeasonsStatistics = (
    seasonsStatistics: DivisionSeasonsStatistics
) => ({
    type: GET_DIVISION_SEASONS_STATISTICS,
    seasonsStatistics,
});

/** handles sets active division */
export const setActiveDivision = (activeDivision: CurrentDivisionSeasons) => ({
    type: SET_ACTIVE_DIVISION,
    activeDivision,
});

const client = new DivisionsClient();
const service = new DivisionsService(client);

/** thunk that handles gets current seasons divisions */
export const listOfCurrentDivisionSeasons = () =>
    async function(dispatch: Dispatch) {
        const currentDivisionSeasons =
            await service.getCurrentDivisionSeasons();

        currentDivisionSeasons &&
            dispatch(getCurrentDivisionSeasons(currentDivisionSeasons));
    };

/** thunk that handles gets seasons statistics */
export const divisionSeasonsStatistics = (id: string) =>
    async function(dispatch: Dispatch) {
        const seasonsStatistics = await service.getDivisionSeasonsStatistics(
            id
        );

        dispatch(getDivisionSeasonsStatistics(seasonsStatistics));
    };
