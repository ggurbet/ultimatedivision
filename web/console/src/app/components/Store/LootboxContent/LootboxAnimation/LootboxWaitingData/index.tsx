// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import middleOpening from '@static/img/StorePage/BoxContent/MiddleOpening.webp';
import fallingBox from '@static/img/StorePage/BoxContent/falling.webp';

export const LootboxWaitingData = () => {
    const [isFallenBox, handleFallenBox] = useState<boolean>(true);

    const ANIMATION_LOOTBOX_FALLING_DELAY = 2400;

    useEffect(() => {
        const fallenLootbox = setTimeout(() => handleFallenBox(false), ANIMATION_LOOTBOX_FALLING_DELAY);

        return () => {
            clearTimeout(fallenLootbox);
        };
    });

    return (
        <div>
            {isFallenBox ?
                <img src={fallingBox} alt="falling box" className={'box-animation__box--falling'} />
                :
                <img src={middleOpening} alt="middle opening box" className={'box-animation__box--middle-opening'} />
            }
        </div>
    );
};
