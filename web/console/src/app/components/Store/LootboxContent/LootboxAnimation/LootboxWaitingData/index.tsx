// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { LootboxOpeningPreparation } from './LootboxOpeningPreparation';

import fallingBox from '@static/img/StorePage/BoxContent/fallingBox.gif';

export const LootboxWaitingData = () => {
    const [isFallenBox, handleFallenBox] = useState<boolean>(true);

    const ANIMATION_LOOTBOX_FALLING_DELAY = 780;

    useEffect(() => {
        const falletLootbox = setTimeout(() => handleFallenBox(false), ANIMATION_LOOTBOX_FALLING_DELAY);

        return () => {
            clearTimeout(falletLootbox);
        };
    });

    return (
        <div>
            {isFallenBox ?
                <img src={fallingBox} alt="falling box" className={'box-animation__box--falling'} />
                :
                <LootboxOpeningPreparation />
            }
        </div>
    );
};
