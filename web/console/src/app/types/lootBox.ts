// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Lootbox } from '@/lootbox';

/** Class for lootBox Cards in store */
export class LootboxStats {
    /** LootboxStats implementation */
    constructor(
        public id: string,
        public icon: string,
        public title: string,
        public quantity: number,
        public dropChance: number[],
        public price: string
    ) { }
}
