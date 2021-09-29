// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.


import { Lot, MarketPlacePage } from '@/marketplace';

import { CreatedLot } from '@/app/types/marketplace';
import { Pagination } from '@/app/types/pagination';
import { APIClient } from '.';

/** client for marketplace of api */
export class MarketplaceClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/marketplace';

    /** returns marketplace domain entity with list of lots*/
    public async list({ selectedPage, limit }: Pagination): Promise<MarketPlacePage> {
        const path = `${this.ROOT_PATH}?page=${selectedPage}&limit=${limit}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const marketplaceJSON = await response.json();

        return new MarketPlacePage(
            marketplaceJSON.lots.map((lot: Lot) => new Lot(
                lot.id,
                lot.itemId,
                lot.type,
                lot.userId,
                lot.shopperId,
                lot.status,
                lot.startPrice,
                lot.maxPrice,
                lot.currentPrice,
                lot.startTime,
                lot.endTime,
                lot.period,
                lot.card,
            )),
            marketplaceJSON.page,
        );
    };

    /** implements opening lot */
    public async getLotById(id: string): Promise<Lot> {
        const path = `${this.ROOT_PATH}/${id}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const lotJSON = await response.json();
        const lot = lotJSON.lot;

        return new Lot(
            lot.id,
            lot.itemId,
            lot.type,
            lot.userId,
            lot.shopperId,
            lot.status,
            lot.startPrice,
            lot.maxPrice,
            lot.currentPrice,
            lot.startTime,
            lot.endTime,
            lot.period,
            lot.card,
        );
    };

    /** implements creating lot (selling card) */
    public async createLot(lot: CreatedLot): Promise<void> {
        const path = `${this.ROOT_PATH}`;
        const response = await this.http.post(path, JSON.stringify(lot));

        if (!response.ok) {
            await this.handleError(response);
        };
    };

    /** method returns filtered lot list */
    public async filteredList(filterParam: string): Promise<MarketPlacePage> {
        const path = `${this.ROOT_PATH}/lots/?${filterParam}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };

        const marketplaceJSON = await response.json();

        return new MarketPlacePage(
            marketplaceJSON.lots.map((lot: Lot) => new Lot(
                lot.id,
                lot.itemId,
                lot.type,
                lot.userId,
                lot.shopperId,
                lot.status,
                lot.startPrice,
                lot.maxPrice,
                lot.currentPrice,
                lot.startTime,
                lot.endTime,
                lot.period,
                lot.card,
            )),
            marketplaceJSON.page,
        );
    };
};
