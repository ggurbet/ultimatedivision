// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { Pagination } from '@/app/types/pagination';

import { Card, CardWithStats, CardsPage } from '@/card';
/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly card: CardClient;

    /** sets CardClient into card field */
    public constructor(card: CardClient) {
        this.card = card;
    };

    /** gets list of cards by user */
    public async list({ selectedPage, limit }: Pagination): Promise<CardsPage> {
        const response = await this.card.list({ selectedPage, limit });

        return { ...response, cards: response.cards.map((card: Card) => new CardWithStats(card)) };
    };

    /** gets card by id from list of cards */
    public async getCardById(id: string): Promise<CardWithStats> {
        const card = await this.card.getCardById(id);

        return new CardWithStats(card);
    };

    /** gets list of filtered cards */
    public async filteredList(lowRange: string, topRange: string): Promise<CardsPage> {
        const response = await this.card.filteredList(lowRange, topRange);

        return { ...response, cards: response.cards.map((card: Card) => new CardWithStats(card)) };
    };
};
