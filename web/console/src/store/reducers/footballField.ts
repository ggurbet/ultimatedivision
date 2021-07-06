/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { FootballField } from '../../types/footballField';

const FieldSetup = new FootballField()

const Formation = 'Formation';
const Tactics = 'Cactics';
const Captain = 'Captain';


export const handleFormations = (option: string) => {
    return {
        type: Formation,
        action: option
    }
};

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


export const fieldReducer = (cardState = FieldSetup, action: {
    type: string,
    action: string
}) => {

    switch (action.type) {
        case Formation:
            cardState.options.formation = action.action;
            return cardState;
        default:
            return cardState;
    }
};
