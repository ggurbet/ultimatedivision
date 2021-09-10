// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Lootbox } from '@/lootbox';

/** LootboxClient is a lootbox api client */
export class LootboxClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/lootboxes';
    /** buys and opens lootbox */
    public async buy(lootbox: Lootbox): Promise<Response> {
        const response = await this.http.post(this.ROOT_PATH, JSON.stringify({ type: lootbox.type }));
        // TODO: temporary code for further testing.
        if (!response.ok) {
            throw this.handleError;
        }

        const lootboxData = await response.json();

        if (!lootboxData) {
            throw this.handleError;
        }

        return await this.http.post(`${this.ROOT_PATH}/${lootboxData.id}`);
    };
};
