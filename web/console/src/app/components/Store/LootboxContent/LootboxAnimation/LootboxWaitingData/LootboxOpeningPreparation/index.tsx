// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import startOpening from '@static/img/StorePage/BoxContent/StartOpening.gif';
import middleOpening from '@static/img/StorePage/BoxContent/MiddleOpening.gif';

export const LootboxOpeningPreparation = () => {
    const ANIMATION_LOOTBOX_START_OPENING_DELAY = 650;

    const [isStartOpening, handleStartOpening] = useState<boolean>(true);

    useEffect(() => {
        const startOpening = setTimeout(() => handleStartOpening(false), ANIMATION_LOOTBOX_START_OPENING_DELAY);

        return () => {
            clearTimeout(startOpening);
        };
    });

    return (
        <>
            {isStartOpening ?
                <img src={startOpening} alt="start opening box" className={'box-animation__box--start-opening'} />
                :
                <img src={middleOpening} alt="middle opening box" className={'box-animation__box--middle-opening'} />
            }
        </>
    );
};
