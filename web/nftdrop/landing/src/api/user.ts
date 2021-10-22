// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { APIClient } from '.';

/** Client for user controller of api */
export class UserClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/subscribers';
    /** handles the logic of user subscription to news by email */
    public async getNotifications(email: string): Promise<void> {
        const path = `${this.ROOT_PATH}`;
        const response = await this.http.post(path, JSON.stringify({email}));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
};
