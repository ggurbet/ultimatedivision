// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { User } from '@/users';
import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';

/** action types implementation */
export const REGISTER = 'REGISTER';
export const LOGIN = 'LOGIN';
export const SET_USER = 'SET_USER';
/** register action contains type and data for user registration */
export const register = (user: User) => ({
    type: REGISTER,
    user,
});

/** register action contains type and data for user registration */
export const setUser = (user:User) => ({
    type: SET_USER,
    user,
});

const usersClient = new UsersClient();
const usersService = new UsersService(usersClient);

/** thunk that implements user registration */
export const registerUser = (user: User) =>
    async function(dispatch: Dispatch) {
        await usersService.register(user);
        dispatch(register(user));
    };

export const setCurrentUser = () =>
    async function(dispatch: Dispatch) {
        const user = await usersService.getUser();
        dispatch(setUser(user));
    };
