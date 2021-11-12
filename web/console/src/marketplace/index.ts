// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardWithStats } from '@/card';

/** base marketplace Lot type implementation */
export interface Lot {
    id: string;
    itemId: string;
    type: string;
    userId: string;
    shopperId: string;
    status: string;
    startPrice: number;
    maxPrice: number;
    currentPrice: number;
    startTime: string;
    endTime: string;
    period: number;
    card: CardWithStats;
};

/** base MarketPlace domain entity type implementation */
export interface MarketPlacePage {
    lots: Lot[];
    page: {
        offset: number;
        limit: number;
        currentPage: number;
        pageCount: number;
        totalCount: number;
    };
};

/** marketplace reducer state */
export class MarketPlaceState {
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
