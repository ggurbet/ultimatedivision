// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { UsersClient } from '@/api/users';
import { User } from '@/users';
import { UsersService } from '@/users/service';

import {
    LOGIN,
    SET_USER,
} from '../actions/users';

/**
 * UsersState is a representation of users reducer state.
 */
export class UsersState {
    public readonly userService: UsersService;
    public user: User = new User();
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
    switch (action.type) {
    case SET_USER:
        state.user = action.user;
        break;
    default:
        break;
    };

    return { ...state };
};
