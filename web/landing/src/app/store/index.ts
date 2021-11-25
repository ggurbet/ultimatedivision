// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { applyMiddleware, createStore, combineReducers } from 'redux';
import thunkMiddleware from 'redux-thunk';

import { usersReducer } from './reducers/users';

const reducer = combineReducers({
    users: usersReducer,
});

export const store = createStore(reducer, applyMiddleware(thunkMiddleware));

export type RootState = ReturnType<typeof store.getState>;