// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';

import { CHANGE_PASSWORD, LOGIN } from '../actions/users';

/** implementation of user state */
export class UsersState {
    public readonly userService: UserService;
    public user = {
        email: '',
        password: '',
    };
    public constructor(userService: UserService) {
        this.userService = userService;
    };
};

const client = new UserClient();
const service = new UserService(client);

export const usersReducer = (
    state = new UsersState(service),
    action: any = {}
) => {
    const user = state.user;

    switch (action.type) {
        case LOGIN:
            user.email = action.user.email;
            user.password = action.user.password;
            break;
        case CHANGE_PASSWORD:
            user.password = action.passwords.newPassword;
            break;
        default:
            break;
    };

    return { ...state };
};
