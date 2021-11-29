// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Match } from '@/matches';

/** GET_MATCH_SCORE  is action type for update match score. */
export const GET_MATCH_SCORE = 'GET_MATCH_SCORE';

/** getMatchScore dispatch updates match result. */
export const getMatchScore = ({ firstTeam, secondTeam }: Match) => ({
    type: GET_MATCH_SCORE,
    payload: {
        firstTeam,
        secondTeam,
    },
});
