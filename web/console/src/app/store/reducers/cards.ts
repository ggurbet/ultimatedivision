// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { GET_SELLING_CARDS, GET_USER_CARDS } from '@/app/store/actions/cards';
import { CardService } from '@/card/service';
import { CardClient } from '@/api/cards';
import { Card, MarketplaceLot } from '@/card';

/** class for data from backent (test) */
class CardSetup {
    /** class implementation */
    constructor(
        public cardService: CardService,
        public marketplace: MarketplaceLot[],
        public club: Card[],
    ) { };
};

const cardClient = new CardClient();
const cardService = new CardService(cardClient);
export const cardSetup = new CardSetup(cardService, [], []);

export const cardsReducer = (cardState = cardSetup, action: any = {}) => {
    switch (action.type) {
    case GET_USER_CARDS:
        cardState.club = action.cards;
        break;
    case GET_SELLING_CARDS:
        cardState.marketplace = action.cards;
        break;
    default:
        break;
    }

    return { ...cardState };
};
