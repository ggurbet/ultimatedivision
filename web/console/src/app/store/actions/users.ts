// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { User } from '@/users';
import { UsersClient } from '@/api/users';
import { UsersService } from '@/users/service';

/** action types implementation */
export const REGISTER = 'REGISTER';
export const LOGIN = 'LOGIN';
export const CHANGE_PASSWORD = 'CHANGE_PASSWORD';
export const RECOVER_PASSWORD = 'RECOVER_PASSWORD';
export const SET_USER = 'SET_USER';

/** register action contains type and data for user registration */
export const register = (user: User) => ({
    type: REGISTER,
    user,
});
/** login action contains type and data for user login */
export const login = (email: string, password: string) => ({
    type: LOGIN,
    user: {
        email,
        password,
    },
});
/** register action contains type and data for user registration */
export const setUser = (user:User) => ({
    type: SET_USER,
    user,
});
/** changePassword action contains type and data for changes password */
export const changePassword = (password: string, newPassword: string) => ({
    type: CHANGE_PASSWORD,
    passwords: {
        password,
        newPassword,
    },
});
/** recoverPassword action contains type and data for recover password */
export const recoverPassword = (password: string) => ({
    type: RECOVER_PASSWORD,
    password,
});

const usersClient = new UsersClient();
const usersService = new UsersService(usersClient);

/** thunk that implements user registration */
export const registerUser = (user: User) =>
    async function(dispatch: Dispatch) {
        await usersService.register(user);
        dispatch(register(user));
    };

/** thunk that implements user login */
export const loginUser = (email: string, password: string) =>
    async function(dispatch: Dispatch) {
        await usersService.login(email, password);
        dispatch(login(email, password));
    };


export const setCurrentUser = () =>
    async function(dispatch: Dispatch) {
        const user = await usersService.getUser();
        dispatch(setUser(user));
    };
