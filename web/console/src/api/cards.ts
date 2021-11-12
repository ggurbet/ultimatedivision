// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card, CardWithStats, CardsPage } from '@/card';

import { APIClient } from '@/api/index';

import { Pagination } from '@/app/types/pagination';

/** CardClient base implementation */
export class CardClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/cards';

    /** method calls get method from APIClient */
    public async list({ selectedPage, limit }: Pagination): Promise<CardsPage> {
        const path = `${this.ROOT_PATH}?page=${selectedPage}&limit=${limit}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        return await response.json();
    };

    /** method calls get method from APIClient */
    public async getCardById(id: string): Promise<Card> {
        const path = `${this.ROOT_PATH}/${id}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        return await response.json();
    };

    /** method returns filtered card list */
    public async filteredList(lowRange: string, topRange: string): Promise<CardsPage> {
        const path = `${this.ROOT_PATH}/?${lowRange}&${topRange}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        return await response.json();
    };
};
