// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import { App } from '@/App';

import { store } from '@/app/store';

import './index.scss';

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById('root'),
);
