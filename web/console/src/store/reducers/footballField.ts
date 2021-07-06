/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { FootballField } from '../../types/footballField';
import { Card } from '../../store/reducers/footballerCard'

const FieldSetup = new FootballField()

const Formation = 'Formation';
const Tactics = 'Cactics';
const Captain = 'Captain';
const ChoseCard = 'ChoseCard';
const AddCard = 'AddCard';
const RemoveCard = 'RemoveCard';



//Chose type of cards positioning on football field
export const handleFormations = (option: string) => {
    return {
        type: Formation,
        action: option
    }
};

//Adding into cardList in reducer
export const handleCard = (card: Card, index: string | null) => {
    return {
        type: AddCard,
        action: [card, index]
    }
}

//Selection position of card which should be added
export const choseCardPosition = (index: string) => {
    return {
        type: ChoseCard,
        action: index
    }
}

export const handleTactics = (option: string) => {
    return {
        type: Tactics,
        action: option
    }
};

export const handleCaptain = (option: string) => {
    return {
        type: Captain,
        action: option
    }
};


export const fieldReducer = (cardState = FieldSetup, action: any) => {

    switch (action.type) {
        case Formation:
            cardState.options.formation = action.action;
            return cardState;
        case ChoseCard:
            cardState.options.chosedCard = action.action
            return cardState;
        case AddCard:
            const ListOfCards = {
                ...cardState,
            }
            ListOfCards.cardsList[action.action[1]].cardData = action.action[0];
            return ListOfCards;
        default:
            return cardState;
    }
};
