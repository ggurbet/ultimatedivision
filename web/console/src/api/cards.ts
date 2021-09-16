// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { CreatedLot } from '@/card';

/** ClubClient base implementation */
export class CardClient extends APIClient {
    /** To do: add exist endpoint for club */
    private readonly ROOT_PATH: string = '/api/v0';

    /** method calls get method from APIClient */
    public async getCards(): Promise<Response> {
        return await this.http.get(`${this.ROOT_PATH}/cards`);
    }
    /** method calls get method from APIClient */
    public async getFilteredCards(filterParam: string): Promise<Response> {
        return await this.http.get(`${this.ROOT_PATH}/cards/?${filterParam}`);
    }
    /** method post for implementing buying cards */
    public async sellCard(lot: CreatedLot): Promise<Response> {
        return await this.http.post(`${this.ROOT_PATH}/marketplace`, JSON.stringify(lot));
    }
    /** method calls get method from APIClient */
    public async getLots(): Promise<Response> {
        return await this.http.get(`${this.ROOT_PATH}/marketplace`);
    }
    /** method calls get method from APIClient */
    public async getFilteredLots(filterParam: string): Promise<Response> {
        return await this.http.get(`${this.ROOT_PATH}/lots/?${filterParam}`);
    }
}
