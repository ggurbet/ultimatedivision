// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';
import { MarketplaceClient } from '@/api/marketplace';
import { Marketplaces } from '@/marketplace/service';
import { CardWithStats } from '@/card';
import { Lot, MarketPlacePage } from '@/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { Pagination } from '@/app/types/pagination';

export const GET_SELLING_CARDS = ' GET_SELLING_CARDS';
export const MARKETPLACE_CARD = 'OPEN_MARKETPLACE_CARD';

const getLots = (marketplacePage: MarketPlacePage) => ({
    type: GET_SELLING_CARDS,
    marketplacePage,
});
const marketplaceCard = (card: CardWithStats) => ({
    type: MARKETPLACE_CARD,
    card,
});

const marketplaceClient = new MarketplaceClient();
const marketplaces = new Marketplaces(marketplaceClient);
/** thunk for creating user cards list */
export const listOfLots = ({ selectedPage, limit }: Pagination) => async function(dispatch: Dispatch) {
    const marketplace = await marketplaces.list({ selectedPage, limit });
    const lots = marketplace.lots;
    const page = marketplace.page;

    dispatch(getLots({ lots, page }));
};

export const createLot = (lot: CreatedLot) => async function(dispatch: Dispatch) {
    await marketplaces.createLot(lot);
};

/** thunk for opening fotballerCardPage with reload possibility */
export const openMarketplaceCard = (id: string) => async function(dispatch: Dispatch) {
    const lot = await marketplaces.getLotById(id);

    dispatch(marketplaceCard(new CardWithStats(lot.card)));
};

/** thunk returns filtered lots */
export const filteredLots = (lowRange: string, topRange: string) => async function(dispatch: Dispatch) {
    const marketplace = await marketplaces.filteredList(lowRange, topRange);
    const lots = marketplace.lots;
    const page = marketplace.page;

    dispatch(getLots({ lots, page }));
};
