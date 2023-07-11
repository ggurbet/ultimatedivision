// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '@/api';

/**
 * QueueClient is a ws implementation of users API.
 * Exposes queue-related functionality.
 */
export class WebSocketClient extends APIClient {
    /** The WebSocket provides the API for creating and managing
    * a websocket connection to a server and for sending and
    * receiving data on the connection. */
    // TODO: rework functionality.
    public ws: WebSocket = new WebSocket(`${window.location.protocol === 'https:' ? 'wss' : 'ws'}://${window.location.host}/api/v0/connection`);

    public ROOT_PATH: string = '/api/v0/queue';

    /** Sets match queue */
    public async matchQueue(): Promise<void> {
        const path = `${this.ROOT_PATH}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };
    };

    /** Sends action to confirm and reject match, finish search */
    public sendAction(action: string, squadId: string) {
        this.ws.send(JSON.stringify({ action, squadId }));
    };

    /** Sends action with info from unity */
    public sendUnityAction(action: string, match: any) {
        this.ws.send(JSON.stringify({ action, match }));
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public actionAllowAddress(WalletAddress: string, nonce: number) {
        const action: string = 'allowAddress';

        this.ws.send(JSON.stringify({ action, WalletAddress, nonce }));
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public casperActionAllowAddress(casperWallet: string, walletType: string, squadId: string) {
        const action: string = 'allowAddress';

        this.ws.send(JSON.stringify({ action, casperWallet, walletType, squadId }));
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

    /** opens and initialize connection */
    public openConnection() {
        this.ws.onopen = () => {
            this.ws.send('hello');
        };
    }

    /** Closes ws connection. */
    public close() {
        this.ws.close();
    };
};
