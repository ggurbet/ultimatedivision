// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { WebSocketClient } from '../api/websockets';

/** Exposes all queue related logic. */
export class WebSocketService {
    public wsConnectionClient: WebSocketClient = new WebSocketClient();

    /** Changes current conection client. */
    public changeWSConnectionClient() {
        this.wsConnectionClient = new WebSocketClient();
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public actionAllowAddress(wallet: string, nonce: number): void {
        this.wsConnectionClient.actionAllowAddress(wallet, nonce);
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public casperActionAllowAddress(wallet: string, walletType: string, squadId: string): void {
        this.wsConnectionClient.casperActionAllowAddress(wallet, walletType, squadId);
    };

    /** Sends action that indicates that the client is forbidden to add wallet address. */
    public actionForbidAddress(): void {
        this.wsConnectionClient.actionForbidAddress();
    };

    /** Sends action to confirm or reject match. */
    public sendAction(action: string, squadId: string): void {
        this.wsConnectionClient.sendAction(action, squadId);
    };

    /** Sends action, i.e 'startSearch', 'finishSearch', on open webSocket connection. */
    public onOpenConnectionSendAction(action: string, squadId: string): void {
        this.wsConnectionClient.onOpenConnectionSendAction(action, squadId);
    };

    /** Closes ws connection. */
    public close() {
        this.wsConnectionClient.close();
    };

    /** Opens ws connection. */
    public openConnection() {
        this.wsConnectionClient.openConnection();
    };

    /** Sets match queue */
    public matchQueue() {
        this.wsConnectionClient.matchQueue();
    };
};

const webSocketService = new WebSocketService();

/** Sends action to confirm or reject match. */
export const sendAction = (action: string, squadId: string) => {
    webSocketService.sendAction(action, squadId);
};

/** Changes current queue client, and after sends action,
 * i.e 'startSearch', 'finishSearch', on open webSocket connection. */
export const onOpenConnectionSendAction = (action: string, squadId: string) => {
    webSocketService.changeWSConnectionClient();
    webSocketService.onOpenConnectionSendAction(action, squadId);
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const actionAllowAddress = (wallet: string, nonce: number) => {
    webSocketService.actionAllowAddress(wallet, nonce);
};

/** Sets match queue */
export const setMatchQueue = () => {
    webSocketService.matchQueue();
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const casperActionAllowAddress = (wallet: string, walletType: string, squadId: string) => {
    webSocketService.casperActionAllowAddress(wallet, walletType, squadId);
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const actionForbidAddress = () => {
    webSocketService.actionForbidAddress();
};

/** Returns current queue client. */
export const getCurrentWebSocketClient = () => webSocketService.wsConnectionClient;

/** Opens ws connection. */
export const onOpenConnection = () => webSocketService.openConnection();
