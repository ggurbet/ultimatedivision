import { createStore, combineReducers } from 'redux';

import { cardReducer } from './reducers/footballerCard';
import { filterFieldTitlesReducer } from './reducers/filterFieldTitles';

const reducer = combineReducers({
    footballerCard: cardReducer,
    filterFieldTitles: filterFieldTitlesReducer
});

export const store = createStore(reducer);
