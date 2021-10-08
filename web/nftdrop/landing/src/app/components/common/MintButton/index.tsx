// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useState } from 'react';

import MetaMaskOnboarding from '@metamask/onboarding';


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
            onboarding.current &&
            onboarding.current?.startOnboarding();
        }
    };

    return (
        <button className="ultimatedivision-mint-btn"
            data-aos="fade-right"
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
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