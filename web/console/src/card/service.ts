// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';

/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly card: CardClient;
    /** sets ClubClient into club field */
    public constructor(club: CardClient = new CardClient()) {
        this.card = club;
    }
    /** get catds from api */
    public async get() {
        return await this.card.get();
    }
    /** post cards into buyed cardlist */
    public async buy(param: string) {
        return await this.card.buy(param);
    }
}
