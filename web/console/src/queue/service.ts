// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { QueueClient } from '../api/queue';

/** Exposes all queue related logic. */
export class QueueService {
    public queueClient: QueueClient = new QueueClient();

    /** Changes current queue client. */
    public changeQueueClient() {
        this.queueClient = new QueueClient();
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public actionAllowAddress(wallet: string, nonce: number): void {
        this.queueClient.actionAllowAddress(wallet, nonce);
    };

    /** Sends action that indicates that the client allows to add address of wallet. */
    public casperActionAllowAddress(wallet: string, walletType: string, squadId:string): void {
        this.queueClient.casperActionAllowAddress(wallet, walletType, squadId);
    };

    /** Sends action that indicates that the client is forbidden to add wallet address. */
    public actionForbidAddress(): void {
        this.queueClient.actionForbidAddress();
    };

    /** Sends action to confirm or reject match. */
    public sendAction(action: string, squadId: string): void {
        this.queueClient.sendAction(action, squadId);
    };

    /** Sends action, i.e 'startSearch', 'finishSearch', on open webSocket connection. */
    public onOpenConnectionSendAction(action: string, squadId: string): void {
        this.queueClient.onOpenConnectionSendAction(action, squadId);
    };
};

const queueService = new QueueService();

/** Sends action to confirm or reject match. */
export const queueSendAction = (action: string, squadId: string) => {
    queueService.sendAction(action, squadId);
};

/** Changes current queue client, and after sends action,
 * i.e 'startSearch', 'finishSearch', on open webSocket connection. */
export const onOpenConnectionSendAction = (action: string, squadId: string) => {
    queueService.changeQueueClient();
    queueService.onOpenConnectionSendAction(action, squadId);
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const queueActionAllowAddress = (wallet: string, nonce: number) => {
    queueService.actionAllowAddress(wallet, nonce);
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const queueCasperActionAllowAddress = (wallet: string, walletType: string, squadId:string) => {
    queueService.casperActionAllowAddress(wallet, walletType, squadId);
};

/** Sends action that indicates that the client allows to add address of wallet. */
export const actionForbidAddress = () => {
    queueService.actionForbidAddress();
};

/** Returns current queue client. */
export const getCurrentQueueClient = () => queueService.queueClient;
