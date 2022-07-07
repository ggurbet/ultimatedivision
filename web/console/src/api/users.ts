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
    /** changes user password */
    public async changePassword(password: string, newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/change-password`;
        const response = await this.http.post(
            path,
            JSON.stringify({
                password,
                newPassword,
            })
        );

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** checks user token by email confirmation */
    public async checkEmailToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/email/confirm/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** checks user recover token */
    public async checkRecoverToken(token: string | null): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password/${token}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** recovers user password */
    public async recoverPassword(newPassword: string): Promise<void> {
        const path = `${this.ROOT_PATH}/reset-password`;
        const response = await this.http.patch(path, JSON.stringify({ newPassword }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** resets user password by email confirmation */
    public async sendEmailForPasswordReset(email: string): Promise<void> {
        const path = `${this.ROOT_PATH}/password/${email}`;
        const response = await this.http.get(path, JSON.stringify(email));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** sends data to register user with velas wallet */
    public async velasRegister(walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        const path = `${this.ROOT_PATH}/velas/register`;
        const response = await this.http.post(path, JSON.stringify({ walletAddress, accessToken, expiresAt }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
    /** sends address to get nonce to login user */
    public async velasNonce(address: string): Promise<string> {
        const walletType = 'velas_wallet_address';

        const path = `${this.ROOT_PATH}/velas/nonce?address=${address}&walletType=${walletType}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const status = response.json();

        return status;
    }
    /** sends data to login user with velas wallet */
    public async velasLogin(nonce: string, walletAddress: string, accessToken: string, expiresAt: any): Promise<void> {
        const path = `${this.ROOT_PATH}/velas/login`;
        const response = await this.http.post(path, JSON.stringify({ walletAddress, accessToken, expiresAt, nonce }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** gets token to login user with velas wallet */
    public async velasCsrfToken(): Promise<string> {
        const path = 'https://velas.ultimatedivision.com/csrf';
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result.token;
    }

    /** gets creds to fill velas vaclient */
    public async velasVaclientCreds(): Promise<any> {
        const path = `${this.ROOT_PATH}/velas/vaclient`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const result = await response.json();

        return result;
    }
    /** Sends signed message and registers user */
    public async casperRegister(walletAddress: string): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/casper/register`, JSON.stringify(walletAddress));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Gets message from API for sign with casper */
    public async casperNonce(walletAddress: string): Promise<string> {
        const path = `${this.ROOT_PATH}/casper/nonce?address=${walletAddress}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }

        return await response.json();
    }

    /** Sends signed message, and logs-in */
    public async casperLogin(nonce: string, signature: string): Promise<void> {
        const response = await this.http.post(`${this.ROOT_PATH}/casper/login`, JSON.stringify({ nonce, signature }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
