// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Lootbox } from '@/lootbox';

/** Lootbox types */
export type LootboxTypes = 'Regular Box' | 'UD Release Celebration Box';

/** Class for lootBox Cards in store */
export class LootboxStats {
    /** LootboxStats implementation */
    constructor(
        public id: string,
        public icon: string,
        public type: LootboxTypes,
        public quantity: number,
        public dropChance: number[],
        public price: string
    ) { }
}
