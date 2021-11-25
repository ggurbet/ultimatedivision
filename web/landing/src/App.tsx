// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { Routes } from '@/app/router';

export const App = () => (
    <BrowserRouter basename="/landing/">
        {/** TODO: LoadingPage or indicator*/}
        <Suspense fallback={<div>Loading...</div>} >
            <Routes />
        </Suspense>
    </BrowserRouter>
);
