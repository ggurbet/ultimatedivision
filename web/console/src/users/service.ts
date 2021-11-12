// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { User } from '.';
import { UsersClient } from '@/api/users';

/**
 * Exposes all users related logic.
 */
export class UsersService {
    private readonly users: UsersClient;
    /** UsersService contains http implementation of users API  */
    public constructor(users: UsersClient) {
        this.users = users;
    };
    /** handles user registration */
    public async register(user: User): Promise<void> {
        await this.users.register(user);
    };
    /** handles user login */
    public async login(email: string, password: string): Promise<void> {
        await this.users.login(email, password);
    };
    /** changes user password */
    public async changePassword(password: string, newPassword: string): Promise<void> {
        await this.users.changePassword(password, newPassword);
    };
    /** checks user email token */
    public async checkEmailToken(token: string | null): Promise<void> {
        await this.users.checkEmailToken(token);
    };
    /** checks recover token */
    public async checkRecoverToken(token: string | null): Promise<void> {
        await this.users.checkRecoverToken(token);
    };
    /** recovers user password */
    public async recoverPassword(password: string): Promise<void> {
        await this.users.recoverPassword(password);
    };
    /** resets user password by email confirmation */
    public async sendEmailForPasswordReset(email: string): Promise<void> {
        await this.users.sendEmailForPasswordReset(email);
    };
};
