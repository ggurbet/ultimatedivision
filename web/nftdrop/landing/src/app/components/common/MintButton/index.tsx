// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useState } from 'react';

import { ServicePlugin } from '@/app/plugins/service';

import './index.scss';

export const MintButton: React.FC = () => {
    const [text, changeText] = useState('Mint');

    const service = ServicePlugin.create();
    const connect = async () => {
        //@ts-ignore
        const account = await service.connectMetamask();

        account && changeText('Connected');
    };

    return (
        <button className="ultimatedivision-mint-btn"
            data-aos="fade-right"
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
            onClick={connect}
        >
            <span className="ultimatedivision-mint-btn__text">{text}</span>
        </button>
    );
};