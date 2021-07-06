/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

export class FootballField {
    public options = {
        formation: '4-4-2',
        captain: '',
        tactics: '',
    };
    public cardsList = [
        {
            id: 0,
            props: ''
        },
        {
            id: 1,
            props: ''
        },
        {
            id: 2,
            props: ''
        },
        {
            id: 3,
            props: ''
        },
        {
            id: 4,
            props: ''
        },
        {
            id: 5,
            props: ''
        },
        {
            id: 6,
            props: ''
        },
        {
            id: 7,
            props: ''
        },
        {
            id: 8,
            props: ''
        },
        {
            id: 9,
            props: ''
        },
        {
            id: 10,
            props: ''
        },
    ]
}

export class FootballFieldInformationLine {
    constructor(
        public id: string ='',
        public title: string = '',
        public options: string[] = [],
    ) { }
}
