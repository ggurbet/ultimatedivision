// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardClient } from '@/api/cards';
import { Card } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_CARDS';
export const GET_SELLING_CARDS = ' GET_CARDS';

export const getUserCards = (cards: []) => ({
    type: GET_USER_CARDS,
    action: cards,
});
export const getSellingCards = (cards: []) => ({
    type: GET_SELLING_CARDS,
    cards,
});

const client = new CardClient();
const service = new CardService(client);

/** thunk for creating user cards list */
export const userCards = () => async function (dispatch: Dispatch) {
    const response = await service.getUserCards();
    const cards = await response.json();

    await dispatch(getUserCards(cards));
};
/** thunk for creating user cards list */
export const marketplaceCards = () => async function (dispatch: Dispatch) {
    const response = await service.getSellingCards();
    const cards = await response.json();

    await dispatch(getSellingCards(cards));
};
export const sellCard = (id: string) => async function(dispatch: any) {
    await service.sellCard(id);
    dispatch(userCards());
    dispatch(marketplaceCards());
};
