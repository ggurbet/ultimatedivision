// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { GET_MATCH_SCORE } from '../actions/mathes';


import { Goal, Match, MatchTransaction, Team } from '@/matches';

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

/** Describes default contract address. */
const DEFAULT_ADDRESS_CONTRACT: string = '';
/** Desribes default contract address method. */
const DEFALT_CONTRACT_ADDRESS_METHOD: string = '';
/** Describes default nonce contract value. */
const DEFAULT_NONCE_VALUE: number = 0;
/** Describes default hash of signature. */
const SIGNATURE_HASH: string = '';
/** Describes default coins value. */
const COINS_VALUE: string = '';

/** Describes default question to confirm add wallet. */
const CONFIRM_QUESTION: string = '';

const firstTeam = new Team(DEFAULT_FIRST_TEAM_GOALS, DEFAULT_FIRST_TEAM_GOAL_SCORERS, DEFAULT_FIRST_USER_ID);
const secondTeam = new Team(DEFAULT_SECOND_TEAM_GOALS, DEFAULT_SECOND_TEAM_GOAL_SCORERS, DEFAULT_SECOND_USER_ID);

const transaction = new MatchTransaction(
    DEFAULT_NONCE_VALUE,
    SIGNATURE_HASH,
    {
        address: DEFAULT_ADDRESS_CONTRACT,
        addressMethod: DEFALT_CONTRACT_ADDRESS_METHOD,
    },
    COINS_VALUE,
);

/** Exposes matches result that return array of teams. */
const matchResults = [firstTeam, secondTeam];

/** MatchesReducer describes reducer for matches domain entity */
export const matchesReducer = (
    matchesState: Match = new Match(matchResults, CONFIRM_QUESTION, transaction),
    action: any = {}
) => {
    switch (action.type) {
    case GET_MATCH_SCORE:
        return {
            ...matchesState,
            question: action.payload.question,
            matchResults:  action.payload.matchResults,
            transaction: action.payload.transaction,
        };
    default:
        return matchesState;
    }
};
