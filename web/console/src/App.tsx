/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/
import React from 'react';
import { BrowserRouter } from 'react-router-dom';

import { Navbar }
    from './components/Navbar/Navbar';
import { Routes } from './routes/index';

import './App.scss';

/** initial App setup */
export function App() {
    return (
        <BrowserRouter>
            <Navbar />
            <Routes />
        </BrowserRouter>
    );
}

export default App;
