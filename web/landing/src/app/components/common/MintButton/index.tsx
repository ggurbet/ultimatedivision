// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import React from 'react';

import './index.scss';

type MintButtonProps = {
    text: string,
}

export const MintButton: React.FC<MintButtonProps> = (
    { text }) => {

    return (
        <button className="ultimatedivision-mint-btn">
            <span className="ultimatedivision-mint-btn__text">{text}</span>
        </button>
    );
};
