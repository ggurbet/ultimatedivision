// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';
import { Club, Squad } from '@/club';

const DEFAULT_VALUE = 0;

/** class for api methods to declare full path of card inside of club */
export class CardEditIdentificators {
    /** includes club id, squad idm card id and position parameters */
    constructor(
        public clubId: string,
        public squadId: string,
        public cardId: string,
        public position: number = DEFAULT_VALUE
    ) { }
}

/** ClubsClient base implementation */
export class ClubsClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** creates club entity */
    public async createClub(): Promise<string> {
        const response = await this.http.post(`${this.ROOT_PATH}/clubs`);
        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
    /** gets club from api */
    public async getClubs(): Promise<Club[]> {
        const response = await this.http.get(`${this.ROOT_PATH}/clubs`);
        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
    /** creates squad based on exist club id */
    public async createSquad(clubId: string): Promise<string> {
        const response = await this.http.post(`${this.ROOT_PATH}/clubs/${clubId}/squads`);
        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
    /** adds card to sqadCards array */
    public async addCard(path: CardEditIdentificators): Promise<void> {
        const response = await this.http.post(
            `${this.ROOT_PATH}/clubs/${path.clubId}/squads/${path.squadId}/cards/${path.cardId}`,
            JSON.stringify({ position: path.position })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** changes card position inside squadCards array */
    public async changeCardPosition(path: CardEditIdentificators): Promise<void> {
        const response = await this.http.patch(
            `${this.ROOT_PATH}/clubs/${path.clubId}/squads/${path.squadId}/cards/${path.cardId}`,
            JSON.stringify({ position: path.position })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** deletes card from squadCards array */
    public async deleteCard(path: CardEditIdentificators): Promise<void> {
        const response = await this.http.delete(
            `${this.ROOT_PATH}/clubs/${path.clubId}/squads/${path.squadId}/cards/${path.cardId}`
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** updates squad tactic */
    public async updateTactic(squad: Squad, tactic: number): Promise<void> {
        const { captainId, clubId, id } = squad;
        const response = await this.http.patch(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${id}`,
            JSON.stringify({ tactic, captainId })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** updates squad captain */
    public async updateCaptain(squad: Squad, captainId: string): Promise<void> {
        const { tactic, clubId, id } = squad;
        const response = await this.http.patch(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${id}`,
            JSON.stringify({ tactic, captainId })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** updates squad formation */
    public async updateFormation(squad: Squad, formation: number): Promise<void> {
        const { clubId, id } = squad;
        const response = await this.http.put(
            `${this.ROOT_PATH}/clubs/${clubId}/squads/${id}/formation/${formation}`
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** chandes active club */
    public async changeActiveClub(id: string): Promise<void> {
        const response = await this.http.patch(
            `${this.ROOT_PATH}/clubs/${id}`,
            JSON.stringify({ 'status': 1 })
        );
        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
