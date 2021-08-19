// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/cards';

const FIRST_CARD_TYPE = 0;
const SECOND_CARD_TYPE = 1;
const THIRD_CARD_TYPE = 2;
const FOURTH_CARD_TYPE = 3;
const CARDS_AMOUNT = 20;

/** create list of player cards (implementation for test)*/
function cardList(count: number): Card[] {
    const list: Card[] = [];
    while (count) {
        list.push(
            new Card(FIRST_CARD_TYPE),
            new Card(SECOND_CARD_TYPE),
            new Card(THIRD_CARD_TYPE),
            new Card(FOURTH_CARD_TYPE)
        );
        count--;
    }

    return list;
}

export const cardReducer = (cardState = cardList(CARDS_AMOUNT), action: any = {}) => cardState;
