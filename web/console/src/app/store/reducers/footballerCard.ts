// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/app/types/fotballerCard';

/** create list of player cards (implementation for test)*/
function cardList(count: number): Card[] {
    const list: Card[] = [];
    while (count > 0) {
        list.push(new Card(0), new Card(1), new Card(2), new Card(3));
        count--;
    }

    return list;
}

export const cardReducer = (cardState = cardList(20), action: any = {}) => {
    return cardState;
};
