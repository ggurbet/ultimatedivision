// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MarketplaceClient } from '@/api/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { Pagination } from '@/app/types/pagination';
import { CardWithStats } from '@/card';
import { Lot, MarketPlacePage } from '.';

/**
 * exposes all arketplace domain entity related logic
 */
export class Marketplaces {
    protected readonly marketplace: MarketplaceClient;

    /** default marketplaceClient implementation */
    public constructor(marketplace: MarketplaceClient) {
        this.marketplace = marketplace;
    };

    /** returns marketplace domain entity with list of lots */
    public async list({ selectedPage, limit }: Pagination): Promise<MarketPlacePage> {
        const response = await this.marketplace.list({ selectedPage, limit });

        return { ...response, lots: response.lots.map((lot: Lot) => ({ ...lot, card: new CardWithStats(lot.card) })) };
    };

    /** creates lot */
    public async createLot(lot: CreatedLot): Promise<void> {
        await this.marketplace.createLot(lot);
    };

    /** returns list of lots by filter parameters */
    public async filteredList(lowRange: string, topRange: string): Promise<MarketPlacePage> {
        const response = await this.marketplace.filteredList(lowRange, topRange);

        return { ...response, lots: response.lots.map((lot: Lot) => ({ ...lot, card: new CardWithStats(lot.card) })) };
    };

    /** returns lot by id */
    public async getLotById(id: string): Promise<Lot> {
        return await this.marketplace.getLotById(id);
    };
};
