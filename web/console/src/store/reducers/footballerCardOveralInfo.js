
/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import star from '../../img/FootballerCardPage/star.png';
import checked from '../../img/FootballerCardPage/checked.png';

class CardInfoField {
    constructor(key, value, icon) {
        this.label = key;
        this.value = value;
        this.icon = icon;
    }
}

const overalInfo = [
    new CardInfoField('name', 'Albert Ronalculus'),
    new CardInfoField('nation', 'Portugal 🇵🇹'),
    new CardInfoField('skills', 5, star),
    new CardInfoField('weak foot', 5, star),
    new CardInfoField('intl. rep', 5, star),
    new CardInfoField('foot', 'right'),
    new CardInfoField('height', 187),
    new CardInfoField('nation', 87),
    new CardInfoField('revision', 'rare'),
    new CardInfoField('def. wr', 'low'),
    new CardInfoField('arr. wr', 'high'),
    new CardInfoField('added on', '2020-09-10'),
    new CardInfoField('origin', 'na'),
    new CardInfoField('r. Face', 'low'),
    new CardInfoField('b. type', true, checked),
    new CardInfoField('age', '36 years old')
];

/* eslint-disable */
export const cardInfoReducer = (cardState = overalInfo, action) => {
    return cardState;
};
