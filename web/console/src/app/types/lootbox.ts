// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Lootbox types */
export enum LootboxTypes {
    'Regular Box' = 'Regular Box',
    'Cool box' = 'UD Release Celebration Box',
}

/** Class for lootbox Cards in store */
export class LootboxStats {
    /** LootboxStats implementation */
    constructor(
        public id: string,
        public type: LootboxTypes,
        public quantity: number,
        public dropChance: number[],
        public price: string
    ) {}
}
