// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardClient } from '@/api/cards';
import { CardsResponse, CreatedLot, MarkeplaceResponse } from '@/card';

/**
 * exposes all bandwidth related logic
 */
export class CardService {
    protected readonly card: CardClient;
    /** sets ClubClient into club field */
    public constructor(club: CardClient) {
        this.card = club;
    }
    /** getting lot by id */
    public async getCardById(id: string): Promise<Response> {
        return await this.card.getCardById(id);
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
    /** create lot */
    public async createLot(lot: CreatedLot): Promise<Response> {
        return await this.card.createLot(lot);
    }

    /** getting lot by id */
    public async getLotById(id: string): Promise<Response> {
        return await this.card.getLotById(id);
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
