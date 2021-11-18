// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardEditIdentificators, ClubsClient } from '@/api/club';
import Card from '@/app/views/CardPage';
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

export const SET_CLUBS = 'SET_CLUBS';
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

export const setClubs = (clubs: Club[]) => ({
    type: SET_CLUBS,
    clubs,
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

export const createClubs = () =>
    async function(dispatch: Dispatch) {
        const clubId = await service.createClub();
        const squadId = await service.createSquad(clubId);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const getClubs = () =>
    async function(dispatch: Dispatch) {
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const setFormation = (squad: Squad, formation: FormationsType) =>
    async function(dispatch: Dispatch) {
        await service.updateFormation(squad, Formations[formation]);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const setCaptain = (squad: Squad, captainId: string) =>
    async function(dispatch: Dispatch) {
        await service.updateCaptain(squad, captainId);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const setTactic = (squad: Squad, tactic: TacticsType) =>
    async function(dispatch: Dispatch) {
        await service.updateTactic(squad, Tactic[tactic]);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const addCard = (path: CardEditIdentificators) =>
    async function(dispatch: Dispatch) {
        await service.addCard(path);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const deleteCard = (path: CardEditIdentificators) =>
    async function(dispatch: Dispatch) {
        await service.deleteCard(path);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const changeCardPosition = (cardItentificators: CardEditIdentificators) =>
    async function(dispatch: Dispatch) {
        await service.changeCardPosition(cardItentificators);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const swapCards = (currentCard: CardEditIdentificators, existCard: CardEditIdentificators) =>
    async function(dispatch: Dispatch) {
        await service.changeCardPosition(currentCard);
        await service.changeCardPosition(existCard);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };

export const changeActiveClub = (id: string) =>
    async function(dispatch: Dispatch) {
        await service.changeActiveClub(id);
        const clubs = await service.getClubs();
        dispatch(setClubs(clubs));
    };
