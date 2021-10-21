// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useState } from 'react';

import MetaMaskOnboarding from '@metamask/onboarding';
import { ServicePlugin } from '@/app/plugins/service';
import { NFT_ABI, NFT_ABI_SALE } from '@/app/ethers';

import './index.scss';

export const MintButton: React.FC = () => {
    const onboarding = React.useRef<MetaMaskOnboarding>();
    const [connectError, handleError] = useState(false);
    const [text, setButtonText] = useState('Mint');

    React.useEffect(() => {
        if (!onboarding.current) {
            onboarding.current = new MetaMaskOnboarding();
        }
    }, []);

    const connect = async () => {

        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                await window.ethereum.request({ method: 'eth_requestAccounts' });

                setButtonText('Connected');

            } catch (error: any) {
                console.log(error.message);
                handleError(true);

                setTimeout(() => {
                    handleError(false);
                }, 3000);
            }

        } else {
            onboarding.current
                && onboarding.current?.startOnboarding();
        }

        /** for testing */
        const service = ServicePlugin.create();
        const wallet = await service.getWallet();
        const totalSupply = await service.getLastTokenId(wallet, NFT_ABI);
        const contract = await service.sendTransaction(wallet, totalSupply, NFT_ABI_SALE);
    };

    return (
        <button
            className="ultimatedivision-mint-btn"
            onClick={connect}
        >
            {
                connectError
                && <span className="error">Please open metamask manually!</span>
            }
            <span className="ultimatedivision-mint-btn__text">{text}</span>
        </button>
    );
};