// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Suspense, useEffect, useState } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { AboutMenu } from '@components/common/AboutMenu';
import { Notification } from '@components/common/Notification';
import Preloader from '@components/common/Preloader';

import { Routes } from '@/app/routes';
import { useLocalStorage } from './app/hooks/useLocalStorage';
import { getCurrentWebSocketClient, onOpenConnection } from '@/webSockets/service';
import { WebSocketClient } from './api/websockets';

/** initial App setup */
export function App() {
    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();
    const [webSocketClient, setWebSocketClient] = useState<WebSocketClient | null>(null);

    /** Indicates if user is logined in app. */
    // @ts-ignore .
    const isLoggedIn = JSON.parse(getLocalStorageItem('IS_LOGGINED'));

    /** Closes web sockect connection. */
    const closeWebSocketConnection = (e: any) => {
        e.preventDefault();

        /** Updates current websocket client. */
        const newclient = getCurrentWebSocketClient();
        setWebSocketClient(newclient);

        if (isLoggedIn && webSocketClient) {
            webSocketClient.close();
        }
    };

    useEffect(() => {
        window.addEventListener('beforeunload', closeWebSocketConnection);

        return () => {
            window.removeEventListener('beforeunload', closeWebSocketConnection);
        };
    }, []);

    useEffect(() => {
        if (isLoggedIn) {
            onOpenConnection();
        }
    }, []);

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
