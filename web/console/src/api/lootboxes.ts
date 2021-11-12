// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Lootbox } from '@/lootbox';
import { Card } from '@/card';

/** LootboxClient is a lootbox api client */
export class LootboxClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/lootboxes';
    /** buys and opens lootbox */
    public async buy(lootbox: Lootbox): Promise<Card[]> {
        const lootboxResponse = await this.http.post(this.ROOT_PATH, JSON.stringify({ type: lootbox.type }));
        // TODO: temporary code for further testing.
        if (!lootboxResponse.ok) {
            this.handleError(lootboxResponse);
        }

        const lootboxData = await lootboxResponse.json();

        if (!lootboxData) {
            this.handleError(lootboxResponse);
        }

        const responseCards = await this.http.post(`${this.ROOT_PATH}/${lootboxData.id}`);

        if (!responseCards.ok) {
            this.handleError(responseCards);
        }

        const cards = await responseCards.json();

        if (!cards) {
            this.handleError(cards);
        }

        return cards;
    };
};
