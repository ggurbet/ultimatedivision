
//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

/** class for changing dropdown visibility according to component state */
export class DropdownStyle {
    /** property visibility  */
    constructor(
        public vilibility: boolean,
        public height: number
    ) {
        this.vilibility = vilibility;
        this.height = height;
    }

    /** triangle style */
    get triangleRotate() {
        return this.vilibility ? 'rotate(0deg)' : 'rotate(-90deg)';
    }
    /** list height */
    get listHeight() {
        return this.vilibility ? `${this.height}px` : '0';
    }
}
