// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubState } from '@/club';

import {
    ADD_CARD,
    CAPTAIN,
    CARD_POSITION,
    CREATE_CLUB,
    DRAG_START,
    DRAG_TARGET,
    EXCHANGE_CARDS,
    FORMATION,
    REMOVE_CARD,
    SELECTION_VISIBILITY,
    TACTICS,
} from '@/app/store/actions/club';

/** TODO: replace by initial object */
const clubState = new ClubState();

export const clubReducer = (state = clubState, action: any = {}) => {
    const options = state.options;
    const squad = state.squad;
    const cards = state.squadCards;

    switch (action.type) {
    case CREATE_CLUB:
        state = Object.assign(clubState, action.club);
        break;

    // next cases will be replaced
    case FORMATION:
        squad.formation = action.formation;
        break;
    case SELECTION_VISIBILITY:
        options.showCardSeletion = action.isVisible;
        break;
    case CARD_POSITION:
        options.chosedCard = action.index;
        break;
    case ADD_CARD:
        cards[action.fieldCard.index].cardId = action.fieldCard.card.cardId;
        break;
    case REMOVE_CARD:
        cards[action.index].cardId = '';
        break;
    case DRAG_START:
        options.dragStart = action.index;
        break;
    case DRAG_TARGET:
        options.dragTarget = action.index;
        break;
    case EXCHANGE_CARDS:
        const prevCard = cards[action.position.previous];
        cards[action.position.previous] = cards[action.position.current];
        cards[action.position.current] = prevCard;
        break;
    default:
        break;
    }

    return state;
};
