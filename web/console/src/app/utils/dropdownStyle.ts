// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** class for changing dropdown visibility according to component state */
export class DropdownStyle {
    /** property visibility  */
    constructor(
        public vilibility: boolean,
    ) {
        this.vilibility = vilibility;
    }

    /** triangle style */
    get triangleRotate() {
        return this.vilibility ? 'rotate(0deg)' : 'rotate(-90deg)';
    }
    /** list height */
    get listHeight() {
        return this.vilibility ? 'unset' : '0';
    }
}
