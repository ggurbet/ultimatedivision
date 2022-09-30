// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useEffect, useState } from 'react';

import { LootboxKeeping } from './LootboxKeeping';
import { LootboxAnimation } from './LootboxAnimation';

import './index.scss';

export const LootboxContent: React.FC<{
    handleOpenedLootbox: Dispatch<SetStateAction<boolean>>;
    isOpenedLootbox: boolean;
    handleLootboxSelection: Dispatch<SetStateAction<boolean>>;
    isLootboxKeeping: boolean;
    handleLootboxKeeping: Dispatch<SetStateAction<boolean>>;
}> = ({ handleOpenedLootbox, handleLootboxKeeping, isOpenedLootbox, isLootboxKeeping, handleLootboxSelection }) =>
    <div className="box-content">
        {isLootboxKeeping ?
            <LootboxKeeping handleLootboxSelection={handleLootboxSelection} />
            :
            <LootboxAnimation
                isOpenedLootbox={isOpenedLootbox}
                handleOpenedLootbox={handleOpenedLootbox}
                handleLootboxKeeping={handleLootboxKeeping}
            />
        }
    </div>;

