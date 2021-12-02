// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { AboutMenu } from '@components/common/AboutMenu';
import { Notification } from '@components/common/Notification';
import Preloader from '@components/common/Preloader';

import { Routes } from '@/app/routes';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<Preloader />}>
            <BrowserRouter basename="/">
                <Notification />
                <AboutMenu />
                <Routes />
            </BrowserRouter>
        </Suspense>
    );
}

export default App;
