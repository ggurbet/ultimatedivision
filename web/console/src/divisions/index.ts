// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Domain entity Division implementation */
export class Division {
    constructor(
        public id: string,
        public name: number,
        public passingPercent: number,
        public createdAt: Date
    ) {}
}

/** Divisions of current season entity. */
export class CurrentDivisionSeasons {
    constructor(
        public id: string = '',
        public divisionId: string = '',
        public startedAt: Date = new Date(),
        public endedAt: Date = new Date()
    ) {}
}

//* * initial name for divisions state */
const INITIAL_DIVISION_NAME: number = 0;

//* * initial parsing percent for divisions state */
const INITIAL_DIVISION_PERCENT: number = 0;

// TODO: statistics need rewrite (waiting for backend).
/** Division matches statistics entity. */
export class DivisionSeasonsStatistics {
    constructor(
        public division: Division = new Division(
            '',
            INITIAL_DIVISION_NAME,
            INITIAL_DIVISION_PERCENT,
            new Date()
        ),
        public statistics: null | any[] = null
    ) {}
}

// TODO: Can be changed (waiting for backend)
/** divisions reducer initial state. */
export class DivisionsState {
    constructor(
        public currentDivisionsSeasons: CurrentDivisionSeasons[] = [],
        public seasonsStatistics: DivisionSeasonsStatistics = new DivisionSeasonsStatistics(),
        public activeDivision: CurrentDivisionSeasons = new CurrentDivisionSeasons()
    ) {}
}
