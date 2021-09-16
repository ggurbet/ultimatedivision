// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { Card, CardInterface, CardsResponse, CreatedLot, MarkeplaceResponse } from '@/card';

/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly card: CardClient;
    /** sets ClubClient into club field */
    public constructor(club: CardClient) {
        this.card = club;
    }
    /** get filtered cards from api */
    public async getFilteredCards(filterParam: string) {
        const response = await this.card.getFilteredCards(filterParam);

        return await response.json();
    }
    /** get user cards from api */
    public async getCards(): Promise<CardsResponse> {
        const response = await this.card.getCards();

        return await response.json();
    }
    /** sell card */
    public async sellCard(lot: CreatedLot): Promise<Response> {
        return await this.card.sellCard(lot);
    }
    /** get lots from api */
    public async getLots(): Promise<MarkeplaceResponse> {
        const response = await this.card.getLots();

        return await response.json();
    }
    /** get filtered lots from api */
    public async getFilteredLots(filterParam: string) {
        const response = await this.card.getFilteredLots(filterParam);

        return await response.json();
    }
}
