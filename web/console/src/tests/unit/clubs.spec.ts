// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import configureStore from 'redux-mock-store';
import { afterEach, beforeEach, describe, expect, it } from '@jest/globals';
import { useDispatch, useSelector } from "react-redux";
import { cleanup } from "@testing-library/react";

import { ClubsClient } from '@/api/club';
import { SET_CLUBS } from '@/app/store/actions/clubs';
import { Options } from '@/club';

import { Card } from '@/card';

const clubsClient = new ClubsClient();

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

const DEFAULT_VALUE = 0;
const ACTIVE_STATUS_VALUE = 1;

/** Mock squad. */
const MOCK_SQUAD = {
    id: '22222222-0000-0000-0000-000000000000',
    clubId: '11111111-0000-0000-0000-000000000000',
    formation: DEFAULT_VALUE,
    tactic: DEFAULT_VALUE,
    captainId: '00000000-0000-0000-0000-000000000000',
}

/** Mock squad card. */
const MOCK_SQUAD_CARD = {
    squadId: '22222222-0000-0000-0000-000000000000',
    card: new Card(),
    position: DEFAULT_VALUE,
}

/** Mock club. */
const MOCK_CLUB = {
    id: '11111111-0000-0000-0000-000000000000',
    name: 'Club 1',
    createdAt: '2023-02-07T01:13:52.114Z',
    squad: MOCK_SQUAD,
    squadCards: [MOCK_SQUAD_CARD],
    status: ACTIVE_STATUS_VALUE
}

/** Mock initial networks state. */
const initialState = {
    clubsReducer: {
        clubs: [MOCK_CLUB],
        activeClub: MOCK_CLUB,
        options: new Options(),
        isSearchingMatch: false,
    }
};

const reactRedux = { useDispatch, useSelector }
const useDispatchMock = jest.spyOn(reactRedux, "useDispatch");
const useSelectorMock = jest.spyOn(reactRedux, "useSelector");
let updatedStore: any = mockStore(initialState);
const mockDispatch = jest.fn();
useDispatchMock.mockReturnValue(mockDispatch);
updatedStore.dispatch = mockDispatch;

describe('Requests list of clubs.', () => {
    beforeEach(() => {
        successFetchMock([MOCK_CLUB]);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('Requests list clubs.', async () => {
        const clubs = await clubsClient.getClubs();
        expect(clubs).toEqual([MOCK_CLUB]);
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

        it('Must be one club', async () => {
            try {
                await clubsClient.getClubs();
            } catch (error) {
                mockDispatch(SET_CLUBS, {});
                expect(updatedStore.getState().clubsReducer.clubs).toEqual([MOCK_CLUB]);
            }
        });
    })
});
