// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardEditIdentificators, ClubsClient } from '@/api/club';
import { Club, Squad } from '.';

/**
 * exposes all bandwidth related logic
 */
export class ClubService {
    protected readonly club: ClubsClient;

    /** sets ClubClient into club field */
    public constructor(club: ClubsClient) {
        this.club = club;
    };

    /** creates club */
    public async createClub(): Promise<string> {
        return await this.club.createClub();
    };

    /** returns club with existing squads */
    public async getClubs(): Promise<Club[]> {
        return await this.club.getClubs();
    };

    /** creates squad in selected club */
    public async createSquad(clubId: string): Promise<string> {
        return await this.club.createSquad(clubId);
    };

    /** adds card to squad cards list */
    public async addCard(path: CardEditIdentificators): Promise<void> {
        return await this.club.addCard(path);
    };

    /** changes position of existing card */
    public async changeCardPosition(path: CardEditIdentificators): Promise<void> {
        return await this.club.changeCardPosition(path);
    };

    /** deletes card from squad cards list */
    public async deleteCard(path: CardEditIdentificators): Promise<void> {
        return await this.club.deleteCard(path);
    };

    /** updates squad tactic */
    public async updateTactic(squad: Squad, tactic: number): Promise<void> {
        return await this.club.updateTactic(squad, tactic);
    }

    /** updates squad captain */
    public async updateCaptain(squad: Squad, captainId: string): Promise<void> {
        return await this.club.updateCaptain(squad, captainId);
    }

    /** updates squad formation */
    public async updateFormation(squad: Squad, formation: number): Promise<void> {
        return await this.club.updateFormation(squad, formation);
    }
    /** chandes active club */
    public async changeActiveClub(id: string): Promise<void> {
        return await this.club.changeActiveClub(id);
    }
};
