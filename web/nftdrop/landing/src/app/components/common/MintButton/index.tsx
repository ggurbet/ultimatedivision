// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useRef, useState } from 'react';
import MetaMaskOnboarding from '@metamask/onboarding';
import { toast } from 'react-toastify';

import { NFT_ABI, NFT_ABI_SALE } from '@/app/ethers';
import { ServicePlugin } from '@/app/plugins/service';

import './index.scss';

export const MintButton: React.FC = () => {
    const onboarding = useRef<MetaMaskOnboarding>();
    const service = ServicePlugin.create();
    const [text, setButtonText] = useState('Connect');
    const [isConnected, handleConnect] = useState(false);

    useEffect(() => {
        if (!onboarding.current) {
            onboarding.current = new MetaMaskOnboarding();
        }
    }, []);

    const connect = async () => {
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                await window.ethereum.request({ method: 'eth_requestAccounts' });

                handleConnect(true);
            } catch (error: any) {
                toast.error('Please open metamask manually!', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored'
                });

                return;
            }
        } else {
            onboarding.current = new MetaMaskOnboarding();
            onboarding.current?.startOnboarding();
        }

        try {
            const wallet = await service.getWallet();

            await service.getAddress(wallet);
            setButtonText('Mint');
        } catch (error: any) {
            toast.error('You are not in whitelist', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored'
            });
        }
    };

    const sendTransaction = async () => {
        try {
            const wallet = await service.getWallet();

            await service.sendTransaction(wallet, NFT_ABI_SALE);
        } catch (error: any) {
            let errorMessage = 'Failed to connect to contract';
            const presaleErrorCode = -32603;

            if (error.error.code === presaleErrorCode) {
                errorMessage = 'Only one token can be bought on presale';
            };

            toast.error(errorMessage, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored'
            });
        };
    };

    return (
        <button
            className="ultimatedivision-mint-btn"
            onClick={() => {
                !isConnected
                    ? connect()
                    : sendTransaction();
            }}
        >
            <span className="ultimatedivision-mint-btn__text">{text}</span>
        </button>
    );
};