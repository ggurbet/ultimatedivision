// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { LootboxTypes } from '@/app/types/lootBox';

/** domain entity Lootbox implementation */
export class Lootbox {
    /** receives base params: id and type */
    constructor(
        public id: string,
        public type: LootboxTypes,
    ) { };
};
