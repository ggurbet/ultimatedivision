// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';

import { RouteConfig } from '@/app/router';

import { User } from '@/user';
import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';

/** action types implementation */
export const REGISTER = 'REGISTER';
export const LOGIN = 'LOGIN';
export const CHANGE_PASSWORD = 'CHANGE_PASSWORD';
export const RECOVER_PASSWORD = 'RECOVER_PASSWORD';
/** implement registration of new user */
export const register = (user: User) => ({
    type: REGISTER,
    user,
});
/** get registred user by id */
export const login = (email: string, password: string) => ({
    type: LOGIN,
    user: {
        email,
        password,
    }
});
/** changing user password */
export const changePassword = (password: string, newPassword: string) => ({
    type: CHANGE_PASSWORD,
    passwords: {
        password,
        newPassword,
    }
});
/** recover user password */
export const recoverPassword = (password: string) => ({
    type: RECOVER_PASSWORD,
    password
});

const client = new UserClient();
const users = new UserService(client);

/** thunk that implements user registration */
export const registerUser = (user: User) =>
    async function (dispatch: Dispatch) {
        try {
            await users.register(user);
            dispatch(register(user));
            location.pathname = RouteConfig.SignIn.path;
        } catch (error: any) {
            // TODO: rework catching errors
            /* eslint-disable */
            console.log(error.message);
        };
    };

/** thunk that implements user login */
export const loginUser = (email: string, password: string) =>
    async function (dispatch: Dispatch) {
        const whitepaperPath = '/whitepaper';
        try {
            await users.login(email, password);
            dispatch(login(email, password));
            location.pathname = whitepaperPath;
        } catch (error: any) {
            // TODO: rework catching errors
            /* eslint-disable */
            console.log(error.message);
        };
    };

/** thunk that implements user changing password */
export const changeUserPassword = (password: string, newPassword: string) =>
    async function (dispatch: Dispatch) {
        const marketplacePath = '/marketplace';
        try {
            await users.changePassword(password, newPassword);
            dispatch(changePassword(password, newPassword));
            location.pathname = marketplacePath;
        } catch (error: any) {
            // TODO: rework catching errors
            /* eslint-disable */
            console.log(error.message);
        };
    };

/** thunk that implements user reset password */
export const recoverUserPassword = (password: string) =>
    async function (dispatch: Dispatch) {
        try {
            await users.recoverPassword(password);
            dispatch(recoverPassword(password));
            location.pathname = RouteConfig.SignIn.path;
        } catch (error: any) {
            /** TODO: rework catching errros */
            /* eslint-disable */
            console.log(error.message);
        }
    };
