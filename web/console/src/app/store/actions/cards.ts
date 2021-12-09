// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardsClient } from '@/api/cards';
import { Card, CardsPage, CardsQueryParametersField } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_USER_CARDS';
export const USER_CARD = 'OPEN_USER_CARD';

const getCards = (cardsPage: CardsPage) => ({
    type: GET_USER_CARDS,
    cardsPage,
});
const userCard = (card: Card) => ({
    type: USER_CARD,
    card,
});

const cardsClient = new CardsClient();
const cardsService = new CardService(cardsClient);

/** Creates cards query parameters and sets them to CardsService. */
export const createCardsQueryParameters = (queryParameters: CardsQueryParametersField[]) => {
    cardsService.changeCardsQueryParameters(queryParameters);
};

/** thunk for creating user cards list */
export const listOfCards = (selectedPage: number) => async function(dispatch: Dispatch) {
    const response = await cardsService.list(selectedPage);
    const page = response.page;
    const cards = response.cards;

    dispatch(getCards({ cards, page }));
};
/** thunk for opening fotballerCardPage with reload possibility */
export const openUserCard = (id: string) => async function(dispatch: Dispatch) {
    const card = await cardsService.getCardById(id);

    dispatch(userCard(card));
};
