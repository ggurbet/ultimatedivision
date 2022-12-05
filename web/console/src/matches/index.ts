// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

/** Goal is entity that describes scored goal. */
export class Goal {
    /** Goal contains of player card and minute when was scored goal. */
    constructor(public card: Card, public minute: number) { }
}

/** Team describes football team entity. */
export class Team {
    /** Team contains of summary goals number, goals array and userId. */
    constructor(
        public quantityGoals: number,
        public goals: Goal[] | null,
        public userId: string
    ) { }
};

/** Transaction describes transaction entity of match response. */
export class MatchTransaction {
    /** Transaction contains of nonce, signature hash udtContract and value. */
    constructor(
        public nonce: number,
        public signature: string,
        public udtContract: {
            address: string;
            addressMethod: string;
        },
        public value: string | undefined,
    ) { }
};

/** Transaction describes transaction entity of match response. */
export class CasperMatchTransaction {
    /** Transaction contains of nonce, signature hash udtContract and value. */
    constructor(
        public nonce: number,
        public signature: string,
        public value: string,
        public casperTokenContract: {
            address: string;
            addressMethod: string;
        },
        public rpcNodeAddress: string,
    ) { }
};

/** Match exposes match domain entity. */
export class Match {
    /** Contains of matchResults, transaction and question. */
    constructor(
        public matchResults: Team[],
        public question: string,
        public transaction: MatchTransaction,
    ) { };
};
