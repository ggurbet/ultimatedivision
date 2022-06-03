// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useMemo } from "react";
import { useDispatch } from "react-redux";
import { useHistory } from "react-router-dom";
import { toast } from "react-toastify";

import MetaMaskOnboarding from "@metamask/onboarding";

import { useLocalStorage } from "@/app/hooks/useLocalStorage";
import { RouteConfig } from "@/app/routes";
import { ServicePlugin } from "@/app/plugins/service";
import { EthersClient } from "@/api/ethers";
import { NotFoundError } from "@/api";
import { SignedMessage } from "@/app/ethers";
import { UsersClient } from "@/api/users";
import { UsersService } from "@/users/service";

import representLogo from "@static/img/login/represent-logo.gif";
import metamask from "@static/img/login/metamask-icon.svg";
import velas from "@static/img/login/velas-icon.svg";
import casper from "@static/img/login/casper-icon.svg";

import "./index.scss";

// @ts-ignore
import { VAClient } from "@velas/account-client";
// @ts-ignore
import StorageHandler from "../../../../velas/storageHandler";
// @ts-ignore
import KeyStorageHandler from "../../../../velas/keyStorageHandler";

export const SignIn = () => {
    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);
    const ethersService = useMemo(() => ServicePlugin.create(), []);
    const client = useMemo(() => new EthersClient(), []);

    const usersClient = new UsersClient();
    const usersService = new UsersService(usersClient);
    const history = useHistory();

    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    /** generates vaclient with the help of creds  */
    const vaclientService = async () => {
        try {
            const creds = await usersService.velasVaclientCreds();

            const vaclient = new VAClient({
                mode: "redirect",
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
        } catch (e) {
            toast.error(`${e}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: "colored",
            });
        }

        return null;
    };

    const processAuthResult = (e: any, authResult: any) => {
        if (authResult && authResult.access_token_payload) {
            window.history.replaceState({}, document.title, window.location.pathname);
        } else if (e) {
            toast.error(`${e.description}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: "colored",
            });
        }
    };

    const loginVelas = async () => {
        try {
            const csrfToken = await usersService.velasCsrfToken();

            const vaclient = await vaclientService();
            vaclient.authorize(
                {
                    csrfToken: csrfToken,
                    scope: "authorization",
                    challenge: "some_challenge_from_backend",
                },
                processAuthResult
            );
        } catch (error: any) {
            toast.error(`${error}`, {
                position: toast.POSITION.TOP_RIGHT,
                theme: "colored",
            });
        }
    };

    /** Login with matamask. */
    const login: () => Promise<void> = async () => {
        if (!MetaMaskOnboarding.isMetaMaskInstalled()) {
            onboarding.startOnboarding();

            return;
        }
        await window.ethereum.request({
            method: "eth_requestAccounts",
        });
        try {
            const address = await ethersService.getWallet();
            const message = await client.getNonce(address);
            const signedMessage = await ethersService.signMessage(message);
            await client.login(new SignedMessage(message, signedMessage));
            history.push(RouteConfig.MarketPlace.path);
            setLocalStorageItem("IS_LOGGED_IN", true);
        } catch (error: any) {
            if (!(error instanceof NotFoundError)) {
                toast.error("Something went wrong", {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: "colored",
                });

                return;
            }
            try {
                const signedMessage = await ethersService.signMessage("Register with metamask");
                await client.register(signedMessage);
                const address = await ethersService.getWallet();
                const message = await client.getNonce(address);
                const signedNonce = await ethersService.signMessage(message);
                await client.login(new SignedMessage(message, signedNonce));
                history.push(RouteConfig.MarketPlace.path);
                setLocalStorageItem("IS_LOGGED_IN", true);
            } catch (error: any) {
                toast.error("Something went wrong", {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: "colored",
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
                        <div onClick={loginVelas} className="login__content__log-in__item">
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
