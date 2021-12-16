// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardsClient } from '@/api/cards';
import { Card, CardsPage, CardsQueryParametersField } from '@/card';
import { CardService } from '@/card/service';

export const GET_USER_CARDS = ' GET_USER_CARDS';
/** Exposes action type for field cards. */
export const GET_FIELD_CARDS: string = 'GET_FIELD_CARDS';
export const USER_CARD = 'OPEN_USER_CARD';

const getCards = (cardsPage: CardsPage, currentPage: number) => ({
    type: GET_USER_CARDS,
    payload: {
        cardsPage,
        currentPage,
    },
});

/** Exposes an action object for cards on field page.  */
const getFieldCards = (cardsPage: CardsPage, currentPage: number) =>({
    type: GET_FIELD_CARDS,
    payload: {
        cardsPage,
        currentPage,
    },
});

const userCard = (card: Card) => ({
    type: USER_CARD,
    card,
});

const cardsClient = new CardsClient();
const cardsService = new CardService(cardsClient);

const fieldCardsClient = new CardsClient();
const fieldCardsService = new CardService(fieldCardsClient);

/** Clears cards query parameters. */
export const clearCardsQueryParameters = () => {
    cardsService.clearCardsQueryParameters();
};

/** Creates field cards query parameters and sets them to fieldCardsService. */
export const createFieldCardsQueryParameters = (queryParameters: CardsQueryParametersField[]) => {
    fieldCardsService.changeCardsQueryParameters(queryParameters);
};

/** Creates cards query parameters and sets them to CardsService. */
export const createCardsQueryParameters = (queryParameters: CardsQueryParametersField[]) => {
    cardsService.changeCardsQueryParameters(queryParameters);
};

/** FieldCards exposes a middleware for cards entity that dispatches current cards on page . */
export const fieldCards = (selectedPage: number) => async function(dispatch: Dispatch) {
    const response = await fieldCardsService.list(selectedPage);
    const page = response.page;
    const cards = response.cards;
    const currentPage = response.page.currentPage;

    dispatch(getFieldCards({ cards, page}, currentPage ));
};

/** thunk for creating user cards list */
export const listOfCards = (selectedPage: number) => async function(dispatch: Dispatch) {
    const response = await cardsService.list(selectedPage);
    const page = response.page;
    const cards = response.cards;
    const currentPage = response.page.currentPage;

    dispatch(getCards({ cards, page}, currentPage ));
};
/** thunk for opening fotballerCardPage with reload possibility */
export const openUserCard = (id: string) => async function(dispatch: Dispatch) {
    const card = await cardsService.getCardById(id);

    dispatch(userCard(card));
};
