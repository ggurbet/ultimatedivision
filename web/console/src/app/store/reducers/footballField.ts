// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballField } from '@/app/types/footballField';

import {
    ADD_CARD,
    CAPTAIN,
    CARD_POSITION,
    DRAG_START,
    DRAG_TARGET,
    EXCHANGE_CARDS,
    FORMATION,
    REMOVE_CARD,
    SELECTION_VISIBILITY,
    TACTICS,
} from '@/app/store/actions/footballField';

const FieldSetup = new FootballField();


const FITST_ACTION_PARAM = 0;
const SECOND_ACTION_PARAM = 1;


export const fieldReducer = (cardState = FieldSetup, action: any = {}) => {
    const options = cardState.options;
    const cardsList = cardState.cardsList;

    switch (action.type) {
    case FORMATION:
        options.formation = action.action;
        break;
    case SELECTION_VISIBILITY:
        options.showCardSeletion = action.action;
        break;
    case CARD_POSITION:
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
