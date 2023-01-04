// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useHistory } from 'react-router-dom';

// @ts-ignore
import KeyStorageHandler from '@/velas/keyStorageHandler';
// @ts-ignore
import StorageHandler from '@/velas/storageHandler';
// @ts-ignore
import { VAClient } from '@velas/account-client';

import { RouteConfig } from '@/app/routes';
import { InternalError, NotFoundError } from '@/api';
import { VelasClient } from '@/api/velas';
import { VelasService } from '@/velas/service';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';
import { ToastNotifications } from '@/notifications/service';

import ulimatedivisionLogo from '@static/img/registerPage/ultimate.svg';

import './index.scss';

const AuthWrapper = () => {
    const history = useHistory();

    const velasClient = new VelasClient();
    const velasService = new VelasService(velasClient);
    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    /** generates vaclient with the help of creds  */
    const vaclientService = async() => {
        try {
            const vaclientCreds = await velasService.vaclientCreds();

            const vaclient = new VAClient({
                mode: 'redirect',
                clientID: vaclientCreds.clientId,
                redirectUri: vaclientCreds.redirectUri,
                StorageHandler,
                KeyStorageHandler,
                accountProviderHost: vaclientCreds.accountProviderHost,
                networkApiHost: vaclientCreds.networkApiHost,
                transactionsSponsorApiHost: vaclientCreds.transactionsSponsorApiHost,
                transactionsSponsorPubKey: vaclientCreds.transactionsSponsorPubKey,
            });

            return vaclient;
        } catch (error: any) {
            ToastNotifications.notify(error.message);
        }

        return null;
    };

    /** logins via velas  */
    const velasLogin = async(accountKeyEvm: string, accessToken: string, expiresAt: number) => {
        const nonce = await velasService.nonce(accountKeyEvm);

        await velasService.login(nonce, accountKeyEvm, accessToken, expiresAt);

        setLocalStorageItem('IS_LOGGINED', true);
        history.push(RouteConfig.Store.path);
        window.location.reload();
    };

    /** register via velas */
    const velasRegister = async(result: any, authResult: any) => {
        try {
            await velasLogin(result.userinfo.account_key_evm, authResult.access_token, authResult.expires_at);
        } catch (error: any) {
            if (!(error instanceof NotFoundError)) {
                ToastNotifications.notFound();

                return;
            }
            try {
                await velasService.register(
                    result.userinfo.account_key_evm,
                    authResult.access_token,
                    authResult.expires_at
                );
                await velasLogin(result.userinfo.account_key_evm, authResult.access_token, authResult.expires_at);
            } catch (error: any) {
                ToastNotifications.couldNotLogInUserWithVelas();
            }
        }
    };

    const sendAuthData = async(authResult: any) => {
        try {
            const vaclient = await vaclientService();

            await vaclient.userinfo(authResult.access_token, async(err: any, result: any) => {
                if (!err) {
                    await velasRegister(result, authResult);
                }
            });
        } catch (error: any) {
            if (error instanceof InternalError) {
                history.push(RouteConfig.Home.path);
                ToastNotifications.registrationFailed();
            }

            ToastNotifications.notify(error.message);
        }
    };

    const processAuthResult = (error: any, authResult: any) => {
        if (authResult && authResult.access_token_payload) {
            window.history.replaceState({}, document.title, window.location.pathname);

            sendAuthData(authResult);
        } else if (error) {
            window.history.replaceState({}, document.title, window.location.pathname);

            ToastNotifications.notify(error.message);
        }
    };

    const authorization = async() => {
        try {
            const vaclient = await vaclientService();
            vaclient.handleRedirectCallback(processAuthResult);
        } catch (error: any) {
            ToastNotifications.couldNotLogInUserWithVelas();
        }
    };

    useEffect(() => {
        authorization();
    }, []);

    return (
        <div className="auth-wrapper">
            <img src={ulimatedivisionLogo} alt="ultimatedivision-logo" />
        </div>
    );
};
export default AuthWrapper;
