import { createStore, combineReducers } from 'redux';

import { roadmapReducer } from './reducers/roadmap';
import { capabilitiesReducer } from './reducers/capabilities';
import { cardReducer } from './reducers/footballerCard';

const reducer = combineReducers({
    roadmap: roadmapReducer,
    capabilities: capabilitiesReducer,
    footballerCard: cardReducer
});

export const store = createStore(reducer);
