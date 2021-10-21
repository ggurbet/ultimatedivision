// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';

export class EthersClient extends APIClient {
    private readonly ROOT_PATH = 'http://localhost:8086/api/v0/whitelist/'

    public async getAddress(wallet: string) {
        const response = await this.http.get(`${this.ROOT_PATH}${wallet}`);

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }
}