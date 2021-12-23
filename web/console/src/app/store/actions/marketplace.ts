// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { MarketplaceClient } from '@/api/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { Card, CardsQueryParametersField } from '@/card';
import { MarketPlacePage } from '@/marketplace';
import { Marketplaces } from '@/marketplace/service';

export const GET_SELLING_CARDS = ' GET_SELLING_CARDS';
export const MARKETPLACE_CARD = 'OPEN_MARKETPLACE_CARD';

const getLots = (marketplacePage: MarketPlacePage) => ({
    type: GET_SELLING_CARDS,
    marketplacePage,
});
const marketplaceCard = (card: Card) => ({
    type: MARKETPLACE_CARD,
    card,
});

const marketplaceClient = new MarketplaceClient();
const marketplaces = new Marketplaces(marketplaceClient);

/** Returns current cards queryParameters object. */
export const getCurrentLotsQueryParameters = () => marketplaces.getCurrentQueryParameters();

/** Creates lots query parameters and sets them to marketplace service. */
export const createLotsQueryParameters = (queryParameters: CardsQueryParametersField[]) => {
    marketplaces.changeLotsQueryParameters(queryParameters);
};

/** thunk for creating user cards list */
export const listOfLots = (selectedPage: number) => async function(dispatch: Dispatch) {
    const marketplace = await marketplaces.list(selectedPage);
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

    dispatch(marketplaceCard(lot.card));
};
