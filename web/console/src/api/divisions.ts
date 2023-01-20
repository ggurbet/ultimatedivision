// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api/index';

import { CurrentDivisionSeasons, DivisionSeasonsStatistics, SeasonRewardTransaction } from '@/divisions';

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

    /** requests division seasons reward status */
    public async seasonsRewardStatus(): Promise<number> {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/reward/tokens`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        const seasonRewardTokenStatus = await response.json();

        return seasonRewardTokenStatus;
    }

    /** gets division seasons reward */
    public async getDivisionSeasonsReward(): Promise<SeasonRewardTransaction> {
        const response = await this.http.get(
            `${this.ROOT_PATH}/seasons/reward`
        );

        if (!response.ok) {
            await this.handleError(response);
        }

        const seasonReward = await response.json();

        return new SeasonRewardTransaction(
            seasonReward.ID,
            seasonReward.userId,
            seasonReward.seasonID,
            seasonReward.walletAddress,
            seasonReward.casperWalletAddress,
            seasonReward.CasperWalletHash,
            seasonReward.walleType,
            seasonReward.status,
            seasonReward.nonce,
            seasonReward.signature,
            seasonReward.value,
            seasonReward.casperTokenContract,
            seasonReward.rpcNodeAddress
        );
    }
}
