// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

import { CurrentDivisionSeasons, DivisionSeasonsStatistics } from '@/divisions';

/** DivisionsClient base implementation */
export class DivisionsClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0';

    /** gets divisions of current seasons */
    public async getCurrentDivisionSeasons(): Promise<CurrentDivisionSeasons[]> {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/current`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        const currentDivisionSeasons = await response.json();

        return currentDivisionSeasons.map(
            (season: CurrentDivisionSeasons, index: number) =>
                new CurrentDivisionSeasons(
                    season.id,
                    season.divisionId,
                    season.startedAt,
                    season.endedAt
                )
        );
    }

    /** gets division seasons statistics */
    public async getDivisionSeasonsStatistics(id: string): Promise<DivisionSeasonsStatistics> {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/statistics/division/${id}`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        const responseData = await response.json();

        return new DivisionSeasonsStatistics(
            responseData.division,
            responseData.statistics
        );
    }
}
