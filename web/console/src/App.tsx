// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { Navbar } from '@components/common/Navbar';
import { Routes } from '@/app/router';
import { AboutMenu } from '@components/common/AboutMenu';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            {/** TODO: LoadingPage */}
            <BrowserRouter basename="/ud">
                <Navbar />
                <AboutMenu />
                <Routes />
            </BrowserRouter>
        </Suspense>

    );
}

export default App;
