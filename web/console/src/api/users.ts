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
        };
    };
    /** exposes user login logic */
    public async login(email: string, password: string): Promise<void> {
        const path = `${this.ROOT_PATH}/login`;
        const response = await this.http.post(path, JSON.stringify({
            email, password,
        }));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** changes user password */
    public async changePassword(password: string, newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/change-password`;
        const response = await this.http.post(path, JSON.stringify({
            password, newPassword,
        }));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** checks user token by email confirmation */
    public async checkEmailToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/email/confirm/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** checks user recover token */
    public async checkRecoverToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** recovers user password */
    public async recoverPassword(newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password`;
        const response = await this.http.patch(path, JSON.stringify({newPassword}));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
    /** resets user password by email confirmation */
    public async sendEmailForPasswordReset(email: string): Promise<void> {
        const path = `${this.ROOT_PATH}/password/${email}`;
        const response = await this.http.get(path, JSON.stringify(email));

        if (!response.ok) {
            await this.handleError(response);
        };
    };
};
