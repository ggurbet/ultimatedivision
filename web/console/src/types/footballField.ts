/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

/** Base FootballField implementation */
export class FootballField {
    public options = {
        formation: '4-4-2',
        captain: '',
        tactics: '',
        chosedCard: 0,
        dragStart: -1,
        dragTarget: -1,
    };
    /** football team implementation */
    public cardsList = [
        {
            id: 0,
            cardData: null,
        },
        {
            id: 1,
            cardData: null,
        },
        {
            id: 2,
            cardData: null,
        },
        {
            id: 3,
            cardData: null,
        },
        {
            id: 4,
            cardData: null,
        },
        {
            id: 5,
            cardData: null,
        },
        {
            id: 6,
            cardData: null,
        },
        {
            id: 7,
            cardData: null,
        },
        {
            id: 8,
            cardData: null,
        },
        {
            id: 9,
            cardData: null,
        },
        {
            id: 10,
            cardData: null,
        },
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
    ) { }
}
