// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardClient } from '@/api/cards';
import { Card, CardInterface, CreatedLot, MarketplaceLot } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_USER_CARDS';
export const GET_SELLING_CARDS = ' GET_SELLING_CARDS';

export const getCards = (cards: Card[]) => ({
    type: GET_USER_CARDS,
    cards,
});
export const getLots = (cards: Array<Partial<MarketplaceLot>>) => ({
    type: GET_SELLING_CARDS,
    cards,
});

const client = new CardClient();
const service = new CardService(client);

/** thunk for creating user cards list */
export const userCards = () => async function(dispatch: Dispatch) {
    const response = await service.getCards();
    const cards = response.cards;
    dispatch(getCards(cards.map((card: Partial<CardInterface>) => new Card(card))));
};
/** thunk for creating user cards list */
export const marketplaceCards = () => async function(dispatch: Dispatch) {
    const response = await service.getLots();
    const lots = response.lots;
    dispatch(getLots(lots.map((lot: Partial<MarketplaceLot>) => ({ ...lot, card: new Card(lot.card) }))));
};

/** thunk for creating marketplace lot */
export const sellCard = (lot: CreatedLot) => async function(dispatch: any) {
    await service.sellCard(lot);
    dispatch(userCards());
    dispatch(marketplaceCards());
};

/** thunk returns filtered cards */
export const filteredCards = (lowRange: string, topRange: string) => async function(dispatch: Dispatch) {
    const filterParam = `${lowRange}&${topRange}`;
    const response = await service.getFilteredCards(filterParam);
    const cards = await response.json();
    dispatch(getCards(cards.map((card: Partial<CardInterface>) => new Card(card))));
};

/** thunk returns filtered lots */
export const filteredLots = (lowRange: string, topRange: string) => async function(dispatch: Dispatch) {
    const filterParam = `${lowRange}&${topRange}`;
    const response = await service.getFilteredLots(filterParam);
    const lots = await response.json();
    dispatch(getLots(lots.map((lot: Partial<MarketplaceLot>) => ({ ...lot, card: new Card(lot.card) }))));
};
