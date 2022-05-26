// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import MetaMaskOnboarding from '@metamask/onboarding';
import { useMemo } from 'react';
import { useHistory } from 'react-router-dom';
import { toast } from 'react-toastify';

import representLogo from '@static/img/login/represent-logo.gif';
import metamask from '@static/img/login/metamask-icon.svg';
import velas from '@static/img/login/velas-icon.svg';
import casper from '@static/img/login/casper-icon.svg';

import { RouteConfig } from '@/app/routes';
import { ServicePlugin } from '@/app/plugins/service';
import { EthersClient } from '@/api/ethers';
import { NotFoundError } from '@/api';
import { SignedMessage } from '@/app/ethers';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';

import './index.scss';

export const SignIn = () => {
    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);
    const ethersService = useMemo(() => ServicePlugin.create(), []);
    const client = useMemo(() => new EthersClient(), []);

    const history = useHistory();
    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    /** Login with matamask. */
    const login: () => Promise<void> = async() => {
        if (!MetaMaskOnboarding.isMetaMaskInstalled()) {
            onboarding.startOnboarding();

            return;
        }
        await window.ethereum.request({
            method: 'eth_requestAccounts',
        });
        try {
            const address = await ethersService.getWallet();
            const message = await client.getNonce(address);
            const signedMessage = await ethersService.signMessage(message);
            await client.login(new SignedMessage(message, signedMessage));
            history.push(RouteConfig.MarketPlace.path);
            setLocalStorageItem('IS_LOGGED_IN', true);
        } catch (error: any) {
            if (!(error instanceof NotFoundError)) {
                toast.error('Something went wrong', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored',
                });

                return;
            }
            try {
                const signedMessage = await ethersService.signMessage('Register with metamask');
                await client.register(signedMessage);
                const address = await ethersService.getWallet();
                const message = await client.getNonce(address);
                const signedNonce = await ethersService.signMessage(message);
                await client.login(new SignedMessage(message, signedNonce));
                history.push(RouteConfig.MarketPlace.path);
                setLocalStorageItem('IS_LOGGED_IN', true);
            } catch (error: any) {
                toast.error('Something went wrong', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored',
                });
            }
        }
    };

    return (
        <div className="login">
            <div className="login__wrapper">
                <div className="login__represent">
                    <img src={representLogo} alt="utlimate division logo" className="login__represent__logo" />
                </div>
                <div className="login__content">
                    <h1 className="login__content__title">LOGIN</h1>
                    <div className="login__content__log-in">
                        <div onClick={login} className="login__content__log-in__item">
                            <img src={metamask} alt="Metamask logo" className="login__content__log-in__item__logo" />
                            <p className="login__content__log-in__item__text">Connect metamask</p>
                        </div>
                        <div className="login__content__log-in__item">
                            <img src={velas} alt="Velas logo" className="login__content__log-in__item__logo" />
                            <p className="login__content__log-in__item__text">Connect velas account</p>
                        </div>
                        <div className="login__content__log-in__item">
                            <img src={casper} alt="Casper logo" className="login__content__log-in__item__logo" />
                            <p className="login__content__log-in__item__text">Connect casper signer </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
