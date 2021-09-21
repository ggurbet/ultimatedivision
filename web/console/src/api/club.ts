// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

/** ClubClient base implementation */
export class ClubClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** method calls get method from APIClient */
    public async createClub(): Promise<string> {
        const response = await this.http.post(`${this.ROOT_PATH}/clubs`);

        return await response.json();
    }
    /** method calls get method from APIClient */
    public async getClub(): Promise<string> {
        const response = await this.http.get(`${this.ROOT_PATH}/clubs`);

        return await response.json();
    }
    /** method calls get method from APIClient */
    public async createSquad(clubId: string): Promise<string> {
        const response = await this.http.post(`${this.ROOT_PATH}/clubs/${clubId}/squads`);

        return await response.json();
    }
    /** method calls get method from APIClient */
    public async addCard({ clubId, squadId, cardId, position }: { clubId: string, squadId: string, cardId: string, position: number }): Promise<Response> {
        return await this.http.post(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${squadId}/cards/${cardId}`,
            JSON.stringify({ position })
        );
    }
    /** method calls get method from APIClient */
    public async changeCardPosition({ clubId, squadId, cardId, position }: { clubId: string, squadId: string, cardId: string, position: number }): Promise<Response> {
        return await this.http.patch(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${squadId}/cards/${cardId}`,
            JSON.stringify({ position })
        );
    }
    /** method calls get method from APIClient */
    public async deleteCard({ clubId, squadId, cardId }: { clubId: string, squadId: string, cardId: string }): Promise<Response> {
        return await this.http.delete(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${squadId}/cards/${cardId}`
        );
    }
}
