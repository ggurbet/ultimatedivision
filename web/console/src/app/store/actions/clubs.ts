// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardEditIdentificators, ClubsClient } from '@/api/club';
import {
    Club,
    Formations,
    FormationsType,
    Squad,
    Tactic,
    TacticsType,
} from '@/club';
import { ClubService } from '@/club/service';
import { Dispatch } from 'redux';

export const CREATE_CLUB = 'CREATE_CLUB';
export const FORMATION = 'FORMATION';
export const SELECTION_VISIBILITY = 'SELECTION_VISIBILITY';
export const TACTICS = 'TACTICS';
export const CAPTAIN = 'CAPTAIN';
export const CARD_POSITION = 'CARD_POSITION';
export const ADD_CARD = 'ADD_CARD';
export const REMOVE_CARD = 'REMOVE_CARD';
export const DRAG_START = 'DRAG_START';
export const DRAG_TARGET = 'DRAG_TARGET';
export const EXCHANGE_CARDS = 'EXCHANGE_CARDS';

type dragParamType = number | null;
const DEFAULT_CARD_INDEX = null;

const client = new ClubsClient();
const service = new ClubService(client);

export const setClub = (club: Club) => ({
    type: CREATE_CLUB,
    club,
});

export const cardSelectionVisibility = (isVisible: boolean) => ({
    type: SELECTION_VISIBILITY,
    isVisible,
});

/** Selection position of card which should be added */
export const choosePosition = (index: number) => ({
    type: CARD_POSITION,
    index,
});

export const setDragStart = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_START,
    index,
});

export const setDragTarget = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_TARGET,
    index,
});

export const exchangeCards = (
    previous: dragParamType,
    current: dragParamType
) => ({
    type: EXCHANGE_CARDS,
    position: {
        previous,
        current,
    },
});

// Thunks

export const createClub = () =>
    async function (dispatch: Dispatch) {
        const clubId = await service.createClub();
        const squadId = await service.createSquad(clubId);
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const getClub = () =>
    async function (dispatch: Dispatch) {
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const setFormation = (squad: Squad, formation: FormationsType) =>
    async function (dispatch: Dispatch) {
        await service.updateFormation(squad, Formations[formation]);
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const setCaptain = (squad: Squad, captainId: string) =>
    async function (dispatch: Dispatch) {
        await service.updateCaptain(squad, captainId);
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const setTactic = (squad: Squad, tactic: TacticsType) =>
    async function (dispatch: Dispatch) {
        await service.updateTactic(squad, Tactic[tactic]);
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const addCard = (path: CardEditIdentificators) =>
    async function (dispatch: Dispatch) {
        await service.addCard(path);
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const deleteCard = (path: CardEditIdentificators) =>
    async function (dispatch: Dispatch) {
        await service.deleteCard(path);
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const changeCardPosition = (cardItentificators: CardEditIdentificators) =>
    async function (dispatch: Dispatch) {
        await service.changeCardPosition(cardItentificators);
        const club = await service.getClub();
        dispatch(setClub(club));
    };

export const swapCards = (currentCard: CardEditIdentificators, existCard: CardEditIdentificators) =>
    async function (dispatch: Dispatch) {
        await service.changeCardPosition(currentCard);
        await service.changeCardPosition(existCard);
        const club = await service.getClub();
        dispatch(setClub(club));
    };
