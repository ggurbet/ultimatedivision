// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import thunk from 'redux-thunk';
import { applyMiddleware, combineReducers, createStore } from 'redux';
import { cardsReducer } from '@/app/store/reducers/cards';
import { clubReducer } from '@/app/store/reducers/club';
import { lootboxReducer } from '@/app/store/reducers/lootboxes';
import { marketplaceReducer } from '@/app/store/reducers/marketplace';

const reducer = combineReducers({
    cardsReducer,
    clubReducer,
    lootboxReducer,
    marketplaceReducer,
});

export const store = createStore(reducer, applyMiddleware(thunk));

export type RootState = ReturnType<typeof store.getState>;
