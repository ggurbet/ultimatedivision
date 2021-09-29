// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { CardClient } from '@/api/cards';

import { Card, CardsPage, ICard } from '@/card';
import { CardService } from '@/card/service';

import { Pagination } from '@/app/types/pagination';

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

const client = new CardClient();
const service = new CardService(client);

const DEFAULT_PAGE_NUMBER: number = 1;
/** thunk for creating user cards list */
export const listOfCards = ({ selectedPage, limit }: Pagination) => async function(dispatch: Dispatch) {
    const response = await service.list({ selectedPage, limit });

    const cards = response.cards.
        map((card: Partial<ICard>) => new Card(card));
    const page = response.page;
    dispatch(getCards({ cards, page }));
};
/** thunk for opening fotballerCardPage with reload possibility */
export const openUserCard = (id: string) => async function(dispatch: Dispatch) {
    const card = await service.getCardById(id);
    dispatch(userCard(new Card(card)));
};

/** thunk returns filtered cards */
export const filteredCards = (lowRange: string, topRange: string) => async function(dispatch: Dispatch) {
    const filterParam = `${lowRange}&${topRange}`;
    const response = await service.filteredList(filterParam);
    const cards = response.cards.
        map((card: Partial<ICard>) => new Card(card));
    const page = response.page;
    dispatch(getCards({ cards, page }));
};
