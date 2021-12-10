// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardsClient } from '@/api/cards';
import { Card, CardsPage, CardsQueryParametersField } from '@/card';

/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly cards: CardsClient;

    /** sets CardClient into card field */
    public constructor(cards: CardsClient) {
        this.cards = cards;
    };

    /** Changes cards query parameters. */
    public changeCardsQueryParameters(queryParameters: CardsQueryParametersField[]) {
        this.cards.changeCardsQueryParameters(queryParameters);
    };

    /** gets list of cards by user */
    public async list(selectedPage: number): Promise<CardsPage> {
        return await this.cards.list(selectedPage);
    };

    /** gets card by id from list of cards */
    public async getCardById(id: string): Promise<Card> {
        return await this.cards.getCardById(id);
    };
};
