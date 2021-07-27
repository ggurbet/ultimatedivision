import { combineReducers, createStore } from 'redux';

import { cardReducer } from '@/app/store/reducers/footballerCard';
import { fieldReducer } from '@/app/store/reducers/footballField';

const reducer = combineReducers({
    cardReducer,
    fieldReducer,
});

export const store = createStore(reducer);

export type RootState = ReturnType<typeof store.getState>;
