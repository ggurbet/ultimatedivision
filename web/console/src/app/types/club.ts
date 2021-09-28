// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ActionFromReducer } from 'redux';


/** class for each control in option selection on field */
export class FieldControl {
    /** includes id, title and options parameters */
    constructor(
        public id: string = '',
        public title: string = '',
        public action: any = {},
        public options: string[] = [],
    ) { }
};

/* eslint-disable no-magic-numbers */
export enum Formations {
    '4-4-2' = 1,
    '4-2-4' = 2,
    '4-2-2-2' = 3,
    '4-3-1-2' = 4,
    '4-3-3' = 5,
    '4-2-3-1' = 6,
    '4-3-2-1' = 7,
    '4-1-3-2' = 8,
    '5-3-2' = 9,
    '4-5-2' = 10
}
