// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';

import { LootboxCard } from './LootboxCard';

import { LootboxStats, LootboxTypes } from '@/app/types/lootbox';

import './index.scss';

export const LootboxSelection: React.FC<{
    handleOpenedLootbox: Dispatch<SetStateAction<boolean>>;
    handleLootboxSelection: Dispatch<SetStateAction<boolean>>;
    handleLootboxKeeping: Dispatch<SetStateAction<boolean>>;
}> = ({ handleOpenedLootbox, handleLootboxSelection, handleLootboxKeeping }) => {
    const REGULAR_BOX_CARDS_QUANTITY = 5;
    const COOL_BOX_CARDS_QUANTITY = 10;
    /** TODO: remove test code */
    const boxesData = [
        new LootboxStats(
            '1',
            LootboxTypes['Regular Box'],
            REGULAR_BOX_CARDS_QUANTITY,
            // eslint-disable-next-line
            [80, 15, 4, 1],
            '200,000'
        ),
        new LootboxStats(
            '2',
            LootboxTypes['Cool box'],
            COOL_BOX_CARDS_QUANTITY,
            // eslint-disable-next-line
            [70, 20, 8, 2],
            '500,000'
        ),
    ];

    return (
        <div className="box-selection">
            <div className="box-selection__wrapper">
                {boxesData.map((item, index) =>
                    <LootboxCard
                        lootBoxStats={item}
                        key={index}
                        handleOpenedLootbox={handleOpenedLootbox}
                        handleLootboxKeeping={handleLootboxKeeping}
                        handleLootboxSelection={handleLootboxSelection}
                    />
                )}
            </div>
        </div>
    );
};
