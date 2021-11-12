// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';

import {
    CHANGE_PASSWORD,
    LOGIN,
    RECOVER_PASSWORD,
} from '../actions/users';

/**
 * UsersState is a representation of users reducer state.
 */
export class UsersState {
    public readonly userService: UsersService;
    public user = {
        email: '',
        password: '',
    };
    /** UsersState contains service implementation of users  */
    public constructor(userService: UsersService) {
        this.userService = userService;
    };
};

const usersClient = new UsersClient();
const usersService = new UsersService(usersClient);

export const usersReducer = (
    state = new UsersState(usersService),
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
    case RECOVER_PASSWORD:
        user.password = action.password;
        break;
    default:
        break;
    };

    return { ...state };
};
