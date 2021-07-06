import { createStore, combineReducers } from 'redux';

import { cardReducer } from './reducers/footballerCard';
import { fieldReducer } from './reducers/footballField';

const reducer = combineReducers({
    cardReducer,
    fieldReducer
});

export const store = createStore(reducer);

export type RootState = ReturnType<typeof store.getState>
