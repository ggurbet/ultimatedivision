// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** base class for type pagination implements */
export class Pagination {
    /** default pagination type implementation */
    constructor(
        public selectedPage: number,
        public limit: number,
    ) { };
};
