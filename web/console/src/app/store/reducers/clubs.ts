// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ClubState } from '@/club';

import {
    ADD_CARD,
    CARD_POSITION,
    CREATE_CLUB,
    DRAG_START,
    SELECTION_VISIBILITY,
} from '@/app/store/actions/clubs';

/** TODO: replace by initial object */
const clubState = new ClubState();

export const clubsReducer = (state = clubState, action: any = {}) => {
    const options = state.options;
    const cards = state.squadCards;

    switch (action.type) {
    case CREATE_CLUB:
        state = Object.assign(clubState, action.club);
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
    case DRAG_START:
        options.dragStart = action.index;
        break;
    default:
        break;
    }

    return state;
};
