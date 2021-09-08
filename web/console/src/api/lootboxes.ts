// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Lootbox } from '@/lootbox';

/** LootboxClient is a lootbox api client */
export class LootboxClient extends APIClient {
    private readonly ROOT_PATH: string = '/lootboxes';
    /** buys and opens lootbox */
    public async buy(lootbox: Lootbox): Promise<Response> {
        /** TODO: delete post http method after back-end reworks */
        await this.http.post(this.ROOT_PATH, JSON.stringify({ type: lootbox.type }));

        return await this.http.delete(`${this.ROOT_PATH}/${lootbox.uuid}`);
    };
};
