// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { LootboxClient } from '@/api/lootboxes';
import { LootboxService } from '@/lootbox/service';
import { BUY_LOOTBOX } from '../actions/lootboxes';

/** Lootbox state base implementation */
export class LootboxState {
    public readonly lootboxService: LootboxService;
    public lootbox = [];
    /** receives lootbox service */
    public constructor(lootboxService: LootboxService) {
        this.lootboxService = lootboxService;
    };
};

const lootboxClient = new LootboxClient();
const lootboxService = new LootboxService(lootboxClient);
const lootboxState = new LootboxState(lootboxService);

export const lootboxReducer = (
    state = lootboxState,
    action: any = {}
) => {
    switch (action.type) {
    case BUY_LOOTBOX:
        state.lootbox = action.lootbox;
        break;
    default:
        break;
    };

    return { ...state };
};
