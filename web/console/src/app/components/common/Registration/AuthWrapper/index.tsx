// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import { toast } from 'react-toastify';

// @ts-ignore
import KeyStorageHandler from '../../../../velas/keyStorageHandler';
// @ts-ignore
import StorageHandler from '../../../../velas/storageHandler';
// @ts-ignore
import { VAClient } from '@velas/account-client';

import { RouteConfig } from '@/app/routes';
import { InternalError, NotFoundError } from '@/api';
import { VelasClient } from '@/api/velas';
import { VelasService } from '@/velas/service';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';

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
        } catch (e) {
            toast.error(`${e}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
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
        } catch (e) {
            if (!(e instanceof NotFoundError)) {
                toast.error('Something went wrong', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored',
                });

                return;
            }
            try {
                await velasService.register(
                    result.userinfo.account_key_evm,
                    authResult.access_token,
                    authResult.expires_at
                );
                await velasLogin(result.userinfo.account_key_evm, authResult.access_token, authResult.expires_at);
            } catch (e) {
                toast.error('Something went wrong', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored',
                });
            }
        }
    };

    const sendAuthData = async(authResult: any) => {
        try {
            const vaclient = await vaclientService();

            await vaclient.userinfo(authResult.access_token, async(err: any, result: any) => {
                if (err) {
                    toast.error('Something went wrong', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    });
                } else {
                    await velasRegister(result, authResult);
                }
            });
        } catch (error) {
            if (error instanceof InternalError) {
                history.push(RouteConfig.Home.path);
                toast.error('Registration failed', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored',
                });
            }

            toast.error('Something went wrong', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    const processAuthResult = (e: any, authResult: any) => {
        if (authResult && authResult.access_token_payload) {
            window.history.replaceState({}, document.title, window.location.pathname);

            sendAuthData(authResult);
        } else if (e) {
            window.history.replaceState({}, document.title, window.location.pathname);

            toast.error(`${e.description}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    const authorization = async() => {
        try {
            const vaclient = await vaclientService();
            vaclient.handleRedirectCallback(processAuthResult);
        } catch (e) {
            toast.error(`${e}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
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
