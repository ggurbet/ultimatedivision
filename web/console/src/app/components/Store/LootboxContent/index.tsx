// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useEffect, useState } from 'react';

import { LootboxOpening } from './LootboxOpening';
import { LootboxKeeping } from './LootboxKeeping';

import './index.scss';

export const LootboxContent: React.FC<{ handleOpening: Dispatch<SetStateAction<boolean>> }> = ({ handleOpening }) => {
    const [isAnimated, handleAnimation] = useState(true);

    useEffect(() => {
        const TIMEOUT = 8000;
        setTimeout(() => handleAnimation(false), TIMEOUT);
    });

    return (
        <div className="box-content">
            {isAnimated ?
                <LootboxOpening />
                :
                <LootboxKeeping handleOpening={handleOpening} />
            }
        </div>
    );
};

