// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { GlobalWithFetchMock } from 'jest-fetch-mock';

const customGlobal: GlobalWithFetchMock = (global as unknown) as GlobalWithFetchMock;

const WebSocket = require('ws');
global.WebSocket = WebSocket;

customGlobal.fetch = require('jest-fetch-mock');

customGlobal.fetchMock = customGlobal.fetch;

// Disallow warnings and errors from console.
customGlobal.console.warn = (message) => { throw new Error(message); };

customGlobal.console.error = (message) => { throw new Error(message); };
