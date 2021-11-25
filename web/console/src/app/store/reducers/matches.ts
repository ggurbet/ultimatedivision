// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { GET_MATCH_SCORE } from '../actions/mathes';


import { Match } from '@/matches';

const DEFAULT_FIRST_TEAM_GOALS_SCORED: number = 0;
const DEFAULT_SECOND_TEAM_GOALS_SCORED: number = 0;

/** matchesReducer describes reducer for matches domain entity */
export const matchesReducer = (
    matchesState: Match = new Match(DEFAULT_FIRST_TEAM_GOALS_SCORED, DEFAULT_SECOND_TEAM_GOALS_SCORED),
    action: any = {}
) => {
    switch (action.type) {
    case GET_MATCH_SCORE:
        return {
            ...matchesState,
            firstTeamGoalsCrored: action.payload.firstTeamGoalsCrored,
            secondTeamGoalsScored: action.payload.secondTeamGoalsScored,
        };
    default:
        return matchesState;
    }
};
