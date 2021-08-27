// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Card } from '@/card';

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

/** Chose type of cards positioning on football field */
export const setFormation = (option: string) => ({
    type: FORMATION,
    action: option,
});

export const cardSelectionVisibility = (option: boolean) => ({
    type: SELECTION_VISIBILITY,
    action: option,
});

/** Adding into cardList in reducer */
export const addCard = (card: Card, index: number) => ({
    type: ADD_CARD,
    action: [card, index],
});

export const removeCard = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: REMOVE_CARD,
    action: index,
});

/** Selection position of card which should be added */
export const choseCardPosition = (index: number) => ({
    type: CARD_POSITION,
    action: index,
});

export const setDragStart = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_START,
    action: index,
});

export const setDragTarget = (index: dragParamType = DEFAULT_CARD_INDEX) => ({
    type: DRAG_TARGET,
    action: index,
});

export const exchangeCards = (prevPosition: dragParamType, currentPosition: dragParamType) => ({
    type: EXCHANGE_CARDS,
    action: [prevPosition, currentPosition],
});

export const setTactic = (option: string) => ({
    type: TACTICS,
    action: option,
});

export const setCaptain = (option: string) => ({
    type: CAPTAIN,
    action: option,
});
