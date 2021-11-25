// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Ethereum interface */
interface ExternalProvider {
    isMetaMask?: boolean;
    isStatus?: boolean;
    host?: string;
    path?: string;
    sendAsync?: (request: { method: string; params?: any[] }, callback: (error: any, response: any) => void) => void;
    send?: (request: { method: string; params?: any[] }, callback: (error: any, response: any) => void) => void;
    request?: (request: { method: string; params?: any[] }) => Promise<any>;
}

/** adds window.ethereum field type */
declare global {
    interface Window {
        ethereum: ExternalProvider;
    }
}

export {};
