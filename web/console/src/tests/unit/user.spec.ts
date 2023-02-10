// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import configureStore from 'redux-mock-store';
import { afterEach, beforeEach, describe, expect, it } from '@jest/globals';
import { useDispatch, useSelector } from "react-redux";
import { cleanup } from "@testing-library/react";

import { UsersClient } from '@/api/users';
import { CasperNetworkClient } from '@/api/casper';
import { SET_USER } from '@/app/store/actions/users';
import { User } from '@/users';
import { UsersService } from '@/users/service';

const usersClient = new UsersClient();
const casperClient = new CasperNetworkClient();

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

/** Mock initial networks state. */
const initialState = {
    usersReducer: {
        user: {},
        userService: new UsersService(usersClient)
    }
};

/** Mock casper user info. */
const MOCK_USER: User =
    new User(
        "330cab8f2f1a2981eb8d5f4e91ebb7d11325dca70e43a3e89d9957e884c494f0",
        "0202def02b66279e3c4a93f484716077ae0196d78bb6e681898785bd7faceb9f7749",
        "test@test.com",
        "00000000-0000-0000-0000-000000000000",
        "2021-12-17T00:31:52.437252Z",
        "",
        "2022-12-17T00:31:51.874508Z",
        "0x0000000000000000000000000000000000000000",
        "casper-wallet",
    );

const reactRedux = { useDispatch, useSelector }
const useDispatchMock = jest.spyOn(reactRedux, "useDispatch");
const useSelectorMock = jest.spyOn(reactRedux, "useSelector");
let updatedStore: any = mockStore(initialState);
const mockDispatch = jest.fn();
useDispatchMock.mockReturnValue(mockDispatch);
updatedStore.dispatch = mockDispatch;

describe('Requests user.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_USER);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('Casper user', async () => {
        await casperClient.register(MOCK_USER.casperWallet, MOCK_USER.casperWalletHash)
        const user = await usersClient.getUser();

        expect(user).toEqual(MOCK_USER);
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
                await usersClient.getUser();
            } catch (error) {
                mockDispatch(SET_USER, {});
                expect(updatedStore.getState().usersReducer.user).toEqual({});
            }
        });
    })
});

