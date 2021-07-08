/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { FootballField } from '../../types/footballField';
import { Card } from '../../store/reducers/footballerCard';

const FieldSetup = new FootballField();
const FORMATION_TYPE = 'Formation';
const TACTICS_TYPE = 'Cactics';
const CAPTAIN_TYPE = 'Captain';
const CHOSE_CARD_POSITION = 'ChoseCard';
const ADD_CARD = 'AddCard';
const REMOVE_CARD = 'RemoveCard';
const EXCHANGE_CARDS = 'ReplaceCard'

//Chose type of cards positioning on football field
export const handleFormations = (option: string) => {
    return {
        type: FORMATION_TYPE,
        action: option
    }
};

//Adding into cardList in reducer
export const addCard = (card: Card, index: string | null) => {
    return {
        type: ADD_CARD,
        action: [card, index]
    }
}

export const removeCard = (index: number = -1) => {
    return {
        type: REMOVE_CARD,
        action: index
    }
}

//Selection position of card which should be added
export const choseCardPosition = (index: number) => {
    return {
        type: CHOSE_CARD_POSITION,
        action: index
    }
}

export const exchangeCards = (prevPosition: number, currentPosition: number) => {
    return {
        type: EXCHANGE_CARDS,
        action: [prevPosition, currentPosition]
    }
}

export const handleTactics = (option: string) => {
    return {
        type: TACTICS_TYPE,
        action: option
    }
};

export const handleCaptain = (option: string) => {
    return {
        type: CAPTAIN_TYPE,
        action: option
    }
};

export const fieldReducer = (cardState = FieldSetup, action: any) => {

    switch (action.type) {
        case FORMATION_TYPE:
            cardState.options.formation = action.action;
            break;
        case CHOSE_CARD_POSITION:
            cardState.options.chosedCard = action.action
            break;
        case ADD_CARD:
            cardState.cardsList[action.action[1]].cardData = action.action[0];
            break;
        case REMOVE_CARD:
            cardState.cardsList[action.action].cardData = null;
            break;
            case EXCHANGE_CARDS:
            const prevValue = cardState.cardsList[action.action[0]];
            cardState.cardsList[action.action[0]] = cardState.cardsList[action.action[1]];
            cardState.cardsList[action.action[1]] = prevValue;
            break;
    }
    return {...cardState}
};
