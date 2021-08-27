// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { applyMiddleware, combineReducers, createStore } from 'redux';
import thunk from 'redux-thunk';

import { cardsReducer } from '@/app/store/reducers/cards';
import { fieldReducer } from '@/app/store/reducers/footballField';

const reducer = combineReducers({
    cardsReducer,
    fieldReducer,
});

export const store = createStore(reducer, applyMiddleware(thunk));

export type RootState = ReturnType<typeof store.getState>;
