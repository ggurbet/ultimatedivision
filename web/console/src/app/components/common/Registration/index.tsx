// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo } from 'react';
import { useHistory } from 'react-router-dom';

import MetaMaskOnboarding from '@metamask/onboarding';
// @ts-ignore
import KeyStorageHandler from '@/velas/keyStorageHandler';
// @ts-ignore
import StorageHandler from '@/velas/storageHandler';
// @ts-ignore
import { VAClient } from '@velas/account-client';
import { JSEncrypt } from 'jsencrypt';
import { CLPublicKey, Signer } from 'casper-js-sdk';

import { useLocalStorage } from '@/app/hooks/useLocalStorage';
import { RouteConfig } from '@/app/routes';
import { ServicePlugin } from '@/app/plugins/service';
import { EthersClient } from '@/api/ethers';
import { NotFoundError } from '@/api';
import { SignedMessage } from '@/ethers';
import { ToastNotifications } from '@/notifications/service';

import { VelasClient } from '@/api/velas';
import { VelasService } from '@/velas/service';
import { CasperNetworkClient } from '@/api/casper';
import { CasperNetworkService } from '@/casper/service';
import { ACCOUNT_HASH_PREFIX } from '@/casper/types';

import representLogo from '@static/img/login/represent-logo.gif';
import metamask from '@static/img/login/metamask-icon.svg';
import velas from '@static/img/login/velas-icon.svg';
import casper from '@static/img/login/casper-icon.svg';
import closeButton from '@static/img/login/close-icon.svg';

import './index.scss';

export const RegistrationPopup: React.FC<{ closeRegistrationPopup: () => void }> = ({ closeRegistrationPopup }) => {
    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);
    const ethersService = useMemo(() => ServicePlugin.create(), []);
    const client = useMemo(() => new EthersClient(), []);

    const velasClient = new VelasClient();
    const velasService = new VelasService(velasClient);

    const casperClient = new CasperNetworkClient();
    const casperService = new CasperNetworkService(casperClient);

    const history = useHistory();

    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    /** generates vaclient with the help of creds  */
    const vaclientService = async() => {
        try {
            const creds = await velasService.vaclientCreds();

            const vaclient = new VAClient({
                mode: 'redirect',
                clientID: creds.clientId,
                redirectUri: creds.redirectUri,
                StorageHandler,
                KeyStorageHandler,
                accountProviderHost: creds.accountProviderHost,
                networkApiHost: creds.networkApiHost,
                transactionsSponsorApiHost: creds.transactionsSponsorApiHost,
                transactionsSponsorPubKey: creds.transactionsSponsorPubKey,
            });

            return vaclient;
        } catch (error: any) {
            ToastNotifications.notify(error.message);
        }

        return null;
    };

    const processAuthResult = (error: any, authResult: any) => {
        if (authResult && authResult.access_token_payload) {
            window.history.replaceState({}, document.title, window.location.pathname);
        } else if (error) {
            ToastNotifications.notify(error.message);
        }
    };

    const contentVelas = async() => {
        try {
            const csrfToken = await velasService.csrfToken();

            const vaclient = await vaclientService();
            vaclient.authorize(
                {
                    csrfToken: csrfToken,
                    scope: 'authorization',
                    challenge: 'some_challenge_from_backend',
                },
                processAuthResult
            );
        } catch (error: any) {
            ToastNotifications.notify(error.message);
        }
    };

    const loginMetamask = async() => {
        const address = await ethersService.getWallet();
        const message = await client.getNonce(address);
        const signedMessage = await ethersService.signMessage(message);
        await client.login(new SignedMessage(message, signedMessage));
        history.push(RouteConfig.Store.path);

        setLocalStorageItem('IS_LOGGINED', true);
    };

    const loginCasper = async(publicKey: string) => {
        const encrypt = new JSEncrypt();
        const message = await casperService.nonce(publicKey);

        encrypt.setPublicKey(message);
        const encrypted = encrypt.encrypt(publicKey);

        if (encrypted) {
            await casperService.login(message, encrypted);
            history.push(RouteConfig.Store.path);
            window.location.reload();
        }
    };

    const casperRegistration = async() => {
        try {
            const publicKey = await Signer.getActivePublicKey();

            await loginCasper(publicKey);

            setLocalStorageItem('IS_LOGGINED', true);
        } catch (error: any) {
            if (!(error instanceof NotFoundError)) {
                ToastNotifications.notFound();

                return;
            }

            try {
                const publicKey = await Signer.getActivePublicKey();
                const accountHash = CLPublicKey.fromHex(publicKey).toAccountHashStr();
                const convertedAccountHash = accountHash.replace(ACCOUNT_HASH_PREFIX, '');

                await casperService.register(publicKey, convertedAccountHash);

                await loginCasper(publicKey);

                setLocalStorageItem('IS_LOGGINED', true);
            } catch (error: any) {
                ToastNotifications.couldNotLogInUserWithCasper();
            }
        }
    };

    const sendConnectionRequestCasper = () => {
        Signer.sendConnectionRequest();
    };

    /** Login with matamask. */
    const content: () => Promise<void> = async() => {
        if (!MetaMaskOnboarding.isMetaMaskInstalled()) {
            onboarding.startOnboarding();

            return;
        }
        // @ts-ignore
        await window.ethereum.request({
            method: 'eth_requestAccounts',
        });
        try {
            await loginMetamask();
        } catch (error: any) {
            if (!(error instanceof NotFoundError)) {
                ToastNotifications.notFound();

                return;
            }
            try {
                const signedMessage = await ethersService.signMessage('Register with metamask');
                await client.register(signedMessage);
                await loginMetamask();
            } catch (error: any) {
                ToastNotifications.couldNotLogInUserWithMetamask();
            }
        }
    };
    useEffect(() => {
        window.addEventListener('signer:connected', casperRegistration);

        return () => window.removeEventListener('signer:connected', casperRegistration);
    }, []);

    return (
        <div className="registration-pop-up">
            <div className="registration-pop-up__wrapper">
                <div className="registration-pop-up__wrapper__close" onClick={closeRegistrationPopup}>
                    <img src={closeButton} alt="close button" className="registration-pop-up__wrapper__close__icon" />
                </div>
                <div className="registration-pop-up__block">
                    <div className="registration-pop-up__represent">
                        <img
                            src={representLogo}
                            alt="represent logo"
                            className="registration-pop-up__represent__logo"
                        />
                    </div>
                    <div className="registration-pop-up__content">
                        <h1 className="registration-pop-up__content__title">LOGIN</h1>
                        <div className="registration-pop-up__content__block">
                            <div onClick={content} className="registration-pop-up__content__block__item">
                                <img
                                    src={metamask}
                                    alt="Metamask logo"
                                    className="registration-pop-up__content__block__item__logo"
                                />
                                <p className="registration-pop-up__content__block__item__text">Connect metamask</p>
                            </div>
                            <div onClick={contentVelas} className="registration-pop-up__content__block__item">
                                <img
                                    src={velas}
                                    alt="Velas logo"
                                    className="registration-pop-up__content__block__item__logo"
                                />
                                <p className="registration-pop-up__content__block__item__text">Connect velas account</p>
                            </div>
                            <div
                                onClick={sendConnectionRequestCasper}
                                className="registration-pop-up__content__block__item"
                            >
                                <img
                                    src={casper}
                                    alt="Casper logo"
                                    className="registration-pop-up__content__block__item__logo"
                                />
                                <p className="registration-pop-up__content__block__item__text">Connect casper signer</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
