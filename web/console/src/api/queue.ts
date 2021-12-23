// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/**
 * QueueClient is a ws implementation of users API.
 * Exposes queue-related functionality.
 */
export class QueueClient {
    /** The WebSocket provides the API for creating and managing
    * a websocket connection to a server and for sending and
    * receiving data on the connection. */
    public ws: WebSocket = new WebSocket(`ws://${window.location.host}/api/v0/queue`);

    /** Sends action to confirm and reject match, finish search */
    public sendAction(action: string, squadId: string) {
        this.ws.send(JSON.stringify({ action, squadId }));
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public actionAllowAddress(WalletAddress: string, nonce: number) {
        const action: string = 'allowAddress';

        this.ws.send(JSON.stringify({ action, WalletAddress, nonce }));
    };

    /** Sends action that indicates that the client is forbidden to add wallet address. */
    public actionForbidAddress() {
        const action: string = 'forbidAddress';

        this.ws.send(JSON.stringify(action));
    };

    /** TODO: this will be deleted after ./queue/chore.go solution. */
    /** Sends action, i.e 'startSearch', 'finishSearch', on open webSocket connection. */
    public onOpenConnectionSendAction(action: string, squadId: string) {
        this.ws.onopen = () => {
            this.sendAction(action, squadId);
        };
    };
};
