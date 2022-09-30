// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useEffect, useState } from 'react';
import { useSelector } from 'react-redux';

import { PlayerCard } from '@components/common/PlayerCard';
import { RootState } from '@/app/store';
import { Card } from '@/card';

import boxOpening from '@static/img/StorePage/BoxContent/FinishOpening.gif';
import box from '@static/img/StorePage/BoxContent/box.png';

export const LootboxOpening: React.FC<{
    handleOpenedLootbox: Dispatch<SetStateAction<boolean>>;
    isOpeningLootbox: boolean;
    handleOpeningLootbox: Dispatch<SetStateAction<boolean>>;
    handleLootboxKeeping: Dispatch<SetStateAction<boolean>>;
}> = ({ handleOpenedLootbox, handleOpeningLootbox, isOpeningLootbox, handleLootboxKeeping }) => {
    const FIRST_CARD = 0;
    const ANIMATION_LOOTBOX_OPENING_DELAY = 900;
    const ANIMATION_LOOTBOX_CARD_APPEARNCE_DELAY = 5000;

    const cards: Card[] = useSelector((state: RootState) => state.lootboxReducer.lootbox);

    const handleEndOfOpening = () => {
        handleOpenedLootbox(false);
        handleLootboxKeeping(true);
    };

    useEffect(() => {
        const openingLootBox = setTimeout(() => handleOpeningLootbox(false), ANIMATION_LOOTBOX_OPENING_DELAY);
        const endOfOpeningLootbox = setTimeout(handleEndOfOpening, ANIMATION_LOOTBOX_CARD_APPEARNCE_DELAY);

        return () => {
            clearTimeout(openingLootBox);
            clearTimeout(endOfOpeningLootbox);
        };
    });

    return (
        <>
            <div>
                {!isOpeningLootbox ?
                    <img src={box} alt="box opened" className={'box-animation__box--opened'} />
                    :
                    <img src={boxOpening} alt="box opening" className={'box-animation__box--opening'} />
                }
            </div>
            <div className="box-animation__card__container">
                <div className="box-animation__card__container__backlight">
                    <PlayerCard className={'box-animation__card'} id={cards[FIRST_CARD].id} />
                </div>
            </div>
        </>
    );
};
