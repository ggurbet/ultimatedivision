// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

/** Domain entity Division implementation */
export class Division {
    /** default division implementation */
    constructor(
        public id: string,
        public name: number,
        public passingPercent: number,
        public createdAt: Date
    ) { }
}

/** Divisions of current season entity. */
export class CurrentDivisionSeasons {
    /** default currentDivisionSeasons implementation */
    constructor(
        public id: string = '',
        public divisionId: string = '',
        public startedAt: Date = new Date(),
        public endedAt: Date = new Date()
    ) { }
}

/** initial name for divisions state */
const INITIAL_DIVISION_NAME: number = 0;

/** initial parsing percent for divisions state */
const INITIAL_DIVISION_PERCENT: number = 0;

// TODO: statistics need rewrite (waiting for backend).
/** Division matches statistics entity. */
export class DivisionSeasonsStatistics {
    /** default divisionSeasonsStatistics implementation */
    constructor(
        public division: Division = new Division(
            '',
            INITIAL_DIVISION_NAME,
            INITIAL_DIVISION_PERCENT,
            new Date()
        ),
        public statistics: null | any[] = null
    ) { }
}

// TODO: Can be changed (waiting for backend).
/** divisions reducer initial state. */
export class DivisionsState {
    /** default divisionState implementation */
    constructor(
        public currentDivisionsSeasons: CurrentDivisionSeasons[] = [],
        public seasonsStatistics: DivisionSeasonsStatistics = new DivisionSeasonsStatistics(),
        public activeDivision: CurrentDivisionSeasons = new CurrentDivisionSeasons()
    ) { }
}

/** Initial status for season reward */
const INITIAL_SEASON_REWARD_STATUS: number = 0;

/** Initial seasonId for season reward */
const INITIAL_SEASON_REWARD_SEASON_ID: number = 0;

/** initial nonce for season reward */
const INITIAL_SEASON_REWARD_NONCE: number = 0;

/** Seasons reward transaction entity. */
export class SeasonRewardTransaction {
    /** default SeasonRewardTransaction implementation */
    constructor(
        public Id: string = '00000000-0000-0000-0000-000000000000',
        public userId: string = '00000000-0000-0000-0000-000000000000',
        public seasonId: number = INITIAL_SEASON_REWARD_SEASON_ID,
        public walletAddress: string = '0x0000000000000000000000000000000000000000',
        public casperWalletAddress: string = '',
        public CasperWalletHash: string = '',
        public walleType: string = '',
        public status: number = INITIAL_SEASON_REWARD_STATUS,
        public nonce: number = INITIAL_SEASON_REWARD_NONCE,
        public signature: string = '',
        public value: string = '0',
        public casperTokenContract: string = '',
        public rpcNodeAddress: string = ''
    ) { }
}
