// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Club, Squad } from '@/club';

/** ClubClient base implementation */
export class ClubClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** method calls get method from APIClient */
    public async createClub(): Promise<string> {
        const response = await this.http.post(`${this.ROOT_PATH}/clubs`);
        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
    /** method calls get method from APIClient */
    public async getClub(): Promise<Club> {
        const response = await this.http.get(`${this.ROOT_PATH}/clubs`);
        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
    /** method calls get method from APIClient */
    public async createSquad(clubId: string): Promise<string> {
        const response = await this.http.post(`${this.ROOT_PATH}/clubs/${clubId}/squads`);
        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
    /** method calls get method from APIClient */
    public async addCard({ squad, cardId, position }: { squad: Squad; cardId: string; position: number }): Promise<void> {
        const response = await this.http.post(
            `${this.ROOT_PATH}/clubs/${squad.clubId}/squads/${squad.id}/cards/${cardId}`,
            JSON.stringify({ position })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** method calls get method from APIClient */
    public async changeCardPosition({ clubId, squadId, cardId, position }: { clubId: string; squadId: string; cardId: string; position: number }): Promise<void> {
        const response = await this.http.patch(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${squadId}/cards/${cardId}`,
            JSON.stringify({ position })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** method calls get method from APIClient */
    public async deleteCard({ clubId, squadId, cardId }: { clubId: string; squadId: string; cardId: string }): Promise<void> {
        const response = await this.http.delete(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${squadId}/cards/${cardId}`
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** method updates squad position, formation and captain */
    public async updateSquad(squad: Squad): Promise<void> {
        const { tactic, formation, captainId, clubId, id } = squad;
        const response = await this.http.patch(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${id}`,
            JSON.stringify({ formation, tactic, captainId })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
