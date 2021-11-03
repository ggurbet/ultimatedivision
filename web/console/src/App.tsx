// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { Routes } from '@/app/routes';
import { AboutMenu } from '@components/common/AboutMenu';

import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.min.css';

/** initial App setup */
export function App() {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            {/** TODO: LoadingPage */}
            <BrowserRouter basename="/">
                <ToastContainer
                    position="top-right"
                    autoClose={5000}
                    hideProgressBar
                    newestOnTop={false}
                    closeOnClick={false}
                    rtl={false}
                    pauseOnFocusLoss
                    pauseOnHover
                />
                <AboutMenu />
                <Routes />
            </BrowserRouter>
        </Suspense>

    );
}

export default App;
