import { createStore, combineReducers } from 'redux';

import { cardReducer } from './reducers/footballerCard';

const reducer = combineReducers({
    footballerCard: cardReducer,
});

export const store = createStore(reducer);

export type RootState = ReturnType<typeof store.getState>
