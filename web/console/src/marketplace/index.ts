// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

/** base marketplace Lot type implementation */
export class Lot {
    /** default lot initial values */
    constructor(
        public id: string,
        public itemId: string,
        public type: string,
        public userId: string,
        public shopperId: string,
        public status: string,
        public startPrice: number,
        public maxPrice: number,
        public currentPrice: number,
        public startTime: string,
        public endTime: string,
        public period: number,
        public card: Card,
    ) { };
};
/** base MarketPlace domain entity type implementation */
export class MarketPlacePage {
    /** default MarketPlace initial values */
    constructor(
        public lots: Lot[],
        public page: {
            offset: number;
            limit: number;
            currentPage: number;
            pageCount: number;
            totalCount: number;
        },
    ) { };
};
