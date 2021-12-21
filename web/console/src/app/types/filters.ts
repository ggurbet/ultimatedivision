// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Exposes default index which does not exist in array. */
const DEFAULT_FILTER_ITEM_INDEX = -1;

/** Class exposes default parameter in useContext */
export class Context {
    constructor(
        public activeFilterIndex: number = DEFAULT_FILTER_ITEM_INDEX,
        public setActiveFilterIndex: any = () => { }
    ) { }
}
