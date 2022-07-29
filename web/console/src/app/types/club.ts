// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

const DEFAULT_NUMBER = 0;
const FOUR_ELEMENTS_COLUMN = 4;
const FIVE_ELEMENTS_COLUMN = 5;

export enum amountColumnsElements {
    'default' = DEFAULT_NUMBER,
    'four-elements' = FOUR_ELEMENTS_COLUMN,
    'five-elements' = FIVE_ELEMENTS_COLUMN,
}

/** class for each control in option selection on field */
export class Control {
    /** includes id, title and options parameters */
    constructor(
        public id: string = '',
        public title: string = '',
        public action: any = {},
        public options: string[] = [],
        public columnElements: number = DEFAULT_NUMBER,
        public currentValue: any = '',
        public fieldId: string = '',
        public fieldName: string = '',
        public fieldText: string = ''
    ) {}
}
