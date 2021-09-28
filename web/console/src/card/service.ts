// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { Pagination } from '@/app/types/pagination';

import { Card, Cards } from '@/card';
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
    public async list({ selectedPage, limit }: Pagination): Promise<Cards> {
        return await this.card.list({ selectedPage, limit });
    };
    /** gets card by id from list of cards */
    public async getCardById(id: string): Promise<Card> {
        return await this.card.getCardById(id);
    };
    /** gets list of filtered cards */
    public async filteredList(filterParam: string): Promise<Cards> {
        return await this.card.filteredList(filterParam);
    };
};
