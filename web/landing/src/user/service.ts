// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { User } from '.';
import { UserClient } from '@/api/user';

/** exposes all user related logic  */
export class UserService {
    private readonly users: UserClient;
    public constructor(users: UserClient) {
        this.users = users;
    };
    /** handles user registration */
    public async register(user: User): Promise<void> {
        return await this.users.register(user);
    };
    /** return registred user */
    public async login(email: string, password: string): Promise<void> {
        return await this.users.login(email, password);
    };
    /** handles user changing password */
    public async changePassword(password: string, newPassword: string): Promise<void> {
        return await this.users.changePassword(password, newPassword);
    };
};
