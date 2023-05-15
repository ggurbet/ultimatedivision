// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';
import { Lot, MarketPlacePage } from '@/marketplace';
import { GET_SELLING_CARDS, MARKETPLACE_CARD } from '../actions/marketplace';

/** Markeplace state base implementation */
class MarketplaceState {
    /** default state implementation */
    constructor(
        public marketplacePage: MarketPlacePage,
        public lot: Lot,
    ) { };
};

const DEFAULT_PRICES = 0;
const DEFAULT_PERIOD = 0;
const DEFAULT_OFFSET_VALUE: number = 0;
const DEFAULT_LIMIT_VALUE: number = 24;
const FIRST_PAGE: number = 1;
const PAGES_COUNT: number = 1;
const LOTS_TOTAL_COUNT: number = 1;

export const page = {
    offset: DEFAULT_OFFSET_VALUE,
    limit: DEFAULT_LIMIT_VALUE,
    currentPage: FIRST_PAGE,
    pageCount: PAGES_COUNT,
    totalCount: LOTS_TOTAL_COUNT,
};

const marketplacePage = new MarketPlacePage([], page);
const lot: Lot = {
    cardId: '00000000-0000-0000-0000-000000000000',
    type: 'card',
    userId: '00000000-0000-0000-0000-000000000000',
    shopperId: '00000000-0000-0000-0000-000000000000',
    status: 'active',
    currentPrice: DEFAULT_PRICES,
    maxPrice: DEFAULT_PRICES,
    startPrice: DEFAULT_PRICES,
    startTime: '',
    endTime: '',
    period: DEFAULT_PERIOD,
    card: new Card(),
};


export const marketplaceReducer = (marketplaceState: MarketplaceState = new MarketplaceState(marketplacePage, lot), action: any = {}) => {
    switch (action.type) {
    case GET_SELLING_CARDS:
        return {
            ...marketplaceState,
            marketplacePage: action.marketplacePage,
        };
    case MARKETPLACE_CARD:
        return {
            ...marketplaceState,
            lot: action.lot,
        };
    default:
        return marketplaceState;
    }
};
