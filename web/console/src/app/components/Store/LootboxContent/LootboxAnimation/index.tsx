// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useState } from 'react';

import { LootboxOpening } from './LootboxOpening';
import { LootboxWaitingData } from './LootboxWaitingData';

import './index.scss';

export const LootboxAnimation: React.FC<{
    isOpenedLootbox: boolean;
    handleOpenedLootbox: Dispatch<SetStateAction<boolean>>;
    handleLootboxKeeping: Dispatch<SetStateAction<boolean>>;
}> = ({ isOpenedLootbox, handleOpenedLootbox, handleLootboxKeeping }) => {
    const [isOpeningLootbox, handleOpeningLootbox] = useState<boolean>(true);

    return (
        <div className="box-animation">
            <div
                className={`box-animation__box-container ${
                    !isOpeningLootbox ? 'box-animation__box-container--opened' : ''
                }`}
            >
                {!isOpenedLootbox ?
                    <LootboxWaitingData />
                    :
                    <LootboxOpening
                        handleOpenedLootbox={handleOpenedLootbox}
                        handleOpeningLootbox={handleOpeningLootbox}
                        isOpeningLootbox={isOpeningLootbox}
                        handleLootboxKeeping={handleLootboxKeeping}
                    />
                }
            </div>
        </div>
    );
};
