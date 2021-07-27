/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { Navbar } from '@components/Navbar';
import { Routes } from '@/app/routes';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            {/** TODO: LoadingPage */}
            <BrowserRouter>
                <Navbar />
                <Routes />
            </BrowserRouter>
        </Suspense>

    );
}

export default App;
