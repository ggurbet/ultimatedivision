// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { UserClient } from '@/api/user';

/** Exposes all user related logic. */
export class UserService {
    private readonly users: UserClient;
    public constructor(users: UserClient) {
        this.users = users;
    };

    /** Handles the logic of user subscription to news by email. */
    public async getNotifications(email: string): Promise<void> {
        await this.users.getNotifications(email);
    };
};
