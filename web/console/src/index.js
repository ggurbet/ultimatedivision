/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';

import { App } from './App';

import { store } from './store';

import './index.scss';

ReactDOM.render(
    <Provider store={store}>
        <BrowserRouter>
            <App />
        </BrowserRouter>
    </Provider>,
    document.getElementById('root'),
);
