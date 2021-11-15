// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** class for each control in option selection on field */
export class Control {
    /** includes id, title and options parameters */
    constructor(
        public id: string = '',
        public title: string = '',
        public action: any = {},
        public options: string[] = [],
    ) { }
};
