// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { User } from '@/users';
import { APIClient } from '.';

/**
 * UsersClient is a http implementation of users API.
 * Exposes all users-related functionality.
 */
export class UsersClient extends APIClient {
    private readonly ROOT_PATH: string = '/api/v0/auth';
    /** exposes user registration logic */
    public async register(user: User): Promise<void> {
        const path = `${this.ROOT_PATH}/register`;
        const response = await this.http.post(path, JSON.stringify(user));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** exposes user login logic */
    public async login(email: string, password: string): Promise<void> {
        const path = `${this.ROOT_PATH}/login`;
        const response = await this.http.post(
            path,
            JSON.stringify({
                email,
                password,
            })
        );

        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
