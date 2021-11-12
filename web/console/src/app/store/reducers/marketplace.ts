// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardWithStats } from '@/card';
import { MarketPlacePage, MarketPlaceState } from '@/marketplace';
import { GET_SELLING_CARDS, MARKETPLACE_CARD } from '../actions/marketplace';

/** Markeplace state base implementation */
class MarketplaceState {
    /** default state implementation */
    constructor(
        public marketplacePage: MarketPlacePage,
        public card: CardWithStats,
    ) { };
};

const DEFAULT_OFFSET_VALUE: number = 0;
const DEFAULT_LIMIT_VALUE: number = 24;
const FIRST_PAGE: number = 1;
const PAGES_COUNT: number = 1;
const LOTS_TOTAL_COUNT: number = 1;

const page = {
    offset: DEFAULT_OFFSET_VALUE,
    limit: DEFAULT_LIMIT_VALUE,
    currentPage: FIRST_PAGE,
    pageCount: PAGES_COUNT,
    totalCount: LOTS_TOTAL_COUNT,
};

const marketplacePage = new MarketPlaceState([], page);
const card = new CardWithStats();

export const marketplaceReducer = (marketplaceState: MarketplaceState = new MarketplaceState(marketplacePage, card), action: any = {}) => {
    switch (action.type) {
        case GET_SELLING_CARDS:
            return {
                ...marketplaceState,
                marketplacePage: action.marketplacePage,
            };
        case MARKETPLACE_CARD:
            return {
                ...marketplaceState,
                card: action.card,
            };
        default:
            return marketplaceState;
    }
};
