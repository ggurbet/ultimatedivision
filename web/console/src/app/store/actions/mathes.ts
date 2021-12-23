// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Match } from '@/matches';

/** GET_MATCH_SCORE  is action type for update match score and transaction. */
export const GET_MATCH_SCORE = 'GET_MATCH_SCORE';

/** GetMatchScore dispatch updates match result and transaction values. */
export const getMatchScore = (match: Match) => ({
    type: GET_MATCH_SCORE,
    payload: match,
});
