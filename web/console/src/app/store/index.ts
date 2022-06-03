// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import thunk from "redux-thunk";
import { applyMiddleware, combineReducers, createStore } from "redux";
import { cardsReducer } from "@/app/store/reducers/cards";
import { clubsReducer } from "@/app/store/reducers/clubs";
import { lootboxReducer } from "@/app/store/reducers/lootboxes";
import { marketplaceReducer } from "@/app/store/reducers/marketplace";
import { divisionsReducer } from "@/app/store/reducers/divisions";
import { matchesReducer } from "@/app/store/reducers/matches";
import { usersReducer } from "./reducers/users";

const reducer = combineReducers({
    cardsReducer,
    clubsReducer,
    lootboxReducer,
    marketplaceReducer,
    usersReducer,
    divisionsReducer,
    matchesReducer,
});

export const store = createStore(reducer, applyMiddleware(thunk));

export type RootState = ReturnType<typeof store.getState>;
