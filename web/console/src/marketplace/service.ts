// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MarketplaceClient } from '@/api/marketplace';
import { CardsQueryParametersField } from '@/card';
import { Lot, MarketPlacePage } from '@/marketplace';
import { CreatedLot } from '@/app/types/marketplace';
import { OfferTransaction } from '@/casper/types';

/**
 * exposes all arketplace domain entity related logic
 */
export class Marketplaces {
    protected readonly marketplace: MarketplaceClient;

    /** default marketplaceClient implementation */
    public constructor(marketplace: MarketplaceClient) {
        this.marketplace = marketplace;
    };

    /** Returns current lots queryParameters object. */
    public getCurrentQueryParameters() {
        return this.marketplace.queryParameters;
    };

    /** Changes lots query parameters. */
    public changeLotsQueryParameters(queryParameters: CardsQueryParametersField[]) {
        this.marketplace.changeLotsQueryParameters(queryParameters);
    };

    /** returns marketplace domain entity with list of lots */
    public async list(selectedPage: number): Promise<MarketPlacePage> {
        return await this.marketplace.list(selectedPage);
    };

    /** creates lot */
    public async createLot(lot: CreatedLot): Promise<void> {
        await this.marketplace.createLot(lot);
    };

    /** returns lot by id */
    public async getLotById(id: string): Promise<Lot> {
        return await this.marketplace.getLotById(id);
    };

    /** places a bid */
    public async placeBid(lotId: string, amount: number): Promise<void> {
        return await this.marketplace.placeBid(lotId, amount);
    };

    /** places a bid */
    public async endTime(lotId: string): Promise<boolean> {
        return await this.marketplace.endTime(lotId);
    };

    /** returns lot data */
    public async lotData(cardId: string): Promise<any> {
        return await this.marketplace.lotData(cardId);
    };

    /** returns make offer data */
    public async offer(cardId: string): Promise<OfferTransaction> {
        return await this.marketplace.offer(cardId);
    };
};
