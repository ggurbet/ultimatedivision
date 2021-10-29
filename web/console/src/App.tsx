// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { Routes } from '@/app/routes';
import { AboutMenu } from '@components/common/AboutMenu';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            {/** TODO: LoadingPage */}
            <BrowserRouter basename="/">
                <AboutMenu />
                <Routes />
            </BrowserRouter>
        </Suspense>

    );
}

export default App;
