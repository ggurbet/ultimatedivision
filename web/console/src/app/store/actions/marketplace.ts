// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { MarketplaceClient } from '@/api/marketplace';

import { Marketplaces } from '@/marketplace/service';

import { Card } from '@/card';
import { Lot, MarketPlace } from '@/marketplace';

import { CreatedLot } from '@/app/types/marketplace';
import { Pagination } from '@/app/types/pagination';

export const GET_SELLING_CARDS = ' GET_SELLING_CARDS';
export const MARKETPLACE_CARD = 'OPEN_MARKETPLACE_CARD';

const getLots = (marketplace: MarketPlace) => ({
    type: GET_SELLING_CARDS,
    marketplace,
});
const marketplaceCard = (card: Card) => ({
    type: MARKETPLACE_CARD,
    card,
});

const marketplaceClient = new MarketplaceClient();
const marketplaces = new Marketplaces(marketplaceClient);
/** thunk for creating user cards list */
export const listOfLots = ({ selectedPage, limit }: Pagination) => async function(dispatch: Dispatch) {
    const marketplace = await marketplaces.list({ selectedPage, limit });
    const lots = marketplace.lots.
        map((lot: Lot) => ({ ...lot, card: new Card(lot.card) }));
    const page = marketplace.page;
    dispatch(getLots({ lots, page }));
};

export const createLot = (lot: CreatedLot) => async function(dispatch: Dispatch) {
    await marketplaces.createLot(lot);
};

/** thunk for opening fotballerCardPage with reload possibility */
export const openMarketplaceCard = (id: string) => async function(dispatch: Dispatch) {
    const lot = await marketplaces.getLotById(id);

    dispatch(marketplaceCard(new Card(lot.card)));
};

/** thunk returns filtered lots */
export const filteredLots = (lowRange: string, topRange: string) => async function(dispatch: Dispatch) {
    const filterParam = `${lowRange}&${topRange}`;
    const marketplace = await marketplaces.filteredList(filterParam);
    const lots = marketplace.lots.
        map((lot: Lot) => ({ ...lot, card: new Card(lot.card) }));
    const page = marketplace.page;
    dispatch(getLots({ lots, page }));
};
