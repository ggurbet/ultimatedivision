// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { WebSocketAPICLient } from './webSocketClient';

/**
 * QueueClient is a ws implementation of users API.
 * Exposes queue-related functionality.
 */
export class QueueClient extends WebSocketAPICLient {
    /** sends action to confirm and reject match, finish search */
    public sendAction(action: string, squadId: string) {
        this.ws.send(JSON.stringify({ action, squadId }));
    };

    /** starts searching match on first open webSocket connection. */
    public startSearch(action: string, squadId: string) {
        this.ws.onopen = () => {
            this.sendAction(action, squadId);
        };
    };
};
