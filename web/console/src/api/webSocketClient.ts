// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/**
 * WebSocketAPICLient is base client that holds webSocket.
 */
export class WebSocketAPICLient {
    /** The WebSocket provides the API for creating and managing
     * a websocket connection to a server and for sending and
     * receiving data on the connection. */
    public readonly ws: WebSocket = new WebSocket('ws://localhost:8088/api/v0/queue');
};
