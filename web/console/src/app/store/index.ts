import { combineReducers, createStore } from 'redux';

import { cardReducer } from '@store/reducers/footballerCard';
import { fieldReducer } from '@store/reducers/footballField';

const reducer = combineReducers({
    cardReducer,
    fieldReducer,
});

export const store = createStore(reducer);

export type RootState = ReturnType<typeof store.getState>;
