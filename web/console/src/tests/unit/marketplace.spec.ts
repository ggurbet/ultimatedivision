// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import configureStore from 'redux-mock-store';
import { afterEach, beforeEach, describe, expect, it } from '@jest/globals';
import { useDispatch, useSelector } from "react-redux";
import { cleanup } from "@testing-library/react";

import { MarketplaceClient } from '@/api/marketplace';
import { GET_SELLING_CARDS } from '@/app/store/actions/marketplace';
import { Lot, MarketPlacePage } from '@/marketplace';
import { Card } from '@/card';

const marketplaceClient = new MarketplaceClient();

const mockStore = configureStore();

const successFetchMock = async (body: any) => {
    globalThis.fetch = () =>
        Promise.resolve({
            json: () => Promise.resolve(body),
            ok: true,
            status: 200,
        }) as Promise<Response>;
};

const failedFetchMock = async () => {
    globalThis.fetch = () => {
        throw new Error();
    };
};

const mockedGlobalFetch = globalThis.fetch;

const SELECTED_PAGE = 1;
const DEFAULT_VALUE = 0;

/** Mock lot. */
const MOCK_LOT: Lot = {
    cardId: '23056596-a25b-4580-a719-dd7ac13b79bb',
    type: '',
    userId: '27056596-a25b-4580-a719-dd7ac13b79bb',
    shopperId: '27056596-a25b-8880-a719-dd7ac13b79bb',
    status: '',
    startPrice: DEFAULT_VALUE,
    maxPrice: DEFAULT_VALUE,
    currentPrice: DEFAULT_VALUE,
    startTime: '',
    endTime: '',
    period: DEFAULT_VALUE,
    card: new Card(),
}

/** Mock divisions state. */
const MOCK_DIVISIONS_STATE = {
    marketplacePage: {
        lots: [new Lot(MOCK_LOT)],
        page: {
            offset: 0,
            limit: 0,
            currentPage: 0,
            pageCount: 0,
            totalCount: 0,
        }
    },
    card: new Card()
}

/** Mock initial networks state. */
const initialState = {
    marketplaceReducer: {
        marketplacePage: new MarketPlacePage([new Lot(MOCK_LOT)], {
            offset: 0,
            limit: 0,
            currentPage: 0,
            pageCount: 0,
            totalCount: 0,
        }),
        card: new Card(),
    }
};

const reactRedux = { useDispatch, useSelector }
const useDispatchMock = jest.spyOn(reactRedux, "useDispatch");
const useSelectorMock = jest.spyOn(reactRedux, "useSelector");
let updatedStore: any = mockStore(initialState);
const mockDispatch = jest.fn();
useDispatchMock.mockReturnValue(mockDispatch);
updatedStore.dispatch = mockDispatch;

describe('Requests getting selling cards.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_DIVISIONS_STATE.marketplacePage);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('Requests getting selling cards.', async () => {
        const divisionSeasons = await marketplaceClient.list(SELECTED_PAGE);
        expect(divisionSeasons).toEqual(MOCK_DIVISIONS_STATE.marketplacePage);
    });

    describe('Failed response.', () => {
        beforeEach(() => {
            failedFetchMock();
            useSelectorMock.mockClear();
            useDispatchMock.mockClear();
        });

        afterEach(() => {
            globalThis.fetch = mockedGlobalFetch;
            cleanup();
        });

        it('Must be no cards', async () => {
            try {
                await marketplaceClient.list(SELECTED_PAGE);
            } catch (error) {
                mockDispatch(GET_SELLING_CARDS, {});
                expect(updatedStore.getState().marketplaceReducer.marketplacePage).toEqual(initialState.marketplaceReducer.marketplacePage);
            }
        });
    })
});