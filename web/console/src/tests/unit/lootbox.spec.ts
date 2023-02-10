// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import configureStore from 'redux-mock-store';
import { afterEach, beforeEach, describe, expect, it, beforeAll } from '@jest/globals';
import { useDispatch, useSelector } from "react-redux";
import { cleanup } from "@testing-library/react";

import { LootboxClient } from '@/api/lootboxes';
import { LootboxService } from '@/lootbox/service';
import { BUY_LOOTBOX } from '@/app/store/actions/lootboxes'
import { Lootbox } from '@/lootbox';
import { LootboxTypes } from '@/app/types/lootbox';
import { Card } from '@/card';

const mockStore = configureStore();
const lootboxClient = new LootboxClient();

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

/** Mock initial networks state. */
const initialState = {
    lootBoxReducer: {
        lootboxService: new LootboxService(lootboxClient),
        lootbox: []
    }
};

/** Mock regular box. */
const MOCK_REGULAR_LOOTBOX =
    new Lootbox(
        '00000000-0000-0000-0000-000000000000',
        LootboxTypes['Regular Box']
    )
    ;

const MOCK_REGULAR_BOX_RESPONCE = [
    new Card(),
    new Card(),
    new Card(),
    new Card(),
    new Card(),
]

const reactRedux = { useDispatch, useSelector }
const useDispatchMock = jest.spyOn(reactRedux, "useDispatch");
const useSelectorMock = jest.spyOn(reactRedux, "useSelector");
let updatedStore: any = mockStore(initialState);
const mockDispatch = jest.fn();
useDispatchMock.mockReturnValue(mockDispatch);
updatedStore.dispatch = mockDispatch;

describe('Requests user.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_REGULAR_BOX_RESPONCE);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('Regular box', async () => {
        lootboxClient.buy(MOCK_REGULAR_LOOTBOX)
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

        it('Must be empty user', async () => {
            try {
                await lootboxClient.buy(MOCK_REGULAR_LOOTBOX);
            } catch (error) {
                mockDispatch(BUY_LOOTBOX, {});
                expect(updatedStore.getState().lootBoxReducer.lootbox).toEqual([]);
            }
        });
    })
});
