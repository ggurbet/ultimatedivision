// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import configureStore from 'redux-mock-store';
import { afterEach, beforeEach, describe, expect, it } from '@jest/globals';
import { useDispatch, useSelector } from "react-redux";
import { cleanup } from "@testing-library/react";

import { CardsClient } from '@/api/cards';
import { page } from '@/app/store/reducers/cards';
import { GET_USER_CARDS, USER_CARD } from '@/app/store/actions/cards';
import { Card, CardsPage } from '@/card';

const cardsClient = new CardsClient();

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

/** Mock user card. */
const MOCK_CARD = new Card(
    {
        id: "27056596-a25b-4580-a719-dd7ac13b79bb",
        playerName: "Miquel Deloach",
        quality: "silver",
        pictureType: 0,
        height: 176.62,
        weight: 67.5,
        skinColor: 0,
        hairStyle: 0,
        hairColor: 0,
        accessories: [],
        dominantFoot: "right",
        isTattoos: false,
        status: 0,
        type: "won",
        userId: "c4b97f28-d314-4b60-a2dd-26a9f73ce66e",
        tactics: 49,
        positioning: 41,
        composure: 56,
        aggression: 52,
        vision: 58,
        awareness: 54,
        crosses: 42,
        physique: 7,
        acceleration: 11,
        runningSpeed: 1,
        reactionSpeed: 12,
        agility: 1,
        stamina: 16,
        strength: 8,
        jumping: 7,
        balance: 13,
        technique: 51,
        dribbling: 59,
        ballControl: 59,
        weakFoot: 47,
        skillMoves: 41,
        finesse: 57,
        curve: 44,
        volleys: 45,
        shortPassing: 49,
        longPassing: 45,
        forwardPass: 60,
        offense: 0,
        finishingAbility: 12,
        shotPower: 10,
        accuracy: 2,
        distance: 1,
        penalty: 7,
        freeKicks: 1,
        corners: 3,
        headingAccuracy: 1,
        defence: 57,
        offsideTrap: 58,
        sliding: 62,
        tackles: 64,
        ballFocus: 61,
        interceptions: 57,
        vigilance: 58,
        goalkeeping: 52,
        reflexes: 50,
        diving: 49,
        handling: 42,
        sweeping: 47,
        throwing: 56,

    }
)

/** Mock cards page. */
const MOCK_CARDS_PAGE = new CardsPage([MOCK_CARD], page);

/** Mock cards list. */
const MOCK_CARD_LIST = {
    cards: [
        MOCK_CARD,
        MOCK_CARD,
        MOCK_CARD
    ]
}

/** Mock initial networks state. */
const initialState = {
    cardsReducer: {
        MOCK_CARDS_PAGE: MOCK_CARDS_PAGE,
        card: new Card(),
        currentCardsPage: page.currentPage,
        currentFieldCardsPage: page.currentPage,
    }
};

const reactRedux = { useDispatch, useSelector }
const useDispatchMock = jest.spyOn(reactRedux, "useDispatch");
const useSelectorMock = jest.spyOn(reactRedux, "useSelector");
let updatedStore: any = mockStore(initialState);
const mockDispatch = jest.fn();
useDispatchMock.mockReturnValue(mockDispatch);
updatedStore.dispatch = mockDispatch;

describe('Requests list of cards.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_CARD_LIST);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('Requests list of user card.', async () => {
        const cards = await cardsClient.list(SELECTED_PAGE);
        expect(cards).toEqual(MOCK_CARD_LIST);
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
                await cardsClient.list(page.currentPage);
            } catch (error) {
                mockDispatch(GET_USER_CARDS, {});
                expect(updatedStore.getState().cardsReducer.MOCK_CARDS_PAGE).toEqual(MOCK_CARDS_PAGE);
            }
        });
    })
});

describe('Requests user card.', () => {
    beforeEach(() => {
        successFetchMock(MOCK_CARD);
    });

    afterEach(() => {
        globalThis.fetch = mockedGlobalFetch;
    });

    it('User card', async () => {
        const card = await cardsClient.getCardById('11111111-0000-0000-0000-000000000000');
        expect(card).toEqual(MOCK_CARD);

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

        it('Must be empty card', async () => {
            try {
                await cardsClient.getCardById('00000000-0000-0000-0000-000000000000');
            } catch (error) {
                mockDispatch(USER_CARD, {});
                expect(updatedStore.getState().cardsReducer.card).toEqual(new Card());
            }
        });
    })
});
