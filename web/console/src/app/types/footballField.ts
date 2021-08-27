// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

/** Base FootballField implementation */
export class FootballFieldCard {
    /** class which implements fotballField card */
    constructor(
        public id: number,
        public cardData: null | Card,
    ) { }
}
/* eslint-disable no-magic-numbers */


/** class which implements field setup */
export class FootballField {
    public options = {
        formation: '4-4-2',
        captain: '',
        tactics: '',
        chosedCard: 0,
        showCardSeletion: false,
        dragStart: null,
        dragTarget: null,
    };
    /** football team implementation */
    public cardsList = [
        new FootballFieldCard(0, null),
        new FootballFieldCard(1, null),
        new FootballFieldCard(2, null),
        new FootballFieldCard(3, null),
        new FootballFieldCard(4, null),
        new FootballFieldCard(5, null),
        new FootballFieldCard(6, null),
        new FootballFieldCard(7, null),
        new FootballFieldCard(8, null),
        new FootballFieldCard(9, null),
        new FootballFieldCard(10, null),

    ];
}

/** implementation for each field in
 * FootballFieldInformation component
 */
export class FootballFieldInformationLine {
    /** includes id, title and options parameters */
    constructor(
        public id: string = '',
        public title: string = '',
        public options: string[] = [],
    ) {
        this.id = id;
        this.title = title;
        this.options = options;
    }
};
