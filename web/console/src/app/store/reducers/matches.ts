// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { GET_MATCH_SCORE } from '../actions/mathes';


import { Goal, Match, Team } from '@/matches';

/** Describes default summary goals scored by first team. */
const DEFAULT_FIRST_TEAM_GOALS: number = 0;
/** Describes default summary goals scored by second team. */
const DEFAULT_SECOND_TEAM_GOALS: number = 0;

/** Describes default goal scorers by first team. */
const DEFAULT_FIRST_TEAM_GOAL_SCORERS: Goal[] = [];
/** Describes default goal scorers by second team. */
const DEFAULT_SECOND_TEAM_GOAL_SCORERS: Goal[] = [];

/** Describes default userId value of first player. */
const DEFAULT_FIRST_USER_ID: string = '';

/** Describes default userId valuew of second player. */
const DEFAULT_SECOND_USER_ID: string = '';

const firstTeam = new Team(DEFAULT_FIRST_TEAM_GOALS, DEFAULT_FIRST_TEAM_GOAL_SCORERS, DEFAULT_FIRST_USER_ID);
const secondTeam = new Team(DEFAULT_SECOND_TEAM_GOALS, DEFAULT_SECOND_TEAM_GOAL_SCORERS, DEFAULT_SECOND_USER_ID);

/** matchesReducer describes reducer for matches domain entity */
export const matchesReducer = (
    matchesState: Match = new Match(firstTeam, secondTeam),
    action: any = {}
) => {
    switch (action.type) {
    case GET_MATCH_SCORE:
        return {
            ...matchesState,
            firstTeam: action.payload.firstTeam,
            secondTeam: action.payload.secondTeam,
        };
    default:
        return matchesState;
    }
};
