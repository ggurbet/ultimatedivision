// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballField } from '@/app/types/footballField';
import { Card } from '@/app/types/fotballerCard';

const FieldSetup = new FootballField();
const FORMATION_TYPE = 'Formation';
const CARD_SELECTION_VISIBILITY = 'SelectionVisibility';
const TACTICS_TYPE = 'Cactics';
const CAPTAIN_TYPE = 'Captain';
const CHOSE_CARD_POSITION = 'ChoseCard';
const ADD_CARD = 'AddCard';
const REMOVE_CARD = 'RemoveCard';
const DRAG_START = 'CurrentPossition';
const DRAG_TARGET = 'DragTarget';
const EXCHANGE_CARDS = 'ReplaceCard';

const DEFAULT_CARD_INDEX = null;
const FITST_ACTION_PARAM = 0;
const SECOND_ACTION_PARAM = 1;

type dragParamType = number | null;

/** Chose type of cards positioning on football field */
export const handleFormations = (option: string) => ({
    type: FORMATION_TYPE,
    action: option,
});

export const cardSelectionVisibility = (option: boolean) => ({
    type: CARD_SELECTION_VISIBILITY,
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
    type: CHOSE_CARD_POSITION,
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

export const handleTactics = (option: string) => ({
    type: TACTICS_TYPE,
    action: option,
});

export const handleCaptain = (option: string) => ({
    type: CAPTAIN_TYPE,
    action: option,
});

export const fieldReducer = (cardState = FieldSetup, action: any = {}) => {
    const options = cardState.options;
    const cardsList = cardState.cardsList;

    switch (action.type) {
    case FORMATION_TYPE:
        options.formation = action.action;
        break;
    case CARD_SELECTION_VISIBILITY:
        options.showCardSeletion = action.action;
        break;
    case CHOSE_CARD_POSITION:
        options.chosedCard = action.action;
        break;
    case ADD_CARD:
        cardsList[action.action[SECOND_ACTION_PARAM]].cardData = action.action[FITST_ACTION_PARAM];
        break;
    case REMOVE_CARD:
        cardsList[action.action].cardData = null;
        break;
    case DRAG_START:
        options.dragStart = action.action;
        break;
    case DRAG_TARGET:
        options.dragTarget = action.action;
        break;
    case EXCHANGE_CARDS:
        const prevValue = cardsList[action.action[FITST_ACTION_PARAM]];
        cardsList[action.action[FITST_ACTION_PARAM]] = cardsList[action.action[SECOND_ACTION_PARAM]];
        cardsList[action.action[SECOND_ACTION_PARAM]] = prevValue;
        break;
    default:
        break;
    }

    return { ...cardState };
};
